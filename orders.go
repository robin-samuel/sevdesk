package sevdesk

import (
	"context"
	"fmt"
	"iter"
	"net/http"
	"net/url"
	"time"
)

// OrdersService handles communication with /Order endpoints.
type OrdersService struct {
	c *Client
}

// Order is a sevdesk order (Auftrag / Angebot).
type Order struct {
	ID         ID         `json:"id,omitempty"`
	ObjectName string     `json:"objectName,omitempty"`
	Create     *time.Time `json:"create,omitempty"`
	Update     *time.Time `json:"update,omitempty"`
	SevClient  *Ref       `json:"sevClient,omitempty"`
	CreateUser *Ref       `json:"createUser,omitempty"`

	OrderNumber     string      `json:"orderNumber,omitempty"`
	Contact         *Ref        `json:"contact,omitempty"`
	ContactPerson   *Ref        `json:"contactPerson,omitempty"`
	OrderDate       *time.Time  `json:"orderDate,omitempty"`
	Status          OrderStatus `json:"status,omitempty"` // see [OrderStatus]
	Header          string      `json:"header,omitempty"`
	HeadText        string      `json:"headText,omitempty"`
	FootText        string      `json:"footText,omitempty"`
	AddressCountry  *Ref        `json:"addressCountry,omitempty"`
	DeliveryTerms   string      `json:"deliveryTerms,omitempty"`
	PaymentTerms    string      `json:"paymentTerms,omitempty"`
	Origin          *Ref        `json:"origin,omitempty"`
	Version         Num         `json:"version,omitempty"`
	SmallSettlement Bool        `json:"smallSettlement,omitempty"`
	TaxRate         Decimal     `json:"taxRate,omitempty"`
	TaxRule         *Ref        `json:"taxRule,omitempty"`
	TaxSet          *Ref        `json:"taxSet,omitempty"`
	TaxText         string      `json:"taxText,omitempty"`
	TaxType         string      `json:"taxType,omitempty"`
	OrderType       OrderType   `json:"orderType,omitempty"`
	SendDate        *time.Time  `json:"sendDate,omitempty"`
	Address         string      `json:"address,omitempty"`
	Currency        string      `json:"currency,omitempty"`

	SumNet                      Decimal `json:"sumNet,omitempty"`
	SumTax                      Decimal `json:"sumTax,omitempty"`
	SumGross                    Decimal `json:"sumGross,omitempty"`
	SumDiscounts                Decimal `json:"sumDiscounts,omitempty"`
	SumNetForeignCurrency       Decimal `json:"sumNetForeignCurrency,omitempty"`
	SumTaxForeignCurrency       Decimal `json:"sumTaxForeignCurrency,omitempty"`
	SumGrossForeignCurrency     Decimal `json:"sumGrossForeignCurrency,omitempty"`
	SumDiscountsForeignCurrency Decimal `json:"sumDiscountsForeignCurrency,omitempty"`

	CustomerInternalNote string   `json:"customerInternalNote,omitempty"`
	ShowNet              Bool     `json:"showNet,omitempty"`
	SendType             SendType `json:"sendType,omitempty"`
}

// OrderStatus is the lifecycle state of an order.
type OrderStatus Num

func (s OrderStatus) String() string { return Num(s).String() }

func (s *OrderStatus) UnmarshalJSON(b []byte) error {
	var n Num
	if err := n.UnmarshalJSON(b); err != nil {
		return err
	}
	*s = OrderStatus(n)
	return nil
}

// OrderStatus values.
const (
	// OrderStatusDraft is an unsent draft.
	OrderStatusDraft OrderStatus = 100
	// OrderStatusOpen has been sent to the customer and is awaiting a response.
	OrderStatusOpen OrderStatus = 200
	// OrderStatusRejected was declined by the customer.
	OrderStatusRejected OrderStatus = 300
	// OrderStatusAccepted was accepted by the customer.
	OrderStatusAccepted OrderStatus = 500
	// OrderStatusPartiallyDone is partly fulfilled (e.g. some positions delivered).
	OrderStatusPartiallyDone OrderStatus = 750
	// OrderStatusDone is fully fulfilled.
	OrderStatusDone OrderStatus = 1000
)

// OrderType is the kind of order document.
type OrderType string

// OrderType values.
const (
	// OrderTypeAN is a quote (Angebot).
	OrderTypeAN OrderType = "AN"
	// OrderTypeAB is an order confirmation (Auftragsbestätigung).
	OrderTypeAB OrderType = "AB"
	// OrderTypeLI is a packing list / delivery note (Lieferschein).
	OrderTypeLI OrderType = "LI"
	// OrderTypeECR is a contract note (Eigenbeleg).
	OrderTypeECR OrderType = "ECR"
)

// OrderPos is one line item on an order.
type OrderPos struct {
	ID         ID         `json:"id,omitempty"`
	ObjectName string     `json:"objectName,omitempty"`
	Create     *time.Time `json:"create,omitempty"`
	Update     *time.Time `json:"update,omitempty"`
	SevClient  *Ref       `json:"sevClient,omitempty"`

	Order          *Ref    `json:"order,omitempty"`
	Part           *Ref    `json:"part,omitempty"`
	Quantity       Decimal `json:"quantity,omitempty"`
	Price          Decimal `json:"price,omitempty"`
	PriceNet       Decimal `json:"priceNet,omitempty"`
	PriceTax       Decimal `json:"priceTax,omitempty"`
	PriceGross     Decimal `json:"priceGross,omitempty"`
	Name           string  `json:"name,omitempty"`
	Unity          *Ref    `json:"unity,omitempty"`
	PositionNumber Num     `json:"positionNumber,omitempty"`
	Text           string  `json:"text,omitempty"`
	Discount       Decimal `json:"discount,omitempty"`
	Optional       Bool    `json:"optional,omitempty"`
	TaxRate        Decimal `json:"taxRate,omitempty"`
	SumDiscount    Decimal `json:"sumDiscount,omitempty"`
}

// OrderPosCreate is a position to send when creating an order.
// ObjectName and MapAll are set automatically by [OrdersService.Create].
type OrderPosCreate struct {
	ObjectName string `json:"objectName"` // set by Create
	MapAll     bool   `json:"mapAll"`     // set by Create

	// Part links the position to an inventory item. Optional.
	Part *Ref `json:"part,omitempty"`
	// Quantity of the article/part. Required.
	Quantity Decimal `json:"quantity"`
	// Price is the unit price.
	Price Decimal `json:"price,omitempty"`
	// Name shown on the document. Required.
	Name string `json:"name"`
	// Unity is the unit of measure. Required. Use [UnityRef].
	Unity *Ref `json:"unity"`
	// Text is the long description.
	Text *string `json:"text,omitempty"`
	// Discount is a per-position discount.
	Discount *Decimal `json:"discount,omitempty"`
	// Optional marks the position as optional in a quote.
	Optional *Bool `json:"optional,omitempty"`
	// TaxRate is the VAT rate as a percentage. Required.
	TaxRate Decimal `json:"taxRate"`
}

// CreateOrderParams is the body for [OrdersService.Create].
type CreateOrderParams struct {
	// Order carries the order-level fields. Required.
	Order *OrderCreateFields `json:"order"`
	// Positions are the line items being added.
	Positions []OrderPosCreate `json:"orderPosSave,omitempty"`
	// PositionsDelete removes existing positions by reference.
	PositionsDelete []Ref `json:"orderPosDelete,omitempty"`
}

// OrderCreateFields is the order body for [CreateOrderParams].
// ObjectName and MapAll are set automatically by [OrdersService.Create].
type OrderCreateFields struct {
	ObjectName string `json:"objectName"` // set by Create
	MapAll     bool   `json:"mapAll"`     // set by Create

	// OrderNumber overrides the auto-assigned number.
	OrderNumber *string `json:"orderNumber,omitempty"`
	// Contact is the customer. Required.
	Contact *Ref `json:"contact"`
	// ContactPerson is the sevdesk user shown as the contact. Required.
	ContactPerson *Ref `json:"contactPerson"`
	// OrderDate is the document date. Required.
	OrderDate *time.Time `json:"orderDate"`
	// Status of the order. See [OrderStatusDraft] and adjacent.
	Status *OrderStatus `json:"status,omitempty"`
	// Header is the document title.
	Header *string `json:"header,omitempty"`
	// HeadText is the rich-text intro.
	HeadText *string `json:"headText,omitempty"`
	// FootText is the rich-text outro.
	FootText *string `json:"footText,omitempty"`
	// AddressCountry of the rendered address.
	AddressCountry *Ref `json:"addressCountry,omitempty"`
	// DeliveryTerms is the rendered delivery terms text.
	DeliveryTerms *string `json:"deliveryTerms,omitempty"`
	// PaymentTerms is the rendered payment terms text.
	PaymentTerms *string `json:"paymentTerms,omitempty"`
	// SmallSettlement marks the order as Kleinunternehmer (no VAT shown).
	SmallSettlement *Bool `json:"smallSettlement,omitempty"`
	// TaxRate is the default tax rate when positions don't override it.
	TaxRate *Decimal `json:"taxRate,omitempty"`
	// TaxRule selects the tax rule. Required.
	TaxRule *Ref `json:"taxRule"`
	// TaxText is the line printed below the totals.
	TaxText *string `json:"taxText,omitempty"`
	// TaxType picks the tax mode.
	TaxType *string `json:"taxType,omitempty"`
	// OrderType picks the kind. Required. See [OrderType].
	OrderType OrderType `json:"orderType"`
	// Currency is the ISO 4217 code. Required.
	Currency string `json:"currency"`
	// Address is the rendered address (multi-line string).
	Address *string `json:"address,omitempty"`
	// CustomerInternalNote is private, not shown to the customer.
	CustomerInternalNote *string `json:"customerInternalNote,omitempty"`
	// ShowNet shows net prices on the rendered PDF.
	ShowNet *Bool `json:"showNet,omitempty"`
	// SendType records how the order was sent.
	SendType *SendType `json:"sendType,omitempty"`
}

// UpdateOrderParams is the body for [OrdersService.Update]. Only works for
// drafts. See [OrderCreateFields] for field semantics.
type UpdateOrderParams struct {
	OrderNumber          *string      `json:"orderNumber,omitempty"`
	Contact              *Ref         `json:"contact,omitempty"`
	ContactPerson        *Ref         `json:"contactPerson,omitempty"`
	OrderDate            *time.Time   `json:"orderDate,omitempty"`
	Status               *OrderStatus `json:"status,omitempty"`
	Header               *string      `json:"header,omitempty"`
	HeadText             *string      `json:"headText,omitempty"`
	FootText             *string      `json:"footText,omitempty"`
	AddressCountry       *Ref         `json:"addressCountry,omitempty"`
	DeliveryTerms        *string      `json:"deliveryTerms,omitempty"`
	PaymentTerms         *string      `json:"paymentTerms,omitempty"`
	TaxRate              *Decimal     `json:"taxRate,omitempty"`
	TaxRule              *Ref         `json:"taxRule,omitempty"`
	TaxText              *string      `json:"taxText,omitempty"`
	TaxType              *string      `json:"taxType,omitempty"`
	OrderType            *OrderType   `json:"orderType,omitempty"`
	Currency             *string      `json:"currency,omitempty"`
	Address              *string      `json:"address,omitempty"`
	CustomerInternalNote *string      `json:"customerInternalNote,omitempty"`
	ShowNet              *Bool        `json:"showNet,omitempty"`
	SendType             *SendType    `json:"sendType,omitempty"`
}

// ListOrdersParams filters [OrdersService.List].
type ListOrdersParams struct {
	// Status filters to one status; 0 means any. See [OrderStatus].
	Status OrderStatus
	// OrderNumber matches the exact number.
	OrderNumber string
	// StartDate excludes orders dated before this point.
	StartDate time.Time
	// EndDate excludes orders dated after this point.
	EndDate time.Time
	// Contact narrows to orders for one customer.
	Contact *Ref
}

func (p *ListOrdersParams) query() url.Values {
	if p == nil {
		return nil
	}
	q := url.Values{}
	if p.Status != 0 {
		q.Set("status", p.Status.String())
	}
	if p.OrderNumber != "" {
		q.Set("orderNumber", p.OrderNumber)
	}
	if !p.StartDate.IsZero() {
		q.Set("startDate", fmt.Sprintf("%d", p.StartDate.Unix()))
	}
	if !p.EndDate.IsZero() {
		q.Set("endDate", fmt.Sprintf("%d", p.EndDate.Unix()))
	}
	if p.Contact != nil {
		q.Set("contact[id]", p.Contact.ID.String())
		q.Set("contact[objectName]", p.Contact.ObjectName)
	}
	return q
}

// List returns orders matching the filter.
func (s *OrdersService) List(ctx context.Context, opts *ListOrdersParams) iter.Seq2[Order, error] {
	return listIter[Order](ctx, s.c, "/Order", opts.query())
}

// Get returns the order with the given id.
func (s *OrdersService) Get(ctx context.Context, id ID) (*Order, error) {
	raw, err := s.c.do(ctx, http.MethodGet, fmt.Sprintf("/Order/%d", id), nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeObject[Order](raw)
}

// CreateOrderResult is the return value of [OrdersService.Create].
type CreateOrderResult struct {
	Order     Order      `json:"order"`
	Positions []OrderPos `json:"orderPos"`
}

// Create creates an order with its positions.
func (s *OrdersService) Create(ctx context.Context, params *CreateOrderParams) (*CreateOrderResult, error) {
	if params != nil && params.Order != nil {
		if params.Order.ObjectName == "" {
			params.Order.ObjectName = ObjectOrder
		}
		params.Order.MapAll = true
	}
	for i := range params.Positions {
		if params.Positions[i].ObjectName == "" {
			params.Positions[i].ObjectName = ObjectOrderPos
		}
		params.Positions[i].MapAll = true
	}
	raw, err := s.c.do(ctx, http.MethodPost, "/Order/Factory/saveOrder", nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[CreateOrderResult](raw)
}

// Update modifies an existing order (only drafts).
func (s *OrdersService) Update(ctx context.Context, id ID, params *UpdateOrderParams) (*Order, error) {
	raw, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/Order/%d", id), nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[Order](raw)
}

// Delete removes an order.
func (s *OrdersService) Delete(ctx context.Context, id ID) error {
	_, err := s.c.do(ctx, http.MethodDelete, fmt.Sprintf("/Order/%d", id), nil, nil)
	return err
}

// GetPositions returns the positions of the given order.
func (s *OrdersService) GetPositions(ctx context.Context, id ID, opts *GetPositionsOptions) iter.Seq2[OrderPos, error] {
	return listIter[OrderPos](ctx, s.c, fmt.Sprintf("/Order/%d/getPositions", id), opts.query())
}

// GetDiscounts returns the discounts attached to the given order.
func (s *OrdersService) GetDiscounts(ctx context.Context, id ID, opts *GetPositionsOptions) iter.Seq2[Discount, error] {
	return listIter[Discount](ctx, s.c, fmt.Sprintf("/Order/%d/getDiscounts", id), opts.query())
}

// RelatedObjectsOptions configures [OrdersService.GetRelatedObjects].
type RelatedObjectsOptions struct {
	// IncludeItself, when true, includes the source order in the result.
	IncludeItself bool
	// SortByType groups results by object kind in the response.
	SortByType bool
	// Embed asks sevdesk to inline referenced objects.
	Embed []string
}

func (o *RelatedObjectsOptions) query() url.Values {
	if o == nil {
		return nil
	}
	q := url.Values{}
	if o.IncludeItself {
		q.Set("includeItself", "true")
	}
	if o.SortByType {
		q.Set("sortByType", "true")
	}
	for _, e := range o.Embed {
		q.Add("embed[]", e)
	}
	return q
}

// GetRelatedObjects returns objects (invoices, packing lists, etc.) derived
// from the given order.
func (s *OrdersService) GetRelatedObjects(ctx context.Context, id ID, opts *RelatedObjectsOptions) iter.Seq2[Ref, error] {
	return listIter[Ref](ctx, s.c, fmt.Sprintf("/Order/%d/getRelatedObjects", id), opts.query())
}

// SendViaEmail sends an order by email.
func (s *OrdersService) SendViaEmail(ctx context.Context, id ID, params *SendEmailParams) (*Email, error) {
	raw, err := s.c.do(ctx, http.MethodPost, fmt.Sprintf("/Order/%d/sendViaEmail", id), nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[Email](raw)
}

// CreatePackingList creates a packing list (Lieferschein) from an order.
func (s *OrdersService) CreatePackingList(ctx context.Context, orderID ID, params *CreateOrderParams) (*Order, error) {
	q := url.Values{
		"order[id]":         {orderID.String()},
		"order[objectName]": {ObjectOrder},
	}
	raw, err := s.c.do(ctx, http.MethodPost, "/Order/Factory/createPackingListFromOrder", q, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[Order](raw)
}

// CreateContractNote creates a contract note from an order.
func (s *OrdersService) CreateContractNote(ctx context.Context, orderID ID, params *CreateOrderParams) (*Order, error) {
	q := url.Values{
		"order[id]":         {orderID.String()},
		"order[objectName]": {ObjectOrder},
	}
	raw, err := s.c.do(ctx, http.MethodPost, "/Order/Factory/createContractNoteFromOrder", q, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[Order](raw)
}

// GetPdf returns the rendered PDF of an order.
func (s *OrdersService) GetPdf(ctx context.Context, id ID, download, preventSendBy bool) (*PDFDocument, error) {
	q := url.Values{}
	if download {
		q.Set("download", "true")
	}
	if preventSendBy {
		q.Set("preventSendBy", "true")
	}
	raw, err := s.c.do(ctx, http.MethodGet, fmt.Sprintf("/Order/%d/getPdf", id), q, nil)
	if err != nil {
		return nil, err
	}
	return decodeObject[PDFDocument](raw)
}

// MarkSent records that the order has been sent by the given channel.
func (s *OrdersService) MarkSent(ctx context.Context, id ID, params *MarkSentParams) (*Order, error) {
	raw, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/Order/%d/sendBy", id), nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[Order](raw)
}

// ChangeLayout changes the rendering layout of an order.
func (s *OrdersService) ChangeLayout(ctx context.Context, id ID, params *ChangeLayoutParams) (*RenderResult, error) {
	raw, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/Order/%d/changeParameter", id), nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[RenderResult](raw)
}

// OrderPositionsService handles /OrderPos endpoints.
type OrderPositionsService struct {
	c *Client
}

// List returns order positions filtered by the given order.
func (s *OrderPositionsService) List(ctx context.Context, order *Ref) iter.Seq2[OrderPos, error] {
	var q url.Values
	if order != nil {
		q = url.Values{
			"order[id]":         {order.ID.String()},
			"order[objectName]": {order.ObjectName},
		}
	}
	return listIter[OrderPos](ctx, s.c, "/OrderPos", q)
}

// Get returns the order position with the given id.
func (s *OrderPositionsService) Get(ctx context.Context, id ID) (*OrderPos, error) {
	raw, err := s.c.do(ctx, http.MethodGet, fmt.Sprintf("/OrderPos/%d", id), nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeObject[OrderPos](raw)
}

// UpdateOrderPosParams is the body for [OrderPositionsService.Update].
// See [OrderPosCreate] for field semantics.
type UpdateOrderPosParams struct {
	Part           *Ref     `json:"part,omitempty"`
	Quantity       *Decimal `json:"quantity,omitempty"`
	Price          *Decimal `json:"price,omitempty"`
	Name           *string  `json:"name,omitempty"`
	Unity          *Ref     `json:"unity,omitempty"`
	PositionNumber *int     `json:"positionNumber,omitempty"`
	Text           *string  `json:"text,omitempty"`
	Discount       *Decimal `json:"discount,omitempty"`
	Optional       *Bool    `json:"optional,omitempty"`
	TaxRate        *Decimal `json:"taxRate,omitempty"`
}

// Update modifies an order position.
func (s *OrderPositionsService) Update(ctx context.Context, id ID, params *UpdateOrderPosParams) (*OrderPos, error) {
	raw, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/OrderPos/%d", id), nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[OrderPos](raw)
}

// Delete removes an order position.
func (s *OrderPositionsService) Delete(ctx context.Context, id ID) error {
	_, err := s.c.do(ctx, http.MethodDelete, fmt.Sprintf("/OrderPos/%d", id), nil, nil)
	return err
}

// Discount is a price reduction (or surcharge) on a document.
type Discount struct {
	ID         ID         `json:"id,omitempty"`
	ObjectName string     `json:"objectName,omitempty"`
	Create     *time.Time `json:"create,omitempty"`
	Update     *time.Time `json:"update,omitempty"`
	SevClient  *Ref       `json:"sevClient,omitempty"`
	Object     *Ref       `json:"object,omitempty"`
	Discount   Bool       `json:"discount,omitempty"` // true = discount, false = surcharge
	Text       string     `json:"text,omitempty"`
	Percentage Bool       `json:"percentage,omitempty"` // true = percent, false = absolute
	Value      Decimal    `json:"value,omitempty"`
	IsNet      Bool       `json:"isNet,omitempty"`
}
