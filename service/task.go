package service

import (
	"encoding/json"
	"os/exec"
)

type Task struct {
	Id          uint   `json:"id"`
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

func AddTask(task *Task) error {
	cmd := append([]string{"add"}, task.Description)
	if task.Category == -1 {
		cmd = append(cmd, "category:"+task.Category.Name())
	}

	if _, err := taskwarrior(cmd...); err != nil {
		return err
	}

	return nil
}

func ListTasks(cat Category) ([]Task, error) {
	cmd := []string{"category:" + cat.Name(), "export"}
	var output []byte
	var err error
	if output, err = taskwarrior(cmd...); err != nil {
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(output, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}
