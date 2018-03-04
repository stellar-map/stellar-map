package db

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/stellar-map/stellar-map/go/pkg/entities"
)

var _ = Describe("Ledger", func() {
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

	Context("BatchCreateLedgers", func() {
		It("saves a batch of ledgers", func() {
			ledgers := []*entities.Ledger{&entities.Ledger{
				PagingToken:        "abc",
				LedgerHash:         "abc",
				PreviousLedgerHash: "abc",
				Sequence:           100,
				TransactionCount:   1,
				OperationCount:     3,
				ClosedAt:           time.Now().UTC().Round(time.Millisecond),
				TotalCoins:         1000,
				FeePool:            50,
				BaseFee:            200,
				BaseReserve:        300,
				MaxTxSetSize:       400,
				ProtocolVersion:    5,
			}, &entities.Ledger{
				PagingToken:        "def",
				LedgerHash:         "def",
				PreviousLedgerHash: "def",
				Sequence:           101,
				TransactionCount:   2,
				OperationCount:     4,
				ClosedAt:           time.Now().UTC().Round(time.Millisecond),
				TotalCoins:         2000,
				FeePool:            75,
				BaseFee:            201,
				BaseReserve:        301,
				MaxTxSetSize:       401,
				ProtocolVersion:    5,
			}}

			err := tx.BatchCreateLedgers(ledgers)
			Expect(err).NotTo(HaveOccurred())

			records := []*ledger{}
			err = tx.Order("sequence").Find(&records).Error
			Expect(err).NotTo(HaveOccurred())

			Expect(records).To(HaveLen(len(ledgers)))
			for i, record := range records {
				ledger := ledgers[i]
				Expect(ledgerToEntity(record)).To(Equal(ledger))
			}
		})

		It("does nothing if ledger with sequence already exists", func() {
			ledgers := []*entities.Ledger{&entities.Ledger{
				Sequence:           100,
				LedgerHash:         "abc",
				PreviousLedgerHash: "abc",
				ClosedAt:           time.Now().UTC().Round(time.Millisecond),
			}}
			err := tx.BatchCreateLedgers(ledgers)
			Expect(err).NotTo(HaveOccurred())

			ledgers[0].LedgerHash = "def"
			err = tx.BatchCreateLedgers(ledgers)
			Expect(err).NotTo(HaveOccurred())

			record := &ledger{}
			err = tx.Order("sequence").First(&record).Error
			Expect(err).NotTo(HaveOccurred())
			Expect(record.LedgerHash).To(Equal("abc"))
		})
	})

	Context("PagingToken", func() {
		It("returns the latest paging token for ledgers", func() {
			tx.Create(&ledger{Sequence: 1, LedgerHash: "a", PreviousLedgerHash: "a", PagingToken: "20"})
			tx.Create(&ledger{Sequence: 2, LedgerHash: "b", PreviousLedgerHash: "b", PagingToken: "10"})

			pt, err := tx.PagingToken("ledger")
			Expect(err).NotTo(HaveOccurred())
			Expect(pt).To(Equal("10"))
		})

		It("returns ErrNotFound for missing resources", func() {
			_, err := tx.PagingToken("ledger")
			Expect(err).To(Equal(entities.ErrNotFound))
		})

		It("errors for unknown resources", func() {
			_, err := tx.PagingToken("qwerty")
			Expect(err).To(HaveOccurred())
		})
	})
})
