package jira

type ChangeLog struct {
	StartAt    int        `json:"startAt,omitempty"`
	MaxResults int        `json:"maxResults,omitempty"`
	Total      int        `json:"total,omitempty"`
	Histories  []*History `json:"histories,omitempty"`
}
