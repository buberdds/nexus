package commitment

import (
	"github.com/oasisprotocol/oasis-core/go/common/crypto/hash"
	"github.com/oasisprotocol/oasis-core/go/common/crypto/signature"
	"github.com/oasisprotocol/nexus/coreapi/v21.1.1/roothash/api/block"
)

// ProposedBatchSignatureContext is the context used for signing propose batch
// dispatch messages.
// removed var block

// ProposedBatch is the message sent from the transaction scheduler
// to executor workers after a batch is ready to be executed.
//
// Don't forget to bump CommitteeProtocol version in go/common/version
// if you change anything in this struct.
type ProposedBatch struct {
	// IORoot is the I/O root containing the inputs (transactions) that
	// the executor node should use.
	IORoot hash.Hash `json:"io_root"`

	// StorageSignatures are the storage receipt signatures for the I/O root.
	StorageSignatures []signature.Signature `json:"storage_signatures"`

	// Header is the block header on which the batch should be based.
	Header block.Header `json:"header"`
}

// SignedProposedBatch is a ProposedBatch, signed by
// the transaction scheduler.
type SignedProposedBatch struct {
	signature.Signed
}

// Equal compares vs another SignedProposedBatch for equality.
// removed func

// Open first verifies the blob signature and then unmarshals the blob.
// removed func

// SignProposedBatch signs a ProposedBatch struct using the
// given signer.
// removed func

// GetTransactionScheduler returns the transaction scheduler of the provided
// committee based on the provided round.
// removed func
