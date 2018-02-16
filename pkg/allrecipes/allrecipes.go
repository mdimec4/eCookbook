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

func findTokenText(z *html.Tokenizer, dataAtom atom.Atom, key, value string) (string, error) {
	token := z.Token()
	if token.DataAtom == dataAtom &&
		checkAttr(token.Attr, key, value) {
		// next token should be text of the ingredient span
		tt := z.Next()
		switch tt {
		case html.TextToken:
			token = z.Token()
			return token.Data, nil
		case html.ErrorToken:
			return "", z.Err()
		default:
			return "", errors.New("text was expected")

		}
	}
	return "", nil
}

func getRecipe(url string) error {
	resp, err := http.Get(url)
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
			// did we hit one of the ingredients
			// <span class="recipe-ingred_txt added" ... itemprop="ingredients">
			if ingredient, err := findTokenText(z,
				atom.Span, "itemprop", "ingredients"); err != nil {
				return fmt.Errorf("allrecipes ingredients parser:")
			} else if ingredient != "" {
				fmt.Println("ingredient>>", ingredient)
			}

			// did we hit one of the instructions
			// <span class="recipe-ingred_txt added" ... itemprop="ingredients">
			if instruction, err := findTokenText(z,
				atom.Span, "class", "recipe-directions__list--item"); err != nil {
				return fmt.Errorf("allrecipes instructions parser:")
			} else if instruction != "" {
				fmt.Println("instruction>>", instruction)
			}

		}

	}
	return nil
}

func main() {
	//url := "http://allrecipes.com/recipe/231495/texas-boiled-beer-shrimp/"
	url := "http://allrecipes.com/recipe/11772/spaghetti-pie-i/?clickId=right%20rail0&internalSource=rr_feed_recipe_sb&referringId=231495%20referringContentType%3Drecipe"
	err := getRecipe(url)
	if err != nil {
		fmt.Println(err) // TODO stderr
		return
	}

}
