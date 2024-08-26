package model

import (
	"database/sql"
)

type Stocker struct {
	Id            int
	Name          string
	Code          bool
	Afprijzen     bool
	Splits        bool
	LadenEnLossen bool
    Buddy         bool
	Dagres        bool
	Actie         bool
}

func InsertStocker(db *sql.DB, stocker Stocker) error {
    query := `INSERT INTO Stocker (name, code, afprijzen, splits,
        ladenEnLossen, buddy, dagres, actie) VALUES
        ($1, false, false, false, false, false, false, false)
        RETURNING id;`

    var pk int
    err := db.QueryRow(query, stocker.Name).Scan(&pk)
    if err != nil {
        return err
    }

    return nil
}

func GetStockers(db *sql.DB) ([]Stocker, error) {
    query := "SELECT * FROM Stocker;"
    data, err := db.Query(query)
    if err != nil {
        return nil, err
    }

    stocker := []Stocker{}
    var id int
    var name string
    var code bool
    var afprijzen bool
    var splits bool
    var ladenEnLossen bool
    var buddy bool
    var dagres bool
    var actie bool

    for data.Next() {
        err := data.Scan(&id, &name, &code, &afprijzen, &splits, &ladenEnLossen, &buddy, &dagres, &actie)
        if err != nil {
            return nil, err
        }
        stocker = append(stocker, Stocker{id, name, code, afprijzen, splits, ladenEnLossen, buddy, dagres, actie})
    }

    return stocker, nil
}
