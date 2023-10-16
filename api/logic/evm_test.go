package logic

import (
	_ "embed"
	"encoding/hex"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/oasisprotocol/nexus/analyzer/evmabi"
)

//go:embed test_contracts/artifacts/Varied.json
var artifactVariedJSON []byte
var Varied *abi.ABI

func init() {
	type artifact struct {
		ABI *abi.ABI
	}
	var artifactVaried artifact
	if err := json.Unmarshal(artifactVariedJSON, &artifactVaried); err != nil {
		panic(err)
	}
	Varied = artifactVaried.ABI
}

func TestEVMParseTypes(t *testing.T) {
	data, err := hex.DecodeString("a85acffaffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000000000000000000000000000000000000000000000000001ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000010101010101010101010101010101010101010101010101010101010101010101000000000000000000000000010101010101010101010101010101010101010101010101010101010101010101010101010101010202020200000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000001c00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000024000000000000000000000000000000000000000000000000000000000000002800000000000000000000000000000000000000000000000000000000000000001010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000016100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000016100000000000000000000000000000000000000000000000000000000000000")
	require.NoError(t, err)
	method, args, err := EVMParseData(data, Varied)
	require.NoError(t, err)
	require.Equal(t, Varied.Methods["test"], *method)
	jsonExpected := []string{
		"-1",     // int8
		"1",      // uint8
		"\"-1\"", // int256
		"\"1\"",  // uint256
		"true",   // bool
		"\"0x0101010101010101010101010101010101010101010101010101010101010101\"", // bytes32
		"\"0x0101010101010101010101010101010101010101\"",                         // address
		"\"0x010101010101010101010101010101010101010102020202\"",                 // function (uint16) external returns (uint16)
		"[1,1]",                 // uint16[2]
		"\"0x01\"",              // bytes
		"\"a\"",                 // string
		"[1]",                   // uint16[]
		"{\"n\":1,\"s\":\"a\"}", // O
	}
	for i, input := range method.Inputs {
		transducedArg := evmPreMarshal(args[i], input.Type)
		jsonBytesArg, err1 := json.Marshal(transducedArg)
		require.NoError(t, err1)
		require.Equal(t, jsonExpected[i], string(jsonBytesArg))
	}
}

func TestEVMParseData(t *testing.T) {
	// https://explorer.emerald.oasis.dev/tx/0x1ac7521df4cda38c87cff56b1311ee9362168bd794230415a37f2aff3a554a5f/internal-transactions
	data, err := hex.DecodeString("095ea7b3000000000000000000000000250d48c5e78f1e85f7ab07fec61e93ba703ae6680000000000000000000000000000000000000000000000003782dace9d900000")
	require.NoError(t, err)
	method, args, err := EVMParseData(data, evmabi.ERC20)
	require.NoError(t, err)
	require.Equal(t, evmabi.ERC20.Methods["approve"], *method)
	require.Equal(t, []interface{}{
		ethCommon.HexToAddress("0x250d48c5e78f1e85f7ab07fec61e93ba703ae668"),
		big.NewInt(4000000000000000000),
	}, args)
}

func TestEVMParseEvent(t *testing.T) {
	// https://explorer.emerald.oasis.dev/tx/0x1ac7521df4cda38c87cff56b1311ee9362168bd794230415a37f2aff3a554a5f/logs
	topicsHex := []string{
		"8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925",
		"0000000000000000000000000ecf5262e5b864e1612875f8fc18f151315b5e91",
		"000000000000000000000000250d48c5e78f1e85f7ab07fec61e93ba703ae668",
	}
	topics := make([][]byte, 0, len(topicsHex))
	for _, topicHex := range topicsHex {
		topic, err := hex.DecodeString(topicHex)
		require.NoError(t, err)
		topics = append(topics, topic)
	}
	data, err := hex.DecodeString("0000000000000000000000000000000000000000000000003782dace9d900000")
	require.NoError(t, err)
	event, args, err := EVMParseEvent(topics, data, evmabi.ERC20)
	require.NoError(t, err)
	require.Equal(t, evmabi.ERC20.Events["Approval"], *event)
	require.Equal(t, []interface{}{
		big.NewInt(4000000000000000000),
	}, args)
}
