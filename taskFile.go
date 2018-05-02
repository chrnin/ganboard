package ganboard

// CreateTaskFile https://docs.kanboard.org/en/latest/api/task_file_procedures.html#createtaskfile
func (c *Client) CreateTaskFile(projectID int, taskID int, filename string, blob string) (int, error) {
	query := request{
		Client: c,
		Method: "createTaskFile",
		Params: struct {
			ProjectID int    `json:"project_id,string"`
			TaskID    int    `json:"task_id,string"`
			Filename  string `json:"filename"`
			Blob      string `json:"blob"`
		}{
			ProjectID: projectID,
			TaskID:    taskID,
			Filename:  filename,
			Blob:      blob,
		},
	}
	response, err := query.decodeInt()
	return response, err
}

// GetAllTaskFiles https://docs.kanboard.org/en/latest/api/task_file_procedures.html#getalltaskfiles
func (c *Client) GetAllTaskFiles(taskID int) ([]File, error) {
	query := request{
		Client: c,
		Method: "getAllTaskFiles",
		Params: map[string]int{
			"task_id": taskID,
		},
	}
	response, err := query.decodeFiles()
	return response, err
}

// GetTaskFile https://docs.kanboard.org/en/latest/api/task_file_procedures.html#gettaskfile
func (c *Client) GetTaskFile(fileID int) (File, error) {
	query := request{
		Client: c,
		Method: "getTaskFile",
		Params: map[string]int{
			"file_id": fileID,
		},
	}
	response, err := query.decodeFile()
	return response, err
}

// DownloadTaskFile https://docs.kanboard.org/en/latest/api/task_file_procedures.html#downloadtaskfile
func (c *Client) DownloadTaskFile(fileID int) (string, error) {
	query := request{
		Client: c,
		Method: "downloadTaskFile",
		Params: map[string]int{
			"file_id": fileID,
		},
	}
	response, err := query.decodeString()
	return response, err
}

// RemoveTaskFile https://docs.kanboard.org/en/latest/api/task_file_procedures.html#removetaskfile
func (c *Client) RemoveTaskFile(fileID int) (bool, error) {
	query := request{
		Client: c,
		Method: "removeTaskFile",
		Params: map[string]int{
			"file_id": fileID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// RemoveAllTaskFiles https://docs.kanboard.org/en/latest/api/task_file_procedures.html#removealltaskfiles
func (c *Client) RemoveAllTaskFiles(taskID int) (bool, error) {
	query := request{
		Client: c,
		Method: "removeAllTaskFiles",
		Params: map[string]int{
			"task_id": taskID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}
