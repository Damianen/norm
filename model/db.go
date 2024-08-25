package model

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func DbInit() *sql.DB {
    connStr := os.Getenv("CONNECTION_STRING")

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

func CreateTables(db *sql.DB) {
	query := `DROP TABLE IF EXISTS Stocker CASCADE;
        DROP TABLE IF EXISTS Comments CASCADE;
        DROP TABLE IF EXISTS Aisle CASCADE;
        DROP TABLE IF EXISTS Raport CASCADE;
        DROP TABLE IF EXISTS Shift CASCADE;
        DROP TABLE IF EXISTS Shiftleader CASCADE;
        DROP TABLE IF EXISTS AisleStocker CASCADE;
    `

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	query = `CREATE TABLE Stocker (
	    id SERIAL PRIMARY KEY,
	    name varchar(255) NOT NULL,
	    code bool DEFAULT false,
	    afprijzen bool DEFAULT false,
	    splits bool DEFAULT false,
	    ladenEnLossen bool DEFAULT false,
	    buddy bool DEFAULT false,
	    dagres bool DEFAULT false,
	    actie bool DEFAULT false
	)`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	query = `CREATE TABLE Aisle (
	    id SERIAL PRIMARY KEY,
	    name varchar(255) NOT NULL
	)`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	query = `CREATE TABLE Shiftleader(
	    id SERIAL PRIMARY KEY,
	    name varchar(255) NOT NULL
	)`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	query = `CREATE TABLE Raport(
	    id SERIAL PRIMARY KEY,
	    message varchar(1000) NOT NULL,
	    recepientId int,
	    FOREIGN KEY (recepientId) REFERENCES Shiftleader(id)
	)`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	query = `CREATE TABLE Comments(
	    id SERIAL PRIMARY KEY,
	    dossierNote bool DEFAULT false,
	    message varchar(1000) NOT NULL,
	    stockerId int NOT NULL,
	    FOREIGN KEY (stockerId) REFERENCES Stocker(id)
	)`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	query = `CREATE TABLE Shift(
	    id SERIAL PRIMARY KEY,
	    raportId int,
	    shiftleaderVers int NOT NULL,
	    shiftleaderHB int NOT NULL,
        FOREIGN KEY (shiftleaderVers) REFERENCES Shiftleader(id),
	    FOREIGN KEY (shiftleaderHB) REFERENCES Shiftleader(id)
	)`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	query = `CREATE TABLE AisleStocker(
	    id SERIAL PRIMARY KEY,
	    stockerId int NOT NULL,
	    aisleId int NOT NULL,
	    shiftId int NOT NULL,
	    norm float,
	    normalTime time NOT NULL,
	    timeWorked time NOT NULL,
	    FOREIGN KEY (stockerId) REFERENCES Stocker(id),
	    FOREIGN KEY (aisleId) REFERENCES Aisle(id),
	    FOREIGN KEY (shiftId) REFERENCES Shift(id)
	)`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
