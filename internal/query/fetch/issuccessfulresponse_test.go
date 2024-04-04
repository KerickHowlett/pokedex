package queryfetch

import (
	"net/http"
	"testing"
)

func TestIsSuccessfulResponse(t *testing.T) {
	t.Run("should return true if the response status code is in the 200 range.", func(t *testing.T) {
		if !isSuccessfulResponse(http.StatusOK) {
			t.Error("Expected response to be successful, but got false")
		}
	})

	t.Run("should return false if the response status code is outside the 200 range.", func(t *testing.T) {
		if isSuccessfulResponse(http.StatusInternalServerError) {
			t.Error("Expected response to be unsuccessful, but got true")
		}
	})
}
