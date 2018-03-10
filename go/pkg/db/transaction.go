package db

import (
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"

	"github.com/stellar-map/stellar-map/go/pkg/entities"
)

type transaction struct {
	ID              int32 `gorm:"primary_key"`
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
	Signatures      pq.StringArray
	ValidBefore     time.Time
	ValidAfter      time.Time
	CreatedAt       time.Time
}

func (db *db) BatchCreateTransactions(transactions []*entities.Transaction) error {
	records := make([]*transaction, len(transactions))
	for i, transaction := range transactions {
		records[i] = transactionFromEntity(transaction)
	}

	query := squirrel.
		Insert("transactions").
		Columns(
			"transaction_hash", "ledger_sequence", "ledger_closed_at", "paging_token",
			"account", "account_sequence", "fee_paid", "operation_count", "envelope_xdr",
			"result_xdr", "result_meta_xdr", "fee_meta_xdr", "memo_type", "memo",
			"signatures", "valid_before", "valid_after",
		)

	for _, record := range records {
		query = query.Values(
			record.TransactionHash,
			record.LedgerSequence,
			record.LedgerClosedAt,
			record.PagingToken,
			record.Account,
			record.AccountSequence,
			record.FeePaid,
			record.OperationCount,
			record.EnvelopeXDR,
			record.ResultXDR,
			record.ResultMetaXDR,
			record.FeeMetaXDR,
			record.MemoType,
			record.Memo,
			record.Signatures,
			record.ValidBefore,
			record.ValidAfter,
		)
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	// Ignore rows we've already ingested.
	sql += " ON CONFLICT (transaction_hash) DO NOTHING"

	return db.Exec(sql, args...).Error
}

func transactionFromEntity(entity *entities.Transaction) *transaction {
	return &transaction{
		TransactionHash: entity.TransactionHash,
		LedgerSequence:  entity.LedgerSequence,
		LedgerClosedAt:  entity.LedgerClosedAt,
		PagingToken:     entity.PagingToken,
		Account:         entity.Account,
		AccountSequence: entity.AccountSequence,
		FeePaid:         entity.FeePaid,
		OperationCount:  entity.OperationCount,
		EnvelopeXDR:     entity.EnvelopeXDR,
		ResultXDR:       entity.ResultXDR,
		ResultMetaXDR:   entity.ResultMetaXDR,
		FeeMetaXDR:      entity.FeeMetaXDR,
		MemoType:        entity.MemoType,
		Memo:            entity.Memo,
		Signatures:      entity.Signatures,
		ValidBefore:     entity.ValidBefore,
		ValidAfter:      entity.ValidAfter,
	}
}

func transactionToEntity(record *transaction) *entities.Transaction {
	return &entities.Transaction{
		TransactionHash: record.TransactionHash,
		LedgerSequence:  record.LedgerSequence,
		LedgerClosedAt:  record.LedgerClosedAt,
		PagingToken:     record.PagingToken,
		Account:         record.Account,
		AccountSequence: record.AccountSequence,
		FeePaid:         record.FeePaid,
		OperationCount:  record.OperationCount,
		EnvelopeXDR:     record.EnvelopeXDR,
		ResultXDR:       record.ResultXDR,
		ResultMetaXDR:   record.ResultMetaXDR,
		FeeMetaXDR:      record.FeeMetaXDR,
		MemoType:        record.MemoType,
		Memo:            record.Memo,
		Signatures:      record.Signatures,
		ValidBefore:     record.ValidBefore,
		ValidAfter:      record.ValidAfter,
	}
}
