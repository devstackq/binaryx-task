package models

import "time"

type Account struct {
	CurrencyName string    `json:"currentname"`
	CurrencyCost float64   `json:"currencycost"`
	UUID         string    `json:"uuid"`
	Balance      float64   `json:"balance"`
	CurrencyId   int       `json:"currencyid"`
	Email        string    `json:"email"`
	Recepient    string    `json:"recepient"`
	Amount       float64   `json:"amount"`
	Date         time.Time `json:"senttime"`
}
