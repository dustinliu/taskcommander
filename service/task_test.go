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
	taskList, err := ListTasks(Inbox)
	suite.Nil(err)

	// Check that the returned task list matches the expected tasks
	suite.Equal(tasks[0].Description, taskList[0].Description)
	suite.Equal(tasks[0].Category, taskList[0].Category)
	suite.Equal(tasks[0].Status, taskList[0].Status)

	// Call the ListTasks function again with a different category
	taskList, err = ListTasks(Next)
	suite.Nil(err)

	// Check that the returned task list matches the expected tasks
	suite.Equal(tasks[1].Description, taskList[0].Description)
	suite.Equal(tasks[1].Category, taskList[0].Category)
	suite.Equal(tasks[1].Status, taskList[0].Status)

	// Call the ListTasks function again with a different category
	taskList, err = ListTasks(Someday)
	suite.Nil(err)

	// Check that the returned task list matches the expected tasks
	suite.Equal(tasks[2].Description, taskList[0].Description)
	suite.Equal(tasks[2].Category, taskList[0].Category)
	suite.Equal(tasks[2].Status, taskList[0].Status)
}
