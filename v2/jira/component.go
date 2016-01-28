package jira

type Component struct {
	Self string `json:"self,omitempty"`
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
