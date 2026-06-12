package sevdesk

import (
	"context"
	"fmt"
	"iter"
	"net/http"
	"net/url"
	"time"
)

// AccountingContactsService handles communication with /AccountingContact endpoints.
type AccountingContactsService struct {
	c *Client
}

// AccountingContact ties a contact to debitor/creditor numbers used in DATEV exports.
type AccountingContact struct {
	ID         ID         `json:"id,omitempty"`
	ObjectName string     `json:"objectName,omitempty"`
	Create     *time.Time `json:"create,omitempty"`
	Update     *time.Time `json:"update,omitempty"`
	SevClient  *Ref       `json:"sevClient,omitempty"`

	Contact        *Ref `json:"contact,omitempty"`
	DebitorNumber  Num  `json:"debitorNumber,omitempty"`
	CreditorNumber Num  `json:"creditorNumber,omitempty"`
}

// CreateAccountingContactParams is the body for [AccountingContactsService.Create].
type CreateAccountingContactParams struct {
	// Contact this accounting record belongs to. Required.
	Contact *Ref `json:"contact"`
	// DebitorNumber is the DATEV debitor account assigned to the contact
	// (used for outgoing invoices).
	DebitorNumber *int `json:"debitorNumber,omitempty"`
	// CreditorNumber is the DATEV creditor account assigned to the contact
	// (used for incoming vouchers).
	CreditorNumber *int `json:"creditorNumber,omitempty"`
}

// UpdateAccountingContactParams is the body for [AccountingContactsService.Update].
// See [CreateAccountingContactParams] for field semantics.
type UpdateAccountingContactParams = CreateAccountingContactParams

// List returns accounting contacts; pass a contact filter to scope.
func (s *AccountingContactsService) List(ctx context.Context, contact *Ref) iter.Seq2[AccountingContact, error] {
	var q url.Values
	if contact != nil {
		q = url.Values{
			"contact[id]":         {contact.ID.String()},
			"contact[objectName]": {contact.ObjectName},
		}
	}
	return listIter[AccountingContact](ctx, s.c, "/AccountingContact", q)
}

// Get returns the accounting contact with the given id.
func (s *AccountingContactsService) Get(ctx context.Context, id ID) (*AccountingContact, error) {
	raw, err := s.c.do(ctx, http.MethodGet, fmt.Sprintf("/AccountingContact/%d", id), nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeObject[AccountingContact](raw)
}

// Create creates a new accounting contact.
func (s *AccountingContactsService) Create(ctx context.Context, params *CreateAccountingContactParams) (*AccountingContact, error) {
	raw, err := s.c.do(ctx, http.MethodPost, "/AccountingContact", nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[AccountingContact](raw)
}

// Update modifies an accounting contact.
func (s *AccountingContactsService) Update(ctx context.Context, id ID, params *UpdateAccountingContactParams) (*AccountingContact, error) {
	raw, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/AccountingContact/%d", id), nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[AccountingContact](raw)
}

// Delete removes an accounting contact.
func (s *AccountingContactsService) Delete(ctx context.Context, id ID) error {
	_, err := s.c.do(ctx, http.MethodDelete, fmt.Sprintf("/AccountingContact/%d", id), nil, nil)
	return err
}
