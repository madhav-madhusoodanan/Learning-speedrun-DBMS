package utils

import (
	"time"
)

// Global map to store number-to-strings mapping
var rpcMap = make(map[int64][]string)

// Initialize adds strings to the map for a given number key
func InitializeRPC() {
	// Example initialization
	rpcMap[10] = []string{"https://1rpc.io/op", "https://optimism.llamarpc.com", "https://mainnet.optimism.io"}
	rpcMap[8453] = []string{"https://1rpc.io/base", "https://base.drpc.org", "https://mainnet.base.org"}
	rpcMap[31337] = []string{"http://127.0.0.1:8545", "http://127.0.0.1:8545"}
}

// GetRandomString returns a random string from the array associated with the given number
// Returns empty string if key doesn't exist
func GetRPC(chainId int64) string {
	// Check if the key exists and has strings
	strings, exists := rpcMap[chainId]
	if !exists || len(strings) == 0 {
		return ""
	}
	
	// Generate random index
	maxLen := len(strings)
	index := time.Now().UTC().UnixMilli() % int64(maxLen)

	return strings[index]
}
