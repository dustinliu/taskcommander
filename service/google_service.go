package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/adrg/xdg"
	"github.com/dustinliu/taskcommander/core"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	tasks "google.golang.org/api/tasks/v1"
)

const (
	client_id_key     = "TC_CLIENT_ID"
	client_secret_key = "TC_CLIENT_SECRET"
	client_id         = "50757396817-rfurg5f6dsdtag6dohrorqvkhtq39o48.apps.googleusercontent.com"

	tokenFileName = "token.json"
	opReference   = "op://Personal/Google Task oauth/credential"
)

var _ TaskService = (*GoogleTaskService)(nil)

var taskLists = map[Category]*tasks.TaskList{
	CategoryInbox:   {Title: "Inbox"},
	CategoryNext:    {Title: "Next"},
	CategorySomeday: {Title: "Someday"},
}

type GoogleTaskService struct {
	srv *tasks.Service
}

func NewGoogleTaskService() (*GoogleTaskService, error) {
	s := &GoogleTaskService{}

	return s, nil
}

func (g *GoogleTaskService) NewTask() Task {
	return newGoogleTask(&tasks.Task{})
}

func (g *GoogleTaskService) AddTask(task Task) (Task, error) {
	category := task.GetCategory()
	core.GetLogger().Debugf("add task to category: %s", category.Name())
	gtask, err := task.(*googleTask).getGoogleTask()
	if err != nil {
		return nil, fmt.Errorf("failed to get google task: %w", err)
	}

	t, err := g.srv.Tasks.Insert(taskLists[category].Id, gtask).Do()
	if err != nil {
		return nil, fmt.Errorf("failed to add task: %w", err)
	}

	return newGoogleTask(t), nil
}

func (g *GoogleTaskService) ListProjects() ([]string, error) {
	return nil, nil
}

func (g *GoogleTaskService) ListTodoTasks() ([]Task, error) {
	tasks := []Task{}
	for _, taskList := range taskLists {
		gtasks, err := g.srv.Tasks.List(taskList.Id).Do()
		if err != nil {
			return nil, fmt.Errorf("failed to get tasks: %w", err)
		}
		for _, gtask := range gtasks.Items {
			tasks = append(tasks, newGoogleTask(gtask))
		}
	}
	return tasks, nil
}

func (g *GoogleTaskService) ListTags() ([]string, error) {
	return nil, nil
}

func (g *GoogleTaskService) OAuth2(urlHandler func(string) error) error {
	oauthConfig, err := getOauthConfig()
	if err != nil {
		return fmt.Errorf("failed to get oauth config: %w", err)
	}

	tokenFile, err := xdg.StateFile(filepath.Join(core.AppName, tokenFileName))
	if err != nil {
		return fmt.Errorf("failed to get token file path: %w", err)
	}

	if _, err := os.Stat(tokenFile); err != nil {
		server := &http.Server{Addr: ":9873"}
		authChan := make(chan error, 1)
		defer close(authChan)
		defer server.Shutdown(context.Background()) // nolint:errcheck // don't care error

		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		stateToken := fmt.Sprintf("%d", r.Int())
		http.HandleFunc("/taskcommander/oauth", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Connection", "close")

			authCode := r.URL.Query().Get("code")
			state := r.URL.Query().Get("state")
			if state != stateToken {
				http.Error(w, "invalid state", http.StatusBadRequest)
				authChan <- errors.New("invalid state")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			err := fetchOauthToken(authCode, tokenFile, *oauthConfig)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Auth done, you can close this page now")
			authChan <- err
		})

		go func() {
			if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
				log.Fatalf("HTTP server error: %v", err)
			}
		}()

		if err := urlHandler(oauthConfig.AuthCodeURL(stateToken, oauth2.AccessTypeOffline)); err != nil {
			return fmt.Errorf("failed to handle url: %w", err)
		}

		err := <-authChan
		if err != nil {
			return fmt.Errorf("failed to auth: %w", err)
		}
	}

	srv, err := getService(oauthConfig, tokenFile)
	if err != nil {
		return fmt.Errorf("failed to create service: %w", err)
	}
	g.srv = srv

	if err := g.ensureTaskListExist(); err != nil {
		return fmt.Errorf("failed to init task list: %w", err)
	}
	return nil
}

// TODO: find another way to store the secret
func getOauthConfig() (*oauth2.Config, error) {
	b, err := os.ReadFile(core.GetConfig().Gtask.CredentialFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read client secret file: %w", err)
	}

	config, err := google.ConfigFromJSON(b, tasks.TasksScope)
	if err != nil {
		return nil, fmt.Errorf("failed to parse client secret file to config: %w", err)
	}

	//config.ClientID = os.Getenv(client_id_key)
	//if config.ClientID == "" {
	//return nil, fmt.Errorf("client id not found, please set the environment variable %s", client_id_key)
	//}
	config.ClientID = client_id

	secret, err := getSecret()
	if err != nil {
		return nil, fmt.Errorf("failed to get secret: %w", err)
	}
	config.ClientSecret = secret

	return config, nil
}

func fetchOauthToken(authCode, tokenFile string, oauthConfig oauth2.Config) error {
	tok, err := oauthConfig.Exchange(context.TODO(), authCode)
	if err != nil {
		return fmt.Errorf("unable to retrieve token from web: %w", err)
	}

	if err := saveToken(tokenFile, tok); err != nil {
		return fmt.Errorf("failed to save new token: %w", err)
	}
	return nil
}

func (g *GoogleTaskService) getTaskLists() ([]*tasks.TaskList, error) {
	lists, err := g.srv.Tasklists.List().Do()
	if err != nil {
		return nil, err
	}

	core.GetLogger().Debugf("task lists: %v", lists)
	return lists.Items, nil
}

func (g *GoogleTaskService) ensureTaskListExist() error {
	core.GetLogger().Debug("create necessary task lists")
	lists, err := g.getTaskLists()
	if err != nil {
		return err
	}

	for _, taskList := range taskLists {
		list := g.findTaskList(lists, taskList)
		if list == nil {
			core.GetLogger().Debugf("create task list: %s", taskList.Title)
			list, err = g.srv.Tasklists.Insert(taskList).Do()
			if err != nil {
				return fmt.Errorf("failed to create task list: %w", err)
			}
		}
		core.GetLogger().Debugf("tasklist found, id: [%s], title: [%s]", list.Id, list.Title)
		taskList.Id = list.Id
	}
	return nil
}

func (g *GoogleTaskService) findTaskList(
	lists []*tasks.TaskList,
	taskList *tasks.TaskList,
) *tasks.TaskList {
	for _, list := range lists {
		if list.Title == taskList.Title {
			return list
		}
	}
	return nil
}

// TODO: get credentials from config
func getService(config *oauth2.Config, tokenFile string) (*tasks.Service, error) {
	if _, err := os.Stat(tokenFile); err != nil {
		return nil, fmt.Errorf("token file not found: %w", err)
	}

	token, err := tokenFromFile(tokenFile)
	if err != nil {
		return nil, fmt.Errorf("failed to get token from file: %w", err)
	}

	tokenSource := config.TokenSource(context.Background(), token)
	if !token.Valid() {
		newToken, err := tokenSource.Token()
		if err != nil {
			return nil, fmt.Errorf("failed to get new token: %w", err)
		}
		if err := saveToken(tokenFile, newToken); err != nil {
			return nil, fmt.Errorf("failed to save new token: %w", err)
		}
	}

	client := oauth2.NewClient(context.Background(), tokenSource)
	srv, err := tasks.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("failed to create service: %w", err)
	}

	return srv, nil
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func saveToken(path string, token *oauth2.Token) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("failed to write token file: %w", err)
	}
	defer f.Close()
	encoder := json.NewEncoder(f)
	if err := encoder.Encode(token); err != nil {
		return fmt.Errorf("failed to encode token: %w", err)
	}
	return nil
}

func getSecret() (string, error) {
	op, err := exec.LookPath("op")
	if err != nil {
		return "", fmt.Errorf("failed to find 1password cli: %w", err)
	}
	secret, err := exec.Command(op, "read", opReference).Output()
	if err != nil {
		return "", fmt.Errorf("failed to get secret: %w", err)
	}

	return strings.TrimSpace(string(secret)), nil
}
