-- State initialization for the Damask Upgrade.
-- https://github.com/oasisprotocol/mainnet-artifacts/releases/tag/2022-04-11

BEGIN;

-- Create Damask Upgrade Schema with `chain-id`.
CREATE SCHEMA IF NOT EXISTS chain;
GRANT USAGE ON SCHEMA chain TO PUBLIC;

-- Custom types
CREATE DOMAIN public.uint_numeric NUMERIC(1000,0) CHECK(VALUE >= 0);
CREATE DOMAIN public.uint63 BIGINT CHECK(VALUE >= 0);
CREATE DOMAIN public.uint31 INTEGER CHECK(VALUE >= 0);
CREATE DOMAIN public.hex64 TEXT CHECK(VALUE ~ '^[0-9a-f]{64}$');
-- base64(ed25519 public key); from https://github.com/oasisprotocol/oasis-core/blob/f95186e3f15ec64bdd36493cde90be359bd17da8/go/common/crypto/signature/signature.go#L90-L90
CREATE DOMAIN public.base64_ed25519_pubkey TEXT CHECK(VALUE ~ '^[A-Za-z0-9+/]{43}=$');
CREATE DOMAIN public.oasis_addr TEXT CHECK(length(VALUE) = 46 AND VALUE ~ '^oasis1');

-- Block Data
CREATE TABLE chain.blocks
(
  height     UINT63 PRIMARY KEY,
  block_hash HEX64 NOT NULL,
  time       TIMESTAMP WITH TIME ZONE NOT NULL,
  num_txs    UINT31 NOT NULL,

  -- State Root Info
  namespace TEXT NOT NULL,
  version   UINT63 NOT NULL,
  type      TEXT NOT NULL,  -- "invalid" | "state-root" | "io-root"; From https://github.com/oasisprotocol/oasis-core/blob/f95186e3f15ec64bdd36493cde90be359bd17da8/go/storage/mkvs/node/node.go#L68-L68
  root_hash HEX64 NOT NULL,

  beacon     BYTEA,
  metadata   JSONB
);
CREATE INDEX ix_blocks_time ON chain.blocks (time);

CREATE TABLE chain.transactions
(
  block UINT63 NOT NULL REFERENCES chain.blocks(height) DEFERRABLE INITIALLY DEFERRED,
  tx_index  UINT31 NOT NULL,

  tx_hash   HEX64 NOT NULL,
  nonce      UINT63 NOT NULL,
  fee_amount UINT_NUMERIC,
  max_gas    UINT_NUMERIC, -- uint64 in go; because the value might conceivably be >2^63, we use UINT_NUMERIC over UINT63 here.
  method     TEXT NOT NULL,
  sender     oasis_addr NOT NULL,
  body       BYTEA NOT NULL,

  -- Error Fields
  -- This includes an encoding of no error.
  module  TEXT,
  code    UINT31 NOT NULL,  -- From https://github.com/oasisprotocol/oasis-core/blob/f95186e3f15ec64bdd36493cde90be359bd17da8/go/consensus/api/transaction/results/results.go#L20-L20
  message TEXT,

  -- We require a composite primary key since duplicate transactions (with identical hashes) can
  -- be included within blocks for this chain.
  PRIMARY KEY (block, tx_index)
);
-- Queries by sender and/or tx_hash are available via the API.
CREATE INDEX ix_transactions_sender ON chain.transactions (sender);
CREATE INDEX ix_transactions_tx_hash ON chain.transactions (tx_hash);

CREATE TABLE chain.events
(
  tx_block UINT63 NOT NULL,
  tx_index  UINT31,

  type    TEXT NOT NULL,  -- Enum with many values, see https://github.com/oasisprotocol/nexus/blob/89b68717205809b491d7926533d096444611bd6b/analyzer/api.go#L171-L171
  body    JSONB,
  tx_hash   HEX64, -- could be fetched from `transactions` table; denormalized for efficiency
  related_accounts TEXT[],

  FOREIGN KEY (tx_block, tx_index) REFERENCES chain.transactions(block, tx_index) DEFERRABLE INITIALLY DEFERRED
);
CREATE INDEX ix_events_related_accounts ON chain.events USING gin(related_accounts);
CREATE INDEX ix_events_tx_block ON chain.events (tx_block);  -- for fetching events without filters
CREATE INDEX ix_events_tx_hash ON chain.events (tx_hash);
CREATE INDEX ix_events_type ON chain.events (type, tx_block);  -- tx_block is for sorting the events of a given type by recency

-- Beacon Backend Data

CREATE TABLE chain.epochs
(
  id           UINT63 PRIMARY KEY,
  -- Earliest known height that belongs to the epoch.
  start_height UINT63 NOT NULL,
  -- Max known height that belongs to the epoch.
  end_height   UINT63 NOT NULL CHECK (end_height >= start_height),
  UNIQUE (start_height, end_height)
);
CREATE INDEX ix_epochs_id ON chain.epochs (id);

-- Registry Backend Data
CREATE TABLE chain.entities
(
  id      base64_ed25519_pubkey PRIMARY KEY,
  address oasis_addr NOT NULL, -- Deterministically derived from the ID.
  meta    JSONB  -- Signed statements about the entity from https://github.com/oasisprotocol/metadata-registry
);

CREATE TABLE chain.nodes
(
  -- `id` technically REFERENCES chain.claimed_nodes(node_id) because node had to be pre-claimed; see chain.claimed_nodes.
  -- However, postgres does not allow foreign keys to a non-unique column.
  id         base64_ed25519_pubkey PRIMARY KEY,
  -- Owning entity. The entity has likely claimed this node (see chain.claimed_nodes) previously. However
  -- historically (as per @Yawning), we also allowed node registrations that are signed with the entity signing key,
  -- in which case, the node would be allowed to register without having been pre-claimed by the entity.
  -- For those cases, (id, entity_id) is not a foreign key into chain.claimed_nodes.
  -- Similarly, an entity can un-claim a node after the node registered, but the node can remain be registered for a while.
  entity_id  base64_ed25519_pubkey NOT NULL REFERENCES chain.entities(id),
  expiration UINT63 NOT NULL, -- The epoch in which this node expires.

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
  voting_power     UINT63 DEFAULT 0

  -- TODO: Track node status.
);

-- Claims of entities that they own nodes. Each entity claims 0 or more nodes when it registers.
-- A node can only register if it declares itself to be owned by an entity that previously claimed it.
CREATE TABLE chain.claimed_nodes
(
  entity_id base64_ed25519_pubkey NOT NULL REFERENCES chain.entities(id) DEFERRABLE INITIALLY DEFERRED,
  node_id   base64_ed25519_pubkey NOT NULL,  -- No REFERENCES because the node likely does not exist (in the DB) yet when the entity claims it.

  PRIMARY KEY (entity_id, node_id)
);

CREATE TABLE chain.runtimes
(
  id           HEX64 PRIMARY KEY,
  suspended    BOOLEAN NOT NULL DEFAULT false,  -- not tracked as of Dec 2022
  kind         TEXT NOT NULL,  -- "invalid" | "compute" | "manager"; see https://github.com/oasisprotocol/oasis-core/blob/f95186e3f15ec64bdd36493cde90be359bd17da8/go/registry/api/runtime.go#L54-L54
  tee_hardware TEXT NOT NULL,  -- "invalid" | "intel-sgx"; see https://github.com/oasisprotocol/oasis-core/blob/f95186e3f15ec64bdd36493cde90be359bd17da8/go/common/node/node.go#L474-L474
  key_manager  HEX64
);

CREATE TABLE chain.runtime_nodes
(
  runtime_id HEX64 NOT NULL REFERENCES chain.runtimes(id) DEFERRABLE INITIALLY DEFERRED,
  node_id    base64_ed25519_pubkey NOT NULL REFERENCES chain.nodes(id) DEFERRABLE INITIALLY DEFERRED,

  PRIMARY KEY (runtime_id, node_id)
);

-- Staking Backend Data

CREATE TABLE chain.accounts
(
  address oasis_addr PRIMARY KEY,

  -- General Account
  general_balance UINT_NUMERIC DEFAULT 0,
  nonce           UINT63 NOT NULL DEFAULT 0,

  -- Escrow Account
  -- TODO: Use UINT_NUMERIC for the next four columns. Their values should always be >=0;
  -- however in Cobalt, the emitted events didn't allow perfect tracking of shares, so
  -- a dead-reckoning analyzer can arrive at negative values (https://github.com/oasisprotocol/nexus/pull/370).
  escrow_balance_active         NUMERIC(1000,0) NOT NULL DEFAULT 0,
  escrow_total_shares_active    NUMERIC(1000,0) NOT NULL DEFAULT 0,
  escrow_balance_debonding      NUMERIC(1000,0) NOT NULL DEFAULT 0,
  escrow_total_shares_debonding NUMERIC(1000,0) NOT NULL DEFAULT 0

  -- TODO: Track commission schedule and staking accumulator.
);

CREATE TABLE chain.allowances
(
  owner       oasis_addr NOT NULL REFERENCES chain.accounts(address) DEFERRABLE INITIALLY DEFERRED,
  -- When creating an allowance for the purpose of subsequently depositing funds to a
  -- paratime account A in paratime P (i.e. the expected use case for allowances), `beneficiary` is
  -- the "staking account" of P. The staking account is a special account derived from the paratime ID:
  --  - derivation: https://github.com/oasisprotocol/oasis-core/blob/f95186e3f15ec64bdd36493cde90be359bd17da8/go/staking/api/address.go#L96-L96
  --  - precomputed accounts: https://github.com/oasisprotocol/oasis-wallet-web/blob/34fdf495de5ca0d585addf0073f6a71bba556588/src/config.ts#L89-L139
  beneficiary oasis_addr,
  allowance   UINT_NUMERIC,

  PRIMARY KEY (owner, beneficiary)
);

CREATE TABLE chain.commissions
(
  address  oasis_addr PRIMARY KEY NOT NULL REFERENCES chain.accounts(address) DEFERRABLE INITIALLY DEFERRED,
  schedule JSONB
);

CREATE TABLE chain.delegations
(
  delegatee oasis_addr NOT NULL REFERENCES chain.accounts(address) DEFERRABLE INITIALLY DEFERRED,
  delegator oasis_addr NOT NULL REFERENCES chain.accounts(address) DEFERRABLE INITIALLY DEFERRED,
  shares    UINT_NUMERIC NOT NULL,

  PRIMARY KEY (delegatee, delegator)
);
CREATE INDEX ix_delegations_delegator ON chain.delegations(delegator);

CREATE TABLE chain.debonding_delegations
(
  id         BIGSERIAL PRIMARY KEY,  -- index-internal ID
  delegatee  oasis_addr NOT NULL REFERENCES chain.accounts(address) DEFERRABLE INITIALLY DEFERRED,
  delegator  oasis_addr NOT NULL REFERENCES chain.accounts(address) DEFERRABLE INITIALLY DEFERRED,
  shares     UINT_NUMERIC NOT NULL,
  debond_end UINT63 NOT NULL  -- EpochTime, i.e. number of epochs since base epoch
);

-- Scheduler Backend Data

CREATE TABLE chain.committee_members
(
  node      TEXT NOT NULL,
  valid_for UINT63 NOT NULL,
  runtime   TEXT NOT NULL,
  kind      TEXT NOT NULL,
  role      TEXT NOT NULL,

  PRIMARY KEY (node, runtime, kind, role)
);

-- Governance Backend Data

CREATE TABLE chain.proposals
(
  id            UINT63 PRIMARY KEY,
  submitter     oasis_addr NOT NULL,
  state         TEXT NOT NULL DEFAULT 'active',  -- "active" | "passed" | "rejected" | "failed"; see https://github.com/oasisprotocol/oasis-core/blob/f95186e3f15ec64bdd36493cde90be359bd17da8/go/governance/api/proposal.go#L29-L29
  executed      BOOLEAN NOT NULL DEFAULT false,
  deposit       UINT_NUMERIC NOT NULL,

  -- If this proposal is a new proposal.
  handler            TEXT,
  cp_target_version  TEXT,
  rhp_target_version TEXT,
  rcp_target_version TEXT,
  upgrade_epoch      UINT63,

  -- If this proposal cancels an existing proposal.
  cancels UINT63 REFERENCES chain.proposals(id) DEFAULT NULL,

  created_at    UINT63 NOT NULL,  -- EpochTime, i.e. number of epochs since base epoch
  closes_at     UINT63 NOT NULL,  -- EpochTime, i.e. number of epochs since base epoch
  invalid_votes UINT_NUMERIC NOT NULL DEFAULT 0 -- uint64 in go; because the value might conceivably be >2^63, we use UINT_NUMERIC over UINT63 here.
);

CREATE TABLE chain.votes
(
  proposal UINT63 NOT NULL REFERENCES chain.proposals(id) DEFERRABLE INITIALLY DEFERRED,
  voter    oasis_addr NOT NULL,
  vote     TEXT,  -- "yes" | "no" | "abstain"; see https://github.com/oasisprotocol/oasis-core/blob/f95186e3f15ec64bdd36493cde90be359bd17da8/go/registry/api/runtime.go#L54-L54

  PRIMARY KEY (proposal, voter)
);

-- Related Accounts Data

CREATE TABLE chain.accounts_related_transactions
(
  account_address oasis_addr NOT NULL,
  tx_block UINT63 NOT NULL,
  tx_index UINT31 NOT NULL,
  FOREIGN KEY (tx_block, tx_index) REFERENCES chain.transactions(block, tx_index) DEFERRABLE INITIALLY DEFERRED
);
CREATE INDEX ix_accounts_related_transactions_address_block_index ON chain.accounts_related_transactions (account_address);

-- Indexing Progress Management
CREATE TABLE chain.processed_blocks  -- Moved to analysis.processed_blocks in 06_analysis_schema.up.sql
(
  height         UINT63 NOT NULL,
  analyzer       TEXT NOT NULL,

  processed_time TIMESTAMP WITH TIME ZONE, -- NULL if the block is not yet processed.
  locked_time     TIMESTAMP WITH TIME ZONE NOT NULL,

  PRIMARY KEY (analyzer, height)
);

CREATE INDEX ix_processed_blocks_analyzer_height_locked_unprocessed ON chain.processed_blocks (analyzer, height, locked_time) WHERE processed_time IS NULL; -- Index for efficient query of unprocessed blocks.
CREATE INDEX ix_processed_blocks_analyzer_height_locked_processed ON chain.processed_blocks (analyzer, height, locked_time, processed_time) WHERE processed_time IS NOT NULL; -- Index for efficient query of processed blocks.

-- Keeps track of chains for which we've already processed the genesis data.
CREATE TABLE chain.processed_geneses (
    chain_context TEXT NOT NULL PRIMARY KEY,  -- identifies the genesis data; derived from its hash
    processed_time TIMESTAMP WITH TIME ZONE NOT NULL
);

-- Grant others read-only use. This does NOT apply to future tables in the schema.
GRANT SELECT ON ALL TABLES IN SCHEMA chain TO PUBLIC;
GRANT EXECUTE ON ALL FUNCTIONS IN SCHEMA chain TO PUBLIC;

COMMIT;
