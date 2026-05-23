package sevdesk

import "net/http"

const (
	defaultBaseURL   = "https://my.sevdesk.de/api/v1"
	defaultUserAgent = "sevdesk-go/0.1.0"
)

// Client is the entry point to the sevdesk API. Construct one with [New], then
// reach resources through the service fields:
//
//	c := sevdesk.New("API_KEY")
//	c.Contacts.Get(ctx, 42)
//	c.Invoices.List(ctx, nil)
type Client struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string
	userAgent  string

	Contacts                   *ContactsService
	ContactAddresses           *ContactAddressesService
	CommunicationWays          *CommunicationWaysService
	AccountingContacts         *AccountingContactsService
	ContactCustomFields        *ContactCustomFieldsService
	ContactCustomFieldSettings *ContactCustomFieldSettingsService
	Vouchers                   *VouchersService
	CheckAccounts              *CheckAccountsService
	Transactions               *TransactionsService
	Invoices                   *InvoicesService
	CreditNotes                *CreditNotesService
	Orders                     *OrdersService
	OrderPositions             *OrderPositionsService
	Parts                      *PartsService
	Tags                       *TagsService
	PrivateTransactionRules    *PrivateTransactionRulesService
}

// New returns a Client authenticated with the given API key.
func New(apiKey string, opts ...Option) *Client {
	c := &Client{
		httpClient: http.DefaultClient,
		baseURL:    defaultBaseURL,
		apiKey:     apiKey,
		userAgent:  defaultUserAgent,
	}
	for _, opt := range opts {
		opt(c)
	}
	c.Contacts = &ContactsService{c: c}
	c.ContactAddresses = &ContactAddressesService{c: c}
	c.CommunicationWays = &CommunicationWaysService{c: c}
	c.AccountingContacts = &AccountingContactsService{c: c}
	c.ContactCustomFields = &ContactCustomFieldsService{c: c}
	c.ContactCustomFieldSettings = &ContactCustomFieldSettingsService{c: c}
	c.Vouchers = &VouchersService{c: c}
	c.CheckAccounts = &CheckAccountsService{c: c}
	c.Transactions = &TransactionsService{c: c}
	c.Invoices = &InvoicesService{c: c}
	c.CreditNotes = &CreditNotesService{c: c}
	c.Orders = &OrdersService{c: c}
	c.OrderPositions = &OrderPositionsService{c: c}
	c.Parts = &PartsService{c: c}
	c.Tags = &TagsService{c: c}
	c.PrivateTransactionRules = &PrivateTransactionRulesService{c: c}
	return c
}
