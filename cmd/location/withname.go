package location

func WithName(name string) func(*Location) {
	return func(p *Location) {
		p.Name = name
	}
}
