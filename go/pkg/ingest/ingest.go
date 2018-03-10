package ingest

import (
	"context"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/stellar/go/amount"
	"github.com/stellar/go/clients/horizon"

	"github.com/stellar-map/stellar-map/go/pkg/entities"
)

const (
	// Beginning starts streamining from the earliest page.
	Beginning = Cursor("0")
	// Continue starts streaming from the latest ingested page.
	Continue = Cursor("continue")
	// Now starts streaming new pages.
	Now = Cursor("now")
)

// Cursor represents the where to begin streaming resources from.
type Cursor string

const (
	defaultBatchSize      = 1000
	defaultBatchWaitDelay = 5 * time.Second
)

// Ingester is responsible for ingesting Stellar resources.
type Ingester interface {
	Ledgers(ctx context.Context, cursor Cursor) error
}

// HorizonIngester is an ingester that ingests from a Horizon server.
type HorizonIngester struct {
	client horizon.ClientInterface
	repo   entities.Repo
	logger *log.Logger

	BatchSize      int
	BatchWaitDelay time.Duration
}

// NewHorizon returns a new HorizonIngestor configured with the provided arguments.
func NewHorizon(client horizon.ClientInterface, repo entities.Repo, logger *log.Logger) *HorizonIngester {
	return &HorizonIngester{
		client:         client,
		repo:           repo,
		logger:         logger,
		BatchSize:      defaultBatchSize,
		BatchWaitDelay: defaultBatchWaitDelay,
	}
}

// Ledgers ingests the ledger resource.
func (i *HorizonIngester) Ledgers(ctx context.Context, cursor Cursor) error {
	horizonCursor, err := i.horizonCursorForResource("ledger", cursor)
	if err != nil {
		return errors.Wrap(err, "error parsing cursor")
	}

	ledgers := make([]*entities.Ledger, 0, i.BatchSize)
	timer := time.Now().Add(i.BatchWaitDelay)

	err = i.client.StreamLedgers(ctx, horizonCursor, func(l horizon.Ledger) {
		ledger, err := ledgerToEntity(&l)
		if err != nil {
			i.logger.WithError(err).Error("failed to parse ledger")
			return
		}

		ledgers = append(ledgers, ledger)

		count := len(ledgers)
		if time.Now().After(timer) || count >= i.BatchSize {
			if err := i.repo.BatchCreateLedgers(ledgers); err != nil {
				fields := log.Fields{"start": ledgers[0].Sequence, "end": ledgers[count-1].Sequence}
				i.logger.WithError(err).WithFields(fields).Error("failed to create ledgers")
			}
			i.logger.WithField("count", count).Info("ingested new records")

			ledgers = make([]*entities.Ledger, 0, i.BatchSize)
			timer = time.Now().Add(i.BatchWaitDelay)
		}
	})

	if err != nil {
		return errors.Wrap(err, "error streaming ledgers")
	}
	return nil
}

// Transactions ingests the transaction resource.
func (i *HorizonIngester) Transactions(ctx context.Context, cursor Cursor) error {
	horizonCursor, err := i.horizonCursorForResource("transactions", cursor)
	if err != nil {
		return errors.Wrap(err, "error parsing cursor")
	}

	transactions := make([]*entities.Transaction, 0, i.BatchSize)
	timer := time.Now().Add(i.BatchWaitDelay)

	err = i.client.StreamAllTransactions(ctx, horizonCursor, func(t horizon.Transaction) {
		transaction, err := transactionToEntity(&t)
		if err != nil {
			i.logger.WithError(err).Error("failed to parse transaction")
			return
		}

		transactions = append(transactions, transaction)

		count := len(transactions)
		if time.Now().After(timer) || count >= i.BatchSize {
			if err := i.repo.BatchCreateTransactions(transactions); err != nil {
				i.logger.WithError(err).Error("failed to create transactions")
			}
			i.logger.WithField("count", count).Info("ingested new records")

			transactions = make([]*entities.Transaction, 0, i.BatchSize)
			timer = time.Now().Add(i.BatchWaitDelay)
		}
	})

	if err != nil {
		return errors.Wrap(err, "error streaming transactions")
	}
	return nil
}

func (i *HorizonIngester) horizonCursorForResource(resource string, cursor Cursor) (*horizon.Cursor, error) {
	if cursor == Continue {
		c, err := i.repo.PagingToken(resource)
		if err != nil {
			if err != entities.ErrNotFound {
				return nil, err
			}
			cursor = Beginning
		}
		cursor = Cursor(c)
	}

	horizonCursor := horizon.Cursor(cursor)
	return &horizonCursor, nil
}

func ledgerToEntity(ledger *horizon.Ledger) (*entities.Ledger, error) {
	totalCoins, err := amount.ParseInt64(ledger.TotalCoins)
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing total_coins %s", ledger.TotalCoins)
	}

	feePool, err := amount.ParseInt64(ledger.FeePool)
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing fee_pool %s", ledger.FeePool)
	}

	return &entities.Ledger{
		Sequence:           ledger.Sequence,
		LedgerHash:         ledger.Hash,
		PreviousLedgerHash: ledger.PrevHash,
		PagingToken:        ledger.PT,
		TransactionCount:   ledger.TransactionCount,
		OperationCount:     ledger.OperationCount,
		ClosedAt:           ledger.ClosedAt,
		TotalCoins:         totalCoins,
		FeePool:            feePool,
		BaseFee:            ledger.BaseFee,
		BaseReserve:        ledger.BaseReserve,
		MaxTxSetSize:       ledger.MaxTxSetSize,
		ProtocolVersion:    ledger.ProtocolVersion,
	}, nil
}

func transactionToEntity(transaction *horizon.Transaction) (*entities.Transaction, error) {
	validBefore, err := parseTime(transaction.ValidBefore)
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing valid_before %s", transaction.ValidBefore)
	}

	validAfter, err := parseTime(transaction.ValidAfter)
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing valid_after %s", transaction.ValidAfter)
	}

	return &entities.Transaction{
		TransactionHash: transaction.Hash,
		LedgerSequence:  transaction.Ledger,
		LedgerClosedAt:  transaction.LedgerCloseTime,
		PagingToken:     transaction.PagingToken,
		Account:         transaction.Account,
		AccountSequence: transaction.AccountSequence,
		FeePaid:         transaction.FeePaid,
		OperationCount:  transaction.OperationCount,
		EnvelopeXDR:     transaction.EnvelopeXdr,
		ResultXDR:       transaction.ResultXdr,
		ResultMetaXDR:   transaction.ResultMetaXdr,
		FeeMetaXDR:      transaction.FeeMetaXdr,
		MemoType:        transaction.MemoType,
		Memo:            transaction.Memo,
		Signatures:      transaction.Signatures,
		ValidBefore:     validBefore,
		ValidAfter:      validAfter,
	}, nil
}

func parseTime(s string) (time.Time, error) {
	if s == "" {
		return time.Time{}, nil
	}
	return time.Parse(time.RFC3339, s)
}
