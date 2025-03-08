package structs

type LiFiTransaction struct {
	Type          string        `json:"type"`
	ID            string        `json:"id"`
	Tool          string        `json:"tool"`
	ToolDetails   ToolDetails   `json:"toolDetails"`
	Action        Action        `json:"action"`
	Estimate      Estimate      `json:"estimate"`
	IncludedSteps []IncludedStep `json:"includedSteps"`
	Integrator    string        `json:"integrator"`
	TransactionRequest TransactionRequest `json:"transactionRequest"`
}

type ToolDetails struct {
	Key     string `json:"key"`
	Name    string `json:"name"`
	LogoURI string `json:"logoURI"`
}

type Token struct {
	Address  string `json:"address"`
	ChainID  int    `json:"chainId"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
	Name     string `json:"name"`
	CoinKey  string `json:"coinKey"`
	LogoURI  string `json:"logoURI"`
	PriceUSD string `json:"priceUSD"`
}

type Action struct {
	FromToken    Token   `json:"fromToken"`
	FromAmount   string  `json:"fromAmount"`
	ToToken      Token   `json:"toToken"`
	FromChainID  int     `json:"fromChainId"`
	ToChainID    int     `json:"toChainId"`
	Slippage     float64 `json:"slippage"`
	FromAddress  string  `json:"fromAddress"`
	ToAddress    string  `json:"toAddress"`
}

type FeeCost struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Token       Token  `json:"token"`
	Amount      string `json:"amount"`
	AmountUSD   string `json:"amountUSD"`
	Percentage  string `json:"percentage"`
	Included    bool   `json:"included"`
}

type GasCost struct {
	Type      string `json:"type"`
	Price     string `json:"price"`
	Estimate  string `json:"estimate"`
	Limit     string `json:"limit"`
	Amount    string `json:"amount"`
	AmountUSD string `json:"amountUSD"`
	Token     Token  `json:"token"`
}

type Estimate struct {
	Tool             string    `json:"tool"`
	ApprovalAddress  string    `json:"approvalAddress"`
	ToAmountMin      string    `json:"toAmountMin"`
	ToAmount         string    `json:"toAmount"`
	FromAmount       string    `json:"fromAmount"`
	FeeCosts         []FeeCost `json:"feeCosts"`
	GasCosts         []GasCost `json:"gasCosts"`
	ExecutionDuration int       `json:"executionDuration"`
	FromAmountUSD    string    `json:"fromAmountUSD"`
	ToAmountUSD      string    `json:"toAmountUSD"`
}

type IncludedStep struct {
	ID           string        `json:"id"`
	Type         string        `json:"type"`
	Action       Action        `json:"action"`
	Estimate     Estimate      `json:"estimate"`
	Tool         string        `json:"tool"`
	ToolDetails  ToolDetails   `json:"toolDetails"`
}

type TransactionRequest struct {
	Data     string   `json:"data"`
	To       string   `json:"to"`
	Value    string   `json:"value"`
	From     string   `json:"from"`
	ChainID  int      `json:"chainId"`
	GasPrice string   `json:"gasPrice"`
	GasLimit string   `json:"gasLimit"`
}