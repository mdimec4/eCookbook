package main

import (
	"errors"
)

// Database represents general DB interface, which is useful if we later
// decide to replace one DB back-end with another.
// It is also useful for mocking in unit tests.
type Database interface {
	CreateRecepie(r Recepie) error
	UpdateRecepie(r Recepie) error
	ListRecepies() ([]Recepie, error)
	GetRecepie(recepieID string) (Recepie, error)
	DeleteRecepie(recepieID string) error
}

type registerDbFn func(params map[string]string) (Database, error)

var (
	databaseBackends map[string]registerDbFn = map[string]registerDbFn{}
)

// RegisterDatabase allows DB back-end to register itself as a
// possible back-end.
func RegisterDatabase(dbBackName string, rdb registerDbFn) error {
	if _, ok := databaseBackends[dbBackName]; ok {
		return errors.New("database back-end already exists")
	}
	if rdb == nil {
		return errors.New("bad DB register functon (nil)")
	}
	databaseBackends[dbBackName] = rdb
	return nil
}

// NewDatabaseConnection will initalize database back-end, selected with use of dbBackName parameter.
// params will be used for DB back-end initialization.
func NewDatabaseConnection(dbBackName string, params map[string]string) (Database, error) {
	f, ok := databaseBackends[dbBackName]
	if !ok {
		return nil, errors.New("database backend doesn't exist")
	}
	return f(params)
}
