package structs

type SwapRequest struct {
	SourceChain        int32   `json:"source_chain" binding:"required"`
	DestinationChain   int32   `json:"destination_chain" binding:"required"`
	SourceToken        string  `json:"source_token" binding:"required,len=42"` // Ethereum address length
	DestinationToken   string  `json:"destination_token" binding:"required,len=42"`
	SourceAmount       string  `json:"source_amount" binding:"required"`       // Using string to handle large numbers
	DestinationAmount  string  `json:"destination_amount" binding:"required"`
	DollarValue        string  `json:"dollar_value" binding:"required"`
	SourceAddress      string  `json:"source_address" binding:"required,len=42"`
	DestinationAddress string  `json:"destination_address" binding:"required,len=42"`
	TransactionHash    string  `json:"transaction_hash" binding:"required,len=66"` // 0x + 32 bytes hex
}

// RequestTokenDetail represents the JSON structure for token details in the request
type RequestTokenDetail struct {
	TokenAddress string `json:"token" binding:"required"`
	Amount       string `json:"amount" binding:"required"`     // Amount as string to handle big integers
}

// RequestRoute represents the JSON structure for route in the request
type RequestRoute struct {
	ProcessorIndex string   `json:"processorIndex" binding:"required"` // ProcessorIndex as string to handle big integers
	SellTokenDetails []RequestTokenDetail `json:"sellTokenDetails" binding:"required"`
	Payload       string `json:"payload" binding:"required"`     // Hex string
	Metadata      string `json:"metadata" binding:"required"`    // Hex string
}

// SwapRequest represents the complete swap request structure
type SwapRequestRampX struct {
	UserAddress   string            `json:"userAddress" binding:"required"`
	SellTokens    []RequestTokenDetail `json:"sellTokens" binding:"required"`
	BuyTokens     []RequestTokenDetail `json:"buyTokens" binding:"required"`
	Routes        []RequestRoute    `json:"routes" binding:"required"`
	UserSignature string            `json:"userSignature" binding:"required"` // Hex string
	ChainID       int64             `json:"chainId" binding:"required"`
}