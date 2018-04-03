package main

const manualEntryBackendName string = "manual-entry"

type manualEntryBackend struct {
}

// handleNewRecipe will create unique ID for recipe,
func (meb manualEntryBackend) handleNewRecipe(recipe Recipe) (Recipe, error) {
	id, err := uniqueRecipeID(recipe.Title)
	if err != nil {
		return Recipe{}, err
	}
	recipe.RecipeID = id
	if recipe.Backend == "" {
		recipe.Backend = manualEntryBackendName
	}
	return recipe, nil
}
