package parser

import "strings"

// ParserUtil is a function type that represents options for configuring a
// Sanitizer.
type ParserUtil func(*ParserUtils)

// ParserUtils is a struct that provides function dependencies for sanitizing
// input strings.
//
// The Fields field is a function that takes a string as input and returns a
// slice of strings. It is used to specify the fields that should be sanitized
// by the ParserUtils.
//
// The ToLower field is a function that takes a string as input and returns the
// lowercase version of it. It is used to convert input strings to lowercase.
//
// The TrimSpace field is a function that takes a string as input and returns the
// trimmed string. It is used to remove leading and trailing white spaces from a
// string.
type ParserUtils struct {
	Fields    func(string) []string
	ToLower   func(string) string
	TrimSpace func(string) string
}

// NewParser creates a new instance of the Sanitizer struct with the provided
// options.
//
// It accepts variadic SanitizerOptions as parameters.
//
// The returned Sanitizer object is used to sanitize input strings.
//
// Parameters:
//   - options: The options to configure the Sanitizer.
//
// Returns:
//   - A new instance of the Sanitizer struct.
//
// Example usage:
//
//	sanitizer := NewParser()
//	sanitizedOutput := sanitizer.ParseInput("  HELLO  ")
//	fmt.Println(sanitizedOutput) // Output: [hello]
//
// Now the sanitizer can be used to sanitize input strings.
func NewParser(options ...ParserUtil) *ParserUtils {
	sanitizer := &ParserUtils{
		Fields:    strings.Fields,
		ToLower:   strings.ToLower,
		TrimSpace: strings.TrimSpace,
	}

	for _, option := range options {
		option(sanitizer)
	}

	return sanitizer
}

// WithFields is a function that returns a SanitizerOptions function which sets
// the Fields field of the Sanitizer struct.
//
// The Fields field is a function that takes a string as input and returns a
// slice of strings. It is used to specify the fields that should be sanitized
// by the Sanitizer.
//
// The returned SanitizerOptions function can be used to set the Fields field of
// a Sanitizer instance.
//
// Parameters:
//   - fields: The function to be called for splitting input strings into fields.
//
// Returns:
//   - An option function that sets the Fields field of the Sanitizer.
//
// Example usage:
//
//	options := WithFields(strings.Fields)
//	sanitizer := NewSanitizer(options)
//
// Now the sanitizer will use the strings.Fields function to split input strings
// into fields.
func WithFields(fields func(string) []string) ParserUtil {
	return func(s *ParserUtils) {
		s.Fields = fields
	}
}

// WithToLower sets the function to convert input strings to lowercase for the
// Sanitizer. It takes a function as a parameter that accepts a string and
// returns the lowercase version of it.
//
// Parameters:
//   - toLower: The function to be called for converting input strings to
//     lowercase.
//
// Returns:
//   - An option function that sets the ToLower field of the Sanitizer.
//
// Example usage:
//
//	options := WithToLower(strings.ToLower)
//	sanitizer := NewSanitizer(options)
//
// Now the sanitizer will convert input strings to lowercase using the
// strings.ToLower function.
func WithToLower(toLower func(string) string) ParserUtil {
	return func(s *ParserUtils) {
		s.ToLower = toLower
	}
}

// WithTrimSpace is a function that returns a SanitizerOptions function, which
// sets the TrimSpace function of a Sanitizer. The TrimSpace function is used to
// remove leading and trailing white spaces from a string. The trimSpace
// parameter is a function that takes a string as input and returns the trimmed
// string.
//
// Parameters:
//   - trimSpace: The function to be called for trimming leading and trailing
//     spaces from a string.
//
// Returns:
//   - An option function that sets the TrimSpace field of the Sanitizer.
//
// Example usage:
//
//	opts := WithTrimSpace(strings.TrimSpace)
//	s := NewSanitizer(opts)
//
// Now s.TrimSpace will be set to strings.TrimSpace
func WithTrimSpace(trimSpace func(string) string) ParserUtil {
	return func(s *ParserUtils) {
		s.TrimSpace = trimSpace
	}
}

// ParseInput sanitizes the input string by trimming leading and trailing
// spaces, converting it to lowercase, and splitting it into individual words.
//
// It returns a slice of strings representing the sanitized words.
//
// Parameters:
//   - input: The input string to be sanitized.
//
// Returns:
//   - A slice of strings representing the sanitized words.
//
// Example usage:
//
//	sanitizer := NewSanitizer()
//	sanitizedOutput := sanitizer.ParseInput("  HELLO  ")
func (s *ParserUtils) ParseInput(input string) (words []string) {
	trimmedInput := s.TrimSpace(input)
	lowercasedInput := s.ToLower(trimmedInput)

	words = s.Fields(lowercasedInput)

	return words
}
