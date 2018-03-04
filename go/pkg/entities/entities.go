package entities

import "time"

// Ledger references information stored within a single ledger within
// the Stellar blockchain.
type Ledger struct {
	ID                    string
	PagingToken           string
	Hash                  string
	PrevHash              string
	Sequence              int32
	TransactionCount      int32
	OperationCount        int32
	ClosedAt              time.Time
	TotalCoins            string
	FeePool               string
	BaseFee               int32
	BaseReserve           int32
	MaxTransactionSetSize int32
	ProtocolVersion       int32
}

// Transaction represents a single transction on the Stellar blockchain.
type Transaction struct {
	ID              string
	PagingToken     string
	Hash            string
	Ledger          int32
	LedgerClosedAt  time.Time
	Account         string
	AccountSequence string
	FeePaid         int32
	OperationCount  int32
	EnvelopeXdr     string
	ResultXdr       string
	ResultMetaXdr   string
	FeeMetaXdr      string
	MemoType        string
	Memo            string
	Signatures      []string
	ValidAfter      string
	ValidBefore     string
}

// Repo is the data access layer to save and fetch data from storage.
type Repo interface {
	BatchCreateLedgers(ledgers []*Ledger) error
}
