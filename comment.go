package ganboard

import "encoding/json"

// CreateComment https://docs.kanboard.org/en/latest/api/comment_procedures.html#createcomment
func (c *Client) CreateComment(taskID int, userID int, content string) (int, error) {
	query := request{
		Client: c,
		Method: "createComment",
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
	response, err := query.decodeInt()
	return response, err
}

// GetComment https://docs.kanboard.org/en/latest/api/comment_procedures.html#getcomment
func (c *Client) GetComment(commentID int) (Comment, error) {
	query := request{
		Client: c,
		Method: "getComment",
		Params: map[string]int{
			"comment_id": commentID,
		},
	}
	response, err := query.decodeComment()
	return response, err
}

// GetAllComments https://docs.kanboard.org/en/latest/api/comment_procedures.html#getallcomments
func (c *Client) GetAllComments(taskID int) ([]Comment, error) {
	query := request{
		Client: c,
		Method: "getAllComments",
		Params: map[string]int{
			"task_id": taskID,
		},
	}
	response, err := query.decodeComments()
	return response, err
}

// UpdateComment https://docs.kanboard.org/en/latest/api/comment_procedures.html#updatecomment
func (c *Client) UpdateComment(commentID int, content string) (bool, error) {
	query := request{
		Client: c,
		Method: "updateComment",
		Params: struct {
			CommentID int    `json:"id,string"`
			Content   string `json:"content"`
		}{
			CommentID: commentID,
			Content:   content,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// RemoveComment https://docs.kanboard.org/en/latest/api/comment_procedures.html#removecomment
func (c *Client) RemoveComment(commentID int) (bool, error) {
	query := request{
		Client: c,
		Method: "updateComment",
		Params: map[string]int{
			"comment_id": commentID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
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

func (r *request) decodeComments() ([]Comment, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return nil, err
	}

	body := struct {
		JSONRPC string    `json:"jsonrpc"`
		ID      int       `json:"id"`
		Result  []Comment `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}

func (r *request) decodeComment() (Comment, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return Comment{}, err
	}

	body := struct {
		JSONRPC string  `json:"jsonrpc"`
		ID      int     `json:"id"`
		Result  Comment `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}
