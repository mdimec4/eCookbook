package main
import (
	"net/http"
)

type Recepie struct {
	RecepieID string `json:"recepie_id"`
	Publisher string `json:"publisher"`
	SourceURL string `json:"source_url"`
	Title string `json:"title"`
	ImageURL string `json:"image_url"`
	Tags []string `json:"tags"`
	Ingredients []string `json:"ingredients"`
	Instructions []struct {
		Number int `json:"number"`
		ImageURL string `json:"image_url"`
		Instruction string `json:"instruction"`
	} `json:"instructions"`
	Tips []string `json:"tips"`
}

/*
func getRecepie(recepieID uint) (Recepie, error) {
}

func getRecepieList() ([]Recepie, error) {
}

func update(recepie Recepie) error {
}

func newRecepie(recepie Recepie) (uint, error) {
}

func GetRecepiesList(w http.ResponseWriter, r *http.Request) {
}

func GetRecepie(w http.ResponseWriter, r *http.Request) {
}

func PostNewRecepie(w http.ResponseWriter, r *http.Request) {
}

*/
