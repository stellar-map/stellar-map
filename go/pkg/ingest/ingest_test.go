package ingest_test

import (
	"bytes"
	"context"
	"io/ioutil"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/stellar/go/clients/horizon"
	"github.com/stretchr/testify/mock"

	"github.com/stellar-map/stellar-map/go/pkg/entities"
	"github.com/stellar-map/stellar-map/go/pkg/ingest"
	"github.com/stellar-map/stellar-map/go/pkg/mocks"
)

var _ = Describe("Ingest", func() {
	var (
		ctx    = context.Background()
		cursor = ingest.Now

		mockCtrl *gomock.Controller
		repo     *mocks.MockRepo
		client   *horizon.MockClient
		logBuf   *bytes.Buffer
		logger   *log.Logger
		ingester *ingest.HorizonIngester

		horizonCursor *horizon.Cursor
		horizonLedger *horizon.Ledger
		ledger        *entities.Ledger
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		repo = mocks.NewMockRepo(mockCtrl)
		client = &horizon.MockClient{}
		logBuf = &bytes.Buffer{}
		logger = log.New()
		logger.Out = logBuf
		ingester = ingest.NewHorizon(client, repo, logger)

		hc := horizon.Cursor(cursor)
		horizonCursor = &hc
		horizonLedger = &horizon.Ledger{
			Sequence:   12345,
			PT:         "PT",
			Hash:       "ledgerhash",
			TotalCoins: "100.0000001",
			FeePool:    "0",
		}
		ledger = &entities.Ledger{
			Sequence:    12345,
			PagingToken: "PT",
			LedgerHash:  "ledgerhash",
			TotalCoins:  1000000001,
			FeePool:     0,
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Ledgers", func() {
		It("creates ledger with the repo", func() {
			client.On("StreamLedgers", ctx, horizonCursor, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
				handler := args[2].(horizon.LedgerHandler)
				handler(*horizonLedger)
			})
			repo.EXPECT().BatchCreateLedgers([]*entities.Ledger{ledger}).Return(nil)

			ingester.BatchSize = 1
			err := ingester.Ledgers(ctx, cursor)
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns an error if StreamLedgers fail", func() {
			client.On("StreamLedgers", ctx, horizonCursor, mock.Anything).Return(errors.New("fail"))

			err := ingester.Ledgers(ctx, cursor)
			Expect(err).To(HaveOccurred())
		})

		It("logs an error if parsing the received ledger fails", func() {
			client.On("StreamLedgers", ctx, horizonCursor, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
				horizonLedger.TotalCoins = "abcd"
				handler := args[2].(horizon.LedgerHandler)
				handler(*horizonLedger)
			})

			err := ingester.Ledgers(ctx, cursor)
			Expect(err).NotTo(HaveOccurred())

			logBytes, err := ioutil.ReadAll(logBuf)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(logBytes)).To(ContainSubstring("cannot parse amount"))
		})

		It("logs an error if creating with the repo fails", func() {
			client.On("StreamLedgers", ctx, horizonCursor, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
				handler := args[2].(horizon.LedgerHandler)
				handler(*horizonLedger)
			})
			repo.EXPECT().BatchCreateLedgers([]*entities.Ledger{ledger}).Return(errors.New("repo failed"))

			ingester.BatchSize = 1
			err := ingester.Ledgers(ctx, cursor)
			Expect(err).NotTo(HaveOccurred())

			logBytes, err := ioutil.ReadAll(logBuf)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(logBytes)).To(ContainSubstring("repo failed"))
		})
	})
})
