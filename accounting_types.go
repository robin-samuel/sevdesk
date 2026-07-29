package sevdesk

import (
	"context"
	"fmt"
	"iter"
	"net/http"
	"time"
)

// AccountingTypesService exposes sevdesk's /AccountingType endpoints.
//
// AccountingType is the booking account of the sevdesk 1.0 bookkeeping system.
// sevdesk maintains the catalogue — your job is to GET the existing entries and
// pick the right ID for [VoucherPosCreate.AccountingType].
//
// The German subset (125 entries with SKR03 / SKR04 numbers) is also bundled as
// named [*Ref] variables — e.g. [AccountingTypePetrol], usable without any API
// call.
//
// On sevdesk-Update 2.0 this catalogue is superseded by AccountDatev: use
// [ReceiptGuidanceService] to discover accounts and set
// [VoucherPosCreate.AccountDatev]. Existing AccountingType ids whose SKR number
// survived into the receipt guidance keep working — sevdesk maps them — so 1.0
// code doesn't break on migration, but the account's allowed tax rules and
// rates are enforced from then on. [Client.BookkeepingVersion] tells you which
// system a key is on.
type AccountingTypesService struct {
	c *Client
}

// AccountingType is a booking account in the sevdesk 1.0 system.
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
func (s *AccountingTypesService) List(ctx context.Context) iter.Seq2[AccountingType, error] {
	return listIter[AccountingType](ctx, s.c, "/AccountingType", nil)
}

// Get returns the AccountingType with the given id.
func (s *AccountingTypesService) Get(ctx context.Context, id ID) (*AccountingType, error) {
	raw, err := s.c.do(ctx, http.MethodGet, fmt.Sprintf("/AccountingType/%d", id), nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeObject[AccountingType](raw)
}
