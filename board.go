package ganboard

import (
	"encoding/json"
)

// GetBoard https://docs.kanboard.org/en/latest/api/board_procedures.html#getboard
func (c *Client) GetBoard(projectID int) ([]Board, error) {
	query := request{
		Client: c,
		Method: "getBoard",
		Params: map[string]int{
			"project_id": projectID,
		},
	}
	response, err := query.decodeBoards()
	return response, err
}

// Board type
type Board struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Columns     []Column `json:"columns"`
	Description string   `json:"description"`
	IsActive    int      `json:"is_active,string"`
	NbColumns   int      `json:"nb_columns"`
	NbSwimlanes int      `json:"nb_swimlanes"`
	NbTasks     int      `json:"nb_tasks"`
	Position    int      `json:"position,string"`
	ProjectID   int      `json:"project_id,string"`
	Score       int      `json:"score"`
}

// Column type
type Column struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Position    int    `json:"int,string"`
	ProjectID   int    `json:"project_id,string"`
	TaskLimit   int    `json:"task_limit,string"`
	Description string `json:"description,omitempty"`
	Tasks       []Task `json:"tasks"`
	NbTasks     int    `json:"nb_tasks"`
	Score       int    `json:"score"`
}

func (r *request) decodeBoards() ([]Board, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return nil, err
	}

	body := struct {
		JSONRPC string  `json:"jsonrpc"`
		ID      int     `json:"id"`
		Result  []Board `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}
