package handlers

import "fmt"

const (
	contentType     string = "Content-Type"
	applicationJson string = "application/json"
)

// createPayloadError creates the payload error with given message
func createPayloadError(message string) map[string]any {
	return map[string]any{
		"message": message,
	}
}

// ValidContentType validates that the ctSrc content type is the same as ctExpected. Returns true if they are equal, or false with a payload error otherwise
func validateContentType(ctSrc string, ctExpected string) (payload map[string]any, ok bool) {
	if ctExpected != ctSrc {
		return createPayloadError(fmt.Sprintf("%s expected '%s', but it was '%s'", contentType, ctExpected, ctSrc)), false
	}
	return nil, true
}
