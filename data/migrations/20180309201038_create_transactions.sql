-- +migrate Up
CREATE TABLE transactions (
  id SERIAL NOT NULL PRIMARY KEY,
  transaction_hash TEXT NOT NULL,
  ledger_sequence INT NOT NULL,
  ledger_closed_at TIMESTAMPTZ,
  paging_token TEXT,
  account TEXT,
  account_sequence TEXT,
  fee_paid INT,
  operation_count INT,
  envelope_xdr TEXT,
  result_xdr TEXT,
  result_meta_xdr TEXT,
  fee_meta_xdr TEXT,
  memo_type TEXT,
  memo TEXT,
  signatures TEXT[],
  valid_before TIMESTAMPTZ,
  valid_after TIMESTAMPTZ,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE UNIQUE INDEX ON transactions (transaction_hash);
CREATE INDEX ON transactions (ledger_sequence);
CREATE INDEX ON transactions (created_at);

-- +migrate Down
DROP TABLE transactions;
