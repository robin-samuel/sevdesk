package sevdesk

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// InvoicesService handles communication with /Invoice endpoints.
type InvoicesService struct {
	c *Client
}

// Invoice is a sevdesk invoice.
type Invoice struct {
	ID         ID         `json:"id,omitempty"`
	ObjectName string     `json:"objectName,omitempty"`
	Create     *time.Time `json:"create,omitempty"`
	Update     *time.Time `json:"update,omitempty"`
	SevClient  *Ref       `json:"sevClient,omitempty"`
	CreateUser *Ref       `json:"createUser,omitempty"`

	InvoiceNumber string     `json:"invoiceNumber,omitempty"`
	Contact       *Ref       `json:"contact,omitempty"`
	ContactPerson *Ref       `json:"contactPerson,omitempty"`
	InvoiceDate   *time.Time `json:"invoiceDate,omitempty"`
	Header        string     `json:"header,omitempty"`
	HeadText      string     `json:"headText,omitempty"`
	FootText      string     `json:"footText,omitempty"`
	TimeToPay     Num        `json:"timeToPay,omitempty"`
	DiscountTime  Num        `json:"discountTime,omitempty"`
	Discount      Decimal    `json:"discount,omitempty"`

	Address                string `json:"address,omitempty"`
	AddressName            string `json:"addressName,omitempty"`
	AddressName2           string `json:"addressName2,omitempty"`
	AddressStreet          string `json:"addressStreet,omitempty"`
	AddressZip             string `json:"addressZip,omitempty"`
	AddressCity            string `json:"addressCity,omitempty"`
	AddressCountry         *Ref   `json:"addressCountry,omitempty"`
	AddressGender          string `json:"addressGender,omitempty"`
	AddressParentName      string `json:"addressParentName,omitempty"`
	AddressParentName2     string `json:"addressParentName2,omitempty"`
	DeliveryAddressCountry *Ref   `json:"deliveryAddressCountry,omitempty"`

	PayDate           *time.Time    `json:"payDate,omitempty"`
	DeliveryDate      *time.Time    `json:"deliveryDate,omitempty"`
	DeliveryDateUntil *time.Time    `json:"deliveryDateUntil,omitempty"`
	Status            InvoiceStatus `json:"status,omitempty"` // see [InvoiceStatus]
	SmallSettlement   Bool          `json:"smallSettlement,omitempty"`
	TaxRate           Decimal       `json:"taxRate,omitempty"`
	TaxText           string        `json:"taxText,omitempty"`
	TaxType           string        `json:"taxType,omitempty"`
	TaxRule           *Ref          `json:"taxRule,omitempty"`
	TaxSet            *Ref          `json:"taxSet,omitempty"`
	DunningLevel      Num           `json:"dunningLevel,omitempty"`
	SendDate          *time.Time    `json:"sendDate,omitempty"`
	SendType          SendType      `json:"sendType,omitempty"`
	InvoiceType       InvoiceType   `json:"invoiceType,omitempty"`
	PaymentMethod     *Ref          `json:"paymentMethod,omitempty"`
	CostCentre        *Ref          `json:"costCentre,omitempty"`
	Origin            *Ref          `json:"origin,omitempty"`

	AccountIntervall   string     `json:"accountIntervall,omitempty"`
	AccountNextInvoice *time.Time `json:"accountNextInvoice,omitempty"`
	AccountLastInvoice *time.Time `json:"accountLastInvoice,omitempty"`
	AccountStartDate   *time.Time `json:"accountStartDate,omitempty"`
	AccountEndDate     *time.Time `json:"accountEndDate,omitempty"`

	ReminderTotal    Decimal    `json:"reminderTotal,omitempty"`
	ReminderDebit    Decimal    `json:"reminderDebit,omitempty"`
	ReminderDeadline *time.Time `json:"reminderDeadline,omitempty"`
	ReminderCharge   Decimal    `json:"reminderCharge,omitempty"`

	Currency string `json:"currency,omitempty"`

	SumNet                          Decimal `json:"sumNet,omitempty"`
	SumTax                          Decimal `json:"sumTax,omitempty"`
	SumGross                        Decimal `json:"sumGross,omitempty"`
	SumDiscounts                    Decimal `json:"sumDiscounts,omitempty"`
	SumDiscountNet                  Decimal `json:"sumDiscountNet,omitempty"`
	SumDiscountGross                Decimal `json:"sumDiscountGross,omitempty"`
	SumNetForeignCurrency           Decimal `json:"sumNetForeignCurrency,omitempty"`
	SumTaxForeignCurrency           Decimal `json:"sumTaxForeignCurrency,omitempty"`
	SumGrossForeignCurrency         Decimal `json:"sumGrossForeignCurrency,omitempty"`
	SumDiscountsForeignCurrency     Decimal `json:"sumDiscountsForeignCurrency,omitempty"`
	SumDiscountNetForeignCurrency   Decimal `json:"sumDiscountNetForeignCurrency,omitempty"`
	SumDiscountGrossForeignCurrency Decimal `json:"sumDiscountGrossForeignCurrency,omitempty"`
	SumNetAccounting                Decimal `json:"sumNetAccounting,omitempty"`
	SumTaxAccounting                Decimal `json:"sumTaxAccounting,omitempty"`
	SumGrossAccounting              Decimal `json:"sumGrossAccounting,omitempty"`
	PaidAmount                      Decimal `json:"paidAmount,omitempty"`

	CustomerInternalNote                string     `json:"customerInternalNote,omitempty"`
	EInvoiceReference                   string     `json:"einvoiceReference,omitempty"`
	ShowNet                             Bool       `json:"showNet,omitempty"`
	Enshrined                           *time.Time `json:"enshrined,omitempty"`
	SendPaymentReceivedNotificationDate *time.Time `json:"sendPaymentReceivedNotificationDate,omitempty"`
}

// InvoiceStatus is the lifecycle state of an invoice.
type InvoiceStatus Num

func (s InvoiceStatus) String() string { return Num(s).String() }

func (s *InvoiceStatus) UnmarshalJSON(b []byte) error {
	var n Num
	if err := n.UnmarshalJSON(b); err != nil {
		return err
	}
	*s = InvoiceStatus(n)
	return nil
}

// InvoiceStatus values.
const (
	// InvoiceStatusDraft is an unsent draft.
	InvoiceStatusDraft InvoiceStatus = 100
	// InvoiceStatusOpen has been sent and is awaiting payment.
	InvoiceStatusOpen InvoiceStatus = 200
	// InvoiceStatusPartiallyPaid means some money has been received.
	InvoiceStatusPartiallyPaid InvoiceStatus = 750
	// InvoiceStatusPaid is fully paid.
	InvoiceStatusPaid InvoiceStatus = 1000
)

// InvoiceType picks the document kind.
type InvoiceType string

// InvoiceType values.
const (
	// InvoiceTypeRegular is a standard invoice (Rechnung).
	InvoiceTypeRegular InvoiceType = "RE"
	// InvoiceTypeRecurring is a recurring-invoice template (Wiederkehrende Rechnung).
	InvoiceTypeRecurring InvoiceType = "WKR"
	// InvoiceTypeCancellation reverses another invoice (Stornorechnung).
	InvoiceTypeCancellation InvoiceType = "SR"
	// InvoiceTypeReminder is a dunning notice (Mahnung).
	InvoiceTypeReminder InvoiceType = "MA"
	// InvoiceTypeECommerce is an e-commerce invoice.
	InvoiceTypeECommerce InvoiceType = "ER"
	// InvoiceTypePartial bills a partial amount against an order (Teilrechnung).
	InvoiceTypePartial InvoiceType = "TR"
	// InvoiceTypeFinal closes out a series of partial invoices (Abschlussrechnung).
	InvoiceTypeFinal InvoiceType = "AR"
)

// InvoicePos is one line item on an invoice.
type InvoicePos struct {
	ID         ID         `json:"id,omitempty"`
	ObjectName string     `json:"objectName,omitempty"`
	Create     *time.Time `json:"create,omitempty"`
	Update     *time.Time `json:"update,omitempty"`
	SevClient  *Ref       `json:"sevClient,omitempty"`

	Invoice            *Ref    `json:"invoice,omitempty"`
	Part               *Ref    `json:"part,omitempty"`
	Quantity           Decimal `json:"quantity,omitempty"`
	Price              Decimal `json:"price,omitempty"`
	PriceNet           Decimal `json:"priceNet,omitempty"`
	PriceTax           Decimal `json:"priceTax,omitempty"`
	PriceGross         Decimal `json:"priceGross,omitempty"`
	Name               string  `json:"name,omitempty"`
	Unity              *Ref    `json:"unity,omitempty"`
	PositionNumber     Num     `json:"positionNumber,omitempty"`
	Text               string  `json:"text,omitempty"`
	Discount           Decimal `json:"discount,omitempty"`
	TaxRate            Decimal `json:"taxRate,omitempty"`
	SumDiscount        Decimal `json:"sumDiscount,omitempty"`
	SumNetAccounting   Decimal `json:"sumNetAccounting,omitempty"`
	SumTaxAccounting   Decimal `json:"sumTaxAccounting,omitempty"`
	SumGrossAccounting Decimal `json:"sumGrossAccounting,omitempty"`
}

// InvoicePosCreate is a position to send when creating an invoice.
// ObjectName and MapAll are set automatically by [InvoicesService.Create].
type InvoicePosCreate struct {
	ObjectName string `json:"objectName"` // set by Create
	MapAll     bool   `json:"mapAll"`     // set by Create

	// Part links the position to an inventory item. Optional — free-form
	// positions just set Name without a Part.
	Part *Ref `json:"part,omitempty"`
	// Quantity of the article/part. Required.
	Quantity Decimal `json:"quantity"`
	// Price is the unit price. Set either Price (net) or both PriceNet+PriceGross.
	Price Decimal `json:"price,omitempty"`
	// PriceNet is the unit price without tax.
	PriceNet *Decimal `json:"priceNet,omitempty"`
	// PriceGross is the unit price including tax.
	PriceGross *Decimal `json:"priceGross,omitempty"`
	// Name shown on the document. Required.
	Name string `json:"name"`
	// Unity is the unit of measure (each, kg, hour, etc.). Required.
	// Use [UnityRef] to build it.
	Unity *Ref `json:"unity"`
	// Text is the position's long description (rich text allowed).
	Text *string `json:"text,omitempty"`
	// Discount is a per-position discount, as percentage when InvoicePosCreate
	// is used inline (e.g. 10 for 10%).
	Discount *Decimal `json:"discount,omitempty"`
	// TaxRate is the VAT rate as a percentage (e.g. 19). Required.
	TaxRate Decimal `json:"taxRate"`
}

// InvoiceDiscountCreate is a discount/surcharge to attach to an invoice
// (also used for credit notes). ObjectName and MapAll are set automatically.
type InvoiceDiscountCreate struct {
	ObjectName string `json:"objectName"` // set by Create
	MapAll     bool   `json:"mapAll"`     // set by Create

	// Discount: true = discount, false = surcharge.
	Discount bool `json:"discount"`
	// Text shown next to the discount on the rendered document.
	Text string `json:"text"`
	// Percentage: true if Value is a percent, false if it's an absolute amount.
	Percentage bool `json:"percentage"`
	// Value of the discount/surcharge.
	Value Decimal `json:"value"`
}

// CreateInvoiceParams is the body for [InvoicesService.Create].
type CreateInvoiceParams struct {
	// Invoice carries the invoice-level fields. Required.
	Invoice *InvoiceCreateFields `json:"invoice"`
	// Positions are the line items being added.
	Positions []InvoicePosCreate `json:"invoicePosSave,omitempty"`
	// PositionsDelete removes existing positions by reference.
	PositionsDelete []Ref `json:"invoicePosDelete,omitempty"`
	// Discounts are document-level discounts or surcharges.
	Discounts []InvoiceDiscountCreate `json:"discountSave,omitempty"`
	// DiscountsDelete removes existing discounts by reference.
	DiscountsDelete []Ref `json:"discountDelete,omitempty"`
	// Filename of a file previously uploaded — attaches it to the invoice.
	Filename string `json:"filename,omitempty"`
}

// InvoiceCreateFields is the invoice body for [CreateInvoiceParams].
// ObjectName and MapAll are set automatically by [InvoicesService.Create].
type InvoiceCreateFields struct {
	ObjectName string `json:"objectName"` // set by Create
	MapAll     bool   `json:"mapAll"`     // set by Create

	// InvoiceNumber overrides the auto-assigned number.
	InvoiceNumber *string `json:"invoiceNumber,omitempty"`
	// Contact is the customer being invoiced. Required.
	Contact *Ref `json:"contact"`
	// ContactPerson is the sevdesk user shown as the contact on the invoice. Required.
	ContactPerson *Ref `json:"contactPerson"`
	// InvoiceDate is the document date. Required.
	InvoiceDate *time.Time `json:"invoiceDate"`
	// Header is the invoice title (often "Invoice <Number>").
	Header *string `json:"header,omitempty"`
	// HeadText is the rich-text intro shown above the positions table.
	HeadText *string `json:"headText,omitempty"`
	// FootText is the rich-text outro shown below the totals.
	FootText *string `json:"footText,omitempty"`
	// TimeToPay is the payment deadline in days from InvoiceDate.
	TimeToPay *int `json:"timeToPay,omitempty"`
	// Discount is an old-style flat discount percentage (rarely used —
	// prefer Discounts on [CreateInvoiceParams]).
	Discount *int `json:"discount,omitempty"`
	// Address is the rendered billing address (multi-line string).
	Address *string `json:"address,omitempty"`
	// AddressCountry of the billing address.
	AddressCountry *Ref `json:"addressCountry,omitempty"`
	// DeliveryDate is when the goods/services were delivered.
	DeliveryDate *time.Time `json:"deliveryDate,omitempty"`
	// Status of the invoice. See [InvoiceStatusDraft] and adjacent.
	Status *InvoiceStatus `json:"status,omitempty"`
	// SmallSettlement marks the invoice as Kleinunternehmer (no VAT shown).
	SmallSettlement *Bool `json:"smallSettlement,omitempty"`
	// TaxRate is the default tax rate when positions don't override it.
	TaxRate *Decimal `json:"taxRate,omitempty"`
	// TaxRule selects the tax rule (domestic, reverse charge, etc.). Required.
	TaxRule *Ref `json:"taxRule"`
	// TaxText is the line printed below the totals (e.g. "VAT 19%").
	TaxText *string `json:"taxText,omitempty"`
	// TaxType picks the tax mode. Common values: "default", "eu", "noteu",
	// "custom", "ss" (Kleinunternehmer).
	TaxType *string `json:"taxType,omitempty"`
	// InvoiceType picks the document kind. Required. See [InvoiceType].
	InvoiceType InvoiceType `json:"invoiceType"`
	// Currency is the ISO 4217 code (e.g. "EUR"). Required.
	Currency string `json:"currency"`
	// PaymentMethod links a default payment method.
	PaymentMethod *Ref `json:"paymentMethod,omitempty"`
	// CostCentre assigns the invoice to a cost centre.
	CostCentre *Ref `json:"costCentre,omitempty"`
	// ShowNet shows net prices on the rendered PDF.
	ShowNet *Bool `json:"showNet,omitempty"`
	// SendDate marks when the invoice was sent.
	SendDate *time.Time `json:"sendDate,omitempty"`
	// SendType records how the invoice was sent. See [SendType].
	SendType *SendType `json:"sendType,omitempty"`
	// CustomerInternalNote is private, not shown to the customer.
	CustomerInternalNote *string `json:"customerInternalNote,omitempty"`
	// PropertyIsEInvoice flags this as a structured e-invoice (XRechnung / ZUGFeRD).
	PropertyIsEInvoice *Bool `json:"propertyIsEInvoice,omitempty"`
	// Origin references the source object (e.g. an Order this invoice came from).
	Origin *Ref `json:"origin,omitempty"`
}

// UpdateInvoiceParams is the body for [InvoicesService.Update]. Only works for
// drafts. See [InvoiceCreateFields] for field semantics.
type UpdateInvoiceParams struct {
	InvoiceNumber        *string        `json:"invoiceNumber,omitempty"`
	Contact              *Ref           `json:"contact,omitempty"`
	ContactPerson        *Ref           `json:"contactPerson,omitempty"`
	InvoiceDate          *time.Time     `json:"invoiceDate,omitempty"`
	Header               *string        `json:"header,omitempty"`
	HeadText             *string        `json:"headText,omitempty"`
	FootText             *string        `json:"footText,omitempty"`
	TimeToPay            *int           `json:"timeToPay,omitempty"`
	Discount             *int           `json:"discount,omitempty"`
	Address              *string        `json:"address,omitempty"`
	AddressCountry       *Ref           `json:"addressCountry,omitempty"`
	DeliveryDate         *time.Time     `json:"deliveryDate,omitempty"`
	Status               *InvoiceStatus `json:"status,omitempty"`
	SmallSettlement      *Bool          `json:"smallSettlement,omitempty"`
	TaxRate              *Decimal       `json:"taxRate,omitempty"`
	TaxRule              *Ref           `json:"taxRule,omitempty"`
	TaxText              *string        `json:"taxText,omitempty"`
	TaxType              *string        `json:"taxType,omitempty"`
	Currency             *string        `json:"currency,omitempty"`
	PaymentMethod        *Ref           `json:"paymentMethod,omitempty"`
	CostCentre           *Ref           `json:"costCentre,omitempty"`
	ShowNet              *Bool          `json:"showNet,omitempty"`
	SendDate             *time.Time     `json:"sendDate,omitempty"`
	SendType             *SendType      `json:"sendType,omitempty"`
	CustomerInternalNote *string        `json:"customerInternalNote,omitempty"`
}

// ListInvoicesParams filters [InvoicesService.List].
type ListInvoicesParams struct {
	// Status filters to one status; 0 (the default) means any.
	// See [InvoiceStatus].
	Status InvoiceStatus
	// InvoiceNumber matches the exact invoice number.
	InvoiceNumber string
	// StartDate excludes invoices dated before this point.
	StartDate time.Time
	// EndDate excludes invoices dated after this point.
	EndDate time.Time
	// Contact narrows to invoices for one customer.
	Contact *Ref
}

func (p *ListInvoicesParams) query() url.Values {
	if p == nil {
		return nil
	}
	q := url.Values{}
	if p.Status != 0 {
		q.Set("status", p.Status.String())
	}
	if p.InvoiceNumber != "" {
		q.Set("invoiceNumber", p.InvoiceNumber)
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

// List returns invoices matching the given filter.
func (s *InvoicesService) List(ctx context.Context, opts *ListInvoicesParams) ([]Invoice, error) {
	raw, err := s.c.do(ctx, http.MethodGet, "/Invoice", opts.query(), nil)
	if err != nil {
		return nil, err
	}
	return decodeList[Invoice](raw)
}

// Get returns the invoice with the given id.
func (s *InvoicesService) Get(ctx context.Context, id ID) (*Invoice, error) {
	raw, err := s.c.do(ctx, http.MethodGet, fmt.Sprintf("/Invoice/%d", id), nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeObject[Invoice](raw)
}

// GetPositionsOptions configures position-listing methods like
// [InvoicesService.GetPositions] and [OrdersService.GetPositions].
type GetPositionsOptions struct {
	// Limit caps the number of positions returned.
	Limit int
	// Offset skips this many positions before returning results.
	Offset int
	// Embed asks sevdesk to inline referenced objects (e.g. "part", "unity").
	Embed []string
}

func (o *GetPositionsOptions) query() url.Values {
	if o == nil {
		return nil
	}
	q := url.Values{}
	if o.Limit > 0 {
		q.Set("limit", fmt.Sprintf("%d", o.Limit))
	}
	if o.Offset > 0 {
		q.Set("offset", fmt.Sprintf("%d", o.Offset))
	}
	for _, e := range o.Embed {
		q.Add("embed[]", e)
	}
	return q
}

// GetPositions returns the positions of the given invoice.
func (s *InvoicesService) GetPositions(ctx context.Context, id ID, opts *GetPositionsOptions) ([]InvoicePos, error) {
	raw, err := s.c.do(ctx, http.MethodGet, fmt.Sprintf("/Invoice/%d/getPositions", id), opts.query(), nil)
	if err != nil {
		return nil, err
	}
	return decodeList[InvoicePos](raw)
}

// CreateInvoiceResult is the return value of [InvoicesService.Create].
type CreateInvoiceResult struct {
	Invoice   Invoice      `json:"invoice"`
	Positions []InvoicePos `json:"invoicePos"`
	Filename  string       `json:"filename,omitempty"`
}

// Create creates an invoice with its positions and discounts in one atomic call.
func (s *InvoicesService) Create(ctx context.Context, params *CreateInvoiceParams) (*CreateInvoiceResult, error) {
	if params != nil && params.Invoice != nil {
		if params.Invoice.ObjectName == "" {
			params.Invoice.ObjectName = ObjectInvoice
		}
		params.Invoice.MapAll = true
	}
	for i := range params.Positions {
		if params.Positions[i].ObjectName == "" {
			params.Positions[i].ObjectName = ObjectInvoicePos
		}
		params.Positions[i].MapAll = true
	}
	for i := range params.Discounts {
		if params.Discounts[i].ObjectName == "" {
			params.Discounts[i].ObjectName = ObjectDiscounts
		}
		params.Discounts[i].MapAll = true
	}
	raw, err := s.c.do(ctx, http.MethodPost, "/Invoice/Factory/saveInvoice", nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[CreateInvoiceResult](raw)
}

// Update modifies an existing invoice. Only works for drafts.
func (s *InvoicesService) Update(ctx context.Context, id ID, params *UpdateInvoiceParams) (*Invoice, error) {
	raw, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/Invoice/%d", id), nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[Invoice](raw)
}

// CreateFromOrderParams selects an order to derive the invoice from.
type CreateFromOrderParams struct {
	// Order is the source order. Required.
	Order *Ref `json:"order"`
	// Type, when set, makes this a partial or final invoice for the order.
	// See [PartialInvoiceType]. Leave empty for a regular invoice.
	Type PartialInvoiceType `json:"type,omitempty"`
	// Amount is the partial amount (used with Type and PartialType).
	Amount Decimal `json:"amount,omitempty"`
	// PartialType controls whether Amount is a percentage or absolute value.
	// See [PartialType].
	PartialType PartialType `json:"partialType,omitempty"`
}

// PartialInvoiceType picks the kind of partial invoice when deriving one
// from an order via [InvoicesService.CreateFromOrder].
type PartialInvoiceType string

// PartialInvoiceType values.
const (
	// PartialInvoiceTypePartial makes a partial invoice (Teilrechnung).
	PartialInvoiceTypePartial PartialInvoiceType = "TR"
	// PartialInvoiceTypeFinal makes a closing invoice (Abschlussrechnung).
	PartialInvoiceTypeFinal PartialInvoiceType = "AR"
)

// PartialType controls how [CreateFromOrderParams.Amount] is interpreted.
type PartialType string

// PartialType values.
const (
	// PartialTypePercentage means [CreateFromOrderParams.Amount] is a percent
	// of the order total.
	PartialTypePercentage PartialType = "percentage"
	// PartialTypeAmount means [CreateFromOrderParams.Amount] is an absolute value.
	PartialTypeAmount PartialType = "amount"
)

// CreateFromOrder creates an invoice from an existing order.
func (s *InvoicesService) CreateFromOrder(ctx context.Context, params *CreateFromOrderParams) (*Invoice, error) {
	raw, err := s.c.do(ctx, http.MethodPost, "/Invoice/Factory/createInvoiceFromOrder", nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[Invoice](raw)
}

// CreateReminder creates a reminder (Mahnung) for an unpaid invoice.
func (s *InvoicesService) CreateReminder(ctx context.Context, invoiceID ID) (*Invoice, error) {
	q := url.Values{
		"invoice[id]":         {invoiceID.String()},
		"invoice[objectName]": {ObjectInvoice},
	}
	body := map[string]any{
		"invoice": map[string]any{
			"id":         invoiceID,
			"objectName": ObjectInvoice,
		},
	}
	raw, err := s.c.do(ctx, http.MethodPost, "/Invoice/Factory/createInvoiceReminder", q, body)
	if err != nil {
		return nil, err
	}
	return decodeObject[Invoice](raw)
}

// IsPartiallyPaid reports whether the invoice has received any partial payment.
func (s *InvoicesService) IsPartiallyPaid(ctx context.Context, id ID) (bool, error) {
	raw, err := s.c.do(ctx, http.MethodGet, fmt.Sprintf("/Invoice/%d/getIsPartiallyPaid", id), nil, nil)
	if err != nil {
		return false, err
	}
	return string(raw) == "true", nil
}

// Cancel creates a cancellation invoice for the given invoice.
func (s *InvoicesService) Cancel(ctx context.Context, id ID) (*Invoice, error) {
	raw, err := s.c.do(ctx, http.MethodPost, fmt.Sprintf("/Invoice/%d/cancelInvoice", id), nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeObject[Invoice](raw)
}

// RenderResult contains the PDF metadata returned by [InvoicesService.Render].
type RenderResult struct {
	Thumbs []string `json:"thumbs,omitempty"`
	Pages  Num      `json:"pages,omitempty"`
	DocID  string   `json:"docId,omitempty"`
}

// Render forces re-rendering of an invoice's PDF.
func (s *InvoicesService) Render(ctx context.Context, id ID, forceReload bool) (*RenderResult, error) {
	body := map[string]bool{"forceReload": forceReload}
	raw, err := s.c.do(ctx, http.MethodPost, fmt.Sprintf("/Invoice/%d/render", id), nil, body)
	if err != nil {
		return nil, err
	}
	// /render returns metadata directly under "objects", or sometimes flat —
	// decodeObject handles both array and bare object.
	return decodeObject[RenderResult](raw)
}

// SendEmailParams is the body for the SendViaEmail methods on Invoice,
// CreditNote, and Order services.
type SendEmailParams struct {
	// ToEmail is the recipient. Required.
	ToEmail string `json:"toEmail"`
	// Subject of the email. Required.
	Subject string `json:"subject"`
	// Text body. Plain text or HTML.
	Text string `json:"text,omitempty"`
	// Copy, when true, sends a copy of the email to the sender.
	Copy *bool `json:"copy,omitempty"`
	// AdditionalAttachments is a comma-separated list of file paths
	// (sevdesk-managed temp files) to attach beyond the document itself.
	AdditionalAttachments string `json:"additionalAttachments,omitempty"`
	// CcEmail is a comma-separated list of CC recipients.
	CcEmail string `json:"ccEmail,omitempty"`
	// BccEmail is a comma-separated list of BCC recipients.
	BccEmail string `json:"bccEmail,omitempty"`
	// SendXML, when true, attaches the e-invoice XML alongside the PDF
	// (invoice-only field).
	SendXML *bool `json:"sendXml,omitempty"`
}

// SendViaEmail sends an invoice by email.
func (s *InvoicesService) SendViaEmail(ctx context.Context, id ID, params *SendEmailParams) (*Email, error) {
	raw, err := s.c.do(ctx, http.MethodPost, fmt.Sprintf("/Invoice/%d/sendViaEmail", id), nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[Email](raw)
}

// PDFDocument is the response from [InvoicesService.GetPdf] and similar endpoints.
type PDFDocument struct {
	Filename      string `json:"filename,omitempty"`
	MimeType      string `json:"mimeType,omitempty"`
	Base64Encoded bool   `json:"base64encoded,omitempty"`
	Content       string `json:"content,omitempty"` // base64 string
}

// GetPdf returns the rendered PDF of an invoice as base64-encoded content.
func (s *InvoicesService) GetPdf(ctx context.Context, id ID, download, preventSendBy bool) (*PDFDocument, error) {
	q := url.Values{}
	if download {
		q.Set("download", "true")
	}
	if preventSendBy {
		q.Set("preventSendBy", "true")
	}
	raw, err := s.c.do(ctx, http.MethodGet, fmt.Sprintf("/Invoice/%d/getPdf", id), q, nil)
	if err != nil {
		return nil, err
	}
	return decodeObject[PDFDocument](raw)
}

// GetXML returns the e-invoice XML for the given invoice.
func (s *InvoicesService) GetXML(ctx context.Context, id ID) (string, error) {
	raw, err := s.c.do(ctx, http.MethodGet, fmt.Sprintf("/Invoice/%d/getXml", id), nil, nil)
	if err != nil {
		return "", err
	}
	return string(raw), nil
}

// SendType is the channel by which a document was sent.
type SendType string

// SendType values.
const (
	// SendTypeVPR marks the document as printed.
	SendTypeVPR SendType = "VPR"
	// SendTypeVPDF marks the document as downloaded as a PDF.
	SendTypeVPDF SendType = "VPDF"
	// SendTypeVM marks the document as emailed.
	SendTypeVM SendType = "VM"
	// SendTypeVP marks the document as sent by postal mail.
	SendTypeVP SendType = "VP"
)

// MarkSentParams is the body for MarkSent methods on Invoice, CreditNote,
// and Order services.
type MarkSentParams struct {
	// SendType records how the document was sent. Required. See [SendType].
	SendType SendType `json:"sendType"`
	// SendDraft, when true, also marks an underlying draft as sent.
	SendDraft *bool `json:"sendDraft,omitempty"`
}

// MarkSent records that the invoice has been sent by the given channel.
func (s *InvoicesService) MarkSent(ctx context.Context, id ID, params *MarkSentParams) (*Invoice, error) {
	raw, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/Invoice/%d/sendBy", id), nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[Invoice](raw)
}

// Enshrine locks the invoice so it can no longer be edited.
func (s *InvoicesService) Enshrine(ctx context.Context, id ID) error {
	_, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/Invoice/%d/enshrine", id), nil, nil)
	return err
}

// BookInvoiceParams is the body for [InvoicesService.Book] and
// [CreditNotesService.Book].
type BookInvoiceParams struct {
	// Amount paid. Can be partial — book again to add more.
	Amount Decimal `json:"amount"`
	// Date of the booking (usually today).
	Date *time.Time `json:"date"`
	// Type of booking. See [BookingType].
	Type BookingType `json:"type"`
	// CheckAccount the payment hits. Required.
	CheckAccount *Ref `json:"checkAccount"`
	// CheckAccountTransaction links this booking to an existing transaction
	// (e.g. matched from a bank import).
	CheckAccountTransaction *Ref `json:"checkAccountTransaction,omitempty"`
	// CreateFeed controls whether sevdesk creates an activity feed entry.
	CreateFeed *bool `json:"createFeed,omitempty"`
}

// InvoiceBookResult is the log entry returned by [InvoicesService.Book].
type InvoiceBookResult struct {
	ID          ID         `json:"id,omitempty"`
	ObjectName  string     `json:"objectName,omitempty"`
	Create      *time.Time `json:"create,omitempty"`
	Invoice     *Ref       `json:"creditNote,omitempty"` // sevdesk reuses this name
	FromStatus  string     `json:"fromStatus,omitempty"`
	ToStatus    string     `json:"toStatus,omitempty"`
	AmountPaid  Decimal    `json:"ammountPayed,omitempty"` // sevdesk typo
	BookingDate *time.Time `json:"bookingDate,omitempty"`
	SevClient   *Ref       `json:"sevClient,omitempty"`
}

// Book records a payment against an invoice.
func (s *InvoicesService) Book(ctx context.Context, id ID, params *BookInvoiceParams) (*InvoiceBookResult, error) {
	raw, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/Invoice/%d/bookAmount", id), nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[InvoiceBookResult](raw)
}

// ResetToOpen reverts a paid or enshrined invoice back to open status.
func (s *InvoicesService) ResetToOpen(ctx context.Context, id ID) (*Invoice, error) {
	raw, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/Invoice/%d/resetToOpen", id), nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeObject[Invoice](raw)
}

// ResetToDraft reverts an invoice back to draft status.
func (s *InvoicesService) ResetToDraft(ctx context.Context, id ID) (*Invoice, error) {
	raw, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/Invoice/%d/resetToDraft", id), nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeObject[Invoice](raw)
}

// LayoutKey picks which layout aspect to change via
// [InvoicesService.ChangeLayout] and similar.
type LayoutKey string

// LayoutKey values.
const (
	// LayoutKeyTemplate picks a rendering template.
	LayoutKeyTemplate LayoutKey = "template"
	// LayoutKeyLetterpaper picks the letterpaper (Briefpapier) background.
	LayoutKeyLetterpaper LayoutKey = "letterPaper"
	// LayoutKeyLanguage picks the document language code.
	LayoutKeyLanguage LayoutKey = "language"
	// LayoutKeyPayPal picks the PayPal payment configuration.
	LayoutKeyPayPal LayoutKey = "payPal"
)

// ChangeLayoutParams is the body for ChangeLayout methods on Invoice,
// CreditNote, and Order services.
type ChangeLayoutParams struct {
	// Key picks which layout aspect to change. Required. See [LayoutKey].
	Key LayoutKey `json:"key"`
	// Value is the id of the template/letterpaper/language/payPal config
	// to apply.
	Value string `json:"value"`
}

// ChangeLayout changes the rendering layout of an invoice (template, paper, language).
func (s *InvoicesService) ChangeLayout(ctx context.Context, id ID, params *ChangeLayoutParams) (*RenderResult, error) {
	raw, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/Invoice/%d/changeParameter", id), nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[RenderResult](raw)
}

// Email is the response shape for endpoints that send emails.
type Email struct {
	ID         ID         `json:"id,omitempty"`
	ObjectName string     `json:"objectName,omitempty"`
	Create     *time.Time `json:"create,omitempty"`
	Update     *time.Time `json:"update,omitempty"`
	Object     *Ref       `json:"object,omitempty"`
	From       string     `json:"from,omitempty"`
	To         string     `json:"to,omitempty"`
	Subject    string     `json:"subject,omitempty"`
	Text       string     `json:"text,omitempty"`
	CC         string     `json:"cc,omitempty"`
	BCC        string     `json:"bcc,omitempty"`
	Arrived    *time.Time `json:"arrived,omitempty"`
	SevClient  *Ref       `json:"sevClient,omitempty"`
}

// ListInvoicePositionsParams filters [InvoicesService.ListPositions].
type ListInvoicePositionsParams struct {
	// ID looks up a single position by id.
	ID ID
	// Invoice narrows positions to one invoice.
	Invoice *Ref
	// Part narrows positions to those referencing a specific inventory item.
	Part *Ref
}

func (p *ListInvoicePositionsParams) query() url.Values {
	if p == nil {
		return nil
	}
	q := url.Values{}
	if p.ID != 0 {
		q.Set("id", p.ID.String())
	}
	if p.Invoice != nil {
		q.Set("invoice[id]", p.Invoice.ID.String())
		q.Set("invoice[objectName]", p.Invoice.ObjectName)
	}
	if p.Part != nil {
		q.Set("part[id]", p.Part.ID.String())
		q.Set("part[objectName]", p.Part.ObjectName)
	}
	return q
}

// ListPositions returns invoice positions across invoices (use Invoice filter
// for a single invoice; [InvoicesService.GetPositions] is usually nicer).
func (s *InvoicesService) ListPositions(ctx context.Context, opts *ListInvoicePositionsParams) ([]InvoicePos, error) {
	raw, err := s.c.do(ctx, http.MethodGet, "/InvoicePos", opts.query(), nil)
	if err != nil {
		return nil, err
	}
	return decodeList[InvoicePos](raw)
}
