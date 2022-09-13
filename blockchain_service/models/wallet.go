package wallet


type Wallet struct {
	WalletAddress   string  `json:"WalletAddress"`
	CurrencyBalance float64 `json:"CurrencyBalance"`
	CurrencyCode    string  `json:"CurrencyCode"`
}

type AllWallets struct {
	Wallets []*Wallet
}
