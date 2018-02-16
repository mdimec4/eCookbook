package main

import (
	"fmt"
	"golang.org/x/net/html" //https://stackoverflow.com/questions/30109061/golang-parse-html-extract-all-content-with-body-body-tags
	//"io/ioutil"
	//"golang.org/x/net/html/atom"
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

func loopSiblings(n *html.Node, f func(c *html.Node) bool) bool {
	for c := n; c != nil; c = c.NextSibling {
		fmt.Println("ls")
		if !f(c) {
			return false
		}
	}
	return true
}

func loopChildren(n *html.Node, f func(*html.Node) bool) bool {
	for c := n.FirstChild; c != nil; c = c.FirstChild {
		fmt.Println("lc")
		if !f(c) {
			return false
		}
	}
	return true
}

func crawlNodes(nIn *html.Node, f func(*html.Node) bool) {
	var crawlNodesHelper func(cont bool, n *html.Node) bool

	crawlNodesHelper = func(cont bool, n *html.Node) bool {
		fmt.Println(0)
		loopSiblings(n, func(ns *html.Node) bool {
			fmt.Println(1)
			return loopChildren(ns, func(cn *html.Node) bool {
				fmt.Println(2)
				return crawlNodesHelper(f(cn), cn)
			})
		})
		return true
	}

	crawlNodesHelper(true, nIn)
}

func main() {
	// var ingredients []string
	documentNode, err := getRecipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		htmlNode := documentNode.FirstChild

		//find body
		var bodyNode *html.Node
		loopSiblings(htmlNode.NextSibling.FirstChild, func(n *html.Node) bool {
			if n.DataAtom == atom.Body {
				bodyNode = n
				return false
			}
			return true
		})

		loopSiblings(bodyNode.FirstChild, func(n *html.Node) bool {
			fmt.Println(n.DataAtom, n)
			return true
		})
	*/
	crawlNodes(documentNode, func(n *html.Node) bool {
		if n == nil {
			fmt.Println(n, " it is nil")
		}
		fmt.Println(n)
		return true
	})
}
