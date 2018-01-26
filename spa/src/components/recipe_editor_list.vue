<template>
  <div class="recipe_editor_list">
    <h1>Edit recipes</h1>
    <div v-for="recipe in recipes" v-bind:key="recipe.recipe_id" id="container">
      <router-link v-bind:to="{name: 'RecipeEditor', params: { id: recipe.recipe_id}}"  class="container-part"><p>{{recipe.title}}</p></router-link>
      <a v-if="recipe.source_url && recipe.source_url !=''" v-bind:href="recipe.source_url" class="container-part">source</a>
      <div v-else  class="container-part">-</div>
      <div v-on:click="removeRecipe(recipe)" class="container-part"><icon style="color: red;" scale=1 name="remove"></icon></div>
    </div>
    <router-link  v-bind:to="{name: 'RecipeEditorNew'}"><button>Add New</button></router-link>
  </div>
</template>

<script>
import { getRecipes, deleteRecipe } from './api'
export default {
  name: 'recipe_editor_list',
  data () {
    return {
      recipes: []
    }
  },
  mounted: function () {
    getRecipes().then((recipes) => {
      // recipes[0].source_url = 'http://google.com' /* XXX test */
      this.recipes = recipes
      // console.log(this.recipes)
    }, (err) => console.error('promise: ', err))
  },
  methods: {
    removeRecipe: function (recipe) {
      deleteRecipe(recipe.recipe_id).then((status) => {
        console.log(status)
        // use filter to also remove succesfuly removed recipe
        this.recipes = this.recipes.filter((recipeFilter) => {
          return recipeFilter.recipe_id !== recipe.recipe_id
        })
      }, (err) => console.error('promise: ', err))
    }
  }
}

</script>
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
};
#container {
  border: solid;
  border-width: 1px 1px 0px 0px;
  border-color:black;
  width: calc(100% - 1px);
  height: calc(10vh - 2px);
};
.container-part {
  float:left;
  border: solid;
  border-width: 0px 0px 1px 1px;
  border-color:black;
  width: calc(33.33% - 1px); /* subtract 1px border width */
  height: 100%;
};
</style>
