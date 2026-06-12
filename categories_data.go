package sevdesk

// Category references for the system-managed [Category] catalogue.
// Drop directly into a [*Ref] field, e.g.
// [CreateContactParams.Category] or [ContactAddress.Category].
//
// Snapshot taken from sevdesk's /Category endpoint. IDs are permanent.

// Contact categories.
var (
	// CategoryContactSupplier — Lieferant.
	CategoryContactSupplier = &Ref{ID: 2, ObjectName: ObjectCategory}
	// CategoryContactCustomer — Kunde.
	CategoryContactCustomer = &Ref{ID: 3, ObjectName: ObjectCategory}
	// CategoryContactPartner — Partner.
	CategoryContactPartner = &Ref{ID: 4, ObjectName: ObjectCategory}
	// CategoryContactProspectCustomer — Interessent.
	CategoryContactProspectCustomer = &Ref{ID: 28, ObjectName: ObjectCategory}
	// CategoryContactPaymentProvider — Zahlungsdienstleister.
	CategoryContactPaymentProvider = &Ref{ID: 355842, ObjectName: ObjectCategory}
)

// ContactAddress categories.
var (
	// CategoryAddressWork — Arbeit.
	CategoryAddressWork = &Ref{ID: 43, ObjectName: ObjectCategory}
	// CategoryAddressPrivat — Privat.
	CategoryAddressPrivat = &Ref{ID: 44, ObjectName: ObjectCategory}
	// CategoryAddressEmpty.
	CategoryAddressEmpty = &Ref{ID: 45, ObjectName: ObjectCategory}
	// CategoryAddressInvoiceAddress — Rechnungsanschrift.
	CategoryAddressInvoiceAddress = &Ref{ID: 47, ObjectName: ObjectCategory}
	// CategoryAddressDeliveryAddress — Lieferanschrift.
	CategoryAddressDeliveryAddress = &Ref{ID: 48, ObjectName: ObjectCategory}
	// CategoryAddressPickupAddress — Abholanschrift.
	CategoryAddressPickupAddress = &Ref{ID: 49, ObjectName: ObjectCategory}
)

// Part categories.
var (
	// CategoryPartArticle — Standard.
	CategoryPartArticle = &Ref{ID: 1, ObjectName: ObjectCategory}
	// CategoryPartService — Dienstleistung.
	CategoryPartService = &Ref{ID: 97430, ObjectName: ObjectCategory}
)

// Document categories.
var (
	// CategoryDocumentOtherDocuments — Sonstige Dokumente.
	CategoryDocumentOtherDocuments = &Ref{ID: 16, ObjectName: ObjectCategory}
	// CategoryDocumentDeliveryNotes — Lieferscheine.
	CategoryDocumentDeliveryNotes = &Ref{ID: 53, ObjectName: ObjectCategory}
	// CategoryDocumentOrderConfirmation — Auftragsbestätigung.
	CategoryDocumentOrderConfirmation = &Ref{ID: 54, ObjectName: ObjectCategory}
	// CategoryDocumentQuickFile — Schnellablage.
	CategoryDocumentQuickFile = &Ref{ID: 56, ObjectName: ObjectCategory}
)

// Task categories.
var (
	// CategoryTaskPhoneCall — Anruf.
	CategoryTaskPhoneCall = &Ref{ID: 31, ObjectName: ObjectCategory}
	// CategoryTaskEmail — E-Mail.
	CategoryTaskEmail = &Ref{ID: 38, ObjectName: ObjectCategory}
	// CategoryTaskNone — Keine.
	CategoryTaskNone = &Ref{ID: 39, ObjectName: ObjectCategory}
	// CategoryTaskAppointement — Termin.
	CategoryTaskAppointement = &Ref{ID: 40, ObjectName: ObjectCategory}
	// CategoryTaskOffer — Angebot.
	CategoryTaskOffer = &Ref{ID: 41, ObjectName: ObjectCategory}
	// CategoryTaskFax — Fax.
	CategoryTaskFax = &Ref{ID: 42, ObjectName: ObjectCategory}
)

// Project categories.
var (
	// CategoryProjectNewProject — Neuprojekt.
	CategoryProjectNewProject = &Ref{ID: 8, ObjectName: ObjectCategory}
	// CategoryProjectExpansionProject — Erweiterungsprojekt.
	CategoryProjectExpansionProject = &Ref{ID: 9, ObjectName: ObjectCategory}
	// CategoryProjectGuaranteeProject — Garantieprojekt.
	CategoryProjectGuaranteeProject = &Ref{ID: 10, ObjectName: ObjectCategory}
)

// ProjectTime categories.
var (
	// CategoryProjectTimeRework — Nacharbeit.
	CategoryProjectTimeRework = &Ref{ID: 57, ObjectName: ObjectCategory}
)
