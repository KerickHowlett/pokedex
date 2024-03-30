package mocklocation

type MockLocation struct {
	name string
}

func (l *MockLocation) GetName() string {
	return l.name
}

func (l *MockLocation) SetName(name string) {
	l.name = name
}

type MockLocationOption func(*MockLocation)

func WithName(name string) MockLocationOption {
	return func(l *MockLocation) {
		l.name = name
	}
}

func NewMockLocation(options ...MockLocationOption) *MockLocation {
	location := &MockLocation{
		name: "Pallet Town",
	}

	for _, option := range options {
		option(location)
	}

	return location
}
