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
	RegisterDatabase("postgres", NewPostgresConnectionRegistration)
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

func NewPostgresConnectionRegistration(params map[string]string) (Database, error) {
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

	rows, err := postgresDB.Database.Query("SELECT data FROM recipe LIMIT 1")
	if err != nil {
		_, err = postgresDB.Database.Exec(
			"CREATE TABLE recipe (recipe_id text PRIMARY KEY, data jsonb)")
		if err != nil {
			return nil, err
		}
	} else {
		rows.Close()
	}
	return postgresDB, nil
}

/*
double insert should yield this:
ERROR:  duplicate key value violates unique constraint "recipe_recipe_id_key"
DETAIL:  Key (recipe_id)=(333) already exists.
*/
func (postgres *PostgresBackend) CreateRecipe(r Recipe) error {
	id := r.RecipeID
	recJson, err := json.Marshal(r)
	if err != nil {
		return err
	}
	_, err = postgres.Database.Exec(
		"INSERT INTO recipe (recipe_id, data) VALUES ($1, $2)",
		id, string(recJson))
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			fmt.Fprintf(os.Stderr, "CreateRecipe(): %s\n", err)
			return errRecipeExist
		}
		return err
	}
	return nil
}

func (postgres *PostgresBackend) UpdateRecipe(r Recipe) error {
	var rs sql.Result
	id := r.RecipeID
	recJson, err := json.Marshal(r)
	if err != nil {
		return err
	}
	rs, err = postgres.Database.Exec(
		"UPDATE recipe SET data = $2 WHERE recipe_id = $1",
		id, string(recJson))
	if err != nil {
		return err
	}
	if cnt, err := rs.RowsAffected(); err != nil {
		return err
	} else if cnt < 1 {
		return errRecipeExistNot
	}
	return nil
}

/*
func (postgres *PostgresBackend) UpdateOrCreateRecipe(r Recipe) error {
	id := r.RecipeID
	recJson, err := json.Marshal(r)
	if err != nil {
		return err
	}
	_, err = postgres.Database.Exec(`
		INSERT INTO recipe (recipe_id, data) VALUES ($1, $2)
		ON CONFLICT (recipe_id) DO UPDATE SET data = excluded.data`,
		id, string(recJson))
	if err != nil {
		return err
	}
	return nil
}
*/

func (postgres *PostgresBackend) ListRecipes() ([]Recipe, error) {
	var (
		data    string
		recipe  Recipe
		recipes []Recipe
	)
	rows, err := postgres.Database.Query("SELECT data FROM recipe")
	if err != nil {
		if err == sql.ErrNoRows {
			return []Recipe{}, nil
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&data)
		if err != nil {
			if err == sql.ErrNoRows {
				return []Recipe{}, nil
			}
			return nil, err
		}
		err = json.Unmarshal([]byte(data), &recipe)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}
	return recipes, nil
}

func (postgres *PostgresBackend) GetRecipe(recipeID string) (Recipe, error) {
	var (
		data   string
		recipe Recipe
	)
	row := postgres.Database.QueryRow(
		"SELECT data FROM recipe WHERE recipe_id = $1",
		recipeID)
	err := row.Scan(&data)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errRecipeExistNot
		}
		return Recipe{}, err
	}
	err = json.Unmarshal([]byte(data), &recipe)
	if err != nil {
		return Recipe{}, err
	}

	return recipe, nil
}

func (postgres *PostgresBackend) DeleteRecipe(recipeID string) error {
	rs, err := postgres.Database.Exec(
		"DELETE FROM recipe WHERE recipe_id = $1",
		recipeID)
	if err != nil {
		return err
	}
	if cnt, err := rs.RowsAffected(); err != nil {
		return err
	} else if cnt < 1 {
		return errRecipeExistNot
	}
	return nil
}
