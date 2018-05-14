package ganboard

import (
	"fmt"
	"testing"
)

var (
	projectID = 0
	client    = Client{
		Endpoint: "http://localhost/kanboard-1.2.1/jsonrpc.php",
		Username: "jsonrpc",
		Password: "85655db24a48fe6c85340fbedf9948feef798190ddc06ba345e8809d8384",
	}
)

func TestCreateProject(t *testing.T) {
	t.Log("Creating test project, expected ID: 1")

	params := ProjectParams{
		Description: "Test Project for ganboard",
		Identifier:  "TESTPROJECT",
		OwnerID:     1,
		Name:        "Test Project",
	}

	projectID, err := client.CreateProject(params)
	if err != nil {
		t.Errorf("Couldn't create TestProject: %s", err)
	} else {
		if projectID != 1 {
			t.Error("Project ID isn't 1, the testsuit will fail")
		}
	}
}

func TestProjectPermission(t *testing.T) {
	t.Log("Setting project permissions for admin")
	params := ProjectUserParams{
		ProjectID: projectID,
		UserID:    1,
		Role:      "project-manager",
	}
	result, err := client.AddProjectUser(params)

	if err != nil {
		t.Errorf("Couldn't assign role: %s", err)
	} else {
		if result == false {
			t.Errorf("Didn't work either.")
		}
	}
}

func TestCreateTask(t *testing.T) {
	params := TaskParams{
		Title:     "Test task",
		ProjectID: 1,
		DateDue: &time.Date
	}
	idTask, err := client.CreateTask(params)
	if err != nil {
		t.Errorf("Failed to create task: %s", err)
	} else {
		if idTask == 0 {
			t.Error("Failed to create task: reason not provided")
		} else {
			t.Logf("Task Created, id = %d", idTask)
		}
	}
}
func TestRemoveProject(t *testing.T) {
	t.Log("Removing test project")

	result, err := client.RemoveProject(1)

	if err != nil {
		t.Errorf("Couldn't remove TestProject: %s", err)
	} else {
		if result == false {
			t.Errorf("Project not removed. Kanboard didn't say why.")
		}
	}
}
