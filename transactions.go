package sevdesk

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// TransactionsService handles communication with /CheckAccountTransaction endpoints.
type TransactionsService struct {
	c *Client
}

// Transaction is a single movement on a check account.
type Transaction struct {
	ID         ID         `json:"id,omitempty"`
	ObjectName string     `json:"objectName,omitempty"`
	Create     *time.Time `json:"create,omitempty"`
	Update     *time.Time `json:"update,omitempty"`
	SevClient  *Ref       `json:"sevClient,omitempty"`

	ValueDate          *time.Time        `json:"valueDate,omitempty"`
	EntryDate          *time.Time        `json:"entryDate,omitempty"`
	Amount             Decimal           `json:"amount,omitempty"`
	FeeAmount          Decimal           `json:"feeAmount,omitempty"`
	GvCode             string            `json:"gvCode,omitempty"`
	EntryText          string            `json:"entryText,omitempty"`
	PrimaNotaNo        string            `json:"primaNotaNo,omitempty"`
	PaymtPurpose       string            `json:"paymtPurpose,omitempty"`
	PayeePayerBankCode string            `json:"payeePayerBankCode,omitempty"`
	PayeePayerAcctNo   string            `json:"payeePayerAcctNo,omitempty"`
	PayeePayerName     string            `json:"payeePayerName,omitempty"`
	CheckAccount       *Ref              `json:"checkAccount,omitempty"`
	Status             TransactionStatus `json:"status,omitempty"` // see [TransactionStatus]
	Score              Decimal           `json:"score,omitempty"`
	CompareHash        string            `json:"compareHash,omitempty"`
	Enshrined          *time.Time        `json:"enshrined,omitempty"`
	ObonoReceiptUUID   string            `json:"obonoReceiptUuid,omitempty"`
	ExternalID         string            `json:"externalId,omitempty"`
	DeletedAt          *time.Time        `json:"deletedAt,omitempty"`
	RestoredAt         *time.Time        `json:"restoredAt,omitempty"`

	SourceTransaction *Ref `json:"sourceTransaction,omitempty"`
	TargetTransaction *Ref `json:"targetTransaction,omitempty"`
}

// TransactionStatus is the lifecycle state of a check-account transaction.
type TransactionStatus Num

func (s TransactionStatus) String() string { return Num(s).String() }

func (s *TransactionStatus) UnmarshalJSON(b []byte) error {
	var n Num
	if err := n.UnmarshalJSON(b); err != nil {
		return err
	}
	*s = TransactionStatus(n)
	return nil
}

// TransactionStatus values.
const (
	// TransactionStatusCreated has just been imported and is not yet linked
	// to any document.
	TransactionStatusCreated TransactionStatus = 100
	// TransactionStatusLinkedAuto was auto-matched to an invoice or voucher
	// but the user hasn't confirmed the booking.
	TransactionStatusLinkedAuto TransactionStatus = 200
	// TransactionStatusPrivate is marked private (excluded from accounting).
	TransactionStatusPrivate TransactionStatus = 300
	// TransactionStatusAutoBooked was auto-matched and booked without manual
	// review.
	TransactionStatusAutoBooked TransactionStatus = 350
	// TransactionStatusBooked is manually booked.
	TransactionStatusBooked TransactionStatus = 400
)

// CreateTransactionParams is the body for [TransactionsService.Create].
type CreateTransactionParams struct {
	// ValueDate is the date the money was credited/debited. Required.
	ValueDate *time.Time `json:"valueDate"`
	// EntryDate is when sevdesk recorded the transaction (defaults to today).
	EntryDate *time.Time `json:"entryDate,omitempty"`
	// Amount of the transaction. Positive for credits, negative for debits. Required.
	Amount Decimal `json:"amount"`
	// PaymtPurpose is the bank statement's payment purpose line.
	PaymtPurpose *string `json:"paymtPurpose,omitempty"`
	// PayeePayerName is the counterparty name.
	PayeePayerName *string `json:"payeePayerName,omitempty"`
	// CheckAccount the transaction belongs to. Required.
	CheckAccount *Ref `json:"checkAccount"`
	// Status of the transaction. See [TransactionStatusCreated] and adjacent.
	Status *TransactionStatus `json:"status,omitempty"`
}

// UpdateTransactionParams is the body for [TransactionsService.Update].
// See [CreateTransactionParams] for field semantics.
type UpdateTransactionParams struct {
	ValueDate      *time.Time         `json:"valueDate,omitempty"`
	EntryDate      *time.Time         `json:"entryDate,omitempty"`
	Amount         *Decimal           `json:"amount,omitempty"`
	PaymtPurpose   *string            `json:"paymtPurpose,omitempty"`
	PayeePayerName *string            `json:"payeePayerName,omitempty"`
	CheckAccount   *Ref               `json:"checkAccount,omitempty"`
	Status         *TransactionStatus `json:"status,omitempty"`
	// SourceTransaction links this transaction as the source of a rebooking.
	SourceTransaction *Ref `json:"sourceTransaction,omitempty"`
	// TargetTransaction links this transaction as the target of a rebooking.
	TargetTransaction *Ref `json:"targetTransaction,omitempty"`
}

// ListTransactionsParams filters [TransactionsService.List].
type ListTransactionsParams struct {
	// CheckAccount narrows the results to one account.
	CheckAccount *Ref
	// IsBooked, when set, filters to booked (true) or unbooked (false) transactions.
	IsBooked *bool
	// PaymtPurpose does a substring match on the payment purpose.
	PaymtPurpose string
	// StartDate excludes transactions before this point.
	StartDate time.Time
	// EndDate excludes transactions after this point.
	EndDate time.Time
	// PayeePayerName does a substring match on the counterparty.
	PayeePayerName string
	// OnlyCredit, when true, limits the result to credits (positive amounts).
	OnlyCredit *bool
	// OnlyDebit, when true, limits the result to debits (negative amounts).
	OnlyDebit *bool
}

func (p *ListTransactionsParams) query() url.Values {
	if p == nil {
		return nil
	}
	q := url.Values{}
	if p.CheckAccount != nil {
		q.Set("checkAccount[id]", p.CheckAccount.ID.String())
		q.Set("checkAccount[objectName]", p.CheckAccount.ObjectName)
	}
	if p.IsBooked != nil {
		q.Set("isBooked", boolStr(*p.IsBooked))
	}
	if p.PaymtPurpose != "" {
		q.Set("paymtPurpose", p.PaymtPurpose)
	}
	if !p.StartDate.IsZero() {
		q.Set("startDate", p.StartDate.Format(time.RFC3339))
	}
	if !p.EndDate.IsZero() {
		q.Set("endDate", p.EndDate.Format(time.RFC3339))
	}
	if p.PayeePayerName != "" {
		q.Set("payeePayerName", p.PayeePayerName)
	}
	if p.OnlyCredit != nil {
		q.Set("onlyCredit", boolStr(*p.OnlyCredit))
	}
	if p.OnlyDebit != nil {
		q.Set("onlyDebit", boolStr(*p.OnlyDebit))
	}
	return q
}

func boolStr(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

// List returns transactions matching the given filter.
func (s *TransactionsService) List(ctx context.Context, opts *ListTransactionsParams) ([]Transaction, error) {
	raw, err := s.c.do(ctx, http.MethodGet, "/CheckAccountTransaction", opts.query(), nil)
	if err != nil {
		return nil, err
	}
	return decodeList[Transaction](raw)
}

// Get returns the transaction with the given id.
func (s *TransactionsService) Get(ctx context.Context, id ID) (*Transaction, error) {
	raw, err := s.c.do(ctx, http.MethodGet, fmt.Sprintf("/CheckAccountTransaction/%d", id), nil, nil)
	if err != nil {
		return nil, err
	}
	return decodeObject[Transaction](raw)
}

// Create creates a new transaction.
func (s *TransactionsService) Create(ctx context.Context, params *CreateTransactionParams) (*Transaction, error) {
	raw, err := s.c.do(ctx, http.MethodPost, "/CheckAccountTransaction", nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[Transaction](raw)
}

// Update modifies a transaction.
func (s *TransactionsService) Update(ctx context.Context, id ID, params *UpdateTransactionParams) (*Transaction, error) {
	raw, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/CheckAccountTransaction/%d", id), nil, params)
	if err != nil {
		return nil, err
	}
	return decodeObject[Transaction](raw)
}

// Delete removes a transaction.
func (s *TransactionsService) Delete(ctx context.Context, id ID) error {
	_, err := s.c.do(ctx, http.MethodDelete, fmt.Sprintf("/CheckAccountTransaction/%d", id), nil, nil)
	return err
}

// Enshrine locks the transaction so it can no longer be edited.
func (s *TransactionsService) Enshrine(ctx context.Context, id ID) error {
	_, err := s.c.do(ctx, http.MethodPut, fmt.Sprintf("/CheckAccountTransaction/%d/enshrine", id), nil, nil)
	return err
}
