package db

import (
	"time"

	"github.com/Masterminds/squirrel"

	"github.com/stellar-map/stellar-map/go/pkg/entities"
)

type ledger struct {
	Sequence              int32
	LedgerHash            string
	PreviousLedgerHash    string
	PagingToken           string
	TransactionCount      int32
	OperationCount        int32
	ClosedAt              time.Time
	TotalCoins            string
	FeePool               string
	BaseFee               int32
	BaseReserve           int32
	MaxTransactionSetSize int32
	ProtocolVersion       int32
	CreatedAt             time.Time
}

func (db *db) BatchCreateLedgers(ledgers []*entities.Ledger) error {
	records := make([]*ledger, len(ledgers))
	for i, ledger := range ledgers {
		records[i] = ledgerFromEntity(ledger)
	}

	query := squirrel.
		Insert("ledgers").
		Columns(
			"sequence", "ledger_hash", "previous_ledger_hash", "paging_token", "transaction_count",
			"operation_count", "closed_at", "total_coins", "fee_pool", "base_fee", "base_reserve",
			"max_transaction_set_size", "protocol_version", "created_at",
		)

	createdAt := time.Now()
	for _, record := range records {
		query = query.Values(
			record.Sequence,
			record.LedgerHash,
			record.PreviousLedgerHash,
			record.PagingToken,
			record.TransactionCount,
			record.OperationCount,
			record.ClosedAt,
			record.TotalCoins,
			record.FeePool,
			record.BaseFee,
			record.BaseReserve,
			record.MaxTransactionSetSize,
			record.ProtocolVersion,
			createdAt,
		)
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	// Ignore rows we've already ingested.
	sql += " ON CONFLICT (sequence) DO NOTHING"

	return db.Exec(sql, args...).Error
}

func ledgerFromEntity(entity *entities.Ledger) *ledger {
	return &ledger{
		PagingToken:           entity.PagingToken,
		LedgerHash:            entity.Hash,
		PreviousLedgerHash:    entity.PrevHash,
		Sequence:              entity.Sequence,
		TransactionCount:      entity.TransactionCount,
		OperationCount:        entity.OperationCount,
		ClosedAt:              entity.ClosedAt,
		TotalCoins:            entity.TotalCoins,
		FeePool:               entity.FeePool,
		BaseFee:               entity.BaseFee,
		BaseReserve:           entity.BaseReserve,
		MaxTransactionSetSize: entity.MaxTransactionSetSize,
		ProtocolVersion:       entity.ProtocolVersion,
	}
}
