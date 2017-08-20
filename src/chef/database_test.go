package main

import (
    "testing"
    "strings"
)

type mockDB struct {
    params map[string]string
    recipes map[string]Recipe
}

func connectMockDb(params map[string]string) (Database, error) {
    return &mockDB{params: params,
        recipes: map[string]Recipe{}}, nil
}

func (mdb *mockDB) CreateRecipe(r Recipe) error {
    if _, ok := mdb.recipes[r.RecipeID]; ok {
        return errRecipeExist
    }
    mdb.recipes[r.RecipeID] = r
    return nil
}

func (mdb *mockDB) UpdateRecipe(r Recipe) error {
    if _, ok := mdb.recipes[r.RecipeID]; !ok {
        return errRecipeExistNot
    }
    mdb.recipes[r.RecipeID] = r
    return nil
}

func (mdb *mockDB) ListRecipes() ([]Recipe, error) {
    l := make([]Recipe, 0, len(mdb.recipes))
    for _, r := range mdb.recipes {
        l = append(l, r)
    }
    return l, nil
}

func (mdb *mockDB) GetRecipe(id string) (Recipe, error) {
    r, ok := mdb.recipes[id]
    if !ok {
        return Recipe{}, errRecipeExistNot
    }
    return r, nil
}

func (mdb *mockDB) DeleteRecipe(id string) error {
    if _, ok := mdb.recipes[id]; !ok {
        return errRecipeExistNot
    }
    delete(mdb.recipes, id)
    return nil
}

func TestRegisterDatabase(t *testing.T) {
    // reset list for new test
    databaseBackends = map[string]registerDbFn{}
    if err := RegisterDatabase("mock", connectMockDb); err != nil {
        t.Error(err)
    }
    f, ok := databaseBackends["mock"]
    if !ok {
        t.Error("expected mock backend")
    }
    if f == nil {
        t.Error("expected mock connect function")
    }
}

func TestRegisterDatabase_BadName(t *testing.T) {
    // reset list for new test
    databaseBackends = map[string]registerDbFn{}
    err := RegisterDatabase("", connectMockDb)
    if err == nil {
        t.Fatal("err was expected")
    }
    if !strings.Contains(err.Error(), "bad DB back-end name") {
        t.Fatal("expected 'bad DB back-end name' error, got: ", err)
    }
}

func TestRegisterDatabase_BadRegisterFn(t *testing.T) {
    // reset list for new test
    databaseBackends = map[string]registerDbFn{}
    err := RegisterDatabase("mock", nil)
    if err == nil {
        t.Fatal("err was expected")
    }
    if !strings.Contains(err.Error(), "bad DB register function (nil)") {
        t.Fatal("expected 'bad DB register function' error, got: ", err)
    }
}

func TestRegisterDatabase_DubleInsert(t *testing.T) {
    // reset list for new test
    databaseBackends = map[string]registerDbFn{}
    if err := RegisterDatabase("mock", connectMockDb); err != nil {
        t.Error(err)
    }
    err := RegisterDatabase("mock", connectMockDb)
    if err == nil {
        t.Fatal("error was expected")
        return
    }
    if !strings.Contains(err.Error(), "database back-end already exists") {
        t.Fatal("expected 'database back-end already exists' error, got: ", err)
    }
}

func TestNewDatabaseConnection_NoBackend(t *testing.T) {
    // reset list for new test
    databaseBackends = map[string]registerDbFn{}
    db, err := NewDatabaseConnection("doesn't-exist", map[string]string{})
    if err == nil {
        t.Fatal("error was expected")
        return
    }
    if !strings.Contains(err.Error(), "database back-end doesn't exist") {
        t.Fatal("expected 'database back-end doesn't exist' error, got: ", err)
    }
    if db != nil {
        t.Error("should had returned nil db")
    }
}

func TestNewDatabaseConnection(t *testing.T) {
    // reset list for new test
    databaseBackends = map[string]registerDbFn{}

    if err := RegisterDatabase("mock", connectMockDb); err != nil {
        t.Error(err)
    }

    db, err := NewDatabaseConnection("mock", map[string]string{"db": "test_database"})
    if err != nil {
        t.Error(err)
    }
    if db.(*mockDB).params["db"] != "test_database" {
        t.Error("back-end was not set as expected")
    }

    if err := db.CreateRecipe(Recipe{Title: "raspberry pi", RecipeID: "123"}); err != nil {
        t.Error(err)
    }
    r, err := db.GetRecipe("123")
    if err != nil {
        t.Error(err)
    }
    if r.Title != "raspberry pi" {
        t.Error("raspberry pi recepie was expected")
    }
    if err = db.DeleteRecipe("123"); err != nil {
        t.Error(err)
    }
    r, err = db.GetRecipe("123")
    if err == nil {
        t.Error("error was expected")
    }
    if err != errRecipeExistNot {
        t.Error("expected '%s', but got %s", errRecipeExistNot, err)
    }
}
