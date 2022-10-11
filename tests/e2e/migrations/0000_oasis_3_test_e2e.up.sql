-- Indexer state initialization for the Damask Upgrade.
-- https://github.com/oasisprotocol/mainnet-artifacts/releases/tag/2022-04-11

BEGIN;

-- Create Damask Upgrade Schema with `chain-id`.
CREATE SCHEMA IF NOT EXISTS oasis_test;

-- Block Data

CREATE TABLE IF NOT EXISTS oasis_test.blocks
(
  height     BIGINT PRIMARY KEY,
  block_hash TEXT NOT NULL,
  time       TIMESTAMP WITH TIME ZONE NOT NULL,

  -- State Root Info
  namespace TEXT NOT NULL,
  version   BIGINT NOT NULL,
  type      TEXT NOT NULL,
  root_hash TEXT NOT NULL,

  beacon     BYTEA,
  metadata   JSON
);

CREATE TABLE IF NOT EXISTS oasis_test.transactions
(
  block BIGINT NOT NULL REFERENCES oasis_test.blocks(height),

  txn_hash   TEXT NOT NULL,
  txn_index  INTEGER,
  nonce      NUMERIC NOT NULL,
  fee_amount NUMERIC,
  max_gas    NUMERIC,
  method     TEXT NOT NULL,
  sender     TEXT NOT NULL,
  body       BYTEA,

  -- Error Fields
  -- This includes an encoding of no error.
  module  TEXT,
  code    BIGINT,
  message TEXT,

  -- We require a composite primary key since duplicate transactions can
  -- be included within blocks for this chain.
  PRIMARY KEY (block, txn_hash, txn_index)
);

CREATE TABLE IF NOT EXISTS oasis_test.events
(
  backend TEXT NOT NULL,
  type    TEXT NOT NULL,
  body    JSON,

  txn_block  BIGINT NOT NULL,
  txn_hash   TEXT NOT NULL,
  txn_index  INTEGER,

  FOREIGN KEY (txn_block, txn_hash, txn_index) REFERENCES oasis_test.transactions(block, txn_hash, txn_index)
);

-- Beacon Backend Data

CREATE TABLE IF NOT EXISTS oasis_test.epochs
(
  id           BIGINT PRIMARY KEY,
  start_height BIGINT NOT NULL,
  end_height   BIGINT,
  UNIQUE (start_height, end_height)
);

-- Registry Backend Data
CREATE TABLE IF NOT EXISTS oasis_test.entities
(
  id      TEXT PRIMARY KEY,
  address TEXT
);

CREATE TABLE IF NOT EXISTS oasis_test.nodes
(
  id         TEXT PRIMARY KEY,
  entity_id  TEXT NOT NULL REFERENCES oasis_test.entities(id),
  expiration BIGINT NOT NULL,

  -- TLS Info
  tls_pubkey      TEXT NOT NULL,
  tls_next_pubkey TEXT,
  tls_addresses   TEXT ARRAY,

  -- P2P Info
  p2p_pubkey    TEXT NOT NULL,
  p2p_addresses TEXT ARRAY,

  -- Consensus Info
  consensus_pubkey  TEXT NOT NULL,
  consensus_address TEXT,

  -- VRF Info
  vrf_pubkey TEXT,

  roles            TEXT,
  software_version TEXT,

  -- Voting power should only be nonzero for consensus validator nodes.
  voting_power     BIGINT DEFAULT 0,

  -- TODO: Track node status.

  -- Arbitrary additional data.
  extra_data JSON
);

CREATE TABLE IF NOT EXISTS oasis_test.runtimes
(
  id           TEXT PRIMARY KEY,
  suspended    BOOLEAN NOT NULL DEFAULT false,
  kind         TEXT NOT NULL,
  tee_hardware TEXT NOT NULL,
  key_manager  TEXT
);

-- Staking Backend Data

CREATE TABLE IF NOT EXISTS oasis_test.accounts
(
  address TEXT PRIMARY KEY,

  -- General Account
  general_balance NUMERIC DEFAULT 0,
  nonce           BIGINT DEFAULT 0,

  -- Escrow Account
  escrow_balance_active         NUMERIC DEFAULT 0,
  escrow_total_shares_active    NUMERIC DEFAULT 0,
  escrow_balance_debonding      NUMERIC DEFAULT 0,
  escrow_total_shares_debonding NUMERIC DEFAULT 0,

  -- TODO: Track commission schedule and staking accumulator.

  -- Arbitrary additional data.
  extra_data JSON
);

CREATE TABLE IF NOT EXISTS oasis_test.allowances
(
  owner       TEXT NOT NULL,
  beneficiary TEXT NOT NULL,
  allowance   NUMERIC,

  PRIMARY KEY (owner, beneficiary)
);

CREATE TABLE IF NOT EXISTS oasis_test.delegations
(
  delegatee TEXT NOT NULL,
  delegator TEXT NOT NULL,
  shares    NUMERIC NOT NULL,

  PRIMARY KEY (delegatee, delegator)
);

CREATE TABLE IF NOT EXISTS oasis_test.debonding_delegations
(
  delegatee  TEXT NOT NULL,
  delegator  TEXT NOT NULL,
  shares     NUMERIC NOT NULL,
  debond_end BIGINT NOT NULL
);

-- Scheduler Backend Data

CREATE TABLE IF NOT EXISTS oasis_test.committee_members
(
  node      TEXT NOT NULL,
  valid_for BIGINT NOT NULL,
  runtime   TEXT NOT NULL,
  kind      TEXT NOT NULL,
  role      TEXT NOT NULL,

  PRIMARY KEY (node, runtime, kind, role)
);

-- Governance Backend Data

CREATE TABLE IF NOT EXISTS oasis_test.proposals
(
  id            BIGINT PRIMARY KEY,
  submitter     TEXT NOT NULL,
  state         TEXT NOT NULL DEFAULT 'active',
  executed      BOOLEAN NOT NULL DEFAULT false,
  deposit       NUMERIC NOT NULL,

  -- If this proposal is a new proposal.
  handler            TEXT,
  cp_target_version  TEXT,
  rhp_target_version TEXT,
  rcp_target_version TEXT,
  upgrade_epoch      BIGINT,

  -- If this proposal cancels an existing proposal.
  cancels BIGINT REFERENCES oasis_test.proposals(id) DEFAULT NULL,

  created_at    BIGINT NOT NULL,
  closes_at     BIGINT NOT NULL,
  invalid_votes NUMERIC NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS oasis_test.votes
(
  proposal BIGINT NOT NULL REFERENCES oasis_test.proposals(id),
  voter    TEXT NOT NULL,
  vote     TEXT,

  PRIMARY KEY (proposal, voter)
);

-- Indexing Progress Management
CREATE TABLE IF NOT EXISTS oasis_test.processed_blocks
(
  height         BIGINT NOT NULL,
  analyzer       TEXT NOT NULL,
  processed_time TIMESTAMP WITH TIME ZONE NOT NULL,

  PRIMARY KEY (height, analyzer)
);

CREATE TABLE IF NOT EXISTS oasis_test.checkpointed_heights
(
  height BIGINT PRIMARY KEY,
  checkpoint_time TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE oasis_test.nodes DROP CONSTRAINT nodes_entity_id_fkey;

CREATE TABLE IF NOT EXISTS oasis_test.claimed_nodes
(
  entity_id TEXT NOT NULL,
  node_id   TEXT NOT NULL,

  PRIMARY KEY (entity_id, node_id)  
);

ALTER TABLE oasis_test.entities ADD meta JSON;

CREATE TABLE IF NOT EXISTS oasis_test.commissions
(
  address  TEXT PRIMARY KEY NOT NULL REFERENCES oasis_test.accounts(address),
  schedule JSON
);

CREATE INDEX ix_transactions_sender ON oasis_test.transactions (sender);

ALTER TABLE oasis_test.debonding_delegations ADD COLUMN id bigserial PRIMARY KEY;


COMMIT;
