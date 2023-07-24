package v3

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/log"
	"golang.org/x/crypto/sha3"
	"math/big"
)

func keccak256(data []byte) []byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(data)
	return hash.Sum(nil)
}

func getStorageSlotIndex(key int64, mappingSlotIndex *big.Int) []byte {
	// Define the expected arguments
	int16Type, err := abi.NewType("uint256", "16", nil)
	if err != nil {
		log.Error("RonFi getStorageSlotIndex", "err", err)
		return []byte{}
	}

	uint256Type, err := abi.NewType("uint256", "256", nil)
	if err != nil {
		log.Error("RonFi getStorageSlotIndex", "err", err)
		return []byte{}
	}

	args := abi.Arguments{
		abi.Argument{
			Type: int16Type,
		},
		abi.Argument{
			Type: uint256Type,
		},
	}

	// ABI encoding of the input data
	packedData, err := args.Pack(big.NewInt(key), mappingSlotIndex)
	if err != nil {
		log.Error("RonFi getStorageSlotIndex", "err", err)
		return []byte{}
	}

	// Calculate the Keccak-256 hash
	hashed := keccak256(packedData)

	return hashed
}
