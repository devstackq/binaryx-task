package service

import (
	"github.com/devstackq/binaryx/models"
	"github.com/devstackq/binaryx/repository"
)

type WalletService struct {
	repository repository.Wallet
}

func NewWalletService(repo repository.Wallet) *WalletService {
	return &WalletService{repo}
}
func (ws *WalletService) AddCurrency(name string, cost float64) error {

	err := ws.repository.AddCurrency(name, cost)
	if err != nil {
		return err
	}
	return nil
}

func (ws *WalletService) InitBalance(w *models.Account) error {
	//check valid data, if user exist && web token valid ?
	err := ws.repository.InitBalance(w)
	if err != nil {
		return err
	}
	return nil
	//get registered user -> email, realation, Wallet - currency1, currency2 -> setDef vBalamnce 100, save in sql Db
	// get by email uuid -> insert 2 new balance , then - every time update
}
