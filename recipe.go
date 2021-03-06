package main

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"strings"
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
	RecipeID     string        `json:"recipe_id"` // in format "backend:id"
	Backend      string        `json:"backend"`
	Publisher    string        `json:"publisher"`
	SourceURL    string        `json:"source_url"`
	Title        string        `json:"title"`
	ImageURL     string        `json:"image_url"`
	Tags         []string      `json:"tags"`
	Ingredients  []string      `json:"ingredients"`
	Instructions []Instruction `json:"instructions"`
	Tips         []string      `json:"tips"`
}

type Instruction struct {
	ImageURL    string `json:"image_url"`
	Instruction string `json:"instruction"`
}

var (
	db       Database
	backends map[string]backend
)

var (
	errRecipeExist              = errors.New("recipe already exists")
	errRecipeExistNot           = errors.New("recipe doesn't exist")
	errRecipeTitleNotSet        = errors.New("recipe title is not set")
	errRecipeIDNotSet           = errors.New("recipe ID is not set")
	errRecipeSourceURLNotSet    = errors.New("recipe source URL is not set")
	errRecipeIngredientsNotSet  = errors.New("recipe ingredients are not set")
	errRecipeInstructionsNotSet = errors.New("recipe instructions are not set")
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
		return "", errRecipeTitleNotSet
	}
	var crc = crc32.ChecksumIEEE([]byte(title))
	// get random number. max int64 is the top limit
	rand, err := cryptoRandSecure(int64(^(uint64(1) << 63)))
	if err != nil {
		return "", err
	}
	timestamp := time.Now().Unix()
	return fmt.Sprintf("%x-%x-%x", crc, rand, timestamp), nil
}

type backend interface {
	handleNewRecipe(recipe Recipe) (Recipe, error)
}

func findBackend(recipe Recipe) backend {
	// default to manualEntryBackendName for compatibility
	// reasons, since old manual ID's didn't have "backend" entry
	name := manualEntryBackendName
	if recipe.Backend != "" {
		name = recipe.Backend
	}

	if b, ok := backends[name]; ok {
		return b
	}
	return nil
}

func getRecipesList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

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
	// fmt.Println("str ", string(b))
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func getRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Access-Control-Allow-Origin", "*")

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
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func postNewRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe Recipe
	w.Header().Set("Access-Control-Allow-Origin", "*")

	err := json.NewDecoder(r.Body).Decode(&recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if recipe.Backend == "" {
		recipe.Backend = manualEntryBackendName
	}

	b := findBackend(recipe)
	if b == nil {
		http.Error(w, fmt.Sprintf("problem finding back-end named '%s': %s", recipe.Backend, err), http.StatusInternalServerError)
		return
	}
	recipe, err = b.handleNewRecipe(recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if recipe.RecipeID == "" {
		http.Error(w, errRecipeIDNotSet.Error(), http.StatusInternalServerError)
		return
	}
	if recipe.Title == "" {
		http.Error(w, errRecipeTitleNotSet.Error(), http.StatusInternalServerError)
		return
	}
	if len(recipe.Ingredients) == 0 || recipe.Ingredients[0] == "" {
		http.Error(w, errRecipeIngredientsNotSet.Error(), http.StatusInternalServerError)
		return
	}
	if len(recipe.Instructions) == 0 && recipe.Instructions[0].Instruction == "" {
		http.Error(w, errRecipeInstructionsNotSet.Error(), http.StatusInternalServerError)
		return
	}

	err = db.CreateRecipe(recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 201 Created
	// Location
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", r.URL.Path+"/"+recipe.RecipeID)
	w.WriteHeader(http.StatusCreated)
}

func putUpdateRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe Recipe
	w.Header().Set("Access-Control-Allow-Origin", "*")

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

	if recipe.RecipeID == "" {
		http.Error(w, errRecipeIDNotSet.Error(), http.StatusInternalServerError)
		return
	}
	if recipe.Title == "" {
		http.Error(w, errRecipeTitleNotSet.Error(), http.StatusInternalServerError)
		return
	}
	if len(recipe.Ingredients) == 0 || recipe.Ingredients[0] == "" {
		http.Error(w, errRecipeIngredientsNotSet.Error(), http.StatusInternalServerError)
		return
	}
	if len(recipe.Instructions) == 0 && recipe.Instructions[0].Instruction == "" {
		http.Error(w, errRecipeInstructionsNotSet.Error(), http.StatusInternalServerError)
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

func deleteRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Access-Control-Allow-Origin", "*")

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

func sendOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, PUT, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// 204 NoContent
	w.WriteHeader(http.StatusNoContent)
}

const staticPath = "./dist"

func staticHandleFunc(w http.ResponseWriter, r *http.Request) {
	if _, err := os.Stat(staticPath + r.URL.Path); os.IsNotExist(err) {
		// allow paths that will be handled by vue2.js router
		if !(strings.HasPrefix(r.URL.Path, "/device") ||
			strings.HasPrefix(r.URL.Path, "/recipe_editor")) {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		// all non file path url's are redirected to vue2.js app.
		// It's vue2.js router responsibility to gandle them.
		dat, _ := ioutil.ReadFile(staticPath + "/index.html")
		w.Write(dat)
		return
	}
	http.FileServer(http.Dir(staticPath)).ServeHTTP(w, r)
}

func getEnvConf(envName, defVal string) string {
	if v := os.Getenv(envName); v != "" {
		return v
	}
	return defVal
}

func main() {
	var (
		err error
	)
	params := map[string]string{"user": getEnvConf("CHEF_DB_USER", "chef"),
		"dbname":   getEnvConf("CHEF_DB_NAME", "cookbook"),
		"password": getEnvConf("CHEF_DB_PASS", "chef"),
		"sslmode":  getEnvConf("CHEF_DB_SSLMODE", "disable"),
		"host":     getEnvConf("CHEF_DB_HOST", "localhost"),
		"port":     getEnvConf("CHEF_DB_PORT", "5432")}
	db, err = NewDatabaseConnection("postgres", params)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Setting database connection: %s\n", err)
		os.Exit(1)
	}

	// init backends
	backends = map[string]backend{
		manualEntryBackendName: &manualEntryBackend{},
		allRecipesBackendName:  &allRecipesBackend{getEnvConf("ALLRECIPES_MIDDLEWARE_HOST_PORT", "localhost:4007")}}

	m := http.NewServeMux()
	m.HandleFunc("/", staticHandleFunc)
	// TODO implement authentication
	router := mux.NewRouter()
	router.HandleFunc("/api/recipes", getRecipesList).Methods("GET")
	router.HandleFunc("/api/recipes/{id}", getRecipe).Methods("GET")
	router.HandleFunc("/api/recipes", postNewRecipe).Methods("POST")
	router.HandleFunc("/api/recipes/{id}", putUpdateRecipe).Methods("PUT")
	router.HandleFunc("/api/recipes/{id}", deleteRecipe).Methods("DELETE")
	router.HandleFunc("/api/recipes", sendOptions).Methods("OPTIONS")
	router.HandleFunc("/api/recipes/{id}", sendOptions).Methods("OPTIONS")

	m.Handle("/api/", router)
	fmt.Fprintf(os.Stderr, "%v\n", http.ListenAndServe(getEnvConf("CHEF_LISTEN_ADDR", ":4006"), m))
}
