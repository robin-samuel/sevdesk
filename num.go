package sevdesk

import (
	"fmt"
	"strconv"
)

// Num is an integer value that sevdesk often delivers as a JSON string
// (e.g. `"status": "100"`). It accepts both quoted and unquoted forms on the
// wire and always sends as an unquoted JSON number.
type Num int64

func (n Num) String() string { return strconv.FormatInt(int64(n), 10) }

func (n Num) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(int64(n), 10)), nil
}

func (n *Num) UnmarshalJSON(b []byte) error {
	if len(b) == 0 || string(b) == "null" {
		return nil
	}
	s := string(b)
	if s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
	}
	if s == "" {
		return nil
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*n = Num(v)
	return nil
}

// Decimal is a fractional value that sevdesk often delivers as a JSON string
// (e.g. `"taxRate": "19.00"`, `"defaultDiscountAmount": "0.00"`). It accepts
// both quoted and unquoted forms and sends as an unquoted JSON number.
type Decimal float64

func (d Decimal) String() string {
	return strconv.FormatFloat(float64(d), 'f', -1, 64)
}

func (d Decimal) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatFloat(float64(d), 'f', -1, 64)), nil
}

func (d *Decimal) UnmarshalJSON(b []byte) error {
	if len(b) == 0 || string(b) == "null" {
		return nil
	}
	s := string(b)
	if s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
	}
	if s == "" {
		return nil
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	*d = Decimal(v)
	return nil
}

// Bool is a boolean that sevdesk variously delivers as true/false, "true"/"false",
// 0/1, or "0"/"1". It accepts all of those on the wire and always sends as a
// JSON boolean.
//
// For Params, use the package-level pointers [True] and [False] instead of
// constructing one by hand.
type Bool bool

func (b Bool) MarshalJSON() ([]byte, error) {
	if b {
		return []byte("true"), nil
	}
	return []byte("false"), nil
}

func (b *Bool) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "", "null", "false", `"false"`, "0", `"0"`, `""`:
		*b = false
	case "true", `"true"`, "1", `"1"`:
		*b = true
	default:
		return fmt.Errorf("sevdesk: cannot parse %s as Bool", data)
	}
	return nil
}

// True and False are ready-made *Bool values for use in *Params structs:
//
//	&sevdesk.CreateContactParams{ExemptVAT: sevdesk.True}
var (
	True  = ptrBool(true)
	False = ptrBool(false)
)

func ptrBool(v bool) *Bool {
	b := Bool(v)
	return &b
}
