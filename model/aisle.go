package model

import "database/sql"

type Aisle struct {
	Id   int
	Name string
}


func InsertAilse(db *sql.DB, aisle Aisle) error {
    query := `INSERT INTO Aisle (name) VALUES ($1) RETURNING id;`

    var pk int
    err := db.QueryRow(query, aisle.Name).Scan(&pk)
    if err != nil {
        return err
    }

    return nil
}

func GetAisles(db *sql.DB) ([]Aisle, error) {
    query := "SELECT * FROM Aisle;"
    data, err := db.Query(query)
    if err != nil {
        return nil, err
    }

    aisle := []Aisle{}
    var id int
    var name string

    for data.Next() {
        err := data.Scan(&id, &name)
        if err != nil {
            return nil, err
        }
        aisle = append(aisle, Aisle{id, name})
    }

    return aisle, nil
}
