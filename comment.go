package ganboard

import "encoding/json"

// CreateComment https://docs.kanboard.org/en/latest/api/comment_procedures.html#createcomment
func (c *Client) CreateComment(taskID int, userID int, content string) (int, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "createComment",
		ID:      1,
		Params: struct {
			TaskID  int    `json:"task_id,string"`
			UserID  int    `json:"user_id,string"`
			Content string `json:"content"`
		}{
			TaskID:  taskID,
			UserID:  userID,
			Content: content,
		},
	}

	rsp, err := c.Request(r)
	body := responseInt{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetComment https://docs.kanboard.org/en/latest/api/comment_procedures.html#getcomment
func (c *Client) GetComment(commentID int) (Comment, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getComment",
		ID:      1,
		Params: map[string]int{
			"comment_id": commentID,
		},
	}

	rsp, err := c.Request(r)
	body := responseComment{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetAllComments https://docs.kanboard.org/en/latest/api/comment_procedures.html#getallcomments
func (c *Client) GetAllComments(taskID int) ([]Comment, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getAllComments",
		ID:      1,
		Params: map[string]int{
			"task_id": taskID,
		},
	}

	rsp, err := c.Request(r)
	body := responseComments{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// UpdateComment https://docs.kanboard.org/en/latest/api/comment_procedures.html#updatecomment
func (c *Client) UpdateComment(commentID int, content string) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "updateComment",
		ID:      1,
		Params: struct {
			CommentID int    `json:"id,string"`
			Content   string `json:"content"`
		}{
			CommentID: commentID,
			Content:   content,
		},
	}

	rsp, err := c.Request(r)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// RemoveComment https://docs.kanboard.org/en/latest/api/comment_procedures.html#removecomment
func (c *Client) RemoveComment(commentID int) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "updateComment",
		ID:      1,
		Params: map[string]int{
			"comment_id": commentID,
		},
	}

	rsp, err := c.Request(r)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// Comment type
type Comment struct {
	ID           int    `json:"id,string"`
	TaskID       int    `json:"task_id,string"`
	UserID       int    `json:"user_id,string"`
	DateCreation int    `json:"date_creation,string"`
	Comment      string `json:"comment"`
	Username     string `json:"username"`
	Name         string `json:"name"`
}

type responseComment struct {
	JSONRPC string  `json:"jsonrpc"`
	ID      int     `json:"id"`
	Result  Comment `json:"result"`
}

type responseComments struct {
	JSONRPC string    `json:"jsonrpc"`
	ID      int       `json:"id"`
	Result  []Comment `json:"result"`
}
