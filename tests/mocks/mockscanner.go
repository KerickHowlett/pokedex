package mocks

type ScannerMock struct {
	userInput string
	isEnabled bool
}

func (s *ScannerMock) IsEnable() {
	s.isEnabled = true
}

func (s *ScannerMock) Scan() bool {
	return s.isEnabled
}

func (s *ScannerMock) SetIsEnabled(status bool) {
	s.isEnabled = status
}

func (s *ScannerMock) SetUserInput(input string) {
	s.userInput = input
}

func (s *ScannerMock) Text() string {
	return s.userInput
}

func NewScannerMock() *ScannerMock {
	return &ScannerMock{
		isEnabled: true,
		userInput: "",
	}
}
