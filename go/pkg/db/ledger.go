package db

import (
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/stellar-map/stellar-map/go/pkg/entities"
)

type ledger struct {
	Sequence           int32 `gorm:"primary_key"`
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
	CreatedAt          time.Time
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
			"max_tx_set_size", "protocol_version",
		)

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
			record.MaxTxSetSize,
			record.ProtocolVersion,
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
		PagingToken:        entity.PagingToken,
		LedgerHash:         entity.LedgerHash,
		PreviousLedgerHash: entity.PreviousLedgerHash,
		Sequence:           entity.Sequence,
		TransactionCount:   entity.TransactionCount,
		OperationCount:     entity.OperationCount,
		ClosedAt:           entity.ClosedAt,
		TotalCoins:         entity.TotalCoins,
		FeePool:            entity.FeePool,
		BaseFee:            entity.BaseFee,
		BaseReserve:        entity.BaseReserve,
		MaxTxSetSize:       entity.MaxTxSetSize,
		ProtocolVersion:    entity.ProtocolVersion,
	}
}

func ledgerToEntity(record *ledger) *entities.Ledger {
	return &entities.Ledger{
		PagingToken:        record.PagingToken,
		LedgerHash:         record.LedgerHash,
		PreviousLedgerHash: record.PreviousLedgerHash,
		Sequence:           record.Sequence,
		TransactionCount:   record.TransactionCount,
		OperationCount:     record.OperationCount,
		ClosedAt:           record.ClosedAt,
		TotalCoins:         record.TotalCoins,
		FeePool:            record.FeePool,
		BaseFee:            record.BaseFee,
		BaseReserve:        record.BaseReserve,
		MaxTxSetSize:       record.MaxTxSetSize,
		ProtocolVersion:    record.ProtocolVersion,
	}
}

func (db *db) PagingToken(resource string) (string, error) {
	var err error

	switch resource {
	case "ledger":
		l := &ledger{}
		err = db.Order("created_at DESC").First(l).Error
		if err == nil {
			return l.PagingToken, nil
		}
	default:
		return "", errors.Errorf("unknown resource %s", resource)
	}

	if err == gorm.ErrRecordNotFound {
		return "", entities.ErrNotFound
	}
	return "", err
}
