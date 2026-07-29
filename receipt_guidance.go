package sevdesk

import (
	"context"
	"iter"
	"net/http"
	"net/url"
	"strconv"
)

// ReceiptGuidanceService exposes sevdesk's helper endpoints for finding the
// booking account to use on a voucher position.
//
// This is the discovery mechanism of sevdesk-Update 2.0: the bookable accounts
// aren't a documented closed set, and 2.0 dropped custom accounts entirely, so
// only what these endpoints return can be booked. Each guide also states which
// tax rules and tax rates the account combines with — a mismatch between
// account, [VoucherCreateFields.TaxRule] and position TaxRate is the usual
// cause of a 422.
//
// On sevdesk 1.0 the equivalent catalogue is [AccountingTypesService].
type ReceiptGuidanceService struct {
	c *Client
}

// ReceiptGuide describes one DATEV account and how it may be used.
type ReceiptGuide struct {
	// AccountDatevID identifies the account. Use [ReceiptGuide.Ref] to turn it
	// into the [Ref] that [VoucherPosCreate.AccountDatev] wants.
	AccountDatevID ID `json:"accountDatevId,omitempty"`
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

// Ref returns the [Ref] for this account, ready for
// [VoucherPosCreate.AccountDatev].
func (g ReceiptGuide) Ref() *Ref { return AccountDatevRef(g.AccountDatevID) }

// ReceiptGuideTaxRule is one tax rule compatible with a [ReceiptGuide] account.
type ReceiptGuideTaxRule struct {
	// Name is the internal tax-rule identifier (e.g. "USTPFL_UMS_EINN").
	Name string `json:"name,omitempty"`
	// Description is a human-readable label (e.g. "Umsatzsteuerpflichtige Umsätze").
	Description string `json:"description,omitempty"`
	// ID is the tax-rule ID. Use [ReceiptGuideTaxRule.Ref] to get the [Ref]
	// for a TaxRule field; the named rules in [TaxRuleTaxableRevenue] and
	// friends cover the same set.
	ID ID `json:"id,omitempty"`
	// TaxRates are the VAT rates combinable with this rule, as symbolic tokens
	// rather than numbers. A position's TaxRate must be one of them.
	TaxRates []ReceiptGuideTaxRate `json:"taxRates,omitempty"`
}

// ReceiptGuideTaxRate is one VAT rate a [ReceiptGuideTaxRule] permits. sevdesk
// reports these as tokens ("ZERO", "SEVEN", "NINETEEN") rather than numbers.
type ReceiptGuideTaxRate string

// ReceiptGuideTaxRate values.
const (
	ReceiptGuideTaxRateZero     = ReceiptGuideTaxRate("ZERO")
	ReceiptGuideTaxRateSeven    = ReceiptGuideTaxRate("SEVEN")
	ReceiptGuideTaxRateNineteen = ReceiptGuideTaxRate("NINETEEN")
)

// Decimal converts the token to the percentage you put in a position's TaxRate.
// ok is false for a token this SDK doesn't know — the One Stop Shop rules, for
// instance, allow rates that depend on the destination country.
func (r ReceiptGuideTaxRate) Decimal() (rate Decimal, ok bool) {
	switch r {
	case ReceiptGuideTaxRateZero:
		return 0, true
	case ReceiptGuideTaxRateSeven:
		return 7, true
	case ReceiptGuideTaxRateNineteen:
		return 19, true
	}
	return 0, false
}

// Ref returns the [Ref] for this tax rule, ready for a TaxRule field such as
// [VoucherCreateFields.TaxRule].
func (r ReceiptGuideTaxRule) Ref() *Ref { return TaxRuleRef(r.ID) }

// ForAllAccounts returns guidance for every available DATEV account.
func (s *ReceiptGuidanceService) ForAllAccounts(ctx context.Context) iter.Seq2[ReceiptGuide, error] {
	return listOnce[ReceiptGuide](ctx, s.c, "/ReceiptGuidance/forAllAccounts", nil)
}

// ForAccountNumber returns guidance for one DATEV account number
// (e.g. 6815 for "Bürobedarf"). Unlike the other guidance endpoints this one
// answers with a single account.
//
// An account number the client isn't offered fails with [ErrValidation], not
// [ErrNotFound].
func (s *ReceiptGuidanceService) ForAccountNumber(ctx context.Context, accountNumber int) (*ReceiptGuide, error) {
	q := url.Values{"accountNumber": {strconv.Itoa(accountNumber)}}
	raw, err := s.c.do(ctx, http.MethodGet, "/ReceiptGuidance/forAccountNumber", q, nil)
	if err != nil {
		return nil, err
	}
	return decodeObject[ReceiptGuide](raw)
}

// ForTaxRule returns accounts compatible with the given tax rule, identified by
// its internal code — see [TaxRuleNameTaxableRevenue] and friends:
//
//	c.ReceiptGuidance.ForTaxRule(ctx, sevdesk.TaxRuleNameDeductibleExpense)
func (s *ReceiptGuidanceService) ForTaxRule(ctx context.Context, name TaxRuleName) iter.Seq2[ReceiptGuide, error] {
	q := url.Values{"taxRule": {string(name)}}
	return listOnce[ReceiptGuide](ctx, s.c, "/ReceiptGuidance/forTaxRule", q)
}

// ForRevenue returns guidance for all revenue (income) accounts.
func (s *ReceiptGuidanceService) ForRevenue(ctx context.Context) iter.Seq2[ReceiptGuide, error] {
	return listOnce[ReceiptGuide](ctx, s.c, "/ReceiptGuidance/forRevenue", nil)
}

// ForExpense returns guidance for all expense accounts.
func (s *ReceiptGuidanceService) ForExpense(ctx context.Context) iter.Seq2[ReceiptGuide, error] {
	return listOnce[ReceiptGuide](ctx, s.c, "/ReceiptGuidance/forExpense", nil)
}
