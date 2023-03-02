package api

import (
	"github.com/oasisprotocol/oasis-core/go/common"
	"github.com/oasisprotocol/oasis-core/go/common/crypto/signature"
	"github.com/oasisprotocol/oasis-core/go/common/node"
)

// SanityCheck does basic sanity checking on the genesis state.
// removed func

// SanityCheckEntities examines the entities table.
// Returns lookup of entity ID to the entity record for use in other checks.
// removed func

// SanityCheckRuntimes examines the runtimes table.
// removed func

// SanityCheckNodes examines the nodes table.
// Pass lookups of entities and runtimes from SanityCheckEntities
// and SanityCheckRuntimes for cross referencing purposes.
// removed func

// SanityCheckStake ensures entities' stake accumulator claims are consistent
// with general state and entities have enough stake for themselves and all
// their registered nodes and runtimes.
// removed func

// Runtimes lookup used in sanity checks.
type sanityCheckRuntimeLookup struct {
	runtimes          map[common.Namespace]*Runtime
	suspendedRuntimes map[common.Namespace]*Runtime
	allRuntimes       []*Runtime
}

// removed func

// removed func

// removed func

// removed func

// removed func

// Node lookup used in sanity checks.
type sanityCheckNodeLookup struct {
	nodes        map[signature.PublicKey]*node.Node
	nodesByPoint map[string]*node.Node

	nodesList []*node.Node
}

// removed func

// removed func

// removed func
