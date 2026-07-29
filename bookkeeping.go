package sevdesk

import (
	"context"
	"net/http"
)

// BookkeepingVersion is the version of the bookkeeping system behind an API
// key. sevdesk migrated clients from 1.0 to 2.0 individually, and the two
// differ in how positions are booked ([VoucherPosCreate.AccountDatev] vs
// [VoucherPosCreate.AccountingType]) and how VAT is expressed
// (TaxRule vs taxType/taxSet).
type BookkeepingVersion string

// BookkeepingVersion values.
const (
	// BookkeepingV1 is the original system: AccountingType and taxType/taxSet.
	BookkeepingV1 BookkeepingVersion = "1.0"
	// BookkeepingV2 is sevdesk-Update 2.0: AccountDatev and TaxRule.
	BookkeepingV2 BookkeepingVersion = "2.0"
)

// BookkeepingVersion reports which bookkeeping system the authenticated client
// is on, so you can pick between the 1.0 and 2.0 field sets:
//
//	v, err := c.BookkeepingVersion(ctx)
//	if err != nil {
//		return err
//	}
//	pos := sevdesk.VoucherPosCreate{ /* ... */ }
//	if v == sevdesk.BookkeepingV2 {
//		pos.AccountDatev = guide.Ref()
//	} else {
//		pos.AccountingType = sevdesk.AccountingTypePetrol
//	}
//
// Objects created before a client's migration keep their 1.0 representation, so
// a 2.0 answer here does not mean every response will carry 2.0 fields.
func (c *Client) BookkeepingVersion(ctx context.Context) (BookkeepingVersion, error) {
	raw, err := c.do(ctx, http.MethodGet, "/Tools/bookkeepingSystemVersion", nil, nil)
	if err != nil {
		return "", err
	}
	v, err := decodeObject[struct {
		Version BookkeepingVersion `json:"version"`
	}](raw)
	if err != nil {
		return "", err
	}
	return v.Version, nil
}
