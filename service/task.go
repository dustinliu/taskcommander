package service

import (
	"encoding/json"
	"errors"
	"os/exec"
	"time"
)

type Task struct {
	Id          int    `json:"id"`
	UUID        string `json:"uuid"`
	Description string `json:"description"`
	Note        string
	Category    Category `json:"category"`
	Status      string   `json:"status"`
	Due         TaskTime
	Project     string
	Urgency     int      `json:"urgency"`
	Priority    string   `json:"priority"`
	CreatedAt   TaskTime `json:"entry"`
	UpdatedAt   TaskTime `json:"modified"`
}

var taskCmd string

func init() {
	var err error
	if taskCmd, err = exec.LookPath("task"); err != nil {
		panic(err)
	}
}

func NewTask() *Task {
	return &Task{
		Category:  -1,
		CreatedAt: TaskTime(time.Now()),
		UpdatedAt: TaskTime(time.Now()),
	}
}

func AddTask(task *Task) error {
	if task.Category == -1 {
		return errors.New("category is not set")
	}

	cmd := append([]string{"add"}, task.Description, "category:"+task.Category.Name())
	if task.Project != "" {
		cmd = append(cmd, "project:"+task.Project)
	}

	GetLogger().Debug("execute command: ", taskCmd, cmd)
	if _, err := Taskwarrior(cmd...); err != nil {
		return err
	}

	return nil
}

func ListTasks(cat Category) ([]Task, error) {
	cmd := []string{"category:" + cat.Name(), "export"}
	var output []byte
	var err error
	if output, err = Taskwarrior(cmd...); err != nil {
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(output, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}
