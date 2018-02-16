package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func checkAttr(attr []html.Attribute, key, val string) bool {
	for _, a := range attr {
		if a.Key == key && a.Val == val {
			return true
		}
	}
	return false
}

func getRecipe() error {
	resp, err := http.Get("http://allrecipes.com/recipe/231495/texas-boiled-beer-shrimp/")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	z := html.NewTokenizer(resp.Body)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			if z.Err() == io.EOF {
				return nil
			}
			return z.Err()
		case html.StartTagToken:
			token := z.Token()
			// did we hit one of ingredients
			// <span class="recipe-ingred_txt added" ... itemprop="ingredients">
			if token.DataAtom == atom.Span &&
				checkAttr(token.Attr, "itemprop", "ingredients") {
				fmt.Println("-->>", token.DataAtom)
				// next token should be text of the ingredient span
				tt = z.Next()
				switch tt {
				case html.TextToken:
					token = z.Token()
					fmt.Println("next>>", string(z.Raw()))
					fmt.Println("data>>", token.Data)
					fmt.Println("atom>>", token.DataAtom)
				default:
					errors.New("allrecipe html parser: ingredient text was expected")

				}
			}
		}

	}
	return nil
}

func main() {
	// var ingredients []string
	err := getRecipe()
	if err != nil {
		fmt.Println(err) // TODO stderr
		return
	}

}
