package sevdesk

// CommunicationWayKey references — labels you can attach to a
// [CommunicationWay] (e.g. "Work", "Private"). Drop directly into the Key
// field of [CreateCommunicationWayParams] or [UpdateCommunicationWayParams].
//
// Snapshot taken from sevdesk's /CommunicationWayKey endpoint. IDs are permanent.
var (
	// CommunicationWayKeyPrivate — Privat.
	CommunicationWayKeyPrivate = &Ref{ID: 1, ObjectName: ObjectCommunicationWayKey}
	// CommunicationWayKeyWork — Arbeit.
	CommunicationWayKeyWork = &Ref{ID: 2, ObjectName: ObjectCommunicationWayKey}
	// CommunicationWayKeyFax — Fax.
	CommunicationWayKeyFax = &Ref{ID: 3, ObjectName: ObjectCommunicationWayKey}
	// CommunicationWayKeyMobile — Mobil.
	CommunicationWayKeyMobile = &Ref{ID: 4, ObjectName: ObjectCommunicationWayKey}
	// CommunicationWayKeyEmpty.
	CommunicationWayKeyEmpty = &Ref{ID: 5, ObjectName: ObjectCommunicationWayKey}
	// CommunicationWayKeyAutobox — Autobox.
	CommunicationWayKeyAutobox = &Ref{ID: 6, ObjectName: ObjectCommunicationWayKey}
	// CommunicationWayKeyNewsletter — Newsletter.
	CommunicationWayKeyNewsletter = &Ref{ID: 7, ObjectName: ObjectCommunicationWayKey}
	// CommunicationWayKeyInvoiceAddress — Rechnungsadresse.
	CommunicationWayKeyInvoiceAddress = &Ref{ID: 8, ObjectName: ObjectCommunicationWayKey}
)
