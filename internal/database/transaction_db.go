package database

import (
	"database/sql"
	"github.com.br/edumazaro21/fc-ms-wallet/internal/entity"
)

type TransactionDB struct {
	DB *sql.DB
}

func NewTransactionDB(db *sql.DB) *TransactionDB {
	return &TransactionDB{
		DB: db,
	}
}

func (db *TransactionDB) Create(transaction *entity.Transaction) error {
	stmt, err := db.DB.Prepare("INSERT INTO transactions (id, from_id, to_id, amount, created_at) VALUES (?,?,?,?,?) ")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(transaction.ID, transaction.AccountFrom.ID, transaction.AccountTo.ID, transaction.Amount, transaction.CreatedAt)
	return err
}
