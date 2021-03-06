package main

import (
	"errors"
)

// Database represents general DB interface, which is useful if we later
// decide to replace one DB back-end with another.
// It is also useful for mocking in unit tests.
type Database interface {
	CreateRecipe(r Recipe) error
	UpdateRecipe(r Recipe) error
	ListRecipes() ([]Recipe, error)
	GetRecipe(recipeID string) (Recipe, error)
	DeleteRecipe(recipeID string) error
}

type registerDbFn func(params map[string]string) (Database, error)

var (
	databaseBackends = map[string]registerDbFn{}
)

// RegisterDatabase allows DB back-end to register itself as a
// possible back-end.
func RegisterDatabase(dbBackName string, rdb registerDbFn) error {
	if _, ok := databaseBackends[dbBackName]; ok {
		return errors.New("database back-end already exists")
	}
	if rdb == nil {
		return errors.New("bad DB register function (nil)")
	}
	if dbBackName == "" {
		return errors.New("bad DB back-end name")
	}
	databaseBackends[dbBackName] = rdb
	return nil
}

// NewDatabaseConnection will initialized database back-end, selected with use of dbBackName parameter.
// params will be used for DB back-end initialization.
func NewDatabaseConnection(dbBackName string, params map[string]string) (Database, error) {
	f, ok := databaseBackends[dbBackName]
	if !ok {
		return nil, errors.New("database back-end doesn't exist")
	}
	return f(params)
}
