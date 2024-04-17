package entity_link

// EntityLink houses the name and URL of a related entity entry.
//
// Fields:
//   - Name: The name of the link.
//   - URL: The URL of the link.
type EntityLink struct {
	// Name represents the name of the link.
	Name string `json:"name"`
	// URL represents the URL of the link.
	URL string `json:"url"`
}
