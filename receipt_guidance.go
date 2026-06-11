package sevdesk

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// ReceiptGuidanceService exposes sevdesk's helper endpoints for finding the
// right AccountingType to use on a voucher position or credit note.
//
// AccountingType IDs aren't documented as a closed set; you discover them by
// querying these guidance endpoints, e.g. by DATEV account number or by tax
// rule.
type ReceiptGuidanceService struct {
	c *Client
}

// ReceiptGuide describes one DATEV account and how it may be used.
type ReceiptGuide struct {
	// AccountDatevID is the ID to put in your AccountingType reference
	// when booking against this account.
	AccountDatevID Num `json:"accountDatevId,omitempty"`
	// AccountNumber is the DATEV account number (e.g. "8400").
	AccountNumber string `json:"accountNumber,omitempty"`
	// AccountName is the human-readable account name (e.g. "Erlöse 19% USt").
	AccountName string `json:"accountName,omitempty"`
	// Description explains what the account is for.
	Description string `json:"description,omitempty"`
	// AllowedTaxRules lists the tax rules this account can be combined with.
	AllowedTaxRules []ReceiptGuideTaxRule `json:"allowedTaxRules,omitempty"`
	// AllowedReceiptTypes lists which receipt types may post to this account.
	AllowedReceiptTypes []string `json:"allowedReceiptTypes,omitempty"`
}

// ReceiptGuideTaxRule is one tax rule compatible with a [ReceiptGuide] account.
type ReceiptGuideTaxRule struct {
	// Name is the internal tax-rule identifier (e.g. "USTPFL_UMS_EINN").
	Name string `json:"name,omitempty"`
	// Description is a human-readable label (e.g. "Umsatzsteuerpflichtige Umsätze").
	Description string `json:"description,omitempty"`
	// ID is the tax-rule ID; put this in a TaxRule [Ref].
	ID Num `json:"id,omitempty"`
	// TaxRates are the VAT rates combinable with this rule.
	TaxRates []string `json:"taxRates,omitempty"`
}

// ForAllAccounts returns guidance for every available DATEV account.
func (s *ReceiptGuidanceService) ForAllAccounts(ctx context.Context) ([]ReceiptGuide, error) {
	raw, err := s.c.do(ctx, http.MethodGet, "/ReceiptGuidance/forAllAccounts", nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeList[ReceiptGuide](raw)
}

// ForAccountNumber returns guidance for a specific DATEV account number
// (e.g. 8400 for "Erlöse 19% USt").
func (s *ReceiptGuidanceService) ForAccountNumber(ctx context.Context, accountNumber int) ([]ReceiptGuide, error) {
	q := url.Values{"accountNumber": {fmt.Sprintf("%d", accountNumber)}}
	raw, err := s.c.do(ctx, http.MethodGet, "/ReceiptGuidance/forAccountNumber", q, nil)
	if err != nil {
		return nil, err
	}
	return decodeList[ReceiptGuide](raw)
}

// ForTaxRule returns accounts compatible with the given tax-rule name
// (e.g. "USTPFL_UMS_EINN").
func (s *ReceiptGuidanceService) ForTaxRule(ctx context.Context, taxRule string) ([]ReceiptGuide, error) {
	q := url.Values{"taxRule": {taxRule}}
	raw, err := s.c.do(ctx, http.MethodGet, "/ReceiptGuidance/forTaxRule", q, nil)
	if err != nil {
		return nil, err
	}
	return decodeList[ReceiptGuide](raw)
}

// ForRevenue returns guidance for all revenue (income) accounts.
func (s *ReceiptGuidanceService) ForRevenue(ctx context.Context) ([]ReceiptGuide, error) {
	raw, err := s.c.do(ctx, http.MethodGet, "/ReceiptGuidance/forRevenue", nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeList[ReceiptGuide](raw)
}

// ForExpense returns guidance for all expense accounts.
func (s *ReceiptGuidanceService) ForExpense(ctx context.Context) ([]ReceiptGuide, error) {
	raw, err := s.c.do(ctx, http.MethodGet, "/ReceiptGuidance/forExpense", nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeList[ReceiptGuide](raw)
}
