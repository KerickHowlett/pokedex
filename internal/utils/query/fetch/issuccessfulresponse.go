package queryfetch

import "net/http"

func isSuccessfulResponse(statusCode int) bool {
	return statusCode >= http.StatusOK &&
		statusCode < http.StatusMultipleChoices
}
