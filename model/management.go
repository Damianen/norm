package model

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type Management struct {
	Id       int
	Pnl      string
	Password string
	Name     string
	Email    string
	Function string
}

func InsertManagement(db *sql.DB, sl Management) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(sl.Password), 14)
	if err != nil {
		return err
	}

	query := `INSERT INTO management (pnl, password, name, email, function) VALUES ($1, $2, $3, $4, $5);`
	hashedPassword := string(bytes)

	db.QueryRow(query, sl.Pnl, hashedPassword, sl.Name, sl.Email, sl.Function)

	return nil
}

func GetManagementWithPnl(db *sql.DB, pnl string) (Management, error) {
    management := Management{}

	query := "SELECT id, password, name, email, function FROM Management WHERE pnl = $1"
	err := db.QueryRow(query, pnl).Scan(&management.Id, &management.Password, &management.Name, &management.Email, &management.Function)
    if err != nil {
        return Management{}, err
    }

    management.Pnl = pnl
	return management, nil
}

func GetManagement(db *sql.DB) ([]Management, error) {

    query := "SELECT name, pnl, email, function FROM Management"
    data, err := db.Query(query)
    if err != nil {
        return nil, err
    }

    management := []Management{}
    var name string
    var pnl string
    var email string
    var function string

    for data.Next() {
        err := data.Scan(&name, &pnl, &email, &function)
        if err != nil {
            return nil, err
        }
        management = append(management, Management{-1, pnl, "", name, email, function})
    }

    return management, nil
}
