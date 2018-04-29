package ganboard

import (
	"encoding/json"
)

// GetMe https://docs.kanboard.org/en/latest/api/me_procedures.html#getme
func (c *Client) GetMe() (User, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getMe",
		ID:      "1",
	}

	rsp, err := c.Request(request)
	body := responseUser{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetMyDashboard https://docs.kanboard.org/en/latest/api/me_procedures.html#getmydashboard
func (c *Client) GetMyDashboard() (interface{}, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getMyDashboard",
		ID:      "1",
	}

	rsp, err := c.Request(request)
	body := response{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

type responseDashboard struct {
	JSONRPC string    `json:"jsonrpc"`
	ID      int       `json:"id"`
	Result  Dashboard `json:"result"`
}

// Dashboard type
type Dashboard struct {
	Projects []Project   `json:"projects"`
	Tasks    []Task      `json:"tasks"`
	SubTasks interface{} `json:"subtasks"`
}

// Task type
type Task struct {
	ID            int    `json:"id,string"`
	Title         string `json:"title"`
	DateDue       int    `json:"date_due,string"`
	DateCreation  int    `json:"date_creation,string"`
	ProjectID     int    `json:"project_id,string"`
	ColorID       string `json:"color_id"`
	TimeSpent     int    `json:"time_spend,string"`
	TimeEstimated int    `json:"time_estimated,string"`
	ProjectName   string `json:"project_name"`
	URL           string `json:"url"`
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
