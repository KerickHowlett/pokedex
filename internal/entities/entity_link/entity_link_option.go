package entity_link

// EntityLinkOption represents the functional options for an EntityLink.
type EntityLinkOption func(*EntityLink)

// WithName sets the name of the EntityLink.
//
// Parameters:
//   - name: The name of the link.
//
// Returns:
//   - EntityLinkOption: The functional option to set the name of the EntityLink.
//
// Example usage:
//
//	link := entity_link.NewEntityLink(entity_link.WithName("Pikachu"))
func WithName(name string) EntityLinkOption {
	return func(e *EntityLink) {
		e.Name = name
	}
}

// WithURL sets the URL of the EntityLink.
//
// Parameters:
//   - url: The URL of the link.
//
// Returns:
//   - EntityLinkOption: The functional option to set the URL of the EntityLink.
//
// Example usage:
//
//	link := entity_link.NewEntityLink(entity_link.WithURL("https://pokeapi.co/api/v2/pokemon/25/"))
func WithURL(url string) EntityLinkOption {
	return func(e *EntityLink) {
		e.URL = url
	}
}
