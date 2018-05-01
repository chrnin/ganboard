package ganboard

import "encoding/json"

// GetColumns https://docs.kanboard.org/en/latest/api/column_procedures.html#getcolumns
func (c *Client) GetColumns(projectID int) ([]Column, error) {
	query := request{
		Client: c,
		Method: "getColumns",
		Params: map[string]int{
			"project_id": projectID,
		},
	}
	response, err := query.decodeColumns()
	return response, err
}

// GetColumn https://docs.kanboard.org/en/latest/api/column_procedures.html#getcolumn
func (c *Client) GetColumn(columnID int) (Column, error) {
	query := request{
		Client: c,
		Method: "getColumn",
		Params: map[string]int{
			"column_id": columnID,
		},
	}
	response, err := query.decodeColumn()
	return response, err
}

// ChangeColumnPosition https://docs.kanboard.org/en/latest/api/column_procedures.html#changecolumnposition
func (c *Client) ChangeColumnPosition(projectID int, columnID int, position int) (bool, error) {
	query := request{
		Client: c,
		Method: "changeColumnPosition",
		Params: map[string]int{
			"project_id": projectID,
			"column_id":  columnID,
			"position":   position,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// UpdateColumn https://docs.kanboard.org/en/latest/api/column_procedures.html#updatecolumn
func (c *Client) UpdateColumn(params ColumnParams) (bool, error) {
	query := request{
		Client: c,
		Method: "updateColumn",
		Params: params,
	}
	response, err := query.decodeBoolean()
	return response, err
}

// AddColumn https://docs.kanboard.org/en/latest/api/column_procedures.html#addcolumn
func (c *Client) AddColumn(params ColumnParams) (int, error) {
	query := request{
		Client: c,
		Method: "addColumn",
		Params: params,
	}
	response, err := query.decodeInt()
	return response, err
}

// RemoveColumn https://docs.kanboard.org/en/latest/api/column_procedures.html#removecolumn
func (c *Client) RemoveColumn(columnID int) (bool, error) {
	query := request{
		Client: c,
		Method: "addColumn",
		Params: map[string]int{
			"column_id": columnID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// ColumnParams input for UpdateColumn
type ColumnParams struct {
	ColumnID    int    `json:"column_id"`
	Title       string `json:"title"`
	TaskLimit   int    `json:"task_limit,omitempty"`
	Description string `json:"description,omitempty"`
}

func (r *request) decodeColumn() (Column, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return Column{}, err
	}

	body := struct {
		JSONRPC string `json:"jsonrpc"`
		ID      int    `json:"id"`
		Result  Column `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}

func (r *request) decodeColumns() ([]Column, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return nil, err
	}

	body := struct {
		JSONRPC string   `json:"jsonrpc"`
		ID      int      `json:"id"`
		Result  []Column `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}
