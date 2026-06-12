package sevdesk

import (
	"context"
	"fmt"
	"iter"
	"net/http"
	"net/url"
	"time"
)

// TagsService handles /Tag endpoints.
type TagsService struct {
	c *Client
}

// Tag is a label that can be attached to invoices, orders, vouchers, credit notes.
type Tag struct {
	ID         ID         `json:"id,omitempty"`
	ObjectName string     `json:"objectName,omitempty"`
	Create     *time.Time `json:"create,omitempty"`
	SevClient  *Ref       `json:"sevClient,omitempty"`
	Name       string     `json:"name,omitempty"`
}

// TagRelation represents a tag attached to a specific object.
type TagRelation struct {
	ID         ID         `json:"id,omitempty"`
	ObjectName string     `json:"objectName,omitempty"`
	Create     *time.Time `json:"create,omitempty"`
	SevClient  *Ref       `json:"sevClient,omitempty"`
	Tag        *Ref       `json:"tag,omitempty"`
	Object     *Ref       `json:"object,omitempty"`
}

// CreateTagParams is the body for [TagsService.Create]. The tag is attached
// to the given object on creation (sevdesk requires both in one call).
type CreateTagParams struct {
	// Name of the tag. Required.
	Name string `json:"name"`
	// Object the tag is attached to (Invoice, Order, Voucher, or CreditNote).
	// Required.
	Object Ref `json:"object"`
}

// ListTagsParams filters [TagsService.List].
type ListTagsParams struct {
	// ID looks up a single tag by id.
	ID ID
	// Name matches the exact tag name.
	Name string
}

func (p *ListTagsParams) query() url.Values {
	if p == nil {
		return nil
	}
	q := url.Values{}
	if p.ID != 0 {
		q.Set("id", p.ID.String())
	}
	if p.Name != "" {
		q.Set("name", p.Name)
	}
	return q
}

// List returns tags matching the filter.
func (s *TagsService) List(ctx context.Context, opts *ListTagsParams) iter.Seq2[Tag, error] {
	return listIter[Tag](ctx, s.c, "/Tag", opts.query())
}

// Get returns the tag with the given id.
func (s *TagsService) Get(ctx context.Context, id ID) (*Tag, error) {
	raw, err := s.c.do(ctx, http.MethodGet, fmt.Sprintf("/Tag/%d", id), nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeObject[Tag](raw)
}

// Create creates a tag and attaches it to the given object in one call.
func (s *TagsService) Create(ctx context.Context, params *CreateTagParams) (*TagRelation, error) {
	raw, err := s.c.do(ctx, http.MethodPost, "/Tag/Factory/create", nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[TagRelation](raw)
}

// Update renames a tag.
func (s *TagsService) Update(ctx context.Context, id ID, name string) (*Tag, error) {
	body := map[string]string{"name": name}
	raw, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/Tag/%d", id), nil, body)
	if err != nil {
		return nil, err
	}
	return decodeObject[Tag](raw)
}

// Delete removes a tag.
func (s *TagsService) Delete(ctx context.Context, id ID) error {
	_, err := s.c.do(ctx, http.MethodDelete, fmt.Sprintf("/Tag/%d", id), nil, nil)
	return err
}

// Relations returns all tag-to-object relations across the account.
func (s *TagsService) Relations(ctx context.Context) iter.Seq2[TagRelation, error] {
	return listIter[TagRelation](ctx, s.c, "/TagRelation", nil)
}
