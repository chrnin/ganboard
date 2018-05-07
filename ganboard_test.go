package ganboard

import (
	"testing"
)

const (
	AppToken = "85655db24a48fe6c85340fbedf9948feef798190ddc06ba345e8809d8384"
	EndPoint = "http://localhost/kanboard-1.2.1/jsonrpc.php"
)

func TestCreateProject(t *testing.T) {
	t.Log("Creating test project, expected ID: 1")
	client := Client{
		Endpoint: EndPoint,
		Username: "jsonrpc",
		Password: AppToken,
	}
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
	client := Client{
		Endpoint: EndPoint,
		Username: "jsonrpc",
		Password: AppToken,
	}

	params := ProjectUserParams{
		ProjectID: 1,
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
func TestRemoveProject(t *testing.T) {
	t.Log("Removing test project")
	client := Client{
		Endpoint: EndPoint,
		Username: "jsonrpc",
		Password: AppToken,
	}

	result, err := client.RemoveProject(1)

	if err != nil {
		t.Errorf("Couldn't remove TestProject: %s", err)
	} else {
		if result == false {
			t.Errorf("Project not removed. Kanboard didn't say why.")
		}
	}
}
