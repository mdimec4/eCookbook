package main
import (
    "fmt"
	"net/http"
    "os"
    "hash/crc32"
    "errors"
    "crypto/rand"
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

var (
    postgres *PostgresBackend
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

func newRecepie(recepie Recepie) (string, error) {
    if recepie.Title == "" {
        return "", errors.New("Title is not set")
    }
    if recepie.RecepieID == "" {
        var crc1 uint32 = ChecksumIEEE([]byte(recepie.Title))
        var bArr []byte = make([]byte, 10)
        rand.Read(bArr)
        var crc2 uint32 = ChecksumIEEEstring(bArr)
        recepie.RecepieID = fmt.Sprintf("%s-%s", crc1, crc2)
    }
    err := postgres.CreateRecepie(recepie)
    return recepie.RecepieID, err
}

func GetRecepiesList(w http.ResponseWriter, r *http.Request) {
}

func GetRecepie(w http.ResponseWriter, r *http.Request) {
}

func PostNewRecepie(w http.ResponseWriter, r *http.Request) {
}


func main() {
    var (
        err error
    )
    postgres, err = NewPostgresConnection(fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable host=%s port=%s",
                                  "chef", "cookbook", "chef", "localhost", "5432"))
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s", err)
        os.Exit(1)
    }
}
