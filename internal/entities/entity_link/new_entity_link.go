package entity_link

// NewEntityLink creates a new EntityLink with the provided options.
//
// Parameters:
//   - opts: The functional options to set the fields of the EntityLink.
//
// Returns:
//   - *EntityLink: The new EntityLink with the provided options.
//
// Example usage:
//
// link := el.NewEntityLink()
func NewEntityLink(options ...EntityLinkOption) *EntityLink {
	link := &EntityLink{}
	for _, option := range options {
		option(link)
	}
	return link
}
