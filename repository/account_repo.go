package repository

import (
	"database/sql"
	"log"

	"github.com/devstackq/binaryx/models"
)

type WalletRepository struct {
	db *sql.DB
}

func NewWalletRepository(db *sql.DB) *WalletRepository {
	return &WalletRepository{db}
}
func (wr *WalletRepository) Transfer(account *models.Account, recepient *models.Account) error {
	// try 1 sql query, dry
	account.Balance -= account.Amount
	recepient.Balance += account.Amount

	sqlSender := `
		UPDATE wallets
		SET balance = $1
		WHERE uuid = $2 AND currencyid=$3;
		`
	_, err := wr.db.Exec(sqlSender, account.Balance, account.UUID, account.CurrencyId)
	if err != nil {
		return err
	}

	sqlRecepient := `
		UPDATE wallets
		SET balance = $1
		WHERE uuid = $2 AND currencyid=$3;
		`
	_, err = wr.db.Exec(sqlRecepient, recepient.Balance, recepient.UUID, recepient.CurrencyId)
	if err != nil {
		return err
	}
	return nil
}

func (wr *WalletRepository) CheckWallet(account *models.Account) (*models.Account, error) {
	//email - uuid, - wallets -> get balance
	log.Print(account)
	uuid, err := wr.GetUUIDByEmail(account.Email)
	if err != nil {
		return nil, err
	}
	account.UUID = uuid

	sqlStatement := `SELECT balance FROM wallets WHERE uuid=$1 AND currencyid =$2;`
	row := wr.db.QueryRow(sqlStatement, account.UUID, account.CurrencyId)
	err = row.Scan(&account.Balance)
	if err != nil {
		return nil, err
	}
	return account, nil
}
func (wr *WalletRepository) GetAccountsByEmail(email string) ([]models.Account, error) {
	uuid, err := wr.GetUUIDByEmail(email)
	if err != nil {
		return nil, err
	}
	var acc models.Account
	var seqAcc []models.Account
	rows, err := wr.db.Query("SELECT c.name, c.cost, w.balance FROM wallets w LEFT JOIN currencies c ON c.id = w.currencyid  WHERE w.uuid=$1  ", uuid)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		//query getUsername byPostId
		if err := rows.Scan(&acc.CurrencyName, &acc.CurrencyCost, &acc.Balance); err != nil {
			return nil, err
		}
		seqAcc = append(seqAcc, acc)
	}
	return seqAcc, nil
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

func (wr *WalletRepository) GetUUIDByEmail(email string) (string, error) {
	var uuid string
	sqlStatement := `SELECT id FROM users WHERE email=$1;`
	row := wr.db.QueryRow(sqlStatement, email)
	err := row.Scan(&uuid)
	if err != nil {
		return "", err
	}
	return uuid, nil
}

//set new user - 2 currency, btc, eth
func (wr *WalletRepository) InitBalance(acc *models.Account) error {

	sqlStatement := `SELECT id FROM users WHERE email=$1;`
	row := wr.db.QueryRow(sqlStatement, acc.Email)
	err := row.Scan(&acc.UUID)
	if err != nil {
		return err
	}
	acc.UUID, err = wr.GetUUIDByEmail(acc.Email)
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
