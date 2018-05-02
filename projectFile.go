package ganboard

import "encoding/json"

// CreateProjectFile https://docs.kanboard.org/en/latest/api/project_file_procedures.html#project-file-api-procedures
func (c *Client) CreateProjectFile(projectID int, filename string, blob string) (int, error) {
	query := request{
		Client: c,
		Method: "CreateProjectFile",
		Params: struct {
			ProjectID int    `json:"project_id,string"`
			Filename  string `json:"filename"`
			Blob      string `json:"blob"`
		}{
			ProjectID: projectID,
			Filename:  filename,
			Blob:      blob,
		},
	}
	response, err := query.decodeInt()
	return response, err
}

// GetAllProjectFiles https://docs.kanboard.org/en/latest/api/project_file_procedures.html#getallprojectfiles
func (c *Client) GetAllProjectFiles(projectID int) ([]File, error) {
	query := request{
		Client: c,
		Method: "getAllProjectFiles",
		Params: map[string]int{
			"project_id": projectID,
		},
	}
	response, err := query.decodeFiles()
	return response, err
}

// GetProjectFile https://docs.kanboard.org/en/latest/api/project_file_procedures.html#getprojectfile
func (c *Client) GetProjectFile(projectID int, fileID int) (File, error) {
	query := request{
		Client: c,
		Method: "getProjectFile",
		Params: map[string]int{
			"project_id": projectID,
			"file_id":    fileID,
		},
	}
	response, err := query.decodeFile()
	return response, err
}

// DownloadProjectFile https://docs.kanboard.org/en/latest/api/project_file_procedures.html#downloadprojectfile
func (c *Client) DownloadProjectFile(projectID int, fileID int) (string, error) {
	query := request{
		Client: c,
		Method: "downloadProjectFile",
		Params: map[string]int{
			"project_id": projectID,
			"file_id":    fileID,
		},
	}
	response, err := query.decodeString()
	return response, err
}

// RemoveProjectFile https://docs.kanboard.org/en/latest/api/project_file_procedures.html#removeprojectfile
func (c *Client) RemoveProjectFile(projectID int, fileID int) (bool, error) {
	query := request{
		Client: c,
		Method: "removeProjectFile",
		Params: map[string]int{
			"project_id": projectID,
			"file_id":    fileID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// RemoveAllProjectFiles https://docs.kanboard.org/en/latest/api/project_file_procedures.html#removeallprojectfiles
func (c *Client) RemoveAllProjectFiles(projectID int) (bool, error) {
	query := request{
		Client: c,
		Method: "removeAllProjectFiles",
		Params: map[string]int{
			"project_id": projectID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// File type
type File struct {
	ID        int     `json:"id,string"`
	Name      string  `json:"name"`
	Path      string  `json:"path"`
	IsImage   int     `json:"is_image,string"`
	ProjectID int     `json:"project_id,string"`
	Date      int     `json:"date"`
	UserID    int     `json:"user_id,string"`
	Size      float64 `json:"size"`
	Username  string  `json:"username"`
	UserName  string  `json:"user_name"`
}

func (r *request) decodeFiles() ([]File, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return nil, err
	}

	body := struct {
		JSONRPC string `json:"jsonrpc"`
		ID      int    `json:"id"`
		Result  []File `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}

func (r *request) decodeFile() (File, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return File{}, err
	}

	body := struct {
		JSONRPC string `json:"jsonrpc"`
		ID      int    `json:"id"`
		Result  File   `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}
