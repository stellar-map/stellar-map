-- +migrate Up
CREATE TABLE ledgers (
  sequence TEXT NOT NULL PRIMARY KEY,
  ledger_hash TEXT NOT NULL,
  previous_ledger_hash TEXT NOT NULL,
  paging_token TEXT,
  transaction_count INT,
  operation_count INT,
  closed_at TIMESTAMP NOT NULL,
  total_coins TEXT,
  fee_pool TEXT,
  base_fee INT,
  base_reserve INT,
  max_transaction_set_size INT,
  protocol_version INT,
  created_at TIMESTAMP NOT NULL
);
CREATE UNIQUE INDEX ON ledgers (ledger_hash);
CREATE UNIQUE INDEX ON ledgers (previous_ledger_hash);
CREATE INDEX ON ledgers (closed_at);

-- +migrate Down
DROP TABLE ledgers;
