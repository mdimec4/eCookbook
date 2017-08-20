package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

type PostgresBackend struct {
	dsn      string
	Database *sql.DB
}

// Register postgres DB back-end in initialization phase
func init() {
    RegisterDatabase("postgres", NewPostgresConnectionParams)
}

// parsePostgresConnParams() will transform string key value pars from 
// map[strinh]string in to plain string formed as 
// "key1=val1 key2=val2 ... keyN=valN" 
func parsePostgresConnParams(params map[string]string) string {
    str := ""
    for key, val := range params {
        // if not the first key val pair in str,
        // then we need to insert space for separation
        if str != "" {
            str += " "
        }
        str += fmt.Sprintf("%s=%s", key, val)
    }
    return str
}

func NewPostgresConnectionParams(params map[string]string) (Database, error) {
    return NewPostgresConnection(parsePostgresConnParams(params))
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
			"CREATE TABLE recepie (recepie_id text PRIMARY KEY, data jsonb)")
		if err != nil {
			return nil, err
		}
	} else {
		rows.Close()
	}
	return postgresDB, nil
}

/*
double insert should yeald this:
ERROR:  duplicate key value violates unique constraint "recepie_recepie_id_key"
DETAIL:  Key (recepie_id)=(333) already exists.
*/
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
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			fmt.Fprintf(os.Stderr, "CreateRecepie(): %s\n", err)
			return errRecepieExist
		}
		return err
	}
	return nil
}

func (postgres *PostgresBackend) UpdateRecepie(r Recepie) error {
	var rs sql.Result
	id := r.RecepieID
	recJson, err := json.Marshal(r)
	if err != nil {
		return err
	}
	rs, err = postgres.Database.Exec(
		"UPDATE recepie SET data = $2 WHERE recepie_id = $1",
		id, string(recJson))
	if err != nil {
		return err
	}
	if cnt, err := rs.RowsAffected(); err != nil {
		return err
	} else if cnt < 1 {
		return errRecepieExistNot
	}
	return nil
}

/*
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
*/

func (postgres *PostgresBackend) ListRecepies() ([]Recepie, error) {
	var (
		data     string
		recepie  Recepie
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
			if err == sql.ErrNoRows {
				return []Recepie{}, nil
			}
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
		data    string
		recepie Recepie
	)
	row := postgres.Database.QueryRow(
		"SELECT data FROM recepie WHERE recepie_id = $1",
		recepieID)
	err := row.Scan(&data)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errRecepieExistNot
		}
		return Recepie{}, err
	}
	err = json.Unmarshal([]byte(data), &recepie)
	if err != nil {
		return Recepie{}, err
	}

	return recepie, nil
}

func (postgres *PostgresBackend) DeleteRecepie(recepieID string) error {
	rs, err := postgres.Database.Exec(
		"DELETE FROM recepie WHERE recepie_id = $1",
		recepieID)
	if err != nil {
		return err
	}
	if cnt, err := rs.RowsAffected(); err != nil {
		return err
	} else if cnt < 1 {
		return errRecepieExistNot
	}
	return nil
}
