package sevdesk

import (
	"errors"
	"fmt"
	"net/http"
)

// Sentinel errors for the most common sevdesk failure modes. Match with
// [errors.Is]:
//
//	if errors.Is(err, sevdesk.ErrRateLimit) {
//		time.Sleep(backoff)
//	}
var (
	// ErrNotFound matches a 404 response or a single-object lookup that
	// returned zero results.
	ErrNotFound = errors.New("sevdesk: not found")
	// ErrUnauthorized matches a 401 response (bad / missing API key).
	ErrUnauthorized = errors.New("sevdesk: unauthorized")
	// ErrForbidden matches a 403 response (key valid, action not allowed).
	ErrForbidden = errors.New("sevdesk: forbidden")
	// ErrConflict matches a 409 response (e.g. trying to delete a non-draft).
	ErrConflict = errors.New("sevdesk: conflict")
	// ErrValidation matches a 422 response: the request was well-formed but
	// sevdesk rejected its contents. This is the standard answer of
	// sevdesk-Update 2.0 when a booking account, TaxRule and position tax rate
	// don't fit together — [ReceiptGuidanceService] reports the valid
	// combinations. Error.Message carries sevdesk's reason.
	ErrValidation = errors.New("sevdesk: validation failed")
	// ErrRateLimit matches a 429 response.
	ErrRateLimit = errors.New("sevdesk: rate limited")
)

// Error is returned for any non-2xx HTTP response from the sevdesk API.
type Error struct {
	// StatusCode is the HTTP status of the response.
	StatusCode int
	// Method is the HTTP method of the failing request (e.g. "GET").
	Method string
	// Path is the request path (e.g. "/Voucher/42").
	Path string
	// Message is the human-readable error message from sevdesk, or
	// http.StatusText(StatusCode) as a fallback.
	Message string
	// UUID is sevdesk's exceptionUUID, when present — useful for support tickets.
	UUID string
	// Body is the raw response body.
	Body []byte
}

func (e *Error) Error() string {
	prefix := "sevdesk"
	if e.Method != "" && e.Path != "" {
		prefix = fmt.Sprintf("sevdesk %s %s", e.Method, e.Path)
	}
	if e.UUID != "" {
		return fmt.Sprintf("%s: %d %s (%s)", prefix, e.StatusCode, e.Message, e.UUID)
	}
	return fmt.Sprintf("%s: %d %s", prefix, e.StatusCode, e.Message)
}

// Is maps the HTTP status to the package-level sentinel errors so
// [errors.Is] works.
func (e *Error) Is(target error) bool {
	switch target {
	case ErrNotFound:
		return e.StatusCode == http.StatusNotFound
	case ErrUnauthorized:
		return e.StatusCode == http.StatusUnauthorized
	case ErrForbidden:
		return e.StatusCode == http.StatusForbidden
	case ErrConflict:
		return e.StatusCode == http.StatusConflict
	case ErrValidation:
		return e.StatusCode == http.StatusUnprocessableEntity
	case ErrRateLimit:
		return e.StatusCode == http.StatusTooManyRequests
	}
	return false
}
