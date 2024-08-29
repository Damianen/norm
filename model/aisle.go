package model

import (
	"database/sql"
	"strconv"
)

type Aisle struct {
	Id   string
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

func GetAisle(db *sql.DB, idString string) (Aisle, error) {
    id, err := strconv.Atoi(idString)
    if err != nil {
        return Aisle{}, err
    }

    query := "SELECT * FROM Aisle WHERE id = $1;"

    aisle := Aisle{}

    err = db.QueryRow(query, id).Scan(&aisle.Id, &aisle.Name)
    if err != nil {
        return Aisle{}, err
    }

   return aisle, nil
}

func GetAisles(db *sql.DB) ([]Aisle, error) {
    query := "SELECT * FROM Aisle;"
    data, err := db.Query(query)
    if err != nil {
        return nil, err
    }

    aisle := []Aisle{}
    var id string
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

func DeleteAisle(db *sql.DB, aisleId int) error {
    query := "DELETE FROM Aisle WHERE id = $1"
    result, err := db.Exec(query, aisleId)
    if err != nil && result != nil {
        return err
    }
    return nil
}

func UpdateAisle(db *sql.DB, aisle Aisle) error {
    query := "UPDATE Aisle SET name = $1 WHERE id = $2;"
    result, err := db.Exec(query, aisle.Name, aisle.Id)
    if err != nil && result != nil {
        return err
    }
    return nil
}
