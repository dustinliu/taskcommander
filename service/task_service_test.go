package service

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TaskTestSuite struct {
	suite.Suite
}

func TestTaskTestSuite(t *testing.T) {
	suite.Run(t, new(TaskTestSuite))
}

func purgeAll() {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	if err := os.RemoveAll(home_dir + "/.task"); err != nil {
		panic(err)
	}
}

func (suite *TaskTestSuite) SetupTest() {
	purgeAll()
}

func (suite *TaskTestSuite) TestAddTask() {
	task := Task{
		Description: "Test Task",
		Category:    Inbox,
	}

	err := AddTask(&task)
	suite.Nil(err)
}

func (suite *TaskTestSuite) TestListTasks() {
	// Create some test tasks
	tasks := []Task{
		{
			Description: "Test Task 1",
			Category:    Inbox,
			Status:      "pending",
			Tags:        []string{"tag1", "tag2"},
		},
		{
			Description: "Test Task 2",
			Category:    Next,
			Status:      "pending",
		},
		{
			Description: "Test Task 3",
			Category:    Someday,
			Status:      "pending",
		},
	}

	// Add the test tasks to the task list
	for _, task := range tasks {
		err := AddTask(&task)
		suite.Nil(err)
	}

	// Call the ListTasks function
	taskList, err := ListTasksByCategory(Inbox)
	suite.Nil(err)

	// Check that the returned task list matches the expected tasks
	suite.Equal(tasks[0].Description, taskList[0].Description)
	suite.Equal(tasks[0].Category, taskList[0].Category)
	suite.Equal(tasks[0].Status, taskList[0].Status)
	suite.Equal(tasks[0].Tags, taskList[0].Tags)

	// Call the ListTasks function again with a different category
	taskList, err = ListTasksByCategory(Next)
	suite.Nil(err)

	// Check that the returned task list matches the expected tasks
	suite.Equal(tasks[1].Description, taskList[0].Description)
	suite.Equal(tasks[1].Category, taskList[0].Category)
	suite.Equal(tasks[1].Status, taskList[0].Status)

	// Call the ListTasks function again with a different category
	taskList, err = ListTasksByCategory(Someday)
	suite.Nil(err)

	// Check that the returned task list matches the expected tasks
	suite.Equal(tasks[2].Description, taskList[0].Description)
	suite.Equal(tasks[2].Category, taskList[0].Category)
	suite.Equal(tasks[2].Status, taskList[0].Status)
}

func (suite *TaskTestSuite) TestListProjects() {
	// Create some test tasks
	tasks := []Task{
		{
			Description: "Test Task 1",
			Category:    Inbox,
			Status:      "pending",
			Project:     "Test Project 1",
		},
		{
			Description: "Test Task 2",
			Category:    Next,
			Status:      "pending",
		},
		{
			Description: "Test Task 3",
			Category:    Someday,
			Status:      "pending",
			Project:     "Test Project 3",
		},
	}

	// Add the test tasks to the task list
	for _, task := range tasks {
		err := AddTask(&task)
		suite.Nil(err)
	}

	// Call the ListProjects function
	projectList := ListProjects()

	// Check that the returned project list matches the expected projects
	suite.Equal(2, len(projectList))
	suite.Equal(tasks[0].Project, projectList[0])
	suite.Equal(tasks[2].Project, projectList[1])
}

func (suite *TaskTestSuite) TestListTags() {
	// Create some test tasks
	tasks := []Task{
		{
			Description: "Test Task 1",
			Category:    Inbox,
			Status:      "pending",
			Tags:        []string{"tag1", "tag2"},
		},
		{
			Description: "Test Task 2",
			Category:    Next,
			Status:      "pending",
			Tags:        []string{"tag2", "tag3"},
		},
		{
			Description: "Test Task 3",
			Category:    Someday,
			Status:      "pending",
			Tags:        []string{"tag3", "tag4"},
		},
	}

	// Add the test tasks to the task list
	for _, task := range tasks {
		err := AddTask(&task)
		suite.Nil(err)
	}

	// Call the ListTags function
	tagList := ListTags()

	// Check that the returned tag list matches the expected tags
	suite.Equal(4, len(tagList))
	suite.Equal(tasks[0].Tags[0], tagList[0])
	suite.Equal(tasks[0].Tags[1], tagList[1])
	suite.Equal(tasks[1].Tags[1], tagList[2])
	suite.Equal(tasks[2].Tags[1], tagList[3])
}
