package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
)

const allRecipesBackendName string = "allrecipes.com"

var (
	errMiddlewareHostPortNotSet error = errors.New("allrecipes.com middle-ware host is not set")
)

type allRecipesRecipe struct {
	RecipeID    string   `json:"recipe_id"`
	Author      string   `json:"author"`
	SourceURL   string   `json:"source_url"`
	Name        string   `json:"name"`
	ImageURL    string   `json:"image_url"`
	Description string   `json:"description"`
	Ingredients []string `json:"ingredients"`
	Directions  []string `json:"directions"`
	Footnotes   []string `json:"footnotes"`
}

var (
	idSplitRegex *regexp.Regexp = regexp.MustCompile("^([0-9]+)--allrecipes.com$")
)

type allRecipesBackend struct {
	// relies on all allrecipes.com parsing HTTP midleware
	// check for more info: https://github.com/mdimec4/allrecipes
	middlewareHostPort string
}

func newAllRecipesBackend(host string) (*allRecipesBackend, error) {
	if host == "" {
		return nil, errMiddlewareHostPortNotSet
	}
	return &allRecipesBackend{host}, nil
}

// handleNewRecipe will query allrecipes middleware and will  populate
// Recipe missing entries
func (arb allRecipesBackend) handleNewRecipe(recipe Recipe) (Recipe, error) {
	var arRecipe allRecipesRecipe

	if arb.middlewareHostPort == "" {
		return Recipe{}, errMiddlewareHostPortNotSet
	}

	if recipe.Backend == "" {
		recipe.Backend = allRecipesBackendName
	}

	allRecipesID := ""
	// check if recipe.RecipeID maybe has
	// allrecipes.com ID. We can use it, but it needs to be split from suffix first
	if recipe.RecipeID != "" {
		parts := idSplitRegex.FindStringSubmatch(recipe.RecipeID)
		if len(parts) >= 2 {
			allRecipesID = parts[1]
		}

	}
	if allRecipesID == "" && recipe.SourceURL != "" {
		// get recipe ID from allrecipes.com url
		allRecipesURL, err := url.Parse(recipe.SourceURL)
		if err != nil {
			return Recipe{}, fmt.Errorf("allrecipes.com URL parsing problem: %s", err)
		}
		if allRecipesURL.Host != "www.allrecipes.com" {
			return Recipe{}, fmt.Errorf("'www.allrecipes.com' host was expected, we have: '%s'", allRecipesURL.Host)
		}
		arPath := allRecipesURL.Path
		if arPath[0] == '/' {
			arPath = arPath[1:]
		}
		if arPath[len(arPath)-1] == '/' {
			arPath = arPath[:len(arPath)-1]
		}
		parts := strings.Split(arPath, "/")
		// url path scheme is /recipe/{ID}/some name
		if len(parts) < 2 {
			return Recipe{}, fmt.Errorf("failed to get ID from allrecipes.com URL: %s", recipe.SourceURL)
		}
		allRecipesID = parts[1]
	}

	if allRecipesID == "" {
		return Recipe{}, errRecipeIDNotSet
	}

	// get recipe id from url
	u, err := url.Parse("http://" + arb.middlewareHostPort)
	if err != nil {
		return Recipe{}, fmt.Errorf("parse allrecipes.com middle-ware url %s", err)
	}
	u.Path = path.Join("api", "recipe", allRecipesID)

	// parse html
	resp, err := http.Get(u.String())
	if err != nil {
		return Recipe{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK /*200*/ {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading  allrecipes.com middle-ware non 200 response: %s\n", err)
		}
		return Recipe{}, fmt.Errorf("allrecipes.com middle-ware responded with: Status: %s, Body: s", resp.Status, string(b))
	}
	err = json.NewDecoder(resp.Body).Decode(&arRecipe)
	if err != nil {
		return Recipe{}, err
	}

	// translate to our recipe format
	if recipe.RecipeID == "" {
		recipe.RecipeID = arRecipe.RecipeID + "--" + allRecipesBackendName
	}
	recipe.Publisher = arRecipe.Author
	recipe.SourceURL = arRecipe.SourceURL
	recipe.Title = arRecipe.Name
	recipe.ImageURL = arRecipe.ImageURL
	for _, i := range arRecipe.Ingredients {
		recipe.Ingredients = append(recipe.Ingredients, i)
	}
	for _, i := range arRecipe.Directions {
		recipe.Instructions = append(recipe.Instructions,
			Instruction{"", i})
	}
	if arRecipe.Description != "" {
		recipe.Tips = append(recipe.Tips, arRecipe.Description)
	}
	for _, f := range arRecipe.Footnotes {
		recipe.Tips = append(recipe.Tips, f)
	}

	return recipe, nil
}
