package service

type (
	Status   int8
	Category int8
)

const (
	StatusTodo Status = iota
	StatusDone

	StatusInvalid Status = -1
)

const (
	CategoryInbox Category = iota
	CategoryNext
	CategorySomeday
	CategoryFocus

	CategoryInvalid Category = -1
)

var categoryNames = []string{
	"Inbox",
	"Next",
	"Someday",
	"Focus",
}

func (c Category) IsValid() bool {
	return c >= 0 && int(c) < len(categoryNames)
}

func (c Category) Name() string {
	if !c.IsValid() {
		return "Invalid"
	}

	return categoryNames[c]
}

type Task interface {
	GetId() string
	GetTitle() string
	SetTitle(string) Task
	GetNote() string
	SetNote(string) Task
	GetFocus() bool
	SetFocus(bool) Task
	GetStatus() Status
	SetStatus(Status) Task
	GetProject() string
	SetProject(string) Task
	GetTags() []string
	SetTag(string) Task
	GetCategory() Category
	SetCategory(Category) Task
	GetDue() string // RFC3339
	SetDue(string) Task
	GetCompleted() string // RFC3339
	SetCompleted(string) Task
	GetUpdated() string // RFC3339
	Error() error
}

type TaskService interface {
	InitOauth2Needed() bool
	GetOauthAuthUrl() string
	WaitForAuthDone() error
	Init() error
	NewTask() Task
	AddTask(Task) (Task, error)
	ListTasks() ([]Task, error)
	ListTags() ([]string, error)
	ListProjects() ([]string, error)
}

func NewService() (TaskService, error) {
	s, err := NewGoogleTaskService()
	if err != nil {
		return nil, err
	}
	return s, nil
}
