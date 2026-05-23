package sevdesk

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// CommunicationWaysService handles communication with /CommunicationWay endpoints.
type CommunicationWaysService struct {
	c *Client
}

// CommunicationWay is one way to reach a contact (email, phone, web, mobile).
type CommunicationWay struct {
	ID         ID         `json:"id,omitempty"`
	ObjectName string     `json:"objectName,omitempty"`
	Create     *time.Time `json:"create,omitempty"`
	Update     *time.Time `json:"update,omitempty"`
	SevClient  *Ref       `json:"sevClient,omitempty"`

	Contact *Ref                 `json:"contact,omitempty"`
	Type    CommunicationWayType `json:"type,omitempty"`
	Value   string               `json:"value,omitempty"`
	Key     *Ref                 `json:"key,omitempty"`
	Main    Bool                 `json:"main,omitempty"`
}

// CommunicationWayType is the medium of a communication way.
type CommunicationWayType string

// CommunicationWayType values.
const (
	// CommunicationWayTypeEmail is an email address.
	CommunicationWayTypeEmail CommunicationWayType = "EMAIL"
	// CommunicationWayTypePhone is a landline phone number.
	CommunicationWayTypePhone CommunicationWayType = "PHONE"
	// CommunicationWayTypeWeb is a website URL.
	CommunicationWayTypeWeb CommunicationWayType = "WEB"
	// CommunicationWayTypeMobile is a mobile phone number.
	CommunicationWayTypeMobile CommunicationWayType = "MOBILE"
)

// CommunicationWayKey is the purpose/label slot for a communication way
// (e.g. "Work", "Private", "Invoice").
type CommunicationWayKey struct {
	ID              ID         `json:"id,omitempty"`
	ObjectName      string     `json:"objectName,omitempty"`
	Create          *time.Time `json:"create,omitempty"`
	Update          *time.Time `json:"upadate,omitempty"` // sevdesk spec typo: "upadate"
	Name            string     `json:"name,omitempty"`
	TranslationCode string     `json:"translationCode,omitempty"`
}

// CreateCommunicationWayParams is the body for [CommunicationWaysService.Create].
type CreateCommunicationWayParams struct {
	// Contact this communication way belongs to. Required.
	Contact *Ref `json:"contact"`
	// Type categorizes the medium. Required. See [CommunicationWayType].
	Type CommunicationWayType `json:"type"`
	// Value is the actual address (email address, phone number, URL, etc.). Required.
	Value string `json:"value"`
	// Key labels the purpose (e.g. "Work", "Private"). Get available keys with
	// [CommunicationWaysService.Keys].
	Key *Ref `json:"key,omitempty"`
	// Main marks this as the contact's primary communication way of its Type.
	// Use [True] or [False].
	Main *Bool `json:"main,omitempty"`
}

// UpdateCommunicationWayParams is the body for [CommunicationWaysService.Update].
// See [CreateCommunicationWayParams] for field semantics.
type UpdateCommunicationWayParams struct {
	Contact *Ref                  `json:"contact,omitempty"`
	Type    *CommunicationWayType `json:"type,omitempty"`
	Value   *string               `json:"value,omitempty"`
	Key     *Ref                  `json:"key,omitempty"`
	Main    *Bool                 `json:"main,omitempty"`
}

// ListCommunicationWaysParams filters [CommunicationWaysService.List].
type ListCommunicationWaysParams struct {
	// Contact narrows the results to one contact.
	Contact *Ref
	// Type filters to one medium. See [CommunicationWayType].
	Type CommunicationWayType
	// Main, when "true", returns only main communication ways. Sevdesk wants
	// a string here, not a boolean.
	Main string
}

func (p *ListCommunicationWaysParams) query() url.Values {
	if p == nil {
		return nil
	}
	q := url.Values{}
	if p.Contact != nil {
		q.Set("contact[id]", p.Contact.ID.String())
		q.Set("contact[objectName]", p.Contact.ObjectName)
	}
	if p.Type != "" {
		q.Set("type", string(p.Type))
	}
	if p.Main != "" {
		q.Set("main", p.Main)
	}
	return q
}

// List returns communication ways matching the filter.
func (s *CommunicationWaysService) List(ctx context.Context, opts *ListCommunicationWaysParams) ([]CommunicationWay, error) {
	raw, err := s.c.do(ctx, http.MethodGet, "/CommunicationWay", opts.query(), nil)
	if err != nil {
		return nil, err
	}
	return decodeList[CommunicationWay](raw)
}

// Get returns the communication way with the given id.
func (s *CommunicationWaysService) Get(ctx context.Context, id ID) (*CommunicationWay, error) {
	raw, err := s.c.do(ctx, http.MethodGet, fmt.Sprintf("/CommunicationWay/%d", id), nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeObject[CommunicationWay](raw)
}

// Create creates a new communication way for a contact.
func (s *CommunicationWaysService) Create(ctx context.Context, params *CreateCommunicationWayParams) (*CommunicationWay, error) {
	raw, err := s.c.do(ctx, http.MethodPost, "/CommunicationWay", nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[CommunicationWay](raw)
}

// Update modifies a communication way.
func (s *CommunicationWaysService) Update(ctx context.Context, id ID, params *UpdateCommunicationWayParams) (*CommunicationWay, error) {
	raw, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/CommunicationWay/%d", id), nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[CommunicationWay](raw)
}

// Delete removes a communication way.
func (s *CommunicationWaysService) Delete(ctx context.Context, id ID) error {
	_, err := s.c.do(ctx, http.MethodDelete, fmt.Sprintf("/CommunicationWay/%d", id), nil, nil)
	return err
}

// Keys returns all available communication way keys (labels like "Work", "Private").
func (s *CommunicationWaysService) Keys(ctx context.Context) ([]CommunicationWayKey, error) {
	raw, err := s.c.do(ctx, http.MethodGet, "/CommunicationWayKey", nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeList[CommunicationWayKey](raw)
}
