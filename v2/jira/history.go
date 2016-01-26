package jira

type History struct {
	Id      string  `json:"id,omitempty"`
	Author  *User   `json:"author,omitempty"`
	Created string  `json:"created,omitempty"`
	Items   []*Item `json:"items,omitempty"`
}

type Item struct {
	Field      string `json:"field,omitempty"`
	FieldType  string `json:"fieldtype,omitempty"`
	From       string `json:"from,omitempty"`
	FromString string `json:"fromString,omitempty"`
	To         string `json:"to,omitempty"`
	ToString   string `json:"toString,omitempty"`
}
