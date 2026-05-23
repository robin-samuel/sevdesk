package sevdesk

import "time"

// String returns a pointer to the given value. Use in *Params structs:
//
//	&sevdesk.UpdateContactParams{Name: sevdesk.String("Jon Snow")}
func String(v string) *string { return new(v) }

// Int returns a pointer to the given value.
func Int(v int) *int { return new(v) }

// Int64 returns a pointer to the given value.
func Int64(v int64) *int64 { return new(v) }

// Float64 returns a pointer to the given value.
func Float64(v float64) *float64 { return new(v) }

// Time returns a pointer to the given value.
func Time(v time.Time) *time.Time { return new(v) }
