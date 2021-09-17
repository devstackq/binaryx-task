package service

import (
	"log"

	"github.com/devstackq/binaryx/models"
	"github.com/devstackq/binaryx/repository"
)

type WalletService struct {
	repository repository.Wallet
}

func NewWalletService(repo repository.Wallet) *WalletService {
	return &WalletService{repo}

}
func (ws *WalletService) GetAccounts(email string) ([]models.Account, error) {
	seqAcc, err := ws.repository.GetAccountsByEmail(email)
	if err != nil {
		return nil, err
	}
	return seqAcc, nil
}
func (ws *WalletService) AddCurrency(name string, cost float64) error {
	//valid params
	err := ws.repository.AddCurrency(name, cost)
	if err != nil {
		return err
	}
	return nil
}
func (ws *WalletService) TransferMoney(acc *models.Account) error {
	var err error
	//current user check balance
	uuid, err := ws.repository.GetUUIDByEmail(acc.Email)
	if err != nil {
		log.Println(err)
		return err
	}
	acc.UUID = uuid

	acc, err = ws.repository.CheckWallet(acc)
	if err != nil {
		log.Println(err)
		return err
	}
	if acc.Balance >= acc.Amount {
		rec := models.Account{
			Email:      acc.Recepient,
			CurrencyId: acc.CurrencyId,
		}

		uuid, err := ws.repository.GetUUIDByEmail(rec.Email)
		if err != nil {
			log.Println(err)
		}
		rec.UUID = uuid
		//get balance by email
		recepient, err := ws.repository.CheckWallet(&rec)
		if err != nil {
			log.Print(err)
			return err
		}
		//current - amount, recepinet + amount
		err = ws.repository.Transfer(acc, recepient)
		if err != nil {
			log.Println(err)
			return err
		}
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
}
