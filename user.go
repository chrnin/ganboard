package ganboard

import "encoding/json"

// CreateUser https://docs.kanboard.org/en/latest/api/user_procedures.html#createuser
func (c *Client) CreateUser(params UserParams) (int, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "createUser",
		ID:      1,
		Params:  params,
	}

	rsp, err := c.Request(request)
	body := responseInt{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// CreateLdapUser https://docs.kanboard.org/en/latest/api/user_procedures.html#CreateLdapUser
func (c *Client) CreateLdapUser(username string) (int, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "createLdapUser",
		ID:      1,
		Params: map[string]string{
			"username": username,
		},
	}

	rsp, err := c.Request(request)
	body := responseInt{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetUser https://docs.kanboard.org/en/latest/api/user_procedures.html#getuser
func (c *Client) GetUser(userID int) (User, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getUser",
		ID:      1,
		Params: map[string]int{
			"user_id": userID,
		},
	}

	rsp, err := c.Request(request)
	body := responseUser{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetUserByName https://docs.kanboard.org/en/latest/api/user_procedures.html#getuserbyname
func (c *Client) GetUserByName(username string) (User, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getUserByName",
		ID:      1,
		Params: map[string]string{
			"username": username,
		},
	}

	rsp, err := c.Request(request)
	body := responseUser{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetAllUsers https://docs.kanboard.org/en/latest/api/user_procedures.html#getallusers
func (c *Client) GetAllUsers() ([]User, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getAllUsers",
		ID:      1,
	}

	rsp, err := c.Request(request)
	body := responseUsers{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// UpdateUser https://docs.kanboard.org/en/latest/api/user_procedures.html#updateuser
func (c *Client) UpdateUser(params UserParams) (int, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "updateUser",
		ID:      1,
		Params:  params,
	}

	rsp, err := c.Request(request)
	body := responseInt{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// RemoveUser https://docs.kanboard.org/en/latest/api/user_procedures.html#removeuser
func (c *Client) RemoveUser(userID int) (bool, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "removeUser",
		ID:      1,
		Params: map[string]int{
			"user_id": userID,
		},
	}

	rsp, err := c.Request(request)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// DisableUser https://docs.kanboard.org/en/latest/api/user_procedures.html#disableuser
func (c *Client) DisableUser(userID int) (bool, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "disableUser",
		ID:      1,
		Params: map[string]int{
			"user_id": userID,
		},
	}

	rsp, err := c.Request(request)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// EnableUser https://docs.kanboard.org/en/latest/api/user_procedures.html#enableuser
func (c *Client) EnableUser(userID int) (bool, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "enableUser",
		ID:      1,
		Params: map[string]int{
			"user_id": userID,
		},
	}

	rsp, err := c.Request(request)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// IsActiveUser https://docs.kanboard.org/en/latest/api/user_procedures.html#isactiveuser
func (c *Client) IsActiveUser(userID int) (bool, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "isActiveUser",
		ID:      1,
		Params: map[string]int{
			"user_id": userID,
		},
	}

	rsp, err := c.Request(request)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// UserParams input for CreateUser
type UserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"Name,omitempty"`
	Email    string `json:"email,omitempty"`
	Role     string `json:"role,omitempty"`
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

type responseUser struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  User   `json:"result"`
}

type responseUsers struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  []User `json:"result"`
}
