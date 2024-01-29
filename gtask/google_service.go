package gtask

import (
	"context"

	"github.com/dustinliu/taskcommander/service"
	"google.golang.org/api/option"
	tasks "google.golang.org/api/tasks/v1"
)

var taskLists = map[service.Category]*tasks.TaskList{
	service.CategoryInbox:   {Title: "Inbox"},
	service.CategoryNext:    {Title: "Next"},
	service.CategorySomeday: {Title: "Someday"},
}

type GoogleTaskService struct {
	service *tasks.Service
}

func NewGoogleTaskService() (*GoogleTaskService, error) {
	s := &GoogleTaskService{}
	opt := option.WithCredentialsFile("/Users/dustinl/Downloads/google_taskcommand_service_account.json")

	service, err := tasks.NewService(context.Background(), opt)
	if err != nil {
		return nil, err
	}

	s.service = service
	return s, nil
}

func (g *GoogleTaskService) AddTask(task service.Task) error {
	t := task.(*GoogleTask)
	g.service.Tasks.Insert(t.Id, t.Task).Do()

	return nil
}

func (g *GoogleTaskService) ListProjects() ([]string, error) {
	return nil, nil
}

func (g *GoogleTaskService) ListTasksByCategory(cat service.Category) ([]service.Task, error) {
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

	return lists.Items, nil
}

func (g *GoogleTaskService) ensureTaskListExist() error {
	lists, err := g.getTaskLists()
	if err != nil {
		return err
	}

	for _, taskList := range taskLists {
		list := g.findTaskList(lists, taskList)
		if list == nil {
			list, err = g.service.Tasklists.Insert(taskList).Do()
			if err != nil {
				return err
			}
		}
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

//func getClient() (*http.Client, error) {
//b, err := os.ReadFile("/Users/dustinl/Downloads/google_task_oauth.json")
//if err != nil {
//return nil, err
//}

//config, err := google.ConfigFromJSON(b, tasks.TasksScope)
//if err != nil {
//return nil, err
//}

//authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
//}
