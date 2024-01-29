package warrior

import "os/exec"

//import (
//"encoding/json"
//"errors"
//"os/exec"
//)

var taskCmd string

func init() {
	var err error
	if taskCmd, err = exec.LookPath("task"); err != nil {
		panic(err)
	}
}

//func AddTask(task Task) error {
//if task.Category == -1 {
//return errors.New("category is not set")
//}

//cmd := append([]string{"add"}, task.Description, "category:"+task.Category.Name())
//if task.Project != "" {
//cmd = append(cmd, "project:"+task.Project)
//}

//if len(task.Tags) > 0 {
//for _, tags := range task.Tags {
//cmd = append(cmd, "+"+tags)
//}
//}

//GetLogger().Debug("execute command: ", taskCmd, cmd)
//if _, err := Taskwarrior(cmd...); err != nil {
//return err
//}

//return nil
//}

//func ListTasksByCategory(cat Category) ([]Task, error) {
//return execCmd([]string{"category:" + cat.Name(), "export"})
//}

//func ListProjects() []string {
//tasks := dump()
//projects := []string{}
//for _, task := range tasks {
//if task.GetProject() != "" {
//projects = append(projects, task.GetProject())
//}
//}
//return projects
//}

// func ListTags() []string {
// tasks := dump()

//tagsMap := make(map[string]bool)
//tags := []string{}
//for _, task := range tasks {
//for _, tag := range task.GetTags() {
//if tagsMap[tag] {
//continue
//}
//tags = append(tags, tag)
//tagsMap[tag] = true
//}
//}
//return tags
//}

//func dump() []Task {
//cmd := []string{"export"}
//tasks, err := execCmd(cmd)
//if err != nil {
//GetLogger().Warn("failed to get projects: ", taskCmd, cmd, err)
//return []Task{}
//}

//return tasks
//}

//func execCmd(cmd []string) ([]Task, error) {
//output, err := Taskwarrior(cmd...)
//if err != nil {
//return nil, err
//}

//var tasks []Task
//if err := json.Unmarshal(output, &tasks); err != nil {
//return nil, err
//}

//return tasks, nil
//}
