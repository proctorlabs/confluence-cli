package client

//ConfluenceSpace stores the space information
type ConfluenceSpace struct {
	ID   int64  `json:"id,omitempty"`
	Key  string `json:"key,omitempty"`
	Name string `json:"name,omitempty"`
}

//ConfluencePageBodyStorage holds the storage objects of the body
type ConfluencePageBodyStorage struct {
	Value          string `json:"value,omitempty"`
	Representation string `json:"representation,omitempty"`
}

//ConfluencePageBody holds the page contents itself
type ConfluencePageBody struct {
	Storage *ConfluencePageBodyStorage `json:"storage,omitempty"`
}

//ConfluencePage stores the base page object
type ConfluencePage struct {
	Title string              `json:"title,omitempty"`
	Type  string              `json:"type,omitempty"`
	ID    string              `json:"id,omitempty"`
	Space *ConfluenceSpace    `json:"space,omitempty"`
	Body  *ConfluencePageBody `json:"body,omitempty"`
}

func newPage(title, spaceKey string) *ConfluencePage {
	return &ConfluencePage{
		Title: title,
		Type:  "page",
		Space: &ConfluenceSpace{Key: spaceKey},
		Body: &ConfluencePageBody{
			Storage: &ConfluencePageBodyStorage{
				Value:          "",
				Representation: "storage",
			},
		},
	}
}
