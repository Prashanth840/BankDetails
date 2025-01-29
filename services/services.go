package services

import (
	"bankdetails/data"
	"bankdetails/models"
	"time"
)

func CreateAccount(input *models.Account) error {
	str := `insert into accounts(name,balance,status,created_at) values(?,?,?,?)`
	row, err := data.Db.Exec(str, input.Name, input.Balance, input.Status, input.CreatedAt)
	a, _ := row.LastInsertId()
	input.ID = int(a)
	return err
}

func CreateTransaction(input models.Transaction) error {
	input.Timestamp = time.Now()
	str := `insert into transactions(account_id,transaction_type,amount,timestamp) values(?,?,?,?)`
	_, err := data.Db.Exec(str, input.AccountID, input.TransactionType, input.Amount, input.Timestamp)

	return err
}

func GetAccountByID(account_id int) (models.Account, error) {
	var result models.Account
	str := `select id,name,balance,status from accounts where id=?`
	if err := data.Db.QueryRow(str, account_id).Scan(&result.ID, &result.Name, &result.Balance, &result.Status); err != nil {

		return result, err
	}

	return result, nil
}

func UpdateAccount(input models.Account) error {
	str := `update accounts 
			set balance=?
			where id=?`
	_, err := data.Db.Exec(str, input.Balance, input.ID)
	return err
}
