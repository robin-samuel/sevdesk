package sevdesk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// do performs an HTTP request against the sevdesk API, returning the raw
// contents of the response's `objects` field. The returned bytes may be a JSON
// object, a JSON array, or empty.
func (c *Client) do(ctx context.Context, method, path string, query url.Values, body any) (json.RawMessage, error) {
	u := c.baseURL + path
	if len(query) > 0 {
		u += "?" + query.Encode()
	}

	var bodyReader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("sevdesk: marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, u, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("sevdesk: build request: %w", err)
	}
	req.Header.Set("Authorization", c.apiKey)
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sevdesk: %s %s: %w", method, path, err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("sevdesk: read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, parseError(method, path, resp.StatusCode, respBytes)
	}
	if len(respBytes) == 0 {
		return nil, nil
	}

	var env struct {
		Objects json.RawMessage `json:"objects"`
	}
	if err := json.Unmarshal(respBytes, &env); err != nil {
		return nil, fmt.Errorf("sevdesk: decode envelope: %w", err)
	}
	return env.Objects, nil
}

// parseError builds an Error from a non-2xx response.
func parseError(method, path string, status int, body []byte) *Error {
	e := &Error{
		StatusCode: status,
		Method:     method,
		Path:       path,
		Body:       body,
	}
	var wrap struct {
		Error struct {
			Message       string `json:"message"`
			ExceptionUUID string `json:"exceptionUUID"`
		} `json:"error"`
	}
	if json.Unmarshal(body, &wrap) == nil && wrap.Error.Message != "" {
		e.Message = wrap.Error.Message
		e.UUID = wrap.Error.ExceptionUUID
	} else {
		e.Message = http.StatusText(status)
	}
	return e
}

// decodeObject parses a single entity from the `objects` payload. sevdesk
// inconsistently returns single resources as either a bare object or an
// array-of-one, so we handle both.
func decodeObject[T any](raw json.RawMessage) (*T, error) {
	if len(raw) == 0 || string(raw) == "null" {
		return nil, ErrNotFound
	}
	if raw[0] == '[' {
		var arr []T
		if err := json.Unmarshal(raw, &arr); err != nil {
			return nil, fmt.Errorf("sevdesk: decode array: %w", err)
		}
		if len(arr) == 0 {
			return nil, ErrNotFound
		}
		return &arr[0], nil
	}
	var single T
	if err := json.Unmarshal(raw, &single); err != nil {
		return nil, fmt.Errorf("sevdesk: decode object: %w", err)
	}
	return &single, nil
}

// decodeList parses a list of entities from the `objects` payload.
func decodeList[T any](raw json.RawMessage) ([]T, error) {
	if len(raw) == 0 || string(raw) == "null" {
		return nil, nil
	}
	var items []T
	if err := json.Unmarshal(raw, &items); err != nil {
		return nil, fmt.Errorf("sevdesk: decode list: %w", err)
	}
	return items, nil
}
