package ingest

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/stellar/go/clients/horizon"

	"github.com/stellar-map/stellar-map/go/pkg/entities"
)

const (
	defaultBatchSize      = 1000
	defaultBatchWaitDelay = 5 * time.Second
)

// Ingester is responsible for ingesting Stellar resources from a Horizon server.
type Ingester interface {
	Ledgers(ctx context.Context, cursor string) error
	Transactions(ctx context.Context, cursor string) error
}

type ingester struct {
	repo      entities.Repo
	client    *horizon.Client
	batchSize int
}

// New returns a new Ingestor configured with the provided arguments.
func New(repo entities.Repo, testNet, batch bool) Ingester {
	client := horizon.DefaultPublicNetClient
	if testNet {
		client = horizon.DefaultTestNetClient
	}

	batchSize := defaultBatchSize
	if !batch {
		batchSize = 0
	}

	return &ingester{
		repo:      repo,
		client:    client,
		batchSize: batchSize,
	}
}

func (i *ingester) Ledgers(ctx context.Context, cursor string) error {
	logger := log.WithField("name", "ingest.ledgers")
	c := horizon.Cursor(cursor)

	ledgers := make([]*entities.Ledger, 0, i.batchSize)
	timer := time.Now().Add(defaultBatchWaitDelay)

	err := i.client.StreamLedgers(ctx, &c, func(l horizon.Ledger) {
		ledgers = append(ledgers, ledgerToEntity(&l))

		if time.Now().After(timer) || len(ledgers) >= i.batchSize {
			if err := i.repo.BatchCreateLedgers(ledgers); err != nil {
				logger.WithError(err).Error("failed to create ledgers")
			}
			logger.WithField("size", len(ledgers)).Info("ingested new records")

			ledgers = make([]*entities.Ledger, 0, i.batchSize)
			timer = time.Now().Add(defaultBatchWaitDelay)
		}
	})

	if err != nil {
		logger.WithError(err).Error("error streaming ledgers")
	}
	return err
}

func (i *ingester) Transactions(ctx context.Context, cursor string) error {
	return nil
}

func ledgerToEntity(ledger *horizon.Ledger) *entities.Ledger {
	return &entities.Ledger{
		ID:                    ledger.ID,
		PagingToken:           ledger.PT,
		Hash:                  ledger.Hash,
		PrevHash:              ledger.PrevHash,
		Sequence:              ledger.Sequence,
		TransactionCount:      ledger.TransactionCount,
		OperationCount:        ledger.OperationCount,
		ClosedAt:              ledger.ClosedAt,
		TotalCoins:            ledger.TotalCoins,
		FeePool:               ledger.FeePool,
		BaseFee:               ledger.BaseFee,
		BaseReserve:           ledger.BaseReserve,
		MaxTransactionSetSize: ledger.MaxTxSetSize,
		ProtocolVersion:       ledger.ProtocolVersion,
	}
}

func transactionToEntity(transaction *horizon.Transaction) *entities.Transaction {
	return &entities.Transaction{
		ID:              transaction.ID,
		PagingToken:     transaction.PagingToken,
		Hash:            transaction.Hash,
		Ledger:          transaction.Ledger,
		LedgerClosedAt:  transaction.LedgerCloseTime,
		Account:         transaction.Account,
		AccountSequence: transaction.AccountSequence,
		FeePaid:         transaction.FeePaid,
		OperationCount:  transaction.OperationCount,
		EnvelopeXdr:     transaction.EnvelopeXdr,
		ResultXdr:       transaction.ResultXdr,
		ResultMetaXdr:   transaction.ResultMetaXdr,
		FeeMetaXdr:      transaction.FeeMetaXdr,
		MemoType:        transaction.MemoType,
		Memo:            transaction.Memo,
		Signatures:      transaction.Signatures,
		ValidAfter:      transaction.ValidAfter,
		ValidBefore:     transaction.ValidBefore,
	}
}
