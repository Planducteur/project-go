package player

type User struct {
	Id       int `json:"Id"`
	Username string `json:"Username"`
	Password string `json:"Password"`
	PinCode  string `json:"PinCode"`
}

type AllUsers struct {
	Users []*User
}


type Wallet struct {
	PlayerId        int     `json:"PlayerId"`
	PinCode         string  `json:"PinCode"`
	WalletAddress   string  `json:"WalletAddress"`
	CurrencyBalance float64 `json:"CurrencyBalance"`
	CurrencyCode    string  `json:"CurrencyCode"`
}

type AllWallets struct {
	Wallets []*Wallet
}