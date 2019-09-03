// Structs and functions for the internal OurGroceries API i.e. ourgroceries.com/your-lists

package ourgrocer

type command struct {
	Command string `json:"command"`
	TeamID  string `json:"teamId"`
	ListID  string `json:"listId,omitempty"`
	Value   string `json:"value,omitempty"`
}

type getListResponse struct {
	List list `json:"list"`
}

type getListsResponse struct {
	ShoppingLists []ListID `json:"shoppingLists"`
}

type list struct {
	NotesHTML string `json:"notesHtml"`
	VersionID string `json:"versionId"`
	Notes     string `json:"notes"`
	Name      string `json:"name"`
	ID        string `json:"id"`
	ListType  string `json:"listType"`
	Items     []Item `json:"items"`
}
