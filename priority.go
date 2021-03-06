package testrail

// Priority represents a Priority
type Priority struct {
	ID        int    `json:"id"`
	IsDefault bool   `json:"is_default"`
	Name      string `json:"name"`
	Priority  int    `json:"priority"`
	ShortName string `json:"short_name"`
}

// GetPriorities returns a list of available priorities
func (c *Client) GetPriorities() (priorities []Priority, err error) {
	err = c.sendRequest("GET", "get_priorities/", nil, &priorities)
	return
}
