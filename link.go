package ganboard

import "encoding/json"

// GetAllLinks https://docs.kanboard.org/en/latest/api/link_procedures.html#link-api-procedures
func (c *Client) GetAllLinks() ([]Link, error) {
	query := request{
		Client: c,
		Method: "getAllLinks",
	}
	response, err := query.decodeLinks()
	return response, err
}

// GetOppositeLinkID https://docs.kanboard.org/en/latest/api/link_procedures.html#getoppositelinkid
func (c *Client) GetOppositeLinkID(linkID int) (int, error) {
	query := request{
		Client: c,
		Method: "getOppositeLinkId",
	}
	response, err := query.decodeInt()
	return response, err
}

// GetLinkByLabel https://docs.kanboard.org/en/latest/api/link_procedures.html#getlinkbylabel
func (c *Client) GetLinkByLabel(label string) (Link, error) {
	query := request{
		Client: c,
		Method: "getLinkByLabel",
		Params: map[string]string{
			"label": label,
		},
	}
	response, err := query.decodeLink()
	return response, err
}

// GetLinkByID https://docs.kanboard.org/en/latest/api/link_procedures.html#getlinkbyid
func (c *Client) GetLinkByID(linkID int) (Link, error) {
	query := request{
		Client: c,
		Method: "getLinkByID",
		Params: map[string]int{
			"link_id": linkID,
		},
	}
	response, err := query.decodeLink()
	return response, err
}

// CreateLink https://docs.kanboard.org/en/latest/api/link_procedures.html#createlink
// FIXME documentation specifies integers and strings for the same fields
func (c *Client) CreateLink(linkLabel string, oppositeLabel string) (int, error) {
	query := request{
		Client: c,
		Method: "createLink",
		Params: map[string]string{
			"label":          linkLabel,
			"opposite_label": oppositeLabel,
		},
	}
	response, err := query.decodeInt()
	return response, err
}

// UpdateLink https://docs.kanboard.org/en/latest/api/link_procedures.html#updatelink
func (c *Client) UpdateLink(linkID int, oppositeLinkID int, linkLabel string) (bool, error) {
	query := request{
		Client: c,
		Method: "updateLink",
		Params: Link{
			ID:         linkID,
			Label:      linkLabel,
			OppositeID: oppositeLinkID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// RemoveLink https://docs.kanboard.org/en/latest/api/link_procedures.html#updatelink
func (c *Client) RemoveLink(linkID int) (bool, error) {
	query := request{
		Client: c,
		Method: "removeLink",
		Params: Link{
			ID: linkID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// Link type
type Link struct {
	ID         int    `json:"id,string,omitempty"`
	Label      string `json:"label,omitempty"`
	OppositeID int    `json:"opposite_id,omitempty"`
}

func (r *request) decodeLinks() ([]Link, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return nil, err
	}

	body := struct {
		JSONRPC string  `json:"jsonrpc"`
		ID      FlexInt `json:"id"`
		Result  []Link  `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}

func (r *request) decodeLink() (Link, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return Link{}, err
	}

	body := struct {
		JSONRPC string  `json:"jsonrpc"`
		ID      FlexInt `json:"id"`
		Result  Link    `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}
