package sevdesk

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// CheckAccountsService handles communication with /CheckAccount endpoints.
type CheckAccountsService struct {
	c *Client
}

// CheckAccount is a sevdesk payment account (bank account, cash, PayPal, etc.).
type CheckAccount struct {
	ID         ID         `json:"id,omitempty"`
	ObjectName string     `json:"objectName,omitempty"`
	Create     *time.Time `json:"create,omitempty"`
	Update     *time.Time `json:"update,omitempty"`
	SevClient  *Ref       `json:"sevClient,omitempty"`

	Name             string                 `json:"name,omitempty"`
	Type             CheckAccountType       `json:"type,omitempty"`
	ImportType       CheckAccountImportType `json:"importType,omitempty"`
	Currency         string                 `json:"currency,omitempty"`
	CheckAccID       string                 `json:"checkAccId,omitempty"`
	PIN              string                 `json:"pin,omitempty"`
	DefaultAccount   Bool                   `json:"defaultAccount,omitempty"`
	Status           CheckAccountStatus     `json:"status,omitempty"` // see [CheckAccountStatus]
	TranslationCode  string                 `json:"translationCode,omitempty"`
	BankServer       string                 `json:"bankServer,omitempty"`
	Balance          Decimal                `json:"balance,omitempty"`
	AccountingNumber string                 `json:"accountingNumber,omitempty"`
	IBAN             string                 `json:"iban,omitempty"`
	BIC              string                 `json:"bic,omitempty"`
	BaseAccount      Bool                   `json:"baseAccount,omitempty"`
	Priority         Num                    `json:"priority,omitempty"`
	CountryCode      string                 `json:"countryCode,omitempty"`

	AutoMapTransaction   Bool       `json:"autoMapTransaction,omitempty"`
	AutoSyncTransactions Bool       `json:"autoSyncTransactions,omitempty"`
	LastSync             *time.Time `json:"lastSync,omitempty"`
}

// CheckAccountType is the kind of payment account.
type CheckAccountType string

// CheckAccountType values.
const (
	// CheckAccountTypeOnline is a bank-linked account that syncs transactions
	// from the bank (typically via finAPI).
	CheckAccountTypeOnline CheckAccountType = "online"
	// CheckAccountTypeOffline is a clearing account with no automatic sync —
	// entries are made manually or via file import.
	CheckAccountTypeOffline CheckAccountType = "offline"
	// CheckAccountTypeRegister is a cash register account.
	CheckAccountTypeRegister CheckAccountType = "register"
)

// CheckAccountStatus is the lifecycle state of a payment account.
type CheckAccountStatus Num

func (s CheckAccountStatus) String() string { return Num(s).String() }

func (s *CheckAccountStatus) UnmarshalJSON(b []byte) error {
	var n Num
	if err := n.UnmarshalJSON(b); err != nil {
		return err
	}
	*s = CheckAccountStatus(n)
	return nil
}

// CheckAccountStatus values.
const (
	// CheckAccountStatusArchived means the account is hidden and not used
	// for new bookings.
	CheckAccountStatusArchived CheckAccountStatus = 0
	// CheckAccountStatusActive means the account is in active use.
	CheckAccountStatusActive CheckAccountStatus = 100
)

// CheckAccountImportType is the file format used by a file-import account.
type CheckAccountImportType string

// CheckAccountImportType values.
const (
	// CheckAccountImportCSV expects bank statements as comma-separated values.
	CheckAccountImportCSV CheckAccountImportType = "CSV"
	// CheckAccountImportMT940 expects bank statements in SWIFT MT940 format.
	CheckAccountImportMT940 CheckAccountImportType = "MT940"
)

// CreateFileImportAccountParams is the body for [CheckAccountsService.CreateFileImport].
type CreateFileImportAccountParams struct {
	// Name is the display name of the account. Required.
	Name string `json:"name"`
	// ImportType picks the import format. Required. See [CheckAccountImportType].
	ImportType CheckAccountImportType `json:"importType"`
	// AccountingNumber is the DATEV bank account number used for bookings.
	AccountingNumber *int `json:"accountingNumber,omitempty"`
	// IBAN of the real bank account behind this file-import slot.
	IBAN *string `json:"iban,omitempty"`
}

// CreateClearingAccountParams is the body for [CheckAccountsService.CreateClearing].
type CreateClearingAccountParams struct {
	// Name is the display name of the clearing account. Required.
	Name string `json:"name"`
	// AccountingNumber is the DATEV account used when posting through this clearing.
	AccountingNumber *int `json:"accountingNumber,omitempty"`
}

// UpdateCheckAccountParams is the body for [CheckAccountsService.Update].
type UpdateCheckAccountParams struct {
	// Name renames the account.
	Name *string `json:"name,omitempty"`
	// DefaultAccount sets/unsets this account as the default for new bookings.
	// Use [True] or [False].
	DefaultAccount *Bool `json:"defaultAccount,omitempty"`
	// AutoMapTransaction enables sevdesk's auto-matching of incoming transactions
	// to invoices/vouchers.
	AutoMapTransaction *Bool `json:"autoMapTransaction,omitempty"`
	// AccountingNumber is the DATEV account number.
	AccountingNumber *string `json:"accountingNumber,omitempty"`
	// IBAN of the underlying bank account.
	IBAN *string `json:"iban,omitempty"`
	// BIC of the underlying bank.
	BIC *string `json:"bic,omitempty"`
}

// List returns all check accounts.
func (s *CheckAccountsService) List(ctx context.Context) ([]CheckAccount, error) {
	raw, err := s.c.do(ctx, http.MethodGet, "/CheckAccount", nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeList[CheckAccount](raw)
}

// Get returns the check account with the given id.
func (s *CheckAccountsService) Get(ctx context.Context, id ID) (*CheckAccount, error) {
	raw, err := s.c.do(ctx, http.MethodGet, fmt.Sprintf("/CheckAccount/%d", id), nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeObject[CheckAccount](raw)
}

// Update modifies a check account.
func (s *CheckAccountsService) Update(ctx context.Context, id ID, params *UpdateCheckAccountParams) (*CheckAccount, error) {
	raw, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/CheckAccount/%d", id), nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[CheckAccount](raw)
}

// Delete removes a check account.
func (s *CheckAccountsService) Delete(ctx context.Context, id ID) error {
	_, err := s.c.do(ctx, http.MethodDelete, fmt.Sprintf("/CheckAccount/%d", id), nil, nil)
	return err
}

// CreateFileImport creates a new bank account for CSV or MT940 file imports.
func (s *CheckAccountsService) CreateFileImport(ctx context.Context, params *CreateFileImportAccountParams) (*CheckAccount, error) {
	raw, err := s.c.do(ctx, http.MethodPost, "/CheckAccount/Factory/fileImportAccount", nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[CheckAccount](raw)
}

// CreateClearing creates a new clearing account (offline, e.g. for cash flow or PayPal).
func (s *CheckAccountsService) CreateClearing(ctx context.Context, params *CreateClearingAccountParams) (*CheckAccount, error) {
	raw, err := s.c.do(ctx, http.MethodPost, "/CheckAccount/Factory/clearingAccount", nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[CheckAccount](raw)
}

// BalanceAt returns the account balance as of the given date (end of day).
func (s *CheckAccountsService) BalanceAt(ctx context.Context, id ID, date time.Time) (Decimal, error) {
	q := url.Values{"date": {date.Format("2006-01-02")}}
	raw, err := s.c.do(ctx, http.MethodGet, fmt.Sprintf("/CheckAccount/%d/getBalanceAtDate", id), q, nil)
	if err != nil {
		return 0, err
	}
	var bal Decimal
	if err := bal.UnmarshalJSON(raw); err != nil {
		return 0, err
	}
	return bal, nil
}
