# sevdesk

[![Go Reference](https://pkg.go.dev/badge/github.com/robin-samuel/sevdesk.svg)](https://pkg.go.dev/github.com/robin-samuel/sevdesk)
[![Go Version](https://img.shields.io/badge/go-1.26+-00ADD8.svg)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Go SDK for the [sevdesk](https://sevdesk.com/) bookkeeping API.

## Install

```sh
go get github.com/robin-samuel/sevdesk
```

Requires Go 1.26 or later.

## Quick start

```go
package main

import (
	"context"
	"log"

	"github.com/robin-samuel/sevdesk"
)

func main() {
	c := sevdesk.New("YOUR_API_KEY")

	contact, err := c.Contacts.Get(context.Background(), 42)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(contact.Name, contact.CustomerNumber)
}
```

## Services

Each resource hangs off the `Client` as its own service:

| Service | Resource |
| --- | --- |
| `c.Contacts` | Contacts (organizations and people) |
| `c.ContactAddresses` | Postal addresses attached to contacts |
| `c.CommunicationWays` | Email / phone / web / mobile entries |
| `c.AccountingContacts` | DATEV debitor/creditor numbers |
| `c.ContactCustomFields` / `c.ContactCustomFieldSettings` | Custom field values and their definitions |
| `c.Invoices` | Invoices (Rechnungen) — full lifecycle including rendering, sending, booking |
| `c.CreditNotes` | Credit notes (Gutschriften), creatable from invoices or vouchers |
| `c.Orders` / `c.OrderPositions` | Orders, quotes, packing lists, contract notes |
| `c.Vouchers` | Expense / income vouchers, with file upload and recurring support |
| `c.CheckAccounts` | Bank accounts, clearing accounts, cash registers |
| `c.Transactions` | Check-account transactions |
| `c.PrivateTransactionRules` | Rules for auto-marking transactions as private |
| `c.Parts` | Inventory items |
| `c.Tags` | Tags attachable to invoices, orders, vouchers, credit notes |
| `c.AccountingTypes` | List/get booking accounts (bookkeeping **1.0**). 125 German entries are also embedded as named `*Ref` variables (e.g. `sevdesk.AccountingTypePetrol`) for direct use in Params. |
| `c.ReceiptGuidance` | Look up bookable DATEV accounts by number, tax rule, or category (bookkeeping **2.0**). The 652-entry German chart is also bundled — `sevdesk.AccountDatev(6815)` — as are the tax rules, e.g. `sevdesk.TaxRuleDeductibleExpense`. |

Full method reference: [pkg.go.dev/github.com/robin-samuel/sevdesk](https://pkg.go.dev/github.com/robin-samuel/sevdesk).

## Working with optional fields

Request params use pointers for optional fields. Use the helper constructors so call sites stay readable:

```go
_, err := c.Contacts.Create(ctx, &sevdesk.CreateContactParams{
	Category:  sevdesk.CategoryRef(3), // 3 = Customer
	Name:      sevdesk.String("Acme GmbH"),
	Status:    sevdesk.ContactStatusActive,
	ExemptVAT: sevdesk.True,
})
```

For typed enums there's a constant for every documented value (and they're typed, so the compiler catches mismatches):

```go
for invoice, err := range c.Invoices.List(ctx, &sevdesk.ListInvoicesParams{
	Status: sevdesk.InvoiceStatusOpen,
}) {
	if err != nil { return err }
	// ...
}
```

List endpoints return `iter.Seq2[T, error]` and paginate automatically. Drain into a slice with [`sevdesk.Collect`](https://pkg.go.dev/github.com/robin-samuel/sevdesk#Collect):

```go
invoices, err := sevdesk.Collect(c.Invoices.List(ctx, nil))
```

## sevdesk-Update 2.0

sevdesk moves accounts from bookkeeping system 1.0 to 2.0 one at a time, and the two want different fields. Ask which one your key is on:

```go
v, err := c.BookkeepingVersion(ctx) // sevdesk.BookkeepingV1 or BookkeepingV2
```

**Booking accounts.** 2.0 replaces `accountingType` with `accountDatev`, and only accounts returned by the receipt-guidance endpoints are bookable — custom accounts are gone. The German DATEV chart ships with the SDK, so the common case needs no API call and no number lookup:

```go
pos := sevdesk.VoucherPosCreate{
	AccountDatev: sevdesk.AccountDatevBuerobedarf, // 6815 Bürobedarf
	TaxRate:      19,
	Net:          true,
	SumNet:       100,
	SumGross:     119,
}
```

77 accounts a business books to routinely have names — `AccountDatevErloese19USt`, `AccountDatevLaufendeKfzBetriebskosten`, `AccountDatevRechtsUndBeratungskosten`, `AccountDatevReisekostenUnternehmer`. Identifiers keep the German DATEV labels (umlauts as `ue`/`oe`/`ae`/`ss`) so they line up with what you see in sevdesk and in your accountant's chart; the DATEV number and exact label are in each doc comment. All 652 accounts of the chart stay reachable by number:

```go
pos.AccountDatev = sevdesk.AccountDatev(7364) // nil if the number isn't in the chart
```

To discover accounts at runtime — or to find out which tax rules and rates one accepts — ask the guidance endpoints:

```go
guide, err := c.ReceiptGuidance.ForAccountNumber(ctx, 6815)
// guide.Ref() is the *Ref for AccountDatev above
for _, rule := range guide.AllowedTaxRules {
	log.Println(rule.ID, rule.Name, rule.TaxRates) // 9 VORST_ABZUGSF_AUFW [ZERO SEVEN NINETEEN]
}
```

1.0 ids still work where the SKR number survived into the guidance — sevdesk maps them — so `sevdesk.AccountingTypePetrol` and friends keep functioning on a migrated client. Mismatched account/rule/rate combinations answer HTTP 422 (`sevdesk.ErrValidation`).

**VAT rules.** 2.0 replaces `taxType`/`taxSet` with `taxRule` and drops `taxType: "custom"` outright. The documented rules ship as named refs, split by document side:

```go
&sevdesk.InvoiceCreateFields{TaxRule: sevdesk.TaxRuleTaxableRevenue}   // revenue
&sevdesk.VoucherCreateFields{TaxRule: sevdesk.TaxRuleDeductibleExpense} // expense
&sevdesk.InvoiceCreateFields{TaxRule: sevdesk.TaxRuleSmallBusiness}     // Kleinunternehmer
```

Also changed in 2.0, and noted on the affected methods: vouchers can no longer be created as paid (book them instead), credit notes can only be created as drafts, `CreditNotes.CreateFromVoucher` is removed, and `Invoices.ChangeLayout` is removed.

Objects created before a migration keep their 1.0 shape, so responses can still carry `taxType` and `accountingType` on a 2.0 client — both fields stay populated on read models.

## Wire-format quirks

sevdesk's responses have some sharp edges. The SDK papers over them so your Go code stays clean:

- **IDs come back as JSON strings** (`"131941637"`) — exposed as `sevdesk.ID` (an `int64`) with a custom unmarshaler that accepts both forms.
- **Numerics arrive as strings too** (`"100"`, `"19.00"`) — exposed as `sevdesk.Num` and `sevdesk.Decimal` with the same dual-form unmarshaling.
- **Booleans arrive as `"0"` / `"1"`** (and sometimes `"true"` / `"false"`) — exposed as `sevdesk.Bool`. Use the package-level `sevdesk.True` / `sevdesk.False` as `*Bool` params.
- **Related entities use `{id, objectName}` pairs everywhere** — exposed as `sevdesk.Ref`. Use the typed helpers: `sevdesk.ContactRef(id)`, `sevdesk.InvoiceRef(id)`, etc.
- **All responses are wrapped in `{"objects": ...}`** — unwrapped automatically.
- **Country references are pre-built**: `sevdesk.CountryDE`, `sevdesk.CountryUS`, `sevdesk.CountryGB`, … one per ISO 3166-1 alpha-2 code, ready to drop into any `*Ref` field.

## Authentication

API keys come from your sevdesk account under *Settings → User → API token*. Pass it to `New`:

```go
c := sevdesk.New(os.Getenv("SEVDESK_API_KEY"))
```

To target the internal sevdesk instance or a self-hosted setup, override the base URL:

```go
c := sevdesk.New("KEY",
	sevdesk.WithBaseURL("http://sevdesk.local/api/v1"),
	sevdesk.WithHTTPClient(&http.Client{Timeout: 30 * time.Second}),
)
```

## Errors

Non-2xx responses return a `*sevdesk.Error` with the status, method, path, message, and the sevdesk exception UUID (when present). Sentinels cover the common cases — match them with `errors.Is`:

```go
contact, err := c.Contacts.Get(ctx, 999999)
switch {
case errors.Is(err, sevdesk.ErrNotFound):     // 404
case errors.Is(err, sevdesk.ErrUnauthorized): // 401 — bad API key
case errors.Is(err, sevdesk.ErrForbidden):    // 403
case errors.Is(err, sevdesk.ErrConflict):     // 409 — e.g. deleting a non-draft
case errors.Is(err, sevdesk.ErrValidation):   // 422 — e.g. account/tax-rule mismatch
case errors.Is(err, sevdesk.ErrRateLimit):    // 429
}

var apiErr *sevdesk.Error
if errors.As(err, &apiErr) {
	log.Printf("%s %s -> %d %s (%s)",
		apiErr.Method, apiErr.Path, apiErr.StatusCode, apiErr.Message, apiErr.UUID)
}
```

## Status

The high-traffic services (Contacts, Vouchers, Invoices, CreditNotes, CheckAccounts, Transactions) were modelled directly from live API responses so the struct fields match reality, not just the OpenAPI spec. Other services are spec-derived and should be considered slightly less battle-tested — please open an issue if you hit a missing field.

## License

[MIT](LICENSE).
