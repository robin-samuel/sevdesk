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
// # sevdesk-Update 2.0
//
// sevdesk migrates accounts from bookkeeping system 1.0 to 2.0 individually,
// and the two want different fields. [Client.BookkeepingVersion] reports which
// one a key is on.
//
// Booking accounts: 2.0 replaces AccountingType with AccountDatev. Set
// [VoucherPosCreate.AccountDatev] from the bundled German chart — the common
// accounts are named ([AccountDatevBuerobedarf], [AccountDatevErloese19USt], …)
// and any of the 652 is reachable by number with [AccountDatev] — or from
// [ReceiptGuidanceService], which also states the tax rules and rates an account
// allows. Custom accounts are gone; 1.0 ids whose SKR number survived into the
// guidance are mapped automatically.
//
// VAT: 2.0 replaces taxType/taxSet with taxRule, and drops taxType "custom"
// entirely. Use the named rules — [TaxRuleTaxableRevenue] on invoices, orders
// and credit notes, [TaxRuleDeductibleExpense] on vouchers,
// [TaxRuleSmallBusiness] for Kleinunternehmer:
//
//	&sevdesk.VoucherCreateFields{
//		TaxRule: sevdesk.TaxRuleDeductibleExpense,
//		// ...
//	}
//
// A rule that doesn't fit the account or the position's tax rate is rejected
// with HTTP 422, matchable as [ErrValidation].
//
// Objects created before a migration keep their 1.0 representation, so
// responses may carry taxType and accountingType even on a 2.0 client.
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
