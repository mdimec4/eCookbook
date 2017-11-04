<template>
  <div class="recipe_editor_list">
    <h1>Edit recipes</h1>
    <div v-for="recipe in recipes" v-bind:key="recipe.recipe_id">
      <router-link v-bind:to="{name: 'RecipeEditor', params: { id: recipe.recipe_id}}"  style="display:inline; width:20vw;"><p>{{recipe.title}}</p></router-link>
      <a v-if="recipe.source_url && recipe.source_url !=''" v-bind:href="recipe.source_url" style="display:inline; width:20vw;">source</a>
    </div>
    <router-link  v-bind:to="{name: 'RecipeEditorNew'}"><button>Add New</button></router-link>
  </div>
</template>

<script>
import { getRecipes } from './api'
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
  }
}

</script>
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
};
</style>
