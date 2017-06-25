package main

import (

	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
)

type PostgresBackend struct {
	dsn           string
	Database      *sql.DB
}



func NewPostgresConnection(dataSourceName string) (*PostgresBackend, error) {
	var err error
	postgresDB := &PostgresBackend{dsn: dataSourceName}
	postgresDB.Database, err = sql.Open("postgres", dataSourceName)
	postgresDB.Database.SetMaxOpenConns(10)
	if err != nil {
		return nil, err
	}

	rows, err := postgresDB.Database.Query("SELECT data FROM recepie LIMIT 1")
	if err != nil {
		_, err = postgresDB.Database.Exec(
            "CREATE TABLE recepie (recepie_id text UNIQUE, data jsonb)")
		if err != nil {
			return nil, err
		}
	} else {
		rows.Close()
	}
    return postgresDB, nil
}

func (postgres *PostgresBackend) CreateRecepie(r Recepie) error {
	id := r.RecepieID
	recJson, err := json.Marshal(r)
		if err != nil {
		return err
	}
	_, err = postgres.Database.Exec(
		"INSERT INTO recepie (recepie_id, data) VALUES ($1, $2)",
	    id, string(recJson))
	if err != nil {
		return err
	}
	return nil
}

func (postgres *PostgresBackend) Update(r Recepie) error {
	id := r.RecepieID
	recJson, err := json.Marshal(r)
		if err != nil {
		return err
	}
	_, err = postgres.Database.Exec(
		"UPDATE recepie SET data = $2 WHERE recepie_id = $1",
	    id, string(recJson))
	if err != nil {
		return err
	}
	return nil
}

func (postgres *PostgresBackend) UpdateOrCreateRecepie(r Recepie) error {
	id := r.RecepieID
	recJson, err := json.Marshal(r)
		if err != nil {
		return err
	}
	_, err = postgres.Database.Exec(`
		INSERT INTO recepie (recepie_id, data) VALUES ($1, $2)
		ON CONFLICT (recepie_id) DO UPDATE SET data = excluded.data`,
	    id, string(recJson))
	if err != nil {
		return err
	}
	return nil
}

func (postgres *PostgresBackend) ListRecepies() ([]Recepie, error) {
	var (
		data string
		recepie Recepie
		recepies []Recepie
	)
	rows, err := postgres.Database.Query("SELECT data FROM recepie")
	if err != nil {
		if err == sql.ErrNoRows {
			return []Recepie{}, nil
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&data)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(data), &recepie)
		if err != nil {
			return nil, err
		}
		recepies = append(recepies, recepie)
	}
	return recepies, nil
}

func (postgres *PostgresBackend) GetRecepie(recepieID string) (Recepie, error) {
	var (
		data string
		recepie Recepie
	)
	row := postgres.Database.QueryRow(
			"SELECT data FROM recepie WHERE recepie_id = $1",
 			recepieID)
    err := row.Scan(&data)
	if err != nil {
		return Recepie{}, err
	}
	err = json.Unmarshal([]byte(data), &recepie)
	if err != nil {
		return Recepie{}, err
	}

	return recepie, nil
}

func (postgres *PostgresBackend) Delete(recepieID string) error {
	_, err := postgres.Database.Exec(
		"DELETE FROM recepie WHERE recepie_id = $1",
		recepieID)
	if err != nil {
		return err
	}
	return nil
}

