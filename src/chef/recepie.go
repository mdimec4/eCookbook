package main

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"hash/crc32"
	"net/http"
	"os"
	//"strings"
	//"unicode"

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

type Recepie struct {
	RecepieID    string   `json:"recepie_id"`
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
	postgres *PostgresBackend
)

var (
	errRecepieExist    error = errors.New("recipe already exists")
	errRecepieExistNot error = errors.New("recipe doesn't exists")
)

func getRecepie(recepieID string) (Recepie, error) {
	return postgres.GetRecepie(recepieID)
}

func getRecepieList() ([]Recepie, error) {
	return postgres.ListRecepies()
}

func update(recepie Recepie) error {
	return postgres.Update(recepie)
}

/*
func whiteSpaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}
*/

func newRecepie(recepie Recepie) (string, error) {
	if recepie.Title == "" {
		return "", errors.New("Title is not set")
	}
	if recepie.RecepieID == "" {
		var crc1 uint32 = crc32.ChecksumIEEE([]byte(recepie.Title))
		var bArr []byte = make([]byte, 4)
		rand.Read(bArr)
		var num32 uint32 = 0
		num32 = uint32(bArr[3])<<24 | uint32(bArr[2])<<16 | uint32(bArr[1])<<8 | uint32(bArr[0])
		recepie.RecepieID = fmt.Sprintf("%x-%x", crc1, num32)
	}
	err := postgres.CreateRecepie(recepie)
	return recepie.RecepieID, err
}

func deleteRecepie(id string) error {
	return postgres.Delete(id)
}

func GetRecepiesList(w http.ResponseWriter, r *http.Request) {
	l, err := getRecepieList()
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

func GetRecepie(w http.ResponseWriter, r *http.Request) {
}

func PostNewRecepie(w http.ResponseWriter, r *http.Request) {
	var recepie Recepie
	err := json.NewDecoder(r.Body).Decode(&recepie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, err := newRecepie(recepie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//201 Created
	//Location
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", r.URL.Path+"/"+id)
	w.WriteHeader(http.StatusCreated)
}

func PutUpdateRecepie(w http.ResponseWriter, r *http.Request) {
	var recepie Recepie
	err := json.NewDecoder(r.Body).Decode(&recepie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	params := mux.Vars(r)
	if params["id"] != recepie.RecepieID {
		http.Error(w, "URL vs. JSON body 'RecepieID' mismatch", http.StatusConflict)
		return
	}

	err = update(recepie)
	if err != nil {
		if err == errRecepieExistNot {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//201 Created
	//Location
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", r.URL.Path)
	w.WriteHeader(http.StatusCreated)
}

func DeleteRecepie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := deleteRecepie(params["id"])
	if err != nil {
		if err == errRecepieExistNot {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//204 NoContent
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	var (
		err error
	)
	postgres, err = NewPostgresConnection(fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable host=%s port=%s", "chef", "cookbook", "chef", "localhost", "5432")) //TODO use config.json file
	if err != nil {
		fmt.Fprintf(os.Stderr, "Setting database connection: %s\n", err)
		os.Exit(1)
	}
	router := mux.NewRouter()
	router.HandleFunc("/api/recepies", GetRecepiesList).Methods("GET")
	router.HandleFunc("/api/recepies/{id}", GetRecepie).Methods("GET")
	router.HandleFunc("/api/recepies", PostNewRecepie).Methods("POST")
	router.HandleFunc("/api/recepies/{id}", PutUpdateRecepie).Methods("PUT")
	router.HandleFunc("/api/recepies/{id}", DeleteRecepie).Methods("DELETE")
	fmt.Fprintf(os.Stderr, "%v\n", http.ListenAndServe(":4006", router)) //TODO setup from config.json
}
