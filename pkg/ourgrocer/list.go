package ourgrocer

// ListID is a grocery list identifier
type ListID struct {
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
}

// Item is a grocey list item
type Item struct {
	ID         string `json:"id"`
	Value      string `json:"value"`
	CategoryID string `json:"categoryid,omitempty"`
}
