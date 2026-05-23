package sevdesk

// Ref is sevdesk's universal foreign-key shape: every related entity in the API
// is referenced by an {id, objectName} pair (e.g. `{"id": 5, "objectName": "Contact"}`).
type Ref struct {
	ID         ID     `json:"id"`
	ObjectName string `json:"objectName"`
}

// Object names used across the sevdesk API. Use these instead of string literals.
const (
	ObjectContact                 = "Contact"
	ObjectContactAddress          = "ContactAddress"
	ObjectContactCustomField      = "ContactCustomField"
	ObjectContactCustomFieldSet   = "ContactCustomFieldSetting"
	ObjectCommunicationWay        = "CommunicationWay"
	ObjectCommunicationWayKey     = "CommunicationWayKey"
	ObjectAccountingContact       = "AccountingContact"
	ObjectCheckAccount            = "CheckAccount"
	ObjectCheckAccountTransaction = "CheckAccountTransaction"
	ObjectPrivateTransactionRule  = "PrivateTransactionRule"
	ObjectCreditNote              = "CreditNote"
	ObjectCreditNotePos           = "CreditNotePos"
	ObjectInvoice                 = "Invoice"
	ObjectInvoicePos              = "InvoicePos"
	ObjectOrder                   = "Order"
	ObjectOrderPos                = "OrderPos"
	ObjectPart                    = "Part"
	ObjectVoucher                 = "Voucher"
	ObjectVoucherPos              = "VoucherPos"
	ObjectTag                     = "Tag"
	ObjectTagRelation             = "TagRelation"
	ObjectCategory                = "Category"
	ObjectUnity                   = "Unity"
	ObjectSevUser                 = "SevUser"
	ObjectSevClient               = "SevClient"
	ObjectStaticCountry           = "StaticCountry"
	ObjectTaxSet                  = "TaxSet"
	ObjectTaxRule                 = "TaxRule"
	ObjectPaymentMethod           = "PaymentMethod"
	ObjectCostCentre              = "CostCentre"
	ObjectDocument                = "Document"
	ObjectAccountDatev            = "AccountDatev"
	ObjectAccountingType          = "AccountingType"
	ObjectDiscounts               = "Discounts"
)

// ContactRef returns a Ref for a Contact with the given id.
func ContactRef(id ID) *Ref { return &Ref{id, ObjectContact} }

// InvoiceRef returns a Ref for an Invoice with the given id.
func InvoiceRef(id ID) *Ref { return &Ref{id, ObjectInvoice} }

// OrderRef returns a Ref for an Order with the given id.
func OrderRef(id ID) *Ref { return &Ref{id, ObjectOrder} }

// VoucherRef returns a Ref for a Voucher with the given id.
func VoucherRef(id ID) *Ref { return &Ref{id, ObjectVoucher} }

// CreditNoteRef returns a Ref for a CreditNote with the given id.
func CreditNoteRef(id ID) *Ref { return &Ref{id, ObjectCreditNote} }

// PartRef returns a Ref for a Part with the given id.
func PartRef(id ID) *Ref { return &Ref{id, ObjectPart} }

// CheckAccountRef returns a Ref for a CheckAccount with the given id.
func CheckAccountRef(id ID) *Ref { return &Ref{id, ObjectCheckAccount} }

// CheckAccountTransactionRef returns a Ref for a CheckAccountTransaction.
func CheckAccountTransactionRef(id ID) *Ref { return &Ref{id, ObjectCheckAccountTransaction} }

// CategoryRef returns a Ref for a Category with the given id.
func CategoryRef(id ID) *Ref { return &Ref{id, ObjectCategory} }

// UnityRef returns a Ref for a Unity with the given id.
func UnityRef(id ID) *Ref { return &Ref{id, ObjectUnity} }

// CountryRef returns a Ref for a StaticCountry with the given id.
func CountryRef(id ID) *Ref { return &Ref{id, ObjectStaticCountry} }

// UserRef returns a Ref for a SevUser with the given id.
func UserRef(id ID) *Ref { return &Ref{id, ObjectSevUser} }
