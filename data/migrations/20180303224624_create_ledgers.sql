-- +migrate Up
CREATE TABLE ledgers (
  sequence INT NOT NULL PRIMARY KEY,
  ledger_hash TEXT NOT NULL,
  previous_ledger_hash TEXT NOT NULL,
  paging_token TEXT,
  transaction_count INT,
  operation_count INT,
  closed_at TIMESTAMPTZ NOT NULL,
  total_coins BIGINT,
  fee_pool BIGINT,
  base_fee INT,
  base_reserve INT,
  max_tx_set_size INT,
  protocol_version INT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE UNIQUE INDEX ON ledgers (ledger_hash);
CREATE UNIQUE INDEX ON ledgers (previous_ledger_hash);
CREATE INDEX ON ledgers (closed_at);
CREATE INDEX ON ledgers (created_at);

-- +migrate Down
DROP TABLE ledgers;
