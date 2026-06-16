package sevdesk

import (
	"context"
	"encoding/json"
	"fmt"
	"iter"
	"net/http"
	"net/url"
	"time"
)

// ContactsService handles communication with /Contact endpoints.
type ContactsService struct {
	c *Client
}

// Contact is a sevdesk contact (organization or person).
type Contact struct {
	ID         ID         `json:"id,omitempty"`
	ObjectName string     `json:"objectName,omitempty"`
	Create     *time.Time `json:"create,omitempty"`
	Update     *time.Time `json:"update,omitempty"`
	SevClient  *Ref       `json:"sevClient,omitempty"`

	Name           string        `json:"name,omitempty"`
	Status         ContactStatus `json:"status,omitempty"` // see [ContactStatus]
	CustomerNumber string        `json:"customerNumber,omitempty"`
	Parent         *Ref          `json:"parent,omitempty"`
	Surename       string        `json:"surename,omitempty"`
	Familyname     string        `json:"familyname,omitempty"`
	Titel          string        `json:"titel,omitempty"`
	Category       *Ref          `json:"category,omitempty"`
	Description    string        `json:"description,omitempty"`
	AcademicTitle  string        `json:"academicTitle,omitempty"`
	Gender         string        `json:"gender,omitempty"`
	Name2          string        `json:"name2,omitempty"`
	Birthday       string        `json:"birthday,omitempty"` // YYYY-MM-DD
	VATNumber      string        `json:"vatNumber,omitempty"`
	BankAccount    string        `json:"bankAccount,omitempty"`
	BankNumber     string        `json:"bankNumber,omitempty"`

	DefaultCashbackTime       Num     `json:"defaultCashbackTime,omitempty"`
	DefaultCashbackPercent    Decimal `json:"defaultCashbackPercent,omitempty"`
	DefaultTimeToPay          Num     `json:"defaultTimeToPay,omitempty"`
	TaxNumber                 string  `json:"taxNumber,omitempty"`
	TaxOffice                 string  `json:"taxOffice,omitempty"`
	ExemptVAT                 Bool    `json:"exemptVat,omitempty"`
	TaxType                   string  `json:"taxType,omitempty"`
	DefaultDiscountAmount     Decimal `json:"defaultDiscountAmount,omitempty"`
	DefaultDiscountPercentage Bool    `json:"defaultDiscountPercentage,omitempty"`
	BuyerReference            string  `json:"buyerReference,omitempty"`
	GovernmentAgency          Bool    `json:"governmentAgency,omitempty"`
	DefaultShowVAT            Bool    `json:"defaultShowVat,omitempty"`
}

// ContactStatus is the lifecycle stage of a contact.
type ContactStatus Num

func (s ContactStatus) String() string { return Num(s).String() }

func (s *ContactStatus) UnmarshalJSON(b []byte) error {
	var n Num
	if err := n.UnmarshalJSON(b); err != nil {
		return err
	}
	*s = ContactStatus(n)
	return nil
}

// ContactStatus values.
const (
	// ContactStatusLead is a prospect not yet converted to a customer.
	ContactStatusLead ContactStatus = 100
	// ContactStatusPending is a contact in onboarding/verification.
	ContactStatusPending ContactStatus = 500
	// ContactStatusActive is a normal, fully active contact.
	ContactStatusActive ContactStatus = 1000
)

// CreateContactParams is the request body for [ContactsService.Create].
// Only Category is required by the API; all other fields are optional.
type CreateContactParams struct {
	// Category determines the contact kind. Common IDs: 3 (Customer),
	// 4 (Supplier), 2 (Partner). Use [CategoryRef] to build the reference.
	Category *Ref `json:"category"`

	// Name is the organization name. If set, the contact is treated as an
	// organization; for a person, leave empty and set Surename + Familyname.
	Name *string `json:"name,omitempty"`
	// Status controls the contact's lifecycle stage. See [ContactStatus].
	Status ContactStatus `json:"status,omitempty"`
	// CustomerNumber is the unique reference shown to the customer.
	// Auto-assigned if omitted.
	CustomerNumber *string `json:"customerNumber,omitempty"`
	// Parent makes this contact a sub-contact of another contact
	// (department, branch, employee of an organization).
	Parent *Ref `json:"parent,omitempty"`
	// Surename is the first name. Ignored when Name is set.
	Surename *string `json:"surename,omitempty"`
	// Familyname is the last name. Ignored when Name is set.
	Familyname *string `json:"familyname,omitempty"`
	// Titel is the job title (e.g. "CEO"); see AcademicTitle for academic ones.
	Titel *string `json:"titel,omitempty"`
	// Description is free-form internal notes.
	Description *string `json:"description,omitempty"`
	// AcademicTitle prefixes the name on documents (e.g. "Dr.", "Prof.").
	AcademicTitle *string `json:"academicTitle,omitempty"`
	// Gender for persons. Usually "m" or "f".
	Gender *string `json:"gender,omitempty"`
	// Name2 is an additional name line shown below Name on addresses.
	Name2 *string `json:"name2,omitempty"`
	// Birthday in YYYY-MM-DD format.
	Birthday *string `json:"birthday,omitempty"`
	// VATNumber is the contact's VAT ID (e.g. "DE114103514").
	VATNumber *string `json:"vatNumber,omitempty"`
	// BankAccount is the contact's IBAN.
	BankAccount *string `json:"bankAccount,omitempty"`
	// BankNumber is the bank routing code (BIC or Bankleitzahl).
	BankNumber *string `json:"bankNumber,omitempty"`

	// DefaultCashbackTime is the default cash-discount (Skonto) deadline in days.
	DefaultCashbackTime *int `json:"defaultCashbackTime,omitempty"`
	// DefaultCashbackPercent is the default cash-discount percentage.
	DefaultCashbackPercent *float64 `json:"defaultCashbackPercent,omitempty"`
	// DefaultTimeToPay is the default payment deadline in days from invoice date.
	DefaultTimeToPay *int `json:"defaultTimeToPay,omitempty"`
	// TaxNumber is the contact's domestic tax number (Steuernummer).
	TaxNumber *string `json:"taxNumber,omitempty"`
	// TaxOffice is the local tax office (Greek customers only).
	TaxOffice *string `json:"taxOffice,omitempty"`
	// ExemptVAT marks the contact as VAT-exempt. Use [True] or [False].
	ExemptVAT *Bool `json:"exemptVat,omitempty"`
	// DefaultDiscountAmount is the default discount applied to documents.
	DefaultDiscountAmount *float64 `json:"defaultDiscountAmount,omitempty"`
	// DefaultDiscountPercentage: True if DefaultDiscountAmount is a percentage,
	// False if it's an absolute amount.
	DefaultDiscountPercentage *Bool `json:"defaultDiscountPercentage,omitempty"`
	// BuyerReference is the buyer's reference for e-invoicing
	// (Leitweg-ID for German B2G).
	BuyerReference *string `json:"buyerReference,omitempty"`
	// GovernmentAgency marks the contact as a public-sector entity.
	GovernmentAgency *Bool `json:"governmentAgency,omitempty"`
}

// UpdateContactParams is the request body for [ContactsService.Update]. Any
// omitted field is left unchanged. See [CreateContactParams] for field semantics.
type UpdateContactParams struct {
	Name           *string       `json:"name,omitempty"`
	Status         ContactStatus `json:"status,omitempty"`
	CustomerNumber *string       `json:"customerNumber,omitempty"`
	Parent         *Ref          `json:"parent,omitempty"`
	Surename       *string       `json:"surename,omitempty"`
	Familyname     *string       `json:"familyname,omitempty"`
	Titel          *string       `json:"titel,omitempty"`
	Category       *Ref          `json:"category,omitempty"`
	Description    *string       `json:"description,omitempty"`
	AcademicTitle  *string       `json:"academicTitle,omitempty"`
	Gender         *string       `json:"gender,omitempty"`
	Name2          *string       `json:"name2,omitempty"`
	Birthday       *string       `json:"birthday,omitempty"`
	VATNumber      *string       `json:"vatNumber,omitempty"`
	BankAccount    *string       `json:"bankAccount,omitempty"`
	BankNumber     *string       `json:"bankNumber,omitempty"`

	DefaultCashbackTime       *int     `json:"defaultCashbackTime,omitempty"`
	DefaultCashbackPercent    *float64 `json:"defaultCashbackPercent,omitempty"`
	DefaultTimeToPay          *int     `json:"defaultTimeToPay,omitempty"`
	TaxNumber                 *string  `json:"taxNumber,omitempty"`
	TaxOffice                 *string  `json:"taxOffice,omitempty"`
	ExemptVAT                 *Bool    `json:"exemptVat,omitempty"`
	DefaultDiscountAmount     *float64 `json:"defaultDiscountAmount,omitempty"`
	DefaultDiscountPercentage *Bool    `json:"defaultDiscountPercentage,omitempty"`
	BuyerReference            *string  `json:"buyerReference,omitempty"`
	GovernmentAgency          *Bool    `json:"governmentAgency,omitempty"`
}

// SortDirection picks an ordering for List endpoints that support it.
type SortDirection string

// SortDirection values.
const (
	// SortAsc orders results ascending.
	SortAsc SortDirection = "ASC"
	// SortDesc orders results descending.
	SortDesc SortDirection = "DESC"
)

// ListContactsParams filters the result of [ContactsService.List].
type ListContactsParams struct {
	// IncludePersons, when true, returns both organizations AND persons.
	// When false (the default), sevdesk returns only organizations — set this
	// to true to also receive person-type contacts.
	IncludePersons bool
	// Category narrows results to contacts of a given category
	// (e.g. customer, supplier). Use [CategoryRef] to build it.
	Category *Ref
	// City matches contacts whose city equals this value.
	City string
	// Tags narrows results to contacts carrying ALL of the given tags.
	Tags []Ref
	// CustomerNumber matches contacts with this exact customer number.
	CustomerNumber string
	// Parent narrows to sub-contacts of the given parent organization.
	Parent *Ref
	// Name matches contacts whose name, surename, or familyname equals this value.
	Name string
	// Zip matches contacts with this postal code.
	Zip string
	// Country matches contacts in the given country. Use [CountryRef].
	Country *Ref
	// CreateBefore excludes contacts created at or after this point.
	CreateBefore time.Time
	// CreateAfter excludes contacts created at or before this point.
	CreateAfter time.Time
	// UpdateBefore excludes contacts last updated at or after this point.
	UpdateBefore time.Time
	// UpdateAfter excludes contacts last updated at or before this point.
	UpdateAfter time.Time
	// OrderByCustomerNumber sorts results by customer number ascending or
	// descending. Leave empty for the default order.
	OrderByCustomerNumber SortDirection
}

func (p *ListContactsParams) query() url.Values {
	if p == nil {
		return nil
	}
	q := url.Values{}
	if p.IncludePersons {
		q.Set("depth", "1")
	}
	if p.Category != nil {
		q.Set("category[id]", p.Category.ID.String())
		q.Set("category[objectName]", p.Category.ObjectName)
	}
	if p.City != "" {
		q.Set("city", p.City)
	}
	for i, t := range p.Tags {
		q.Set(fmt.Sprintf("tags[%d][id]", i), t.ID.String())
		q.Set(fmt.Sprintf("tags[%d][objectName]", i), t.ObjectName)
	}
	if p.CustomerNumber != "" {
		q.Set("customerNumber", p.CustomerNumber)
	}
	if p.Parent != nil {
		q.Set("parent[id]", p.Parent.ID.String())
		q.Set("parent[objectName]", p.Parent.ObjectName)
	}
	if p.Name != "" {
		q.Set("name", p.Name)
	}
	if p.Zip != "" {
		q.Set("zip", p.Zip)
	}
	if p.Country != nil {
		q.Set("country[id]", p.Country.ID.String())
		q.Set("country[objectName]", p.Country.ObjectName)
	}
	if !p.CreateBefore.IsZero() {
		q.Set("createBefore", fmt.Sprintf("%d", p.CreateBefore.Unix()))
	}
	if !p.CreateAfter.IsZero() {
		q.Set("createAfter", fmt.Sprintf("%d", p.CreateAfter.Unix()))
	}
	if !p.UpdateBefore.IsZero() {
		q.Set("updateBefore", fmt.Sprintf("%d", p.UpdateBefore.Unix()))
	}
	if !p.UpdateAfter.IsZero() {
		q.Set("updateAfter", fmt.Sprintf("%d", p.UpdateAfter.Unix()))
	}
	if p.OrderByCustomerNumber != "" {
		q.Set("orderByCustomerNumber", string(p.OrderByCustomerNumber))
	}
	return q
}

// List returns contacts matching the given filter.
func (s *ContactsService) List(ctx context.Context, opts *ListContactsParams) iter.Seq2[Contact, error] {
	return listIter[Contact](ctx, s.c, "/Contact", opts.query())
}

// Get returns the contact with the given id.
func (s *ContactsService) Get(ctx context.Context, id ID) (*Contact, error) {
	raw, err := s.c.do(ctx, http.MethodGet, fmt.Sprintf("/Contact/%d", id), nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeObject[Contact](raw)
}

// Create creates a new contact. Category is required.
func (s *ContactsService) Create(ctx context.Context, params *CreateContactParams) (*Contact, error) {
	raw, err := s.c.do(ctx, http.MethodPost, "/Contact", nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[Contact](raw)
}

// Update modifies an existing contact.
func (s *ContactsService) Update(ctx context.Context, id ID, params *UpdateContactParams) (*Contact, error) {
	raw, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/Contact/%d", id), nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[Contact](raw)
}

// Delete removes the contact with the given id.
func (s *ContactsService) Delete(ctx context.Context, id ID) error {
	_, err := s.c.do(ctx, http.MethodDelete, fmt.Sprintf("/Contact/%d", id), nil, nil)
	return err
}

// ContactItemCounts is the number of related entities a contact has
// (invoices, orders, vouchers, etc.).
type ContactItemCounts struct {
	Orders      Num `json:"orders"`
	Invoices    Num `json:"invoices"`
	CreditNotes Num `json:"creditNotes"`
	Documents   Num `json:"documents"`
	Persons     Num `json:"persons"`
	Vouchers    Num `json:"vouchers"`
	Letters     Num `json:"letters"`
	Parts       Num `json:"parts"`
	InvoicePos  Num `json:"invoicePos"`
}

// ItemCounts returns the count of orders, invoices, vouchers, etc. associated
// with the given contact.
func (s *ContactsService) ItemCounts(ctx context.Context, id ID) (*ContactItemCounts, error) {
	raw, err := s.c.do(ctx, http.MethodGet, fmt.Sprintf("/Contact/%d/getTabsItemCount", id), nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeObject[ContactItemCounts](raw)
}

// NextCustomerNumber returns the next free customer number suggested by sevdesk.
func (s *ContactsService) NextCustomerNumber(ctx context.Context) (string, error) {
	raw, err := s.c.do(ctx, http.MethodGet, "/Contact/Factory/getNextCustomerNumber", nil, nil)
	if err != nil {
		return "", err
	}
	if len(raw) == 0 {
		return "", nil
	}
	// sevdesk returns the number as a bare JSON string.
	var s2 string
	if err := json.Unmarshal(raw, &s2); err != nil {
		return "", err
	}
	return s2, nil
}

// CustomerNumberAvailable reports whether the given customer number is unused.
func (s *ContactsService) CustomerNumberAvailable(ctx context.Context, number string) (bool, error) {
	q := url.Values{"customerNumber": {number}}
	raw, err := s.c.do(ctx, http.MethodGet, "/Contact/Mapper/checkCustomerNumberAvailability", q, nil)
	if err != nil {
		return false, err
	}
	if len(raw) == 0 {
		return false, nil
	}
	return string(raw) == "true", nil
}

// FindByCustomFieldValue returns contacts whose custom field matches value.
// fieldSetting is the Ref to the ContactCustomFieldSetting; fieldName is
// optional and may be empty.
func (s *ContactsService) FindByCustomFieldValue(ctx context.Context, value string, fieldSetting Ref, fieldName string) iter.Seq2[Contact, error] {
	q := url.Values{
		"value":                          {value},
		"customFieldSetting[id]":         {fieldSetting.ID.String()},
		"customFieldSetting[objectName]": {fieldSetting.ObjectName},
	}
	if fieldName != "" {
		q.Set("customFieldName", fieldName)
	}
	return listIter[Contact](ctx, s.c, "/Contact/Factory/findContactsByCustomFieldValue", q)
}
