package sevdesk

import (
	"context"
	"fmt"
	"iter"
	"net/http"
	"time"
)

// PrivateTransactionRulesService handles /PrivateTransactionRule endpoints.
// These rules automatically mark matching incoming transactions as "private".
type PrivateTransactionRulesService struct {
	c *Client
}

// PrivateTransactionRule defines a transaction-matching pattern. When a new
// transaction matches PaymentPurpose or CounterpartName, it is marked private.
type PrivateTransactionRule struct {
	ID         ID         `json:"id,omitempty"`
	ObjectName string     `json:"objectName,omitempty"`
	Create     *time.Time `json:"create,omitempty"`
	Update     *time.Time `json:"update,omitempty"`
	SevClient  *Ref       `json:"sevClient,omitempty"`

	PaymentPurpose  string `json:"paymentPurpose,omitempty"`
	CounterpartName string `json:"counterpartName,omitempty"`
}

// CreatePrivateTransactionRuleParams is the body for [PrivateTransactionRulesService.Create].
// At least one of PaymentPurpose or CounterpartName must be set.
type CreatePrivateTransactionRuleParams struct {
	ObjectName string `json:"objectName"` // set by Create

	// PaymentPurpose matches transactions whose payment purpose contains this string.
	PaymentPurpose *string `json:"paymentPurpose,omitempty"`
	// CounterpartName matches transactions where the counterparty name contains
	// this string.
	CounterpartName *string `json:"counterpartName,omitempty"`
}

// List returns all private transaction rules.
func (s *PrivateTransactionRulesService) List(ctx context.Context) iter.Seq2[PrivateTransactionRule, error] {
	return listIter[PrivateTransactionRule](ctx, s.c, "/PrivateTransactionRule", nil)
}

// Create creates a new private transaction rule.
func (s *PrivateTransactionRulesService) Create(ctx context.Context, params *CreatePrivateTransactionRuleParams) (*PrivateTransactionRule, error) {
	if params != nil && params.ObjectName == "" {
		params.ObjectName = ObjectPrivateTransactionRule
	}
	raw, err := s.c.do(ctx, http.MethodPost, "/PrivateTransactionRule", nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[PrivateTransactionRule](raw)
}

// Delete removes a private transaction rule.
func (s *PrivateTransactionRulesService) Delete(ctx context.Context, id ID) error {
	_, err := s.c.do(ctx, http.MethodDelete, fmt.Sprintf("/PrivateTransactionRule/%d", id), nil, nil)
	return err
}
