package sevdesk

import (
	"strconv"
)

// ID is an entity identifier. sevdesk returns IDs as JSON strings (`"42"`) but
// accepts them as integers in URL paths and request bodies, so ID is stored as
// int64 with a custom UnmarshalJSON that accepts both forms.
type ID int64

func (id ID) String() string { return strconv.FormatInt(int64(id), 10) }

func (id ID) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(int64(id), 10)), nil
}

func (id *ID) UnmarshalJSON(b []byte) error {
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
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*id = ID(n)
	return nil
}
