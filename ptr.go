package sevdesk

import "time"

// Ptr returns a pointer to v. Handy for any *Params field whose zero value
// is meaningful (so omitempty isn't enough on its own):
//
//	&sevdesk.CreateContactParams{
//		Name:           sevdesk.Ptr("Acme GmbH"),
//		CustomerNumber: sevdesk.Ptr("Customer-1337"),
//		ExemptVAT:      sevdesk.True,
//	}
//
// Mind Go's type inference: sevdesk.Ptr(42) gives *int, not *int64 — use
// [Int64] (or sevdesk.Ptr(int64(42))) when you need a wider integer.
func Ptr[T any](v T) *T { return new(v) }

// String returns a pointer to v. Same as [Ptr] but type-tagged for clarity.
func String(v string) *string { return new(v) }

// Int returns a pointer to v.
func Int(v int) *int { return new(v) }

// Int64 returns a pointer to v. Prefer this over [Ptr] when targeting a
// *int64 field, since Go infers untyped integer literals as *int.
func Int64(v int64) *int64 { return new(v) }

// Float64 returns a pointer to v.
func Float64(v float64) *float64 { return new(v) }

// Time returns a pointer to v.
func Time(v time.Time) *time.Time { return new(v) }
