package sevdesk

import (
	"context"
	"fmt"
	"iter"
	"net/http"
	"time"
)

// ContactAddressesService handles communication with /ContactAddress endpoints.
type ContactAddressesService struct {
	c *Client
}

// ContactAddress is a physical address belonging to a contact.
type ContactAddress struct {
	ID         ID         `json:"id,omitempty"`
	ObjectName string     `json:"objectName,omitempty"`
	Create     *time.Time `json:"create,omitempty"`
	Update     *time.Time `json:"update,omitempty"`
	SevClient  *Ref       `json:"sevClient,omitempty"`

	Contact  *Ref   `json:"contact,omitempty"`
	Street   string `json:"street,omitempty"`
	Zip      string `json:"zip,omitempty"`
	City     string `json:"city,omitempty"`
	Country  *Ref   `json:"country,omitempty"`
	Category *Ref   `json:"category,omitempty"`
	Name     string `json:"name,omitempty"`
	Name2    string `json:"name2,omitempty"`
	Name3    string `json:"name3,omitempty"`
	Name4    string `json:"name4,omitempty"`
}

// CreateContactAddressParams is the body for [ContactAddressesService.Create].
type CreateContactAddressParams struct {
	// Contact this address belongs to. Required.
	Contact *Ref `json:"contact"`
	// Street is the street name and house number.
	Street *string `json:"street,omitempty"`
	// Zip is the postal code.
	Zip *string `json:"zip,omitempty"`
	// City name.
	City *string `json:"city,omitempty"`
	// Country reference. Use [CountryRef] to build it.
	Country *Ref `json:"country,omitempty"`
	// Category classifies the address type (e.g. billing, shipping).
	Category *Ref `json:"category,omitempty"`
	// Name is the first line of the address (overrides the contact's name here).
	Name *string `json:"name,omitempty"`
	// Name2..Name4 are additional address lines shown below Name.
	Name2 *string `json:"name2,omitempty"`
	Name3 *string `json:"name3,omitempty"`
	Name4 *string `json:"name4,omitempty"`
}

// UpdateContactAddressParams is the body for [ContactAddressesService.Update].
// See [CreateContactAddressParams] for field semantics.
type UpdateContactAddressParams = CreateContactAddressParams

// List returns all contact addresses.
func (s *ContactAddressesService) List(ctx context.Context) iter.Seq2[ContactAddress, error] {
	return listIter[ContactAddress](ctx, s.c, "/ContactAddress", nil)
}

// Get returns the contact address with the given id.
func (s *ContactAddressesService) Get(ctx context.Context, id ID) (*ContactAddress, error) {
	raw, err := s.c.do(ctx, http.MethodGet, fmt.Sprintf("/ContactAddress/%d", id), nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeObject[ContactAddress](raw)
}

// Create creates a new contact address.
func (s *ContactAddressesService) Create(ctx context.Context, params *CreateContactAddressParams) (*ContactAddress, error) {
	raw, err := s.c.do(ctx, http.MethodPost, "/ContactAddress", nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[ContactAddress](raw)
}

// Update modifies an existing contact address.
func (s *ContactAddressesService) Update(ctx context.Context, id ID, params *UpdateContactAddressParams) (*ContactAddress, error) {
	raw, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/ContactAddress/%d", id), nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[ContactAddress](raw)
}

// Delete removes a contact address.
func (s *ContactAddressesService) Delete(ctx context.Context, id ID) error {
	_, err := s.c.do(ctx, http.MethodDelete, fmt.Sprintf("/ContactAddress/%d", id), nil, nil)
	return err
}
