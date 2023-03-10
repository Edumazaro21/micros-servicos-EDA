package database

import (
	"database/sql"
	"github.com.br/edumazaro21/fc-ms-wallet/internal/entity"
)

type AccountDB struct {
	DB *sql.DB
}

func NewAccountDB(db *sql.DB) *AccountDB {
	return &AccountDB{
		DB: db,
	}
}

func (db *AccountDB) FindByID(id string) (*entity.Account, error) {
	var account entity.Account
	var client entity.Client
	account.Client = &client

	stmt, err := db.DB.Prepare("SELECT a.id, a.client_id, a.balance, a.created_at, c.id, c.name, c.email, c.created_at, c.updated_at FROM accounts a INNER JOIN clients c ON a.client_id = c.id WHERE a.id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	if err = row.Scan(
		&account.ID,
		&account.Client.ID,
		&account.Balance,
		&account.CreatedAt,
		&client.ID,
		&client.Name,
		&client.Email,
		&client.CreatedAt,
		&client.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return &account, nil
}

func (db *AccountDB) Save(account *entity.Account) error {
	stmt, err := db.DB.Prepare("INSERT INTO accounts (id, client_id, balance, created_at) VALUES (?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(account.ID, account.Client.ID, account.Balance, account.CreatedAt)
	return err
}
