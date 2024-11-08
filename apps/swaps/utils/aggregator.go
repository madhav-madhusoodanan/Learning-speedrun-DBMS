package utils

var aggregators = make(map[int64]string)

// Initialize adds strings to the map for a given number key
func InitializeContracts() {
	// Example initialization
	aggregators[10] = "0x147B8eb97fD247D06C4006D269c90C1908Fb5D54"
	aggregators[8453] = "0x147B8eb97fD247D06C4006D269c90C1908Fb5D54"
	aggregators[31337] = "0xcf7ed3acca5a467e9e704c703e8d87f634fb0fc9"
}

func GetAggregator(chainId int64) string {
	return aggregators[chainId]
}