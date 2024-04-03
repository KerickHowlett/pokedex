package mockscanner

type MockScanner struct {
	userInput string
	isEnabled bool
}

func (s *MockScanner) IsEnable() {
	s.isEnabled = true
}

func (s *MockScanner) Scan() bool {
	return s.isEnabled
}

func (s *MockScanner) SetIsEnabled(status bool) {
	s.isEnabled = status
}

func (s *MockScanner) SetUserInput(input string) {
	s.userInput = input
}

func (s *MockScanner) Text() string {
	return s.userInput
}

func NewMockScanner() *MockScanner {
	return &MockScanner{
		isEnabled: true,
		userInput: "",
	}
}
