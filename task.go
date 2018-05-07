package ganboard

import (
	"encoding/json"
	"time"
)

// CreateTask https://docs.kanboard.org/en/latest/api/task_procedures.html#createtask
func (c *Client) CreateTask(params TaskParams) (int, error) {
	query := request{
		Client: c,
		Method: "createTask",
		Params: params,
	}
	response, err := query.decodeInt()
	return response, err
}

// GetTask https://docs.kanboard.org/en/latest/api/task_procedures.html#gettask
func (c *Client) GetTask(taskID int) (Task, error) {
	query := request{
		Client: c,
		Method: "getTask",
		Params: map[string]int{
			"task_id": taskID,
		},
	}
	response, err := query.decodeTask()
	return response, err
}

// GetTaskByReference https://docs.kanboard.org/en/latest/api/task_procedures.html#gettaskbyreference
func (c *Client) GetTaskByReference(taskID int, reference string) (Task, error) {
	query := request{
		Client: c,
		Method: "getTaskByReference",
		Params: struct {
			TaskID    int    `json:"task_id"`
			Reference string `json:"reference"`
		}{
			TaskID:    taskID,
			Reference: reference,
		},
	}
	response, err := query.decodeTask()
	return response, err
}

// GetAllTasks https://docs.kanboard.org/en/latest/api/task_procedures.html#getalltasks
func (c *Client) GetAllTasks(projectID int, statusID int) ([]Task, error) {
	query := request{
		Client: c,
		Method: "getAllTasks",
		Params: map[string]int{
			"project_id": projectID,
			"status_id":  statusID,
		},
	}
	response, err := query.decodeTasks()
	return response, err
}

// GetOverdueTasks https://docs.kanboard.org/en/latest/api/task_procedures.html#getoverduetasks
func (c *Client) GetOverdueTasks() ([]Task, error) {
	query := request{
		Client: c,
		Method: "getOverdueTasks",
	}
	response, err := query.decodeTasks()
	return response, err
}

// GetOverdueTasksByProject https://docs.kanboard.org/en/latest/api/task_procedures.html#getoverduetasksbyproject
func (c *Client) GetOverdueTasksByProject(projectID int) ([]Task, error) {
	query := request{
		Client: c,
		Method: "getOverdueTasksByProject",
		Params: map[string]int{
			"project_id": projectID,
		},
	}
	response, err := query.decodeTasks()
	return response, err
}

// UpdateTask https://docs.kanboard.org/en/latest/api/task_procedures.html#updatetask
func (c *Client) UpdateTask(params TaskParams) (bool, error) {
	query := request{
		Client: c,
		Method: "updateTask",
		Params: params,
	}
	response, err := query.decodeBoolean()
	return response, err
}

// OpenTask https://docs.kanboard.org/en/latest/api/task_procedures.html#opentask
func (c *Client) OpenTask(taskID int) (bool, error) {
	query := request{
		Client: c,
		Method: "openTask",
		Params: map[string]int{
			"task_id": taskID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// CloseTask https://docs.kanboard.org/en/latest/api/task_procedures.html#closetask
func (c *Client) CloseTask(taskID int) (bool, error) {
	query := request{
		Client: c,
		Method: "closeTask",
		Params: map[string]int{
			"task_id": taskID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// RemoveTask https://docs.kanboard.org/en/latest/api/task_procedures.html#removetask
func (c *Client) RemoveTask(taskID int) (bool, error) {
	query := request{
		Client: c,
		Method: "removeTask",
		Params: map[string]int{
			"task_id": taskID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// MoveTaskPosition https://docs.kanboard.org/en/latest/api/task_procedures.html#movetaskposition
func (c *Client) MoveTaskPosition(projectID int, taskID int, columnID int, position int, swimlaneID int) (bool, error) {
	query := request{
		Client: c,
		Method: "moveTaskPosition",
		Params: map[string]int{
			"project_id":  projectID,
			"task_id":     taskID,
			"column_id":   columnID,
			"position":    position,
			"swimlane_id": swimlaneID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// MoveTaskToProject https://docs.kanboard.org/en/latest/api/task_procedures.html#movetasktoproject
func (c *Client) MoveTaskToProject(params MoveTaskParams) (bool, error) {
	query := request{
		Client: c,
		Method: "moveTaskToProject",
		Params: params,
	}
	response, err := query.decodeBoolean()
	return response, err
}

// DuplicateTaskToProject https://docs.kanboard.org/en/latest/api/task_procedures.html#duplicatetasktoproject
func (c *Client) DuplicateTaskToProject(params MoveTaskParams) (bool, error) {
	query := request{
		Client: c,
		Method: "duplicateTaskToProject",
		Params: params,
	}
	response, err := query.decodeBoolean()
	return response, err
}

// SearchTasks https://docs.kanboard.org/en/latest/api/task_procedures.html#searchtasks
func (c *Client) SearchTasks(projectID int, queryString string) ([]Task, error) {
	query := request{
		Client: c,
		Method: "searchTasks",
		Params: struct {
			ProjectID int    `json:"project_id,string"`
			Query     string `json:"query"`
		}{
			ProjectID: projectID,
			Query:     queryString,
		},
	}
	response, err := query.decodeTasks()
	return response, err
}

// TaskParams input for CreateTask
type TaskParams struct {
	ID                  int        `json:"id,omitempty"`
	Title               string     `json:"title"`
	ProjectID           int        `json:"project_id"`
	ColorID             string     `json:"color_id,omitempty"`
	ColumnID            int        `json:"column_id,string,omitempty"`
	OwnerID             int        `json:"owner_id,string,omitempty"`
	CreatorID           int        `json:"creator_id,omitempty"`
	DateDue             *time.Time `json:"date_due,omitempty"`
	Description         string     `json:"description,omitempty"`
	CategoryID          int        `json:"category_id,string,omitempty"`
	Score               int        `json:"score,string,omitempty"`
	SwimlaneID          int        `json:"swimlane_id,string,omitempty"`
	Priority            int        `json:"priority,omitempty"`
	RecurrenceStatus    int        `json:"recurrence_status,string,omitempty"`
	RecurrenceTrigger   int        `json:"recurrence_trigger,string,omitempty"`
	RecurrenceFactor    int        `json:"recurrence_factor,string,omitempty"`
	RecurrenceTimeframe int        `json:"recurrence_timeframe,string,omitempty"`
	RecurrenceBaseDate  int        `json:"recurrence_basedate,string,omitempty"`
	Tags                []string   `json:"tags,omitempty"`
	DateStarted         *time.Time `json:"date_started,omitempty"`
}

// MoveTaskParams input for MoveTaskToProject
type MoveTaskParams struct {
	TaskID     int `json:"task_id,string"`
	ProjectID  int `json:"project_id,string"`
	SwimlaneID int `json:"swimlane_id,string,omitempty"`
	ColumnID   int `json:"column_id,string,omitempty"`
	CategoryID int `json:"category_id,string,omitempty"`
	OwnerID    int `json:"owner_id,string,omitempty"`
}

// Task type
type Task struct {
	ID                  int    `json:"id,string,omitempty"`
	ProjectID           int    `json:"project_id,string"`
	ColorID             string `json:"color_id,omitempty"`
	ColumnID            int    `json:"column_id,string,omitempty"`
	OwnerID             int    `json:"owner_id,string,omitempty"`
	CreatorID           int    `json:"creator_id,omitempty"`
	DateDue             int    `json:"date_due,string,omitempty"`
	Reference           string `json:"reference"`
	Title               string `json:"title"`
	Description         string `json:"description"`
	DateCreation        int    `json:"date_creation,string"`
	DateCompleted       int    `json:"date_completed,string"`
	DateModification    int    `json:"date_modification,string"`
	DateStarted         int    `json:"date_started,string"`
	TimeEstimated       int    `json:"time_estimated,string"`
	TimeSpent           int    `json:"time_spend,string"`
	Position            int    `json:"position,string"`
	IsActive            int    `json:"is_active,string"`
	Score               int    `json:"score,string"`
	CategoryID          int    `json:"category_id,string"`
	SwimlaneID          int    `json:"swimlane_id,string,omitempty"`
	DateMoved           int    `json:"date_moved"`
	RecurrenceStatus    int    `json:"recurrence_status,string"`
	RecurrenceTrigger   int    `json:"recurrence_trigger,string"`
	RecurrenceFactor    int    `json:"recurrence_factor,string"`
	RecurrenceTimeframe int    `json:"recurrence_timeframe,string"`
	RecurrenceBaseDate  int    `json:"recurrence_basedate,string"`
	RecurrenceParent    int    `json:"recurrence_parent,string"`
	RecurrenceChild     int    `json:"recurrence_child,string"`
	CategoryName        string `json:"category_name"`
	ProjectName         string `json:"project_name"`
	DefaultSwimlane     string `json:"default_swimlane"`
	ColumnTitle         string `json:"column_title"`
	AssigneeUsername    string `json:"assignee_username"`
	AssigneeName        string `json:"assignee_name"`
	CreatorUsername     string `json:"creator_username"`
	CreatorName         string `json:"creator_name"`
	NbComments          int    `json:"nb_comments,string"`
	NbFiles             int    `json:"nb_files,string"`
	NbSubtasks          int    `json:"nb_subtasks,string"`
	NbCompletedSubtasks int    `json:"nb_completed_subtasks,string"`
	NbLinks             int    `json:"nb_links,string"`
	Color               Color  `json:"color"`
}

func (r *request) decodeTasks() ([]Task, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return nil, err
	}

	body := struct {
		JSONRPC string  `json:"jsonrpc"`
		ID      FlexInt `json:"id"`
		Result  []Task  `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}

func (r *request) decodeTask() (Task, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return Task{}, err
	}

	body := struct {
		JSONRPC string  `json:"jsonrpc"`
		ID      FlexInt `json:"id"`
		Result  Task    `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}
