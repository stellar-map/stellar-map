package entities

import (
	"errors"
	"time"
)

var (
	// ErrNotFound is returned by the Repo when the requested resource cannot
	// be found.
	ErrNotFound = errors.New("resource not found")
)

// Ledger references information stored within a single ledger within
// the Stellar blockchain.
type Ledger struct {
	Sequence           int32
	LedgerHash         string
	PreviousLedgerHash string
	PagingToken        string
	TransactionCount   int32
	OperationCount     int32
	ClosedAt           time.Time
	TotalCoins         int64
	FeePool            int64
	BaseFee            int32
	BaseReserve        int32
	MaxTxSetSize       int32
	ProtocolVersion    int32
}

// Transaction represents a single transction on the Stellar blockchain.
type Transaction struct {
	TransactionHash string
	LedgerSequence  int32
	LedgerClosedAt  time.Time
	PagingToken     string
	Account         string
	AccountSequence string
	FeePaid         int32
	OperationCount  int32
	EnvelopeXDR     string
	ResultXDR       string
	ResultMetaXDR   string
	FeeMetaXDR      string
	MemoType        string
	Memo            string
	Signatures      []string
	ValidBefore     time.Time
	ValidAfter      time.Time
}

// Repo is the data access layer to save and fetch data from storage.
type Repo interface {
	BatchCreateLedgers(ledgers []*Ledger) error
	BatchCreateTransactions(transactions []*Transaction) error
	PagingToken(resource string) (string, error)
}
