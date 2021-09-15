package handler

import "github.com/devstackq/binaryx/models"

//счет
type Customer interface {
	GetAccount()
	UpdateAccount()
	TransferMoney()
}

func SetDefaultCurrency(user *models.User) {
	//, get registered user -> email, realation, Wallet - currency1, currency2 -> setDef vBalamnce 100, save in sql Db

}
