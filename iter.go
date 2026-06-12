package sevdesk

import (
	"context"
	"iter"
	"net/http"
	"net/url"
	"strconv"
)

// listIterPageSize is how many records the SDK pulls per request when paging
// through a list endpoint.
const listIterPageSize = 100

// listIter is the internal auto-paginating iterator. It calls path with the
// given baseQuery plus limit/offset, repeating until the API returns a short
// page. Errors from the underlying request stop the iteration.
func listIter[T any](ctx context.Context, c *Client, path string, baseQuery url.Values) iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		var zero T
		for offset := 0; ; offset += listIterPageSize {
			q := url.Values{}
			for k, vv := range baseQuery {
				q[k] = append([]string(nil), vv...)
			}
			q.Set("limit", strconv.Itoa(listIterPageSize))
			q.Set("offset", strconv.Itoa(offset))

			raw, err := c.do(ctx, http.MethodGet, path, q, nil)
			if err != nil {
				yield(zero, err)
				return
			}
			items, err := decodeList[T](raw)
			if err != nil {
				yield(zero, err)
				return
			}
			for _, item := range items {
				if !yield(item, nil) {
					return
				}
			}
			if len(items) < listIterPageSize {
				return
			}
		}
	}
}

// Collect drains a list iterator into a slice. It stops at the first error and
// returns whatever was collected so far together with the error.
//
//	all, err := sevdesk.Collect(c.Invoices.List(ctx, params))
func Collect[T any](seq iter.Seq2[T, error]) ([]T, error) {
	var items []T
	for item, err := range seq {
		if err != nil {
			return items, err
		}
		items = append(items, item)
	}
	return items, nil
}
