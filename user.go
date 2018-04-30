package ganboard

import (
	"encoding/json"
)

// GetMe https://docs.kanboard.org/en/latest/api/me_procedures.html#getme
func (c *Client) GetMe() (User, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getMe",
		ID:      1,
	}

	rsp, err := c.Request(request)
	body := responseUser{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetMyDashboard https://docs.kanboard.org/en/latest/api/me_procedures.html#getmydashboard
// FIXME documentation doesn't fit result.
func (c *Client) GetMyDashboard() (interface{}, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getMyDashboard",
		ID:      1,
	}

	rsp, err := c.Request(request)
	body := responseInt{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body, err
}

// GetMyActivityStream https://docs.kanboard.org/en/latest/api/me_procedures.html#getmyactivitystream
func (c *Client) GetMyActivityStream() ([]Activity, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getMyActivityStream",
		ID:      1,
	}

	rsp, err := c.Request(request)
	body := responseActivityStream{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// CreateMyPrivateProject https://docs.kanboard.org/en/latest/api/me_procedures.html#createmyprivateproject
func (c *Client) CreateMyPrivateProject(params PrivateProjectParams) (int, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "createMyPrivateProject",
		ID:      1,
		Params:  params,
	}

	rsp, err := c.Request(request)
	body := response{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result.(int), err
}

// GetMyProjectList https://docs.kanboard.org/en/latest/api/me_procedures.html#getmyprojectslist
func (c *Client) GetMyProjectList() (map[int]string, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getMyProjectsList",
		ID:      1,
	}

	rsp, err := c.Request(request)
	body := responseProjectList{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetMyOverDueTasks https://docs.kanboard.org/en/latest/api/me_procedures.html#getmyoverduetasks
func (c *Client) GetMyOverDueTasks() ([]Task, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getMyOverDueTasks",
		ID:      1,
	}

	rsp, err := c.Request(request)
	body := responseOverdueTasks{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetMyProjects https://docs.kanboard.org/en/latest/api/me_procedures.html#getmyprojects
func (c *Client) GetMyProjects() ([]Project, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getMyProjects",
		ID:      1,
	}

	rsp, err := c.Request(request)
	body := responseProjects{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// PrivateProjectParams parameters for CreateMyPrivateProject
type PrivateProjectParams struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type responseOverdueTasks struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  []Task `json:"result"`
}

type responseProjectList struct {
	JSONRPC string         `json:"jsonrpc"`
	ID      int            `json:"id"`
	Result  map[int]string `json:"result"`
}

type responseDashboard struct {
	JSONRPC string    `json:"jsonrpc"`
	ID      int       `json:"id"`
	Result  Dashboard `json:"result"`
}

type responseActivityStream struct {
	JSONRPC string     `json:"jsonrpc"`
	ID      int        `json:"id"`
	Result  []Activity `json:"result"`
}

// Dashboard type
type Dashboard struct {
	Projects []Project   `json:"projects"`
	Tasks    []Task      `json:"tasks"`
	SubTasks interface{} `json:"subtasks"`
}

// Task type
type Task struct {
	ID                  int    `json:"id,string"`
	Reference           string `json:"reference"`
	Title               string `json:"title"`
	Description         string `json:"description"`
	DateCreation        int    `json:"date_creation,string"`
	DateCompleted       int    `json:"date_completed,string"`
	DateModification    int    `json:"date_modification,string"`
	DateDue             int    `json:"date_due,string"`
	DateStarted         int    `json:"date_started,string"`
	TimeEstimated       int    `json:"time_estimated,string"`
	TimeSpent           int    `json:"time_spend,string"`
	ColorID             string `json:"color_id"`
	ProjectID           int    `json:"project_id,string"`
	ColumnID            int    `json:"column_id,string"`
	OwnerID             int    `json:"owner_id,string"`
	CreatorID           int    `json:"creator_id,string"`
	Position            int    `json:"position,string"`
	IsActive            int    `json:"is_active,string"`
	Score               int    `json:"score,string"`
	CategoryID          int    `json:"category_id,string"`
	SwimlaneID          int    `json:"swimlane_id,string"`
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
}

type responseUser struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  User   `json:"result"`
}

// User type
type User struct {
	ID                   int    `json:"id,string"`
	UserName             string `json:"username"`
	Role                 string `json:"role"`
	IDLdapUser           bool   `json:"is_ldap_user"`
	Name                 string `json:"name"`
	Email                string `json:"email,omitempty"`
	GoogleID             string `json:"google_id,omitempty"`
	GithubID             string `json:"github_id,omitempty"`
	NotificationsEnabled int    `json:"notifications_enabled,string"`
	Timezone             string `json:"timezone,omitempty"`
	Language             string `json:"language,omitempty"`
	DisableLoginForm     int    `json:"int,string"`
	TwoFactorActivated   bool   `json:"twofactor_activated"`
	TwoFactorSecret      bool   `json:"twofactor_secret"`
	Token                string `json:"token"`
	NotificationsFilter  int    `json:"notifications_filter,string"`
}

// Activity type
// FIXME no information on Changes structure
type Activity struct {
	ID             int           `json:"id,string"`
	DateCreation   int           `json:"date_creation,string"`
	EventName      string        `json:"event_name"`
	CreatorID      int           `json:"creator_id,string"`
	ProjectID      int           `json:"project_id,string"`
	TaskID         int           `json:"task_id,string"`
	AuthorUsername string        `json:"author_username"`
	AuthorName     string        `json:"author_name"`
	Email          string        `json:"email"`
	Task           Task          `json:"task"`
	Changes        []interface{} `json:"changes"`
	Author         string        `json:"author"`
	EventTitle     string        `json:"event_title"`
	EventContent   string        `json:"event_content"`
}
