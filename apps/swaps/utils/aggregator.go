package utils

var aggregators = make(map[int64]string)

// Initialize adds strings to the map for a given number key
func InitializeContracts() {
	// Example initialization
	aggregators[10] = "0xfA07E4342F482Eed28DCdCB201db5DeC690894A0"
	aggregators[8453] = "0x6f1bd0c04b02cB27EFdE27F28cbc464150a6d2eD"
	aggregators[31337] = "0xcf7ed3acca5a467e9e704c703e8d87f634fb0fc9"
}

func GetAggregator(chainId int64) string {
	return aggregators[chainId]
}