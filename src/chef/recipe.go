package main

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"hash/crc32"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

/*
REST API guides
- When should we use PUT and when should we use POST?
  http://restcookbook.com/HTTP%20Methods/put-vs-post/
- HTTP Status Codes
  http://www.restapitutorial.com/httpstatuscodes.html
- Rest api in GO sample
  https://www.thepolyglotdeveloper.com/2016/07/create-a-simple-restful-api-with-golang/
*/

// Recipe represents the recipe in database.
type Recipe struct {
	RecipeID     string   `json:"recipe_id"`
	Publisher    string   `json:"publisher"`
	SourceURL    string   `json:"source_url"`
	Title        string   `json:"title"`
	ImageURL     string   `json:"image_url"`
	Tags         []string `json:"tags"`
	Ingredients  []string `json:"ingredients"`
	Instructions []struct {
		Number      int    `json:"number"`
		ImageURL    string `json:"image_url"`
		Instruction string `json:"instruction"`
	} `json:"instructions"`
	Tips []string `json:"tips"`
}

var (
	db Database
)

var (
	errRecipeExist    error = errors.New("recipe already exists")
	errRecipeExistNot error = errors.New("recipe doesn't exists")
)

// generate random number
func cryptoRandSecure(max int64) (int64, error) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return 0, err
	}
	return nBig.Int64(), nil
}

// uniqueRecipeID will generate unique ID using recipe title.
// And yes title is not the only source of uniqueness,
// so there can be many recipes with the same title
func uniqueRecipeID(title string) (string, error) {
	if title == "" {
		return "", errors.New("title is not set")
	}
	var crc uint32 = crc32.ChecksumIEEE([]byte(title))
	rand, err := cryptoRandSecure(int64(^(uint64(1) << 63)))
	if err != nil {
		return "", err
	}
	timestamp := time.Now().Unix()
	return fmt.Sprintf("%x-%x-%x", crc, rand, timestamp), nil
}

// newRecepy will create unique ID for recipe,
// and insert the recipe into the database.
func newRecipe(recipe Recipe) (string, error) {
	if recipe.RecipeID == "" {
		id, err := uniqueRecipeID(recipe.Title)
		if err != nil {
			return "", err
		}
		recipe.RecipeID = id
	}
	err := db.CreateRecipe(recipe)
	return recipe.RecipeID, err
}

func GetRecipesList(w http.ResponseWriter, r *http.Request) {
	l, err := db.ListRecipes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(l)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("str ", string(b))
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func GetRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	recipe, err := db.GetRecipe(params["id"])
	if err != nil {
		if err == errRecipeExistNot {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("str ", string(b))
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func PostNewRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe Recipe
	err := json.NewDecoder(r.Body).Decode(&recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, err := newRecipe(recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 201 Created
	// Location
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", r.URL.Path+"/"+id)
	w.WriteHeader(http.StatusCreated)
}

func PutUpdateRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe Recipe
	err := json.NewDecoder(r.Body).Decode(&recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	params := mux.Vars(r)
	if params["id"] != recipe.RecipeID {
		http.Error(w, "URL vs. JSON body 'RecipeID' mismatch", http.StatusConflict)
		return
	}

	err = db.UpdateRecipe(recipe)
	if err != nil {
		if err == errRecipeExistNot {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 201 Created
	// Location
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", r.URL.Path)
	w.WriteHeader(http.StatusCreated)
}

func DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := db.DeleteRecipe(params["id"])
	if err != nil {
		if err == errRecipeExistNot {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 204 NoContent
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	var (
		err error
	)
	// TODO use config.json file
	params := map[string]string{"user": "chef", "dbname": "cookbook",
		"password": "chef", "sslmode": "disable", "host": "localhost",
		"port": "5432"}
	db, err = NewDatabaseConnection("postgres", params)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Setting database connection: %s\n", err)
		os.Exit(1)
	}
	// TODO implement authentication
	router := mux.NewRouter()
	router.HandleFunc("/api/recipes", GetRecipesList).Methods("GET")
	router.HandleFunc("/api/recipes/{id}", GetRecipe).Methods("GET")
	router.HandleFunc("/api/recipes", PostNewRecipe).Methods("POST")
	router.HandleFunc("/api/recipes/{id}", PutUpdateRecipe).Methods("PUT")
	router.HandleFunc("/api/recipes/{id}", DeleteRecipe).Methods("DELETE")
	// TODO setup from config.json
	fmt.Fprintf(os.Stderr, "%v\n", http.ListenAndServe(":4006", router))
}
