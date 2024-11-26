package structs

type AddSwapRequest struct {
	SourceChainID        int64   `json:"source_chain_id" binding:"required"`
    SourceTokenAddress   string  `json:"source_token_address" binding:"required"`
    SourceAmount        string  `json:"source_amount" binding:"required"`      // Using string for large numbers
    DestChainID         int64   `json:"destination_chain_id" binding:"required"`
    DestTokenAddress    string  `json:"destination_token_address" binding:"required"`
    DestAmount         string  `json:"destination_amount" binding:"required"` // Using string for large numbers
    DollarValue        float64 `json:"dollar_value" binding:"required"`
    FeeDollarValue     float64 `json:"fee_dollar_value" binding:"required"`
    SourceAddress      string  `json:"source_address" binding:"required,len=42"` // For Ethereum addresses
    DestinationAddress string  `json:"destination_address" binding:"required,len=42"`
    TransactionHash    string  `json:"transaction_hash" binding:"required,len=66"` // For Ethereum tx hash
}