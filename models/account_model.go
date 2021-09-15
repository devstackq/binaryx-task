package models

type Wallet struct {
	UID      string
	Currency struct{} //btc || eth

}

type AccountEth struct {
	Name    string
	Balance int
}

type AccountBtc struct {
	Name    string
	Balance int
}
