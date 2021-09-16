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
func (wr *WalletRepository) AddCurrency(name string, cost float64) error {
	sqlStmt := `
	INSERT INTO currencies (name, cost) VALUES ($1,$2)`
	_, err := wr.db.Exec(sqlStmt, name, cost)
	if err != nil {
		return err
	}
	return nil
}

//next step, get list account: by email, get Id, by uuid - get List wallets left join - currencies,

//currency eth -> wallet1, etc
func (wr *WalletRepository) InitBalance(acc *models.Account) error {

	sqlStatement := `SELECT id FROM users WHERE email=$1;`
	row := wr.db.QueryRow(sqlStatement, acc.Email)
	err := row.Scan(&acc.UUID)
	if err != nil {
		return err
	}

	sqlStmt := `INSERT INTO wallets (balance, currencyid, uuid) VALUES ($1, $2, $3) `
	_, err = wr.db.Exec(sqlStmt, acc.Balance, acc.CurrencyId, acc.UUID)
	if err != nil {
		return err
	}
	return nil
}
