package ganboard

import (
	"encoding/json"
)

// CreateUser https://docs.kanboard.org/en/latest/api/user_procedures.html#createuser
func (c *Client) CreateUser(params UserParams) (int, error) {
	query := request{
		Client: c,
		Method: "createUser",
		Params: params,
	}
	response, err := query.decodeInt()
	return response, err
}

// CreateLdapUser https://docs.kanboard.org/en/latest/api/user_procedures.html#CreateLdapUser
func (c *Client) CreateLdapUser(username string) (int, error) {
	query := request{
		Client: c,
		Method: "createLdapUser",
		Params: map[string]string{
			"username": username,
		},
	}
	response, err := query.decodeInt()
	return response, err
}

// GetUser https://docs.kanboard.org/en/latest/api/user_procedures.html#getuser
func (c *Client) GetUser(userID int) (User, error) {
	query := request{
		Client: c,
		Method: "getUser",
		Params: map[string]int{
			"user_id": userID,
		},
	}
	response, err := query.decodeUser()
	return response, err
}

// GetUserByName https://docs.kanboard.org/en/latest/api/user_procedures.html#getuserbyname
func (c *Client) GetUserByName(username string) (User, error) {
	query := request{
		Client: c,
		Method: "getUserByName",
		Params: map[string]string{
			"username": username,
		},
	}
	response, err := query.decodeUser()
	return response, err
}

// GetAllUsers https://docs.kanboard.org/en/latest/api/user_procedures.html#getallusers
func (c *Client) GetAllUsers() ([]User, error) {
	query := request{
		Client: c,
		Method: "getAllUsers",
	}
	response, err := query.decodeUsers()
	return response, err
}

// UpdateUser https://docs.kanboard.org/en/latest/api/user_procedures.html#updateuser
func (c *Client) UpdateUser(params UserParams) (int, error) {
	query := request{
		Client: c,
		Method: "updateUser",
		Params: params,
	}
	response, err := query.decodeInt()
	return response, err
}

// RemoveUser https://docs.kanboard.org/en/latest/api/user_procedures.html#removeuser
func (c *Client) RemoveUser(userID int) (bool, error) {
	query := request{
		Client: c,
		Method: "removeUser",
		Params: map[string]int{
			"user_id": userID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// DisableUser https://docs.kanboard.org/en/latest/api/user_procedures.html#disableuser
func (c *Client) DisableUser(userID int) (bool, error) {
	query := request{
		Client: c,
		Method: "disableUser",
		Params: map[string]int{
			"user_id": userID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// EnableUser https://docs.kanboard.org/en/latest/api/user_procedures.html#enableuser
func (c *Client) EnableUser(userID int) (bool, error) {
	query := request{
		Client: c,
		Method: "enableUser",
		Params: map[string]int{
			"user_id": userID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// IsActiveUser https://docs.kanboard.org/en/latest/api/user_procedures.html#isactiveuser
func (c *Client) IsActiveUser(userID int) (bool, error) {
	query := request{
		Client: c,
		Method: "isActiveUser",
		Params: map[string]int{
			"user_id": userID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
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
	ID                   int    `json:"id"`
	UserName             string `json:"username"`
	Role                 string `json:"role"`
	IsLdapUser           bool   `json:"is_ldap_user"`
	Name                 string `json:"name"`
	Email                string `json:"email,omitempty"`
	GoogleID             string `json:"google_id,omitempty"`
	GithubID             string `json:"github_id,omitempty"`
	GitlabID             string `json:"gitlab_id,omitempty"`
	NotificationsEnabled int    `json:"notifications_enabled,string"`
	Timezone             string `json:"timezone,omitempty"`
	Language             string `json:"language,omitempty"`
	DisableLoginForm     int    `json:"disable_login_form,string"`
	TwoFactorActivated   bool   `json:"twofactor_activated"`
	TwoFactorSecret      bool   `json:"twofactor_secret"`
	Token                string `json:"token"`
	NotificationsFilter  int    `json:"notifications_filter,string"`
	NbFailedLogin        int    `json:"nb_failed_login,string"`
	LockExpirationDate   int    `json:"lock_expiration_date,string"`
	IsActive             bool   `json:"int,string"`
	AvatarPath           string `json:"avatar_path"`
	APIAccesstoken       string `json:"api_access_token"`
}

func (r *request) decodeUser() (User, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return User{}, err
	}

	body := struct {
		JSONRPC string `json:"jsonrpc"`
		ID      int    `json:"id,string"`
		Result  User   `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}

func (r *request) decodeUsers() ([]User, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return nil, err
	}

	body := struct {
		JSONRPC string `json:"jsonrpc"`
		ID      int    `json:"id,string"`
		Result  []User `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}
