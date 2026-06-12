package sevdesk

import (
	"context"
	"fmt"
	"iter"
	"net/http"
	"net/url"
	"time"
)

// PartsService handles communication with /Part endpoints.
type PartsService struct {
	c *Client
}

// Part is an item in your sevdesk inventory.
type Part struct {
	ID         ID         `json:"id,omitempty"`
	ObjectName string     `json:"objectName,omitempty"`
	Create     *time.Time `json:"create,omitempty"`
	Update     *time.Time `json:"update,omitempty"`
	SevClient  *Ref       `json:"sevClient,omitempty"`

	Name              string     `json:"name,omitempty"`
	PartNumber        string     `json:"partNumber,omitempty"`
	Text              string     `json:"text,omitempty"`
	Category          *Ref       `json:"category,omitempty"`
	Stock             Decimal    `json:"stock,omitempty"`
	StockEnabled      Bool       `json:"stockEnabled,omitempty"`
	Unity             *Ref       `json:"unity,omitempty"`
	Price             Decimal    `json:"price,omitempty"`
	PriceNet          Decimal    `json:"priceNet,omitempty"`
	PriceGross        Decimal    `json:"priceGross,omitempty"`
	PricePartner      Decimal    `json:"pricePartner,omitempty"`
	PriceCustomer     Decimal    `json:"priceCustomer,omitempty"`
	PricePurchase     Decimal    `json:"pricePurchase,omitempty"`
	SecondUnityFactor Decimal    `json:"secondUnityFactor,omitempty"`
	TaxRate           Decimal    `json:"taxRate,omitempty"`
	Image             string     `json:"image,omitempty"`
	Status            PartStatus `json:"status,omitempty"` // see [PartStatus]
	InternalComment   string     `json:"internalComment,omitempty"`
}

// PartStatus is the lifecycle state of an inventory part.
type PartStatus Num

func (s PartStatus) String() string { return Num(s).String() }

func (s *PartStatus) UnmarshalJSON(b []byte) error {
	var n Num
	if err := n.UnmarshalJSON(b); err != nil {
		return err
	}
	*s = PartStatus(n)
	return nil
}

// PartStatus values.
const (
	// PartStatusInactive is archived and not available for new positions.
	PartStatusInactive PartStatus = 50
	// PartStatusActive is in active use.
	PartStatusActive PartStatus = 100
)

// CreatePartParams is the body for [PartsService.Create].
type CreatePartParams struct {
	// Name shown in lists and on documents. Required.
	Name string `json:"name"`
	// PartNumber is your part identifier. Required.
	PartNumber string `json:"partNumber"`
	// Unity is the unit of measure. Required. Use [UnityRef].
	Unity *Ref `json:"unity"`
	// Stock is the initial on-hand quantity.
	Stock Decimal `json:"stock"`
	// TaxRate is the VAT rate as a percentage (e.g. 19). Required.
	TaxRate Decimal `json:"taxRate"`
	// Text is a long description.
	Text *string `json:"text,omitempty"`
	// Category groups the part in your inventory.
	Category *Ref `json:"category,omitempty"`
	// StockEnabled turns stock tracking on for this part.
	StockEnabled *Bool `json:"stockEnabled,omitempty"`
	// Price sets the sale price (typically net).
	Price *Decimal `json:"price,omitempty"`
	// PriceNet is the net sale price.
	PriceNet *Decimal `json:"priceNet,omitempty"`
	// PriceGross is the gross sale price.
	PriceGross *Decimal `json:"priceGross,omitempty"`
	// PricePurchase is the purchase price.
	PricePurchase *Decimal `json:"pricePurchase,omitempty"`
	// Status of the part. See [PartStatusActive] and adjacent.
	Status *PartStatus `json:"status,omitempty"`
	// InternalComment is a private note (not shown to customers).
	InternalComment *string `json:"internalComment,omitempty"`
}

// UpdatePartParams is the body for [PartsService.Update].
// See [CreatePartParams] for field semantics.
type UpdatePartParams struct {
	Name            *string     `json:"name,omitempty"`
	PartNumber      *string     `json:"partNumber,omitempty"`
	Text            *string     `json:"text,omitempty"`
	Category        *Ref        `json:"category,omitempty"`
	Stock           *Decimal    `json:"stock,omitempty"`
	StockEnabled    *Bool       `json:"stockEnabled,omitempty"`
	Unity           *Ref        `json:"unity,omitempty"`
	Price           *Decimal    `json:"price,omitempty"`
	PriceNet        *Decimal    `json:"priceNet,omitempty"`
	PriceGross      *Decimal    `json:"priceGross,omitempty"`
	PricePurchase   *Decimal    `json:"pricePurchase,omitempty"`
	TaxRate         *Decimal    `json:"taxRate,omitempty"`
	Status          *PartStatus `json:"status,omitempty"`
	InternalComment *string     `json:"internalComment,omitempty"`
}

// ListPartsParams filters [PartsService.List].
type ListPartsParams struct {
	// PartNumber matches the exact part number.
	PartNumber string
	// Name matches the exact name.
	Name string
}

func (p *ListPartsParams) query() url.Values {
	if p == nil {
		return nil
	}
	q := url.Values{}
	if p.PartNumber != "" {
		q.Set("partNumber", p.PartNumber)
	}
	if p.Name != "" {
		q.Set("name", p.Name)
	}
	return q
}

// List returns parts matching the given filter.
func (s *PartsService) List(ctx context.Context, opts *ListPartsParams) iter.Seq2[Part, error] {
	return listIter[Part](ctx, s.c, "/Part", opts.query())
}

// Get returns the part with the given id.
func (s *PartsService) Get(ctx context.Context, id ID) (*Part, error) {
	raw, err := s.c.do(ctx, http.MethodGet, fmt.Sprintf("/Part/%d", id), nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeObject[Part](raw)
}

// Create creates a new part.
func (s *PartsService) Create(ctx context.Context, params *CreatePartParams) (*Part, error) {
	raw, err := s.c.do(ctx, http.MethodPost, "/Part", nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[Part](raw)
}

// Update modifies an existing part.
func (s *PartsService) Update(ctx context.Context, id ID, params *UpdatePartParams) (*Part, error) {
	raw, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/Part/%d", id), nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[Part](raw)
}

// Stock returns the current stock count of the given part.
func (s *PartsService) Stock(ctx context.Context, id ID) (Decimal, error) {
	raw, err := s.c.do(ctx, http.MethodGet, fmt.Sprintf("/Part/%d/getStock", id), nil, nil)
	if err != nil {
		return 0, err
	}
	var d Decimal
	if err := d.UnmarshalJSON(raw); err != nil {
		return 0, err
	}
	return d, nil
}
