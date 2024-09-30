BEGIN;

ALTER TABLE history.validators
    ADD COLUMN staking_rewards UINT_NUMERIC;

CREATE TABLE history.escrow_events
(
  tx_block UINT63 NOT NULL,
  epoch UINT63 NOT NULL,
  type TEXT NOT NULL,
  delegatee oasis_addr NOT NULL,
  delegator oasis_addr NOT NULL,
  shares    UINT_NUMERIC,
  amount UINT_NUMERIC,
  debonding_amount UINT_NUMERIC -- for slashing events
);

COMMIT;
