package sevdesk

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// AccountingTypesService exposes sevdesk's /AccountingType endpoints.
//
// AccountingType is the booking account used on voucher positions in the
// sevdesk v1 bookkeeping system. Sevdesk maintains the catalogue — your job
// is to GET the existing entries and pick the right ID for [VoucherPosCreate].
//
// The German subset (125 entries with SKR03 / SKR04 numbers) is also bundled
// as named [*Ref] variables — e.g. [AccountingTypePetrol], usable directly in
// [VoucherPosCreate.AccountingType] without any API call.
//
// On the newer sevdesk v2 (DATEV-based) system, use [ReceiptGuidanceService]
// instead to discover [AccountingType] references via DATEV account numbers.
type AccountingTypesService struct {
	c *Client
}

// AccountingType is a booking account in the sevdesk v1 system.
type AccountingType struct {
	ID         ID         `json:"id,omitempty"`
	ObjectName string     `json:"objectName,omitempty"`
	Create     *time.Time `json:"create,omitempty"`
	Update     *time.Time `json:"update,omitempty"`
	SevClient  *Ref       `json:"sevClient,omitempty"`

	// Name is the human-readable label (e.g. "Erlöse 19%").
	Name string `json:"name,omitempty"`
	// TranslationCode is the i18n key sevdesk uses internally.
	TranslationCode string `json:"translationCode,omitempty"`
	// SKR03 is the German Standardkontenrahmen 03 account number.
	SKR03 string `json:"skr03,omitempty"`
	// SKR04 is the German Standardkontenrahmen 04 account number.
	SKR04 string `json:"skr04,omitempty"`
	// Parent links to a higher-level grouping account.
	Parent *Ref `json:"parent,omitempty"`
	// ShowOnInvoice indicates the account may appear on outgoing invoices
	// (income-side).
	ShowOnInvoice Bool `json:"showOnInvoice,omitempty"`
	// ShowOnExpense indicates the account may appear on incoming vouchers
	// (expense-side).
	ShowOnExpense Bool `json:"showOnExpense,omitempty"`
}

// List returns all AccountingType records.
func (s *AccountingTypesService) List(ctx context.Context) ([]AccountingType, error) {
	raw, err := s.c.do(ctx, http.MethodGet, "/AccountingType", nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeList[AccountingType](raw)
}

// Get returns the AccountingType with the given id.
func (s *AccountingTypesService) Get(ctx context.Context, id ID) (*AccountingType, error) {
	raw, err := s.c.do(ctx, http.MethodGet, fmt.Sprintf("/AccountingType/%d", id), nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeObject[AccountingType](raw)
}
