package sevdesk

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// ContactCustomFieldsService handles /ContactCustomField endpoints (field instances on contacts).
type ContactCustomFieldsService struct {
	c *Client
}

// ContactCustomField is a value of a custom field on a contact.
type ContactCustomField struct {
	ID         ID         `json:"id,omitempty"`
	ObjectName string     `json:"objectName,omitempty"`
	Create     *time.Time `json:"create,omitempty"`
	Update     *time.Time `json:"update,omitempty"`
	SevClient  *Ref       `json:"sevClient,omitempty"`

	Contact                   *Ref                       `json:"contact,omitempty"`
	ContactCustomFieldSetting *ContactCustomFieldSetting `json:"contactCustomFieldSetting,omitempty"`
	Value                     string                     `json:"value,omitempty"`
}

// CreateContactCustomFieldParams is the body for [ContactCustomFieldsService.Create].
type CreateContactCustomFieldParams struct {
	// Contact the custom field is set on. Required.
	Contact *Ref `json:"contact"`
	// ContactCustomFieldSetting is the field definition to attach a value to.
	// Required. List existing settings via [ContactCustomFieldSettingsService.List].
	ContactCustomFieldSetting *Ref `json:"contactCustomFieldSetting"`
	// Value stored in the field. Required.
	Value string `json:"value"`
}

// UpdateContactCustomFieldParams is the body for [ContactCustomFieldsService.Update].
// See [CreateContactCustomFieldParams] for field semantics.
type UpdateContactCustomFieldParams struct {
	Contact                   *Ref    `json:"contact,omitempty"`
	ContactCustomFieldSetting *Ref    `json:"contactCustomFieldSetting,omitempty"`
	Value                     *string `json:"value,omitempty"`
}

// List returns all contact custom field instances.
func (s *ContactCustomFieldsService) List(ctx context.Context) ([]ContactCustomField, error) {
	raw, err := s.c.do(ctx, http.MethodGet, "/ContactCustomField", nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeList[ContactCustomField](raw)
}

// Get returns the contact custom field with the given id.
func (s *ContactCustomFieldsService) Get(ctx context.Context, id ID) (*ContactCustomField, error) {
	raw, err := s.c.do(ctx, http.MethodGet, fmt.Sprintf("/ContactCustomField/%d", id), nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeObject[ContactCustomField](raw)
}

// Create sets a custom field value on a contact.
func (s *ContactCustomFieldsService) Create(ctx context.Context, params *CreateContactCustomFieldParams) (*ContactCustomField, error) {
	raw, err := s.c.do(ctx, http.MethodPost, "/ContactCustomField", nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[ContactCustomField](raw)
}

// Update modifies an existing contact custom field value.
func (s *ContactCustomFieldsService) Update(ctx context.Context, id ID, params *UpdateContactCustomFieldParams) (*ContactCustomField, error) {
	raw, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/ContactCustomField/%d", id), nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[ContactCustomField](raw)
}

// Delete removes a contact custom field value.
func (s *ContactCustomFieldsService) Delete(ctx context.Context, id ID) error {
	_, err := s.c.do(ctx, http.MethodDelete, fmt.Sprintf("/ContactCustomField/%d", id), nil, nil)
	return err
}

// ContactCustomFieldSettingsService handles /ContactCustomFieldSetting endpoints
// (field *definitions* — schemas that ContactCustomField values reference).
type ContactCustomFieldSettingsService struct {
	c *Client
}

// ContactCustomFieldSetting defines a custom field that can be set on contacts.
type ContactCustomFieldSetting struct {
	ID          ID         `json:"id,omitempty"`
	ObjectName  string     `json:"objectName,omitempty"`
	Create      *time.Time `json:"create,omitempty"`
	Update      *time.Time `json:"update,omitempty"`
	SevClient   *Ref       `json:"sevClient,omitempty"`
	Name        string     `json:"name,omitempty"`
	Identifier  string     `json:"identifier,omitempty"`
	Description string     `json:"description,omitempty"`
}

// CreateContactCustomFieldSettingParams is the body for [ContactCustomFieldSettingsService.Create].
type CreateContactCustomFieldSettingParams struct {
	// Name shown for this custom field in the sevdesk UI. Required.
	Name string `json:"name"`
	// Description is an optional hint shown alongside the field.
	Description *string `json:"description,omitempty"`
}

// UpdateContactCustomFieldSettingParams is the body for [ContactCustomFieldSettingsService.Update].
// See [CreateContactCustomFieldSettingParams] for field semantics.
type UpdateContactCustomFieldSettingParams struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// List returns all contact custom field settings.
func (s *ContactCustomFieldSettingsService) List(ctx context.Context) ([]ContactCustomFieldSetting, error) {
	raw, err := s.c.do(ctx, http.MethodGet, "/ContactCustomFieldSetting", nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeList[ContactCustomFieldSetting](raw)
}

// Get returns the contact custom field setting with the given id.
func (s *ContactCustomFieldSettingsService) Get(ctx context.Context, id ID) (*ContactCustomFieldSetting, error) {
	raw, err := s.c.do(ctx, http.MethodGet, fmt.Sprintf("/ContactCustomFieldSetting/%d", id), nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeObject[ContactCustomFieldSetting](raw)
}

// Create defines a new contact custom field.
func (s *ContactCustomFieldSettingsService) Create(ctx context.Context, params *CreateContactCustomFieldSettingParams) (*ContactCustomFieldSetting, error) {
	raw, err := s.c.do(ctx, http.MethodPost, "/ContactCustomFieldSetting", nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[ContactCustomFieldSetting](raw)
}

// Update modifies a contact custom field setting.
func (s *ContactCustomFieldSettingsService) Update(ctx context.Context, id ID, params *UpdateContactCustomFieldSettingParams) (*ContactCustomFieldSetting, error) {
	raw, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/ContactCustomFieldSetting/%d", id), nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[ContactCustomFieldSetting](raw)
}

// Delete removes a contact custom field setting.
func (s *ContactCustomFieldSettingsService) Delete(ctx context.Context, id ID) error {
	_, err := s.c.do(ctx, http.MethodDelete, fmt.Sprintf("/ContactCustomFieldSetting/%d", id), nil, nil)
	return err
}

// ReferenceCount returns the number of ContactCustomField values referencing
// the given setting (useful before deletion).
func (s *ContactCustomFieldSettingsService) ReferenceCount(ctx context.Context, id ID) (int, error) {
	raw, err := s.c.do(ctx, http.MethodGet, fmt.Sprintf("/ContactCustomFieldSetting/%d/getReferenceCount", id), nil, nil)
	if err != nil {
		return 0, err
	}
	var n Num
	if err := n.UnmarshalJSON(raw); err != nil {
		return 0, err
	}
	return int(n), nil
}
