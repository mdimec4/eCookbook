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
			token := z.Token()
			if token.DataAtom == atom.Span &&
				checkAttr(token.Attr, "itemprop", "ingredients") {
				// did we hit one of the ingredients
				// <span class="recipe-ingred_txt added" ... itemprop="ingredients">
				tt := z.Next()
				// next token should be text of the ingredient span
				switch tt {
				case html.TextToken:
					token = z.Token()
					fmt.Println("ingredient>", token.Data)
				case html.ErrorToken:
					return z.Err()
				default:
					return errors.New("allrecipes parser: ingredient text was expected here")
				}
			} else if token.DataAtom == atom.Span &&
				checkAttr(token.Attr, "class", "recipe-directions__list--item") &&
				!checkAttr(token.Attr, "ng-bind", "model.itemNote") {
				// did we hit one of the instructions
				// <span class="recipe-directions__list--item" ...>
				tt := z.Next()
				// next token should be text of the instruction span
				switch tt {
				case html.TextToken:
					token = z.Token()
					fmt.Println("instruction>", token.Data)
				case html.ErrorToken:
					return z.Err()
				default:
					return errors.New("allrecipes parser: instruction text was expected here")
				}
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
