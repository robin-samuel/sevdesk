package sevdesk

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// CreditNotesService handles communication with /CreditNote endpoints.
type CreditNotesService struct {
	c *Client
}

// CreditNote is a sevdesk credit note (Gutschrift).
type CreditNote struct {
	ID         ID         `json:"id,omitempty"`
	ObjectName string     `json:"objectName,omitempty"`
	Create     *time.Time `json:"create,omitempty"`
	Update     *time.Time `json:"update,omitempty"`
	SevClient  *Ref       `json:"sevClient,omitempty"`
	CreateUser *Ref       `json:"createUser,omitempty"`

	CreditNoteNumber string     `json:"creditNoteNumber,omitempty"`
	Contact          *Ref       `json:"contact,omitempty"`
	ContactPerson    *Ref       `json:"contactPerson,omitempty"`
	CreditNoteDate   *time.Time `json:"creditNoteDate,omitempty"`
	Header           string     `json:"header,omitempty"`
	HeadText         string     `json:"headText,omitempty"`
	FootText         string     `json:"footText,omitempty"`
	TimeToPay        Num        `json:"timeToPay,omitempty"`
	DiscountTime     Num        `json:"discountTime,omitempty"`
	Discount         Decimal    `json:"discount,omitempty"`

	Address            string `json:"address,omitempty"`
	AddressName        string `json:"addressName,omitempty"`
	AddressName2       string `json:"addressName2,omitempty"`
	AddressStreet      string `json:"addressStreet,omitempty"`
	AddressZip         string `json:"addressZip,omitempty"`
	AddressCity        string `json:"addressCity,omitempty"`
	AddressCountry     *Ref   `json:"addressCountry,omitempty"`
	AddressGender      string `json:"addressGender,omitempty"`
	AddressParentName  string `json:"addressParentName,omitempty"`
	AddressParentName2 string `json:"addressParentName2,omitempty"`

	PayDate           *time.Time       `json:"payDate,omitempty"`
	DeliveryDate      *time.Time       `json:"deliveryDate,omitempty"`
	DeliveryDateUntil *time.Time       `json:"deliveryDateUntil,omitempty"`
	Status            CreditNoteStatus `json:"status,omitempty"` // see [CreditNoteStatus]
	SmallSettlement   Bool             `json:"smallSettlement,omitempty"`
	TaxRate           Decimal          `json:"taxRate,omitempty"`
	TaxText           string           `json:"taxText,omitempty"`
	TaxType           string           `json:"taxType,omitempty"`
	TaxRule           *Ref             `json:"taxRule,omitempty"`
	TaxSet            *Ref             `json:"taxSet,omitempty"`
	SendDate          *time.Time       `json:"sendDate,omitempty"`
	SendType          SendType         `json:"sendType,omitempty"`
	CreditNoteType    string           `json:"creditNoteType,omitempty"`
	AccountingType    *Ref             `json:"accountingType,omitempty"`
	BookingCategory   BookingCategory  `json:"bookingCategory,omitempty"`

	AccountIntervall      string     `json:"accountIntervall,omitempty"`
	AccountNextCreditNote *time.Time `json:"accountNextCreditNote,omitempty"`
	AccountLastCreditNote *time.Time `json:"accountLastCreditNote,omitempty"`
	AccountEndDate        *time.Time `json:"accountEndDate,omitempty"`

	ReminderTotal    Decimal    `json:"reminderTotal,omitempty"`
	ReminderDebit    Decimal    `json:"reminderDebit,omitempty"`
	ReminderDeadline *time.Time `json:"reminderDeadline,omitempty"`
	ReminderCharge   Decimal    `json:"reminderCharge,omitempty"`

	Currency  string `json:"currency,omitempty"`
	TaxNumber string `json:"taxNumber,omitempty"`
	VATNumber string `json:"vatNumber,omitempty"`

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

	CustomerInternalNote string     `json:"customerInternalNote,omitempty"`
	ShowNet              Bool       `json:"showNet,omitempty"`
	Enshrined            *time.Time `json:"enshrined,omitempty"`
	IsTransferred        Bool       `json:"isTransferred,omitempty"`
}

// CreditNoteStatus is the lifecycle state of a credit note.
type CreditNoteStatus Num

func (s CreditNoteStatus) String() string { return Num(s).String() }

func (s *CreditNoteStatus) UnmarshalJSON(b []byte) error {
	var n Num
	if err := n.UnmarshalJSON(b); err != nil {
		return err
	}
	*s = CreditNoteStatus(n)
	return nil
}

// CreditNoteStatus values.
const (
	// CreditNoteStatusDraft is an unfinalized credit note.
	CreditNoteStatusDraft CreditNoteStatus = 100
	// CreditNoteStatusOpen is finalized and outstanding.
	CreditNoteStatusOpen CreditNoteStatus = 200
	// CreditNoteStatusLinked has been linked to an invoice or voucher.
	CreditNoteStatusLinked CreditNoteStatus = 300
	// CreditNoteStatusPartiallyPaid means some amount has been settled.
	CreditNoteStatusPartiallyPaid CreditNoteStatus = 750
	// CreditNoteStatusPaid is fully settled.
	CreditNoteStatusPaid CreditNoteStatus = 1000
)

// BookingCategory classifies what a credit note is for.
type BookingCategory string

// BookingCategory values.
const (
	// BookingCategoryROYALTYAssigned is a royalty payment assigned to the note.
	BookingCategoryROYALTYAssigned BookingCategory = "ROYALTY_ASSIGNED"
	// BookingCategoryProvision is a commission payment.
	BookingCategoryProvision BookingCategory = "PROVISION"
	// BookingCategoryProvisionAuto is an auto-calculated commission payment.
	BookingCategoryProvisionAuto BookingCategory = "PROVISION_AUTO"
	// BookingCategoryRevenue is a revenue entry.
	BookingCategoryRevenue BookingCategory = "REVENUE"
	// BookingCategoryRevenueDeduct is a revenue deduction (e.g. refund).
	BookingCategoryRevenueDeduct BookingCategory = "REVENUE_DEDUCT"
)

// CreditNotePos is one line item on a credit note.
type CreditNotePos struct {
	ID         ID         `json:"id,omitempty"`
	ObjectName string     `json:"objectName,omitempty"`
	Create     *time.Time `json:"create,omitempty"`
	Update     *time.Time `json:"update,omitempty"`
	SevClient  *Ref       `json:"sevClient,omitempty"`

	CreditNote     *Ref    `json:"creditNote,omitempty"`
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

// CreditNotePosCreate is a position to send when creating a credit note.
// ObjectName and MapAll are set automatically by [CreditNotesService.Create].
type CreditNotePosCreate struct {
	ObjectName string `json:"objectName"` // set by Create
	MapAll     bool   `json:"mapAll"`     // set by Create

	// Part links the position to an inventory item. Optional.
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
	// Unity is the unit of measure. Required. Use [UnityRef].
	Unity *Ref `json:"unity"`
	// Text is the position's long description.
	Text *string `json:"text,omitempty"`
	// Discount is a per-position discount.
	Discount *Decimal `json:"discount,omitempty"`
	// TaxRate is the VAT rate as a percentage. Required.
	TaxRate Decimal `json:"taxRate"`
}

// CreateCreditNoteParams is the body for [CreditNotesService.Create].
type CreateCreditNoteParams struct {
	// CreditNote carries the credit-note-level fields. Required.
	CreditNote *CreditNoteCreateFields `json:"creditNote"`
	// Positions are the line items being added.
	Positions []CreditNotePosCreate `json:"creditNotePosSave,omitempty"`
	// PositionsDelete removes existing positions by reference.
	PositionsDelete []Ref `json:"creditNotePosDelete,omitempty"`
	// Discounts are document-level discounts or surcharges.
	Discounts []InvoiceDiscountCreate `json:"discountSave,omitempty"`
	// DiscountsDelete removes existing discounts by reference.
	DiscountsDelete []Ref `json:"discountDelete,omitempty"`
}

// CreditNoteCreateFields is the credit-note body for [CreateCreditNoteParams].
// ObjectName and MapAll are set automatically by [CreditNotesService.Create].
type CreditNoteCreateFields struct {
	ObjectName string `json:"objectName"` // set by Create
	MapAll     bool   `json:"mapAll"`     // set by Create

	// CreditNoteNumber overrides the auto-assigned number.
	CreditNoteNumber *string `json:"creditNoteNumber,omitempty"`
	// Contact is the customer being credited. Required.
	Contact *Ref `json:"contact"`
	// ContactPerson is the sevdesk user shown as the contact. Required.
	ContactPerson *Ref `json:"contactPerson"`
	// CreditNoteDate is the document date. Required.
	CreditNoteDate *time.Time `json:"creditNoteDate"`
	// Header is the credit note title.
	Header *string `json:"header,omitempty"`
	// HeadText is the rich-text intro.
	HeadText *string `json:"headText,omitempty"`
	// FootText is the rich-text outro.
	FootText *string `json:"footText,omitempty"`
	// AddressCountry of the rendered address.
	AddressCountry *Ref `json:"addressCountry,omitempty"`
	// Address is the rendered address (multi-line string).
	Address *string `json:"address,omitempty"`
	// Status of the credit note. See [CreditNoteStatusDraft] and adjacent.
	Status *CreditNoteStatus `json:"status,omitempty"`
	// SmallSettlement marks the note as Kleinunternehmer (no VAT shown).
	SmallSettlement *Bool `json:"smallSettlement,omitempty"`
	// TaxRate is the default tax rate when positions don't override it.
	TaxRate *Decimal `json:"taxRate,omitempty"`
	// TaxRule selects the tax rule. Required.
	TaxRule *Ref `json:"taxRule"`
	// TaxText is the line printed below the totals.
	TaxText *string `json:"taxText,omitempty"`
	// TaxType picks the tax mode (e.g. "default", "eu", "noteu", "ss").
	TaxType *string `json:"taxType,omitempty"`
	// Currency is the ISO 4217 code. Required.
	Currency string `json:"currency"`
	// BookingCategory is the booking category. See [BookingCategory].
	BookingCategory *BookingCategory `json:"bookingCategory,omitempty"`
	// ShowNet shows net prices on the rendered PDF.
	ShowNet *Bool `json:"showNet,omitempty"`
	// SendType records how the note was sent. See [SendType].
	SendType *SendType `json:"sendType,omitempty"`
	// CustomerInternalNote is private, not shown to the customer.
	CustomerInternalNote *string `json:"customerInternalNote,omitempty"`
	// DeliveryDate is when the goods/services were delivered.
	DeliveryDate *time.Time `json:"deliveryDate,omitempty"`
}

// UpdateCreditNoteParams is the body for [CreditNotesService.Update]. Only
// works for drafts. See [CreditNoteCreateFields] for field semantics.
type UpdateCreditNoteParams struct {
	CreditNoteNumber     *string           `json:"creditNoteNumber,omitempty"`
	Contact              *Ref              `json:"contact,omitempty"`
	ContactPerson        *Ref              `json:"contactPerson,omitempty"`
	CreditNoteDate       *time.Time        `json:"creditNoteDate,omitempty"`
	Header               *string           `json:"header,omitempty"`
	HeadText             *string           `json:"headText,omitempty"`
	FootText             *string           `json:"footText,omitempty"`
	AddressCountry       *Ref              `json:"addressCountry,omitempty"`
	Address              *string           `json:"address,omitempty"`
	Status               *CreditNoteStatus `json:"status,omitempty"`
	SmallSettlement      *Bool             `json:"smallSettlement,omitempty"`
	TaxRate              *Decimal          `json:"taxRate,omitempty"`
	TaxRule              *Ref              `json:"taxRule,omitempty"`
	TaxText              *string           `json:"taxText,omitempty"`
	TaxType              *string           `json:"taxType,omitempty"`
	Currency             *string           `json:"currency,omitempty"`
	BookingCategory      *BookingCategory  `json:"bookingCategory,omitempty"`
	ShowNet              *Bool             `json:"showNet,omitempty"`
	SendType             *SendType         `json:"sendType,omitempty"`
	CustomerInternalNote *string           `json:"customerInternalNote,omitempty"`
	DeliveryDate         *time.Time        `json:"deliveryDate,omitempty"`
}

// ListCreditNotesParams filters [CreditNotesService.List].
type ListCreditNotesParams struct {
	// Status filters by credit-note status. Sevdesk takes a string here;
	// pass the numeric value (e.g. "1000" for paid).
	Status string
	// CreditNoteNumber matches the exact number.
	CreditNoteNumber string
	// StartDate excludes notes dated before this point.
	StartDate time.Time
	// EndDate excludes notes dated after this point.
	EndDate time.Time
	// Contact narrows to notes for one customer.
	Contact *Ref
}

func (p *ListCreditNotesParams) query() url.Values {
	if p == nil {
		return nil
	}
	q := url.Values{}
	if p.Status != "" {
		q.Set("status", p.Status)
	}
	if p.CreditNoteNumber != "" {
		q.Set("creditNoteNumber", p.CreditNoteNumber)
	}
	if !p.StartDate.IsZero() {
		q.Set("startDate", p.StartDate.Format("02.01.2006"))
	}
	if !p.EndDate.IsZero() {
		q.Set("endDate", p.EndDate.Format("02.01.2006"))
	}
	if p.Contact != nil {
		q.Set("contact[id]", p.Contact.ID.String())
		q.Set("contact[objectName]", p.Contact.ObjectName)
	}
	return q
}

// List returns credit notes matching the given filter.
func (s *CreditNotesService) List(ctx context.Context, opts *ListCreditNotesParams) ([]CreditNote, error) {
	raw, err := s.c.do(ctx, http.MethodGet, "/CreditNote", opts.query(), nil)
	if err != nil {
		return nil, err
	}
	return decodeList[CreditNote](raw)
}

// Get returns the credit note with the given id.
func (s *CreditNotesService) Get(ctx context.Context, id ID) (*CreditNote, error) {
	raw, err := s.c.do(ctx, http.MethodGet, fmt.Sprintf("/CreditNote/%d", id), nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeObject[CreditNote](raw)
}

// CreateCreditNoteResult is the return value of [CreditNotesService.Create].
type CreateCreditNoteResult struct {
	CreditNote CreditNote      `json:"creditNote"`
	Positions  []CreditNotePos `json:"creditNotePos"`
}

// Create creates a credit note with its positions and discounts.
func (s *CreditNotesService) Create(ctx context.Context, params *CreateCreditNoteParams) (*CreateCreditNoteResult, error) {
	if params != nil && params.CreditNote != nil {
		if params.CreditNote.ObjectName == "" {
			params.CreditNote.ObjectName = ObjectCreditNote
		}
		params.CreditNote.MapAll = true
	}
	for i := range params.Positions {
		if params.Positions[i].ObjectName == "" {
			params.Positions[i].ObjectName = ObjectCreditNotePos
		}
		params.Positions[i].MapAll = true
	}
	for i := range params.Discounts {
		if params.Discounts[i].ObjectName == "" {
			params.Discounts[i].ObjectName = ObjectDiscounts
		}
		params.Discounts[i].MapAll = true
	}
	raw, err := s.c.do(ctx, http.MethodPost, "/CreditNote/Factory/saveCreditNote", nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[CreateCreditNoteResult](raw)
}

// CreateFromInvoice generates a credit note based on an existing invoice.
func (s *CreditNotesService) CreateFromInvoice(ctx context.Context, invoiceID ID) (*CreateCreditNoteResult, error) {
	body := map[string]any{
		"invoice": map[string]any{
			"id":         invoiceID,
			"objectName": ObjectInvoice,
		},
	}
	raw, err := s.c.do(ctx, http.MethodPost, "/CreditNote/Factory/createFromInvoice", nil, body)
	if err != nil {
		return nil, err
	}
	return decodeObject[CreateCreditNoteResult](raw)
}

// CreateFromVoucher generates a credit note based on an existing voucher.
func (s *CreditNotesService) CreateFromVoucher(ctx context.Context, voucherID ID) (*CreateCreditNoteResult, error) {
	body := map[string]any{
		"voucher": map[string]any{
			"id":         voucherID,
			"objectName": ObjectVoucher,
		},
	}
	raw, err := s.c.do(ctx, http.MethodPost, "/CreditNote/Factory/createFromVoucher", nil, body)
	if err != nil {
		return nil, err
	}
	return decodeObject[CreateCreditNoteResult](raw)
}

// Update modifies an existing credit note (only draft credit notes).
func (s *CreditNotesService) Update(ctx context.Context, id ID, params *UpdateCreditNoteParams) (*CreditNote, error) {
	raw, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/CreditNote/%d", id), nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[CreditNote](raw)
}

// Delete removes a credit note (only draft credit notes).
func (s *CreditNotesService) Delete(ctx context.Context, id ID) error {
	_, err := s.c.do(ctx, http.MethodDelete, fmt.Sprintf("/CreditNote/%d", id), nil, nil)
	return err
}

// SendByPrint marks a credit note as sent and returns the rendered PDF metadata
// in one round-trip (GET /CreditNote/{id}/sendByWithRender).
func (s *CreditNotesService) SendByPrint(ctx context.Context, id ID, sendType string) (*RenderResult, error) {
	q := url.Values{"sendType": {sendType}}
	raw, err := s.c.do(ctx, http.MethodGet, fmt.Sprintf("/CreditNote/%d/sendByWithRender", id), q, nil)
	if err != nil {
		return nil, err
	}
	return decodeObject[RenderResult](raw)
}

// MarkSent records that the credit note has been sent by the given channel.
func (s *CreditNotesService) MarkSent(ctx context.Context, id ID, params *MarkSentParams) (*CreditNote, error) {
	raw, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/CreditNote/%d/sendBy", id), nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[CreditNote](raw)
}

// Enshrine locks the credit note so it can no longer be edited.
func (s *CreditNotesService) Enshrine(ctx context.Context, id ID) error {
	_, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/CreditNote/%d/enshrine", id), nil, nil)
	return err
}

// GetPdf returns the rendered PDF as base64-encoded content.
func (s *CreditNotesService) GetPdf(ctx context.Context, id ID, download, preventSendBy bool) (*PDFDocument, error) {
	q := url.Values{}
	if download {
		q.Set("download", "true")
	}
	if preventSendBy {
		q.Set("preventSendBy", "true")
	}
	raw, err := s.c.do(ctx, http.MethodGet, fmt.Sprintf("/CreditNote/%d/getPdf", id), q, nil)
	if err != nil {
		return nil, err
	}
	return decodeObject[PDFDocument](raw)
}

// SendViaEmail sends a credit note by email.
func (s *CreditNotesService) SendViaEmail(ctx context.Context, id ID, params *SendEmailParams) (*Email, error) {
	raw, err := s.c.do(ctx, http.MethodPost, fmt.Sprintf("/CreditNote/%d/sendViaEmail", id), nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[Email](raw)
}

// Book records a payment against a credit note.
func (s *CreditNotesService) Book(ctx context.Context, id ID, params *BookInvoiceParams) (*InvoiceBookResult, error) {
	raw, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/CreditNote/%d/bookAmount", id), nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[InvoiceBookResult](raw)
}

// ResetToOpen reverts a paid or enshrined credit note back to open status.
func (s *CreditNotesService) ResetToOpen(ctx context.Context, id ID) (*CreditNote, error) {
	raw, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/CreditNote/%d/resetToOpen", id), nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeObject[CreditNote](raw)
}

// ResetToDraft reverts a credit note back to draft status.
func (s *CreditNotesService) ResetToDraft(ctx context.Context, id ID) (*CreditNote, error) {
	raw, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/CreditNote/%d/resetToDraft", id), nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeObject[CreditNote](raw)
}

// ChangeLayout changes the rendering layout (template, paper, language) of a credit note.
func (s *CreditNotesService) ChangeLayout(ctx context.Context, id ID, params *ChangeLayoutParams) (*RenderResult, error) {
	raw, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/CreditNote/%d/changeParameter", id), nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[RenderResult](raw)
}

// ListPositions returns credit note positions; filter by credit note via Params.
func (s *CreditNotesService) ListPositions(ctx context.Context, creditNoteID ID) ([]CreditNotePos, error) {
	q := url.Values{
		"creditNote[id]":         {creditNoteID.String()},
		"creditNote[objectName]": {ObjectCreditNote},
	}
	raw, err := s.c.do(ctx, http.MethodGet, "/CreditNotePos", q, nil)
	if err != nil {
		return nil, err
	}
	return decodeList[CreditNotePos](raw)
}
