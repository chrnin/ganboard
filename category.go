package ganboard

import "encoding/json"

// CreateCategory https://docs.kanboard.org/en/latest/api/category_procedures.html
func (c *Client) CreateCategory(projectID int, name string) (int, error) {
	query := request{
		Client: c,
		Method: "createCategory",
		Params: struct {
			ProjectID int    `json:"project_id"`
			Name      string `json:"name"`
		}{
			ProjectID: projectID,
			Name:      name,
		},
	}
	response, err := query.decodeInt()
	return response, err
}

// GetCategory https://docs.kanboard.org/en/latest/api/category_procedures.html#getcategory
func (c *Client) GetCategory(categoryID int) (Category, error) {
	query := request{
		Client: c,
		Method: "getCategory",
		Params: map[string]int{
			"category_id": categoryID,
		},
	}
	response, err := query.decodeCategory()
	return response, err
}

// GetAllCategories https://docs.kanboard.org/en/latest/api/category_procedures.html#getallcategories
func (c *Client) GetAllCategories(projectID int) ([]Category, error) {
	query := request{
		Client: c,
		Method: "getAllCategory",
		Params: map[string]int{
			"project_id": projectID,
		},
	}
	response, err := query.decodeCategories()
	return response, err
}

// UpdateCategory https://docs.kanboard.org/en/latest/api/category_procedures.html#updatecategory
func (c *Client) UpdateCategory(categoryID int, name string) (bool, error) {
	query := request{
		Client: c,
		Method: "updateCategory",
		Params: struct {
			ID   int    `json:"id,string"`
			Name string `json:"name"`
		}{
			ID:   categoryID,
			Name: name,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// RemoveCategory https://docs.kanboard.org/en/latest/api/category_procedures.html#removecategory
func (c *Client) RemoveCategory(categoryID int) (bool, error) {
	query := request{
		Client: c,
		Method: "removeCategory",
		Params: map[string]int{
			"category_id": categoryID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// Category type
type Category struct {
	ID        int    `json:"id,string"`
	Name      string `json:"name"`
	ProjectID int    `json:"project_id"`
}

func (r *request) decodeCategory() (Category, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return Category{}, err
	}

	body := struct {
		JSONRPC string   `json:"jsonrpc"`
		ID      int      `json:"id"`
		Result  Category `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}

func (r *request) decodeCategories() ([]Category, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return nil, err
	}

	body := struct {
		JSONRPC string     `json:"jsonrpc"`
		ID      int        `json:"id"`
		Result  []Category `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}
