package sevdesk

import (
	"errors"
	"fmt"
	"net/http"
)

// ErrNotFound is returned when the API responds with 404 or when a single-object
// response unexpectedly contains zero results.
var ErrNotFound = errors.New("sevdesk: not found")

// Error is returned for any non-2xx HTTP response from the sevdesk API.
type Error struct {
	StatusCode int
	Message    string
	UUID       string
	Body       []byte
}

func (e *Error) Error() string {
	if e.UUID != "" {
		return fmt.Sprintf("sevdesk: %d %s (%s)", e.StatusCode, e.Message, e.UUID)
	}
	return fmt.Sprintf("sevdesk: %d %s", e.StatusCode, e.Message)
}

// Is lets callers check for ErrNotFound with errors.Is.
func (e *Error) Is(target error) bool {
	return target == ErrNotFound && e.StatusCode == http.StatusNotFound
}
