<template>
  <div class="recipe_editor_list">
    <h1>Edit recipes</h1>
    <router-link v-for="recipe in recipes" v-bind:key="recipe.recipe_id" v-bind:to="{name: 'RecipeEditor', params: { id: recipe.recipe_id}}"><p>{{recipe.title}}</p></router-link>
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
    getRecipes().then((recipes) => { this.recipes = recipes }, (err) => console.error('promise: ', err))
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
