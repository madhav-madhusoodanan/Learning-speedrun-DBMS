package structs

type CheckUsernameRequest struct {
	Username            string        `json:"username"`
}

type CheckUserAddressRequest struct {
	Address            string        `json:"address"`
	ChainId 		   int64        `json:"chainId"`
}

type CreateWalletRequest struct {
	Username           string        `json:"username"`
	Address            string        `json:"address"`
	ChainId 		   int64        `json:"chainId"`
}