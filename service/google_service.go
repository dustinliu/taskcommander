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
	"path/filepath"
	"time"

	"github.com/adrg/xdg"
	"github.com/dustinliu/taskcommander/core"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	tasks "google.golang.org/api/tasks/v1"
)

var _ TaskService = (*GoogleTaskService)(nil)

var taskLists = map[Category]*tasks.TaskList{
	CategoryInbox:   {Title: "Inbox"},
	CategoryNext:    {Title: "Next"},
	CategorySomeday: {Title: "Someday"},
}

type GoogleTaskService struct {
	service *tasks.Service

	tokenFile   string
	oauthConfig *oauth2.Config

	server   *http.Server
	authChan chan error
}

func NewGoogleTaskService() (*GoogleTaskService, error) {
	s := &GoogleTaskService{
		server:   &http.Server{Addr: ":9873"},
		authChan: make(chan error, 1),
	}
	tokenFile, err := xdg.StateFile(filepath.Join(core.AppName, "token.json"))
	if err != nil {
		return nil, fmt.Errorf("failed to get token file path: %w", err)
	}
	s.tokenFile = tokenFile

	b, err := os.ReadFile(core.GetConfig().Gtask.CredentialFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read client secret file: %w", err)
	}
	oauthConfig, err := google.ConfigFromJSON(b, tasks.TasksScope)
	if err != nil {
		return nil, fmt.Errorf("unable to parse client secret file to config: %w", err)
	}
	s.oauthConfig = oauthConfig

	return s, nil
}

func (g *GoogleTaskService) Init() error {
	service, err := getService(g.oauthConfig, g.tokenFile)
	if err != nil {
		return fmt.Errorf("failed to create service: %w", err)
	}
	g.service = service

	if err := g.ensureTaskListExist(); err != nil {
		return fmt.Errorf("failed to init task list: %w", err)
	}
	return nil
}

func (g *GoogleTaskService) InitOauth2Needed() bool {
	if _, err := os.Stat(g.tokenFile); err != nil {
		return true
	}
	return false
}

func (g *GoogleTaskService) GetOauthAuthUrl() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	stateToken := fmt.Sprintf("%d", r.Int())
	server := &http.Server{Addr: ":9873"}
	http.HandleFunc("/taskcommander/oauth", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Connection", "close")

		authCode := r.URL.Query().Get("code")
		state := r.URL.Query().Get("state")
		if state != stateToken {
			http.Error(w, "invalid state", http.StatusBadRequest)
			g.authChan <- errors.New("invalid state")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err := g.fetchOauthToken(authCode)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Auth done, you can close this page now")
		g.authChan <- err
	})

	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	return g.oauthConfig.AuthCodeURL(stateToken, oauth2.AccessTypeOffline)
}

func (g *GoogleTaskService) WaitForAuthDone() error {
	err := <-g.authChan
	close(g.authChan)
	if err := g.server.Shutdown(context.Background()); err != nil {
		return err
	}

	return err
}

func (g *GoogleTaskService) fetchOauthToken(authCode string) error {
	tok, err := g.oauthConfig.Exchange(context.TODO(), authCode)
	if err != nil {
		return fmt.Errorf("unable to retrieve token from web: %w", err)
	}

	if err := saveToken(g.tokenFile, tok); err != nil {
		return fmt.Errorf("failed to save new token: %w", err)
	}
	g.service, err = getService(g.oauthConfig, g.tokenFile)
	if err != nil {
		return fmt.Errorf("failed to create service: %w", err)
	}
	return nil
}

func (g *GoogleTaskService) AddTask(task Task) (Task, error) {
	//t := task.(*GoogleTask)
	//if t, err = g.service.Tasks.Insert(t.Id, t.Task).Do(); err != nil {
	//return fmt.Errorf("failed to add task: %w", err)
	//}

	return nil, nil
}

func (g *GoogleTaskService) ListProjects() ([]string, error) {
	return nil, nil
}

func (g *GoogleTaskService) ListTasksByCategory(cat Category) ([]Task, error) {
	return nil, nil
}

func (g *GoogleTaskService) ListTags() ([]string, error) {
	return nil, nil
}

func (g *GoogleTaskService) getTaskLists() ([]*tasks.TaskList, error) {
	lists, err := g.service.Tasklists.List().Do()
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
			list, err = g.service.Tasklists.Insert(taskList).Do()
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
	service, err := tasks.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("failed to create service: %w", err)
	}

	return service, nil
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
