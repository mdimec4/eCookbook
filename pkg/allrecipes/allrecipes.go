package main

import (
	"fmt"
	"golang.org/x/net/html" //https://stackoverflow.com/questions/30109061/golang-parse-html-extract-all-content-with-body-body-tags
	"io/ioutil"
)

func getRecipe() (*html.Node, error) {
	resp, err := http.Get("http://allrecipes.com/recipe/231495/texas-boiled-beer-shrimp/")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return html.Parse(resp.Body)
}

func main() {
	html, err := getRecipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(html)
}
