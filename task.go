package ganboard

import (
	"encoding/json"
	"time"
)

// CreateTask https://docs.kanboard.org/en/latest/api/task_procedures.html#createtask
func (c *Client) CreateTask(params TaskParams) (int, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "createTask",
		ID:      1,
		Params:  params,
	}

	rsp, err := c.Request(r)
	body := responseInt{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetTask https://docs.kanboard.org/en/latest/api/task_procedures.html#gettask
func (c *Client) GetTask(taskID int) (Task, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getTask",
		ID:      1,
		Params: map[string]int{
			"task_id": taskID,
		},
	}

	rsp, err := c.Request(r)
	body := responseTask{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetTaskByReference https://docs.kanboard.org/en/latest/api/task_procedures.html#gettaskbyreference
func (c *Client) GetTaskByReference(taskID int, reference string) (Task, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getTaskByReference",
		ID:      1,
		Params: struct {
			TaskID    int    `json:"task_id"`
			Reference string `json:"reference"`
		}{
			TaskID:    taskID,
			Reference: reference,
		},
	}

	rsp, err := c.Request(r)
	body := responseTask{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetAllTasks https://docs.kanboard.org/en/latest/api/task_procedures.html#getalltasks
func (c *Client) GetAllTasks(projectID int, statusID int) ([]Task, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getAllTasks",
		ID:      1,
		Params: map[string]int{
			"project_id": projectID,
			"status_id":  statusID,
		},
	}

	rsp, err := c.Request(r)
	body := responseTasks{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetOverdueTasks https://docs.kanboard.org/en/latest/api/task_procedures.html#getoverduetasks
func (c *Client) GetOverdueTasks() ([]Task, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getOverdueTasks",
		ID:      1,
	}

	rsp, err := c.Request(r)
	body := responseTasks{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetOverdueTasksByProject https://docs.kanboard.org/en/latest/api/task_procedures.html#getoverduetasksbyproject
func (c *Client) GetOverdueTasksByProject(projectID int) ([]Task, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getOverdueTasksByProject",
		ID:      1,
		Params: map[string]int{
			"project_id": projectID,
		},
	}

	rsp, err := c.Request(r)
	body := responseTasks{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// UpdateTask https://docs.kanboard.org/en/latest/api/task_procedures.html#updatetask
func (c *Client) UpdateTask(params TaskParams) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "updateTask",
		ID:      1,
		Params:  params,
	}

	rsp, err := c.Request(r)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// OpenTask https://docs.kanboard.org/en/latest/api/task_procedures.html#opentask
func (c *Client) OpenTask(taskID int) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "openTask",
		ID:      1,
		Params: map[string]int{
			"task_id": taskID,
		},
	}

	rsp, err := c.Request(r)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// CloseTask https://docs.kanboard.org/en/latest/api/task_procedures.html#closetask
func (c *Client) CloseTask(taskID int) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "closeTask",
		ID:      1,
		Params: map[string]int{
			"task_id": taskID,
		},
	}

	rsp, err := c.Request(r)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// RemoveTask https://docs.kanboard.org/en/latest/api/task_procedures.html#removetask
func (c *Client) RemoveTask(taskID int) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "removeTask",
		ID:      1,
		Params: map[string]int{
			"task_id": taskID,
		},
	}

	rsp, err := c.Request(r)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// MoveTaskPosition https://docs.kanboard.org/en/latest/api/task_procedures.html#movetaskposition
func (c *Client) MoveTaskPosition(projectID int, taskID int, columnID int, position int, swimlaneID int) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "moveTaskPosition",
		ID:      1,
		Params: map[string]int{
			"project_id":  projectID,
			"task_id":     taskID,
			"column_id":   columnID,
			"position":    position,
			"swimlane_id": swimlaneID,
		},
	}

	rsp, err := c.Request(r)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// MoveTaskToProject https://docs.kanboard.org/en/latest/api/task_procedures.html#movetasktoproject
func (c *Client) MoveTaskToProject(params MoveTaskParams) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "moveTaskToProject",
		ID:      1,
		Params:  params,
	}

	rsp, err := c.Request(r)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// DuplicateTaskToProject https://docs.kanboard.org/en/latest/api/task_procedures.html#duplicatetasktoproject
func (c *Client) DuplicateTaskToProject(params MoveTaskParams) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "duplicateTaskToProject",
		ID:      1,
		Params:  params,
	}

	rsp, err := c.Request(r)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// SearchTasks https://docs.kanboard.org/en/latest/api/task_procedures.html#searchtasks
func (c *Client) SearchTasks(projectID int, query string) ([]Task, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "searchTasks",
		ID:      1,
		Params: struct {
			ProjectID int    `json:"project_id,string"`
			Query     string `json:"query"`
		}{
			ProjectID: projectID,
			Query:     query,
		},
	}

	rsp, err := c.Request(r)
	body := responseTasks{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// TaskParams input for CreateTask
type TaskParams struct {
	ID                  int       `json:"id,omitempty"`
	Title               string    `json:"title"`
	ProjectID           int       `json:"project_id,string"`
	ColorID             string    `json:"color_id,omitempty"`
	ColumnID            int       `json:"column_id,string,omitempty"`
	OwnerID             int       `json:"owner_id,string,omitempty"`
	CreatorID           int       `json:"creator_id,omitempty"`
	DateDue             time.Time `json:"date_due,omitempty"`
	Description         string    `json:"description,omitempty"`
	CategoryID          int       `json:"category_id,string,omitempty"`
	Score               int       `json:"score,string,omitempty"`
	SwimlaneID          int       `json:"swimlane_id,string,omitempty"`
	Priority            int       `json:"priority,omitempty"`
	RecurrenceStatus    int       `json:"recurrence_status,string,omitempty"`
	RecurrenceTrigger   int       `json:"recurrence_trigger,string,omitempty"`
	RecurrenceFactor    int       `json:"recurrence_factor,string,omitempty"`
	RecurrenceTimeframe int       `json:"recurrence_timeframe,string,omitempty"`
	RecurrenceBaseDate  int       `json:"recurrence_basedate,string,omitempty"`
	Tags                []string  `json:"tags,omitempty"`
	DateStarted         time.Time `json:"date_started,omitempty"`
}

// MoveTaskParams input for MoveTaskToProject
type MoveTaskParams struct {
	TaskID     int `json:"task_id,string"`
	ProjectID  int `json:"project_id,string"`
	SwimlaneID int `json:"swimlane_id,string,omitempty"`
	ColumnID   int `json:"column_id,string,omitempty"`
	CategoryID int `json:"category_id,string,omitempty"`
	OwnerID    int `json:"ownser_id,string,omitempty"`
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

type responseTask struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  Task   `json:"result"`
}

type responseTasks struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  []Task `json:"result"`
}
