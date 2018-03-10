package db

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/stellar-map/stellar-map/go/pkg/entities"
)

var _ = Describe("Transaction", func() {
	var (
		tx *db
	)

	BeforeEach(func() {
		gormTx := testDB.DB.Begin()
		Expect(gormTx.Error).NotTo(HaveOccurred())
		tx = &db{DB: gormTx}
	})

	AfterEach(func() {
		err := tx.DB.Rollback().Error
		Expect(err).NotTo(HaveOccurred())
	})

	Context("BatchCreateTransactions", func() {
		It("saves a batch of transactions", func() {
			ledgerClosedAt := time.Date(2018, 5, 9, 0, 0, 0, 0, time.UTC)

			transactions := []*entities.Transaction{&entities.Transaction{
				TransactionHash: "abc",
				LedgerSequence:  1,
				LedgerClosedAt:  ledgerClosedAt,
				PagingToken:     "pt",
				Account:         "acct",
				AccountSequence: "1",
				FeePaid:         1,
				OperationCount:  1,
				EnvelopeXDR:     "a",
				ResultXDR:       "a",
				ResultMetaXDR:   "a",
				FeeMetaXDR:      "a",
				MemoType:        "memo type",
				Memo:            "memo",
				Signatures:      []string{"a", "b", "c"},
				ValidBefore:     ledgerClosedAt,
				ValidAfter:      time.Time{},
			}, &entities.Transaction{
				TransactionHash: "def",
				LedgerSequence:  2,
				LedgerClosedAt:  ledgerClosedAt,
				PagingToken:     "pt",
				Account:         "acct",
				AccountSequence: "1",
				FeePaid:         1,
				OperationCount:  1,
				EnvelopeXDR:     "a",
				ResultXDR:       "a",
				ResultMetaXDR:   "a",
				FeeMetaXDR:      "a",
				MemoType:        "memo type",
				Memo:            "memo",
				Signatures:      []string{"a", "b", "c"},
				ValidBefore:     ledgerClosedAt,
				ValidAfter:      time.Time{},
			}}

			err := tx.BatchCreateTransactions(transactions)
			Expect(err).NotTo(HaveOccurred())

			records := []*transaction{}
			err = tx.Find(&records).Error
			Expect(err).NotTo(HaveOccurred())

			Expect(records).To(HaveLen(len(transactions)))
			for i, record := range records {
				transaction := transactions[i]
				Expect(transactionToEntity(record)).To(Equal(transaction))
			}
		})

		It("does nothing if transaction with hash already exists", func() {
			transactions := []*entities.Transaction{&entities.Transaction{
				TransactionHash: "abc",
				LedgerSequence:  1,
			}}
			err := tx.BatchCreateTransactions(transactions)
			Expect(err).NotTo(HaveOccurred())

			transactions[0].LedgerSequence = 2
			err = tx.BatchCreateTransactions(transactions)
			Expect(err).NotTo(HaveOccurred())

			record := &transaction{}
			err = tx.First(&record).Error
			Expect(err).NotTo(HaveOccurred())
			Expect(record.LedgerSequence).To(Equal(int32(1)))
		})
	})
})
