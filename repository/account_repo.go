package repository

import (
	"database/sql"

	"github.com/devstackq/binaryx/models"
)

type WalletRepository struct {
	db *sql.DB
}

func NewWalletRepository(db *sql.DB) *WalletRepository {
	return &WalletRepository{db}
}

//add new currency, set def cost, then update
func (ar *WalletRepository) AddCurrency(name string, cost float64) error {
	sqlStmt := `
	INSERT INTO currencies (name, cost) VALUES ($1,$2)`
	_, err := ar.db.Exec(sqlStmt, name, cost)
	if err != nil {
		return err
	}
	return nil
}

//currency eth -> wallet1, etc
func (ar *WalletRepository) InitBalance(acc *models.Account) error {
	sqlStmt := `INSERT INTO wallets (balance, currencyid, uuid) VALUES ($1, $2, $3) `
	_, err := ar.db.Exec(sqlStmt, acc.Balance, acc.CurrencyId, acc.UUID)
	if err != nil {
		return err
	}
	return nil
}
