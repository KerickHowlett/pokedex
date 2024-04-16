package query_fetch

import "net/http"

// isSuccessfulResponse checks if the given HTTP status code indicates a successful response.
// It returns true if the status code is in the range of 200 to 299 (inclusive), and false otherwise.
func isSuccessfulResponse(statusCode int) bool {
	return statusCode >= http.StatusOK &&
		statusCode < http.StatusMultipleChoices
}
