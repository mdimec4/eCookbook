package main

import (
	"fmt"
	"golang.org/x/net/html" //https://stackoverflow.com/questions/30109061/golang-parse-html-extract-all-content-with-body-body-tags
	//"io/ioutil"
	"golang.org/x/net/html/atom"
	"net/http"
)

func getRecipe() (*html.Node, error) {
	resp, err := http.Get("http://allrecipes.com/recipe/231495/texas-boiled-beer-shrimp/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	return html.Parse(resp.Body)
}

func loopChildren(parentNode *html.Node, f func(c *html.Node) bool) {
	for c := parentNode.FirstChild; c != nil; c = c.NextSibling {
		if !f(c) {
			break
		}
	}
}

func main() {
	// var ingredients []string
	documentNode, err := getRecipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	htmlNode := documentNode.FirstChild

	//find body
	var bodyNode *html.Node
	loopChildren(htmlNode.NextSibling, func(n *html.Node) bool {
		if n.DataAtom == atom.Body {
			bodyNode = n
			return false
		}
		return true
	})

	loopChildren(bodyNode, func(n *html.Node) bool {
		fmt.Println(n.DataAtom, n)
		return true
	})
}
