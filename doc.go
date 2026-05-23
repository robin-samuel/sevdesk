// Package sevdesk is a Go client for the sevdesk REST API.
//
// # Authentication
//
// Get an API key from the sevdesk UI (Settings → User → API token) and pass it
// to [New]:
//
//	c := sevdesk.New("YOUR_API_KEY")
//
// # Services
//
// Each resource is reachable as a service hanging off the [Client]:
//
//	c.Contacts.Get(ctx, 42)
//	c.Invoices.List(ctx, &sevdesk.ListInvoicesParams{Status: sevdesk.InvoiceStatusOpen})
//	c.Vouchers.Book(ctx, voucherID, &sevdesk.BookVoucherParams{...})
//
// # Optional fields
//
// Request params use pointers for optional fields. The package provides
// helpers so call sites stay readable:
//
//	&sevdesk.CreateContactParams{
//		Category:  sevdesk.CategoryRef(3),
//		Name:      sevdesk.String("Acme GmbH"),
//		ExemptVAT: sevdesk.True,
//	}
//
// # Wire types
//
// sevdesk returns IDs and numeric fields as JSON strings (e.g. `"100"`).
// The SDK exposes these through [ID], [Num], [Decimal], and [Bool], each
// of which accepts both quoted and unquoted forms on the wire.
//
// # Errors
//
// Non-2xx responses return a [*Error] with status, message, and the sevdesk
// exception UUID. [ErrNotFound] is matchable via [errors.Is].
package sevdesk
