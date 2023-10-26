package node

import (
	"fmt"

	"github.com/oasisprotocol/oasis-core/go/common/cbor"
	"github.com/oasisprotocol/oasis-core/go/common/crypto/signature"
	"github.com/oasisprotocol/oasis-core/go/common/sgx"
	"github.com/oasisprotocol/oasis-core/go/common/sgx/ias"
	"github.com/oasisprotocol/oasis-core/go/common/sgx/quote"
)

const (
	// LatestSGXConstraintsVersion is the latest SGX constraints structure version that should be
	// used for all new descriptors.
	LatestSGXConstraintsVersion = 1
)

var emptyFeatures TEEFeatures

// SGXConstraints are the Intel SGX TEE constraints.
type SGXConstraints struct {
	cbor.Versioned

	// Enclaves is the allowed MRENCLAVE/MRSIGNER pairs.
	Enclaves []sgx.EnclaveIdentity `json:"enclaves,omitempty"`

	// Policy is the quote policy.
	Policy *quote.Policy `json:"policy,omitempty"`

	// MaxAttestationAge is the maximum attestation age (in blocks).
	MaxAttestationAge uint64 `json:"max_attestation_age,omitempty"`
}

// sgxConstraintsV0 are the version 0 Intel SGX TEE constraints which only supports IAS.
type sgxConstraintsV0 struct {
	Enclaves             []sgx.EnclaveIdentity       `json:"enclaves,omitempty"`
	AllowedQuoteStatuses []ias.ISVEnclaveQuoteStatus `json:"allowed_quote_statuses,omitempty"`
}

// UnmarshalCBOR is a custom deserializer that handles different structure versions.
func (sc *SGXConstraints) UnmarshalCBOR(data []byte) error {
	// Determine Entity structure version.
	v, err := cbor.GetVersion(data)
	if err != nil {
		v = 0 // Previous SGXConstraints structures were not versioned.
	}
	switch v {
	case 0:
		// Old version only supported the IAS-related constraints.
		var scv0 sgxConstraintsV0
		if err = cbor.Unmarshal(data, &scv0); err != nil {
			return err
		}

		// Convert into new format.
		sc.Versioned = cbor.NewVersioned(0)
		sc.Enclaves = scv0.Enclaves
		sc.Policy = &quote.Policy{
			IAS: &ias.QuotePolicy{
				AllowedQuoteStatuses: scv0.AllowedQuoteStatuses,
			},
		}
		return nil
	case 1:
		// New version, call the default unmarshaler.
		type scv1 SGXConstraints
		return cbor.Unmarshal(data, (*scv1)(sc))
	default:
		return fmt.Errorf("invalid SGX constraints version: %d", v)
	}
}

// MarshalCBOR is a custom serializer that handles different structure versions.
func (sc *SGXConstraints) MarshalCBOR() ([]byte, error) {
	switch sc.V {
	case 0:
		// Old version only supported the IAS-related constraints.
		scv0 := sgxConstraintsV0{
			Enclaves: sc.Enclaves,
		}
		if sc.Policy != nil && sc.Policy.IAS != nil {
			scv0.AllowedQuoteStatuses = sc.Policy.IAS.AllowedQuoteStatuses
		}
		return cbor.Marshal(scv0), nil
	default:
		type scv1 SGXConstraints
		return cbor.Marshal((*scv1)(sc)), nil
	}
}

// ValidateBasic performs basic structure validity checks.
// removed func

// ContainsEnclave returns true iff the allowed enclave list in SGX constraints contain the given
// enclave identity.
// removed func

const (
	// LatestSGXAttestationVersion is the latest SGX attestation structure version that should be
	// used for all new descriptors.
	LatestSGXAttestationVersion = 1
)

// SGXAttestation is an Intel SGX remote attestation.
type SGXAttestation struct {
	cbor.Versioned

	// Quote is an Intel SGX quote.
	Quote quote.Quote `json:"quote"`

	// Height is the runtime's view of the consensus layer height at the time of attestation.
	Height uint64 `json:"height"`

	// Signature is the signature of the attestation by the enclave (RAK).
	Signature signature.RawSignature `json:"signature"`
}

// UnmarshalCBOR is a custom deserializer that handles different structure versions.
func (sa *SGXAttestation) UnmarshalCBOR(data []byte) error {
	// Determine Entity structure version.
	v, err := cbor.GetVersion(data)
	if err != nil {
		v = 0 // Previous SGXAttestation structures were not versioned.
	}
	switch v {
	case 0:
		// Old version only supported the IAS attestation.
		var sav0 ias.AVRBundle
		if err = cbor.Unmarshal(data, &sav0); err != nil {
			return err
		}

		// Convert into new format.
		sa.Versioned = cbor.NewVersioned(0)
		sa.Quote = quote.Quote{
			IAS: &sav0,
		}
		return nil
	case 1:
		// New version, call the default unmarshaler.
		type sav1 SGXAttestation
		return cbor.Unmarshal(data, (*sav1)(sa))
	default:
		return fmt.Errorf("invalid SGX attestation version: %d", v)
	}
}

// MarshalCBOR is a custom serializer that handles different structure versions.
func (sa *SGXAttestation) MarshalCBOR() ([]byte, error) {
	switch sa.V {
	case 0:
		// Old version only supported the IAS attestation.
		return cbor.Marshal(sa.Quote.IAS), nil
	default:
		type sav1 SGXAttestation
		return cbor.Marshal((*sav1)(sa)), nil
	}
}

// ValidateBasic performs basic structure validity checks.
// removed func

// Verify verifies the SGX attestation.
// removed func

// removed func

// HashAttestation hashes the required data that needs to be signed by RAK producing the attestation
// signature. The hash is computed as follows:
//
//	TupleHash[AttestationSignatureContext](reportData, nodeID, height)
// removed func
