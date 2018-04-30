package ganboard

import "encoding/json"

// CreateCategory https://docs.kanboard.org/en/latest/api/category_procedures.html
func (c *Client) CreateCategory(projectID int, name string) (int, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "createCategory",
		ID:      1,
		Params: struct {
			ProjectID int    `json:"project_id"`
			Name      string `json:"name"`
		}{
			ProjectID: projectID,
			Name:      name,
		},
	}

	rsp, err := c.Request(r)
	body := responseInt{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetCategory https://docs.kanboard.org/en/latest/api/category_procedures.html#getcategory
func (c *Client) GetCategory(categoryID int) (Category, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getCategory",
		ID:      1,
		Params: map[string]int{
			"category_id": categoryID,
		},
	}

	rsp, err := c.Request(r)
	body := responseCategory{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetAllCategories https://docs.kanboard.org/en/latest/api/category_procedures.html#getallcategories
func (c *Client) GetAllCategories(projectID int) ([]Category, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getAllCategory",
		ID:      1,
		Params: map[string]int{
			"project_id": projectID,
		},
	}

	rsp, err := c.Request(r)
	body := responseCategories{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// UpdateCategory https://docs.kanboard.org/en/latest/api/category_procedures.html#updatecategory
func (c *Client) UpdateCategory(categoryID int, name string) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "updateCategory",
		ID:      1,
		Params: struct {
			ID   int    `json:"id,string"`
			Name string `json:"name"`
		}{
			ID:   categoryID,
			Name: name,
		},
	}

	rsp, err := c.Request(r)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// RemoveCategory https://docs.kanboard.org/en/latest/api/category_procedures.html#removecategory
func (c *Client) RemoveCategory(categoryID int) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "removeCategory",
		ID:      1,
		Params: map[string]int{
			"category_id": categoryID,
		},
	}

	rsp, err := c.Request(r)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// Category type
type Category struct {
	ID        int    `json:"id,string"`
	Name      string `json:"name"`
	ProjectID int    `json:"project_id"`
}

type responseCategory struct {
	JSONRPC string   `json:"jsonrpc"`
	ID      int      `json:"id"`
	Result  Category `json:"result"`
}

type responseCategories struct {
	JSONRPC string     `json:"jsonrpc"`
	ID      int        `json:"id"`
	Result  []Category `json:"result"`
}
