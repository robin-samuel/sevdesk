package sevdesk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"
)

// VouchersService handles communication with /Voucher endpoints.
type VouchersService struct {
	c *Client
}

// Voucher is a sevdesk voucher (a receipt or expense document).
type Voucher struct {
	ID         ID         `json:"id,omitempty"`
	ObjectName string     `json:"objectName,omitempty"`
	Create     *time.Time `json:"create,omitempty"`
	Update     *time.Time `json:"update,omitempty"`
	SevClient  *Ref       `json:"sevClient,omitempty"`
	CreateUser *Ref       `json:"createUser,omitempty"`

	VoucherDate        *time.Time    `json:"voucherDate,omitempty"`
	Supplier           *Ref          `json:"supplier,omitempty"`
	SupplierName       string        `json:"supplierName,omitempty"`
	SupplierNameAtSave string        `json:"supplierNameAtSave,omitempty"`
	Description        string        `json:"description,omitempty"`
	Document           *Ref          `json:"document,omitempty"`
	PayDate            *time.Time    `json:"payDate,omitempty"`
	Status             VoucherStatus `json:"status,omitempty"` // see [VoucherStatus]
	Currency           string        `json:"currency,omitempty"`

	SumNet             Decimal `json:"sumNet,omitempty"`
	SumTax             Decimal `json:"sumTax,omitempty"`
	SumGross           Decimal `json:"sumGross,omitempty"`
	SumNetAccounting   Decimal `json:"sumNetAccounting,omitempty"`
	SumTaxAccounting   Decimal `json:"sumTaxAccounting,omitempty"`
	SumGrossAccounting Decimal `json:"sumGrossAccounting,omitempty"`
	SumDiscountNet     Decimal `json:"sumDiscountNet,omitempty"`
	SumDiscountGross   Decimal `json:"sumDiscountGross,omitempty"`
	PaidAmount         Decimal `json:"paidAmount,omitempty"`

	SumNetForeignCurrency           Decimal `json:"sumNetForeignCurrency,omitempty"`
	SumTaxForeignCurrency           Decimal `json:"sumTaxForeignCurrency,omitempty"`
	SumGrossForeignCurrency         Decimal `json:"sumGrossForeignCurrency,omitempty"`
	SumDiscountNetForeignCurrency   Decimal `json:"sumDiscountNetForeignCurrency,omitempty"`
	SumDiscountGrossForeignCurrency Decimal `json:"sumDiscountGrossForeignCurrency,omitempty"`

	ShowNet                Bool        `json:"showNet,omitempty"`
	Hidden                 Bool        `json:"hidden,omitempty"`
	SelectedForPaymentFile Bool        `json:"selectedForPaymentFile,omitempty"`
	TaxType                string      `json:"taxType,omitempty"`
	TaxRule                *Ref        `json:"taxRule,omitempty"`
	CreditDebit            CreditDebit `json:"creditDebit,omitempty"`
	VoucherType            VoucherType `json:"voucherType,omitempty"`

	// RecurringIntervall (typo) and RecurringInterval both come back from sevdesk.
	RecurringIntervall   string     `json:"recurringIntervall,omitempty"`
	RecurringInterval    string     `json:"recurringInterval,omitempty"`
	RecurringStartDate   *time.Time `json:"recurringStartDate,omitempty"`
	RecurringNextVoucher *time.Time `json:"recurringNextVoucher,omitempty"`
	RecurringLastVoucher *time.Time `json:"recurringLastVoucher,omitempty"`
	RecurringEndDate     *time.Time `json:"recurringEndDate,omitempty"`

	Enshrined         *time.Time `json:"enshrined,omitempty"`
	SendType          SendType   `json:"sendType,omitempty"`
	InSource          string     `json:"inSource,omitempty"`
	IBAN              string     `json:"iban,omitempty"`
	VATNumber         string     `json:"vatNumber,omitempty"`
	PaymentDeadline   *time.Time `json:"paymentDeadline,omitempty"`
	Tip               Decimal    `json:"tip,omitempty"`
	MileageRate       Decimal    `json:"mileageRate,omitempty"`
	DeliveryDate      *time.Time `json:"deliveryDate,omitempty"`
	DeliveryDateUntil *time.Time `json:"deliveryDateUntil,omitempty"`

	AccountingSpecialCase string `json:"accountingSpecialCase,omitempty"`
	TaxmaroStockAccount   string `json:"taxmaroStockAccount,omitempty"`
	ResultDisdar          string `json:"resultDisdar,omitempty"`
	CostCentre            *Ref   `json:"costCentre,omitempty"`
}

// VoucherStatus is the lifecycle state of a voucher.
type VoucherStatus Num

func (s VoucherStatus) String() string { return Num(s).String() }

func (s *VoucherStatus) UnmarshalJSON(b []byte) error {
	var n Num
	if err := n.UnmarshalJSON(b); err != nil {
		return err
	}
	*s = VoucherStatus(n)
	return nil
}

// VoucherStatus values.
const (
	// VoucherStatusDraft is a voucher in progress, not yet finalized.
	VoucherStatusDraft VoucherStatus = 50
	// VoucherStatusOpen is finalized and awaiting payment.
	VoucherStatusOpen VoucherStatus = 100
	// VoucherStatusPaymentOrdered means a payment has been instructed but
	// not yet confirmed by the bank.
	VoucherStatusPaymentOrdered VoucherStatus = 250
	// VoucherStatusPartiallyPaid means some money has been received/paid.
	VoucherStatusPartiallyPaid VoucherStatus = 750
	// VoucherStatusPaid is fully paid.
	VoucherStatusPaid VoucherStatus = 1000
)

// CreditDebit marks the direction of a voucher (or invoice booking).
type CreditDebit string

// CreditDebit values.
const (
	// VoucherCredit marks an incoming-money voucher (income, customer payment).
	VoucherCredit CreditDebit = "C"
	// VoucherDebit marks an outgoing-money voucher (expense, supplier bill).
	VoucherDebit CreditDebit = "D"
	// VoucherSpend marks petty-cash / employee-spending vouchers.
	VoucherSpend CreditDebit = "S"
)

// VoucherType is the kind of voucher.
type VoucherType string

// VoucherType values.
const (
	// VoucherTypeVoucher is a standard one-off voucher.
	VoucherTypeVoucher VoucherType = "VOU"
	// VoucherTypeRecurring is a template that generates vouchers on a schedule.
	VoucherTypeRecurring VoucherType = "RV"
)

// CreateVoucherParams is the body for [VouchersService.Create]. The endpoint
// (POST /Voucher/Factory/saveVoucher) wraps voucher fields with positions.
type CreateVoucherParams struct {
	// Voucher carries the voucher-level fields. Required.
	Voucher *VoucherCreateFields `json:"voucher"`
	// Positions are the line items being added.
	Positions []VoucherPosCreate `json:"voucherPosSave,omitempty"`
	// PositionsDelete removes existing positions by reference.
	PositionsDelete []Ref `json:"voucherPosDelete,omitempty"`
	// Filename of a file previously uploaded via [VouchersService.UploadFile].
	// Attaches the file to the voucher.
	Filename string `json:"filename,omitempty"`
}

// VoucherCreateFields is the voucher-side of [CreateVoucherParams].
// ObjectName and MapAll are set automatically by [VouchersService.Create].
type VoucherCreateFields struct {
	ObjectName string `json:"objectName"` // set by Create
	MapAll     bool   `json:"mapAll"`     // set by Create

	// Status of the voucher. See [VoucherStatus].
	Status *VoucherStatus `json:"status,omitempty"`
	// TaxType controls how VAT is applied. Common values: "default", "eu",
	// "noteu", "custom", "ss" (Kleinunternehmer).
	TaxType *string `json:"taxType,omitempty"`
	// CreditDebit picks the direction. See [CreditDebit].
	CreditDebit *CreditDebit `json:"creditDebit,omitempty"`
	// VoucherType picks the kind. See [VoucherType].
	VoucherType *VoucherType `json:"voucherType,omitempty"`

	// VoucherDate is the document date printed on the voucher.
	VoucherDate *time.Time `json:"voucherDate,omitempty"`
	// Supplier is the contact issuing the voucher.
	Supplier *Ref `json:"supplier,omitempty"`
	// SupplierName overrides the rendered supplier name (used when Supplier is empty).
	SupplierName *string `json:"supplierName,omitempty"`
	// Description is usually the voucher number printed by the supplier.
	Description *string `json:"description,omitempty"`
	// PayDate is the date the voucher was paid. Setting this marks the voucher as paid.
	PayDate *time.Time `json:"payDate,omitempty"`
	// Currency is the ISO 4217 code (e.g. "EUR").
	Currency *string `json:"currency,omitempty"`
	// PaymentDeadline is the due date.
	PaymentDeadline *time.Time `json:"paymentDeadline,omitempty"`
	// DeliveryDate is when the goods/services were delivered.
	DeliveryDate *time.Time `json:"deliveryDate,omitempty"`
	// DeliveryDateUntil sets the end of a delivery period (use with DeliveryDate).
	DeliveryDateUntil *time.Time `json:"deliveryDateUntil,omitempty"`
	// TaxRule selects the tax rule (e.g. domestic, reverse charge).
	// Required by sevdesk for valid bookings.
	TaxRule *Ref `json:"taxRule,omitempty"`
	// CostCentre assigns the voucher to a cost centre.
	CostCentre *Ref `json:"costCentre,omitempty"`
	// Document links a previously-uploaded document.
	Document *Ref `json:"document,omitempty"`

	// PropertyExchangeRate sets the FX rate to your base currency. Only relevant
	// when Currency differs from your sevdesk account currency.
	PropertyExchangeRate *Decimal `json:"propertyExchangeRate,omitempty"`
	// PropertyForeignCurrencyDeadline is the date the FX rate is fixed at.
	PropertyForeignCurrencyDeadline *time.Time `json:"propertyForeignCurrencyDeadline,omitempty"`
}

// VoucherPosCreate is a single position to add to a voucher.
// ObjectName and MapAll are set automatically by [VouchersService.Create].
type VoucherPosCreate struct {
	ObjectName string `json:"objectName"` // set by Create
	MapAll     bool   `json:"mapAll"`     // set by Create

	// AccountingType assigns the position to a DATEV account.
	// Required. List options via [TODO once added] or sevdesk's UI.
	AccountingType *Ref `json:"accountingType"`
	// TaxRate is the VAT rate as a percentage (e.g. 19 for 19%).
	TaxRate Decimal `json:"taxRate"`
	// Net indicates whether SumNet/SumGross are net (true) or gross (false) amounts.
	Net bool `json:"net"`
	// IsAsset marks this position as a capitalizable asset (Anlagegut).
	IsAsset *bool `json:"isAsset,omitempty"`
	// SumNet is the net amount of the position.
	SumNet Decimal `json:"sumNet"`
	// SumGross is the gross amount including tax.
	SumGross Decimal `json:"sumGross"`
	// Comment is an optional note shown on the position.
	Comment *string `json:"comment,omitempty"`
}

// UpdateVoucherParams is the body for [VouchersService.Update].
// See [VoucherCreateFields] for field semantics.
type UpdateVoucherParams struct {
	VoucherDate       *time.Time     `json:"voucherDate,omitempty"`
	Supplier          *Ref           `json:"supplier,omitempty"`
	SupplierName      *string        `json:"supplierName,omitempty"`
	Description       *string        `json:"description,omitempty"`
	PayDate           *time.Time     `json:"payDate,omitempty"`
	Status            *VoucherStatus `json:"status,omitempty"`
	PaidAmount        *Decimal       `json:"paidAmount,omitempty"`
	TaxType           *string        `json:"taxType,omitempty"`
	TaxRule           *Ref           `json:"taxRule,omitempty"`
	CreditDebit       *CreditDebit   `json:"creditDebit,omitempty"`
	VoucherType       *VoucherType   `json:"voucherType,omitempty"`
	Currency          *string        `json:"currency,omitempty"`
	PaymentDeadline   *time.Time     `json:"paymentDeadline,omitempty"`
	DeliveryDate      *time.Time     `json:"deliveryDate,omitempty"`
	DeliveryDateUntil *time.Time     `json:"deliveryDateUntil,omitempty"`
	Document          *Ref           `json:"document,omitempty"`
	CostCentre        *Ref           `json:"costCentre,omitempty"`

	PropertyExchangeRate            *Decimal   `json:"propertyExchangeRate,omitempty"`
	PropertyForeignCurrencyDeadline *time.Time `json:"propertyForeignCurrencyDeadline,omitempty"`
}

// ListVouchersParams filters [VouchersService.List].
type ListVouchersParams struct {
	// Status filters to one status; 0 (the default) means any.
	// See [VoucherStatus].
	Status VoucherStatus
	// CreditDebit limits to one direction. See [CreditDebit].
	CreditDebit CreditDebit
	// DescriptionLike does a substring match on the voucher description.
	DescriptionLike string
	// StartDate excludes vouchers dated before this point.
	StartDate time.Time
	// EndDate excludes vouchers dated after this point.
	EndDate time.Time
	// Contact narrows to vouchers for one supplier/customer.
	Contact *Ref
}

func (p *ListVouchersParams) query() url.Values {
	if p == nil {
		return nil
	}
	q := url.Values{}
	if p.Status != 0 {
		q.Set("status", p.Status.String())
	}
	if p.CreditDebit != "" {
		q.Set("creditDebit", string(p.CreditDebit))
	}
	if p.DescriptionLike != "" {
		q.Set("descriptionLike", p.DescriptionLike)
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

// BookVoucherParams is the body for [VouchersService.Book].
type BookVoucherParams struct {
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

// BookingType is the type of a payment-booking entry, used by
// [VouchersService.Book], [InvoicesService.Book], and [CreditNotesService.Book].
type BookingType string

// BookingType values.
const (
	// BookTypeNormal is an ordinary payment.
	BookTypeNormal BookingType = "N"
	// BookTypeChargeback reverses a previous payment.
	BookTypeChargeback BookingType = "CB"
	// BookTypeCurrencyFluctuation adjusts a foreign-currency booking for FX changes.
	BookTypeCurrencyFluctuation BookingType = "CF"
	// BookTypeFreeSkonto records a granted cash discount (Skonto).
	BookTypeFreeSkonto BookingType = "FS"
	// BookTypeOverpayment marks a payment larger than the outstanding amount.
	BookTypeOverpayment BookingType = "O"
	// BookTypeOverFluctuation combines an overpayment with an FX adjustment.
	BookTypeOverFluctuation BookingType = "OF"
	// BookTypeMoneyTransferCancel records that a bank transfer was cancelled.
	BookTypeMoneyTransferCancel BookingType = "MTC"
)

// List returns vouchers matching the given filter.
func (s *VouchersService) List(ctx context.Context, opts *ListVouchersParams) ([]Voucher, error) {
	raw, err := s.c.do(ctx, http.MethodGet, "/Voucher", opts.query(), nil)
	if err != nil {
		return nil, err
	}
	return decodeList[Voucher](raw)
}

// Get returns the voucher with the given id.
func (s *VouchersService) Get(ctx context.Context, id ID) (*Voucher, error) {
	raw, err := s.c.do(ctx, http.MethodGet, fmt.Sprintf("/Voucher/%d", id), nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeObject[Voucher](raw)
}

// CreateVoucherResult is what [VouchersService.Create] returns: the created
// voucher and its persisted positions.
type CreateVoucherResult struct {
	Voucher   Voucher      `json:"voucher"`
	Positions []VoucherPos `json:"voucherPos"`
	Filename  string       `json:"filename,omitempty"`
}

// Create creates a voucher with its positions in one atomic call. If you have
// a file to attach, upload it first via [VouchersService.UploadFile] and pass
// the returned filename in params.Filename.
func (s *VouchersService) Create(ctx context.Context, params *CreateVoucherParams) (*CreateVoucherResult, error) {
	if params != nil && params.Voucher != nil {
		if params.Voucher.ObjectName == "" {
			params.Voucher.ObjectName = ObjectVoucher
		}
		params.Voucher.MapAll = true
	}
	for i := range params.Positions {
		if params.Positions[i].ObjectName == "" {
			params.Positions[i].ObjectName = ObjectVoucherPos
		}
		params.Positions[i].MapAll = true
	}
	raw, err := s.c.do(ctx, http.MethodPost, "/Voucher/Factory/saveVoucher", nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[CreateVoucherResult](raw)
}

// Update modifies a voucher.
func (s *VouchersService) Update(ctx context.Context, id ID, params *UpdateVoucherParams) (*Voucher, error) {
	raw, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/Voucher/%d", id), nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[Voucher](raw)
}

// Enshrine locks the voucher so it can no longer be edited (required for
// some compliance regimes).
func (s *VouchersService) Enshrine(ctx context.Context, id ID) error {
	_, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/Voucher/%d/enshrine", id), nil, nil)
	return err
}

// VoucherBookResult is the log entry returned by [VouchersService.Book].
type VoucherBookResult struct {
	ID          ID         `json:"id,omitempty"`
	ObjectName  string     `json:"objectName,omitempty"`
	Create      *time.Time `json:"create,omitempty"`
	Voucher     *Ref       `json:"voucher,omitempty"`
	FromStatus  string     `json:"fromStatus,omitempty"`
	ToStatus    string     `json:"toStatus,omitempty"`
	AmountPaid  Decimal    `json:"amountPayed,omitempty"` // sevdesk typo
	BookingDate *time.Time `json:"bookingDate,omitempty"`
	SevClient   *Ref       `json:"sevClient,omitempty"`
}

// Book records a payment against a voucher.
func (s *VouchersService) Book(ctx context.Context, id ID, params *BookVoucherParams) (*VoucherBookResult, error) {
	raw, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/Voucher/%d/bookAmount", id), nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[VoucherBookResult](raw)
}

// ResetToOpen reverts an enshrined or paid voucher back to open status.
func (s *VouchersService) ResetToOpen(ctx context.Context, id ID) (*Voucher, error) {
	raw, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/Voucher/%d/resetToOpen", id), nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeObject[Voucher](raw)
}

// ResetToDraft reverts a voucher back to draft status.
func (s *VouchersService) ResetToDraft(ctx context.Context, id ID) (*Voucher, error) {
	raw, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/Voucher/%d/resetToDraft", id), nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeObject[Voucher](raw)
}

// VoucherUpload describes a file uploaded via [VouchersService.UploadFile].
// Pass Filename back in [CreateVoucherParams.Filename] when creating the voucher.
type VoucherUpload struct {
	Pages          Num    `json:"pages,omitempty"`
	MimeType       string `json:"mimeType,omitempty"`
	OriginMimeType string `json:"originMimeType,omitempty"`
	Filename       string `json:"filename,omitempty"`
	ContentHash    string `json:"contentHash,omitempty"`
}

// UploadFile uploads a receipt/invoice file (PDF, image) to sevdesk's temp
// storage. The returned Filename is then passed to [VouchersService.Create].
func (s *VouchersService) UploadFile(ctx context.Context, filename string, file io.Reader) (*VoucherUpload, error) {
	var body bytes.Buffer
	mw := multipart.NewWriter(&body)
	part, err := mw.CreateFormFile("file", filename)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(part, file); err != nil {
		return nil, err
	}
	if err := mw.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.c.baseURL+"/Voucher/Factory/uploadTempFile", &body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", s.c.apiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", mw.FormDataContentType())
	if s.c.userAgent != "" {
		req.Header.Set("User-Agent", s.c.userAgent)
	}

	resp, err := s.c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sevdesk: upload: %w", err)
	}
	defer resp.Body.Close()
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, parseError(http.MethodPost, "/Voucher/Factory/uploadTempFile", resp.StatusCode, respBytes)
	}
	var env struct {
		Objects VoucherUpload `json:"objects"`
	}
	if err := json.Unmarshal(respBytes, &env); err != nil {
		return nil, err
	}
	return &env.Objects, nil
}

// VoucherPos is a single line on a voucher.
type VoucherPos struct {
	ID         ID         `json:"id,omitempty"`
	ObjectName string     `json:"objectName,omitempty"`
	Create     *time.Time `json:"create,omitempty"`
	Update     *time.Time `json:"update,omitempty"`
	SevClient  *Ref       `json:"sevClient,omitempty"`

	Voucher                 *Ref    `json:"voucher,omitempty"`
	AccountDatev            *Ref    `json:"accountDatev,omitempty"`
	AccountingType          *Ref    `json:"accountingType,omitempty"`
	EstimatedAccountingType *Ref    `json:"estimatedAccountingType,omitempty"`
	TaxRate                 Decimal `json:"taxRate,omitempty"`
	Net                     Bool    `json:"net,omitempty"`
	IsAsset                 Bool    `json:"isAsset,omitempty"`
	SumNet                  Decimal `json:"sumNet,omitempty"`
	SumTax                  Decimal `json:"sumTax,omitempty"`
	SumGross                Decimal `json:"sumGross,omitempty"`
	SumNetAccounting        Decimal `json:"sumNetAccounting,omitempty"`
	SumTaxAccounting        Decimal `json:"sumTaxAccounting,omitempty"`
	SumGrossAccounting      Decimal `json:"sumGrossAccounting,omitempty"`
	Comment                 string  `json:"comment,omitempty"`
}

// Positions returns the positions of the given voucher.
func (s *VouchersService) Positions(ctx context.Context, voucherID ID) ([]VoucherPos, error) {
	q := url.Values{
		"voucher[id]":         {voucherID.String()},
		"voucher[objectName]": {ObjectVoucher},
	}
	raw, err := s.c.do(ctx, http.MethodGet, "/VoucherPos", q, nil)
	if err != nil {
		return nil, err
	}
	return decodeList[VoucherPos](raw)
}
