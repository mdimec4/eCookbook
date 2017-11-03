<template>
  <div class="recipe_editor_list">
    <h1>Edit recipes</h1>
    <router-link v-for="recipe in recipes" v-bind:key="recipe.recipe_id" v-bind:to="{name: 'RecipeEditor', params: { id: recipe.recipe_id}}"><p>{{recipe.title}}</p></router-link>
     <router-link  v-bind:to="{name: 'RecipeEditorNew'}"><button>Add New</button></router-link>
  </div>
</template>

<script>
import {somethingElse} from './api2'
function getRecipes () {
  var jsonStr = `[
    { 
        "recipe_id": "1-37859", 
        "publisher": "Real Simple", 
        "source_url": "http://food2fork.com/view/37859", 
        "title": "Chicken and Gruyre Turnovers", 
        "image_url": "http://static.food2fork.com/chickenturnover2_300e6667e66.jpg", 
        "tags": [
            "gluten","fish", "vegan"
        ],
        "tips": [ "tip1", "tip2", "tip3"] 
    },
    {
        "recipe_id": "2-4899", 
        "publisher": "Not Simple", 
        "source_url": "http://food2fork.com/view/37859", 
        "title": "Armenian Honey Cake", 
        "image_url": "http://static.food2fork.com/chickenturnover2_300e6667e66.jpg", 
        "tags": [
            "gluten","fish", "vegan"
        ],
        "tips": [ "tip1", "tip2", "tip3"] 
    },
    { 
        "recipe_id": "3-87859", 
        "publisher": "Real Simple", 
        "source_url": "http://food2fork.com/view/37859", 
        "title": "Chicken Rosted", 
        "image_url": "http://static.food2fork.com/chickenturnover2_300e6667e66.jpg", 
        "tags": [
            "gluten","fish", "vegan"
        ],
        "tips": [ "tip1", "tip2", "tip3"] 
    },
    {
        "recipe_id": "4-99899", 
        "publisher": "Not Simple", 
        "source_url": "http://food2fork.com/view/37859", 
        "title": "Mashrooms Cake", 
        "image_url": "http://static.food2fork.com/chickenturnover2_300e6667e66.jpg", 
        "tags": [
            "gluten","fish", "vegan"
        ],
        "tips": [ "tip1", "tip2", "tip3"] 
    },
    {
        "recipe_id": "5-777859", 
        "publisher": "Real Simple", 
        "source_url": "http://food2fork.com/view/37859", 
        "title": "Pork Gruyre Turnovers", 
        "image_url": "http://static.food2fork.com/chickenturnover2_300e6667e66.jpg", 
        "tags": [
            "gluten","fish", "vegan"
        ],
        "tips": [ "tip1", "tip2", "tip3"] 
    },
    {
        "recipe_id": "6-4d3252453899", 
        "publisher": "Not Simple", 
        "source_url": "http://food2fork.com/view/37859", 
        "title": "Specila Cake", 
        "image_url": "http://static.food2fork.com/chickenturnover2_300e6667e66.jpg", 
        "tags": [
            "gluten","fish", "vegan"
        ],
        "tips": [ "tip1", "tip2", "tip3"] 
    },
    { 
        "recipe_id": "7-ergtggsfdgfdsdf", 
        "publisher": "Real Simple", 
        "source_url": "http://food2fork.com/view/37859", 
        "title": "Cookies", 
        "image_url": "http://static.food2fork.com/chickenturnover2_300e6667e66.jpg", 
        "tags": [
            "gluten","fish", "vegan"
        ],
        "tips": [ "tip1", "tip2", "tip3"] 
    },
    {
        "recipe_id": "8-132323434", 
        "publisher": "Not Simple", 
        "source_url": "http://food2fork.com/view/37859", 
        "title": "Chocolatte Cake", 
        "image_url": "http://static.food2fork.com/chickenturnover2_300e6667e66.jpg", 
        "tags": [
            "gluten","fish", "vegan"
        ],
        "tips": [ "tip1", "tip2", "tip3"] 
    },
    { 
        "recipe_id": "9-37859", 
        "publisher": "Real Simple", 
        "source_url": "http://food2fork.com/view/37859", 
        "title": "Solata", 
        "image_url": "http://static.food2fork.com/chickenturnover2_300e6667e66.jpg", 
        "tags": [
            "gluten","fish", "vegan"
        ],
        "tips": [ "tip1", "tip2", "tip3"] 
    },
    {
        "recipe_id": "10-4899", 
        "publisher": "Not Simple", 
        "source_url": "http://food2fork.com/view/37859", 
        "title": "Tortae", 
        "image_url": "http://static.food2fork.com/chickenturnover2_300e6667e66.jpg", 
        "tags": [
            "gluten","fish", "vegan"
        ],
        "tips": [ "tip1", "tip2", "tip3"] 
    },
    { 
        "recipe_id": "11-87859", 
        "publisher": "Real Simple", 
        "source_url": "http://food2fork.com/view/37859", 
        "title": "Piscanec", 
        "image_url": "http://static.food2fork.com/chickenturnover2_300e6667e66.jpg", 
        "tags": [
            "gluten","fish", "vegan"
        ],
        "tips": [ "tip1", "tip2", "tip3"] 
    },
    {
        "recipe_id": "12-99899", 
        "publisher": "Not Simple", 
        "source_url": "http://food2fork.com/view/37859", 
        "title": "Gobice", 
        "image_url": "http://static.food2fork.com/chickenturnover2_300e6667e66.jpg", 
        "tags": [
            "gluten","fish", "vegan"
        ],
        "tips": [ "tip1", "tip2", "tip3"] 
    },
    {
        "recipe_id": "13-777859", 
        "publisher": "Real Simple", 
        "source_url": "http://food2fork.com/view/37859", 
        "title": "Svininja", 
        "image_url": "http://static.food2fork.com/chickenturnover2_300e6667e66.jpg", 
        "tags": [
            "gluten","fish", "vegan"
        ],
        "tips": [ "tip1", "tip2", "tip3"] 
    },
    {
        "recipe_id": "14-4d3252453899", 
        "publisher": "Not Simple", 
        "source_url": "http://food2fork.com/view/37859", 
        "title": "Posebna torta", 
        "image_url": "http://static.food2fork.com/chickenturnover2_300e6667e66.jpg", 
        "tags": [
            "gluten","fish", "vegan"
        ],
        "tips": [ "tip1", "tip2", "tip3"] 
    },
    { 
        "recipe_id": "15-ergtggsfdgfdsdf", 
        "publisher": "Real Simple", 
        "source_url": "http://food2fork.com/view/37859", 
        "title": "Keksi", 
        "image_url": "http://static.food2fork.com/chickenturnover2_300e6667e66.jpg", 
        "tags": [
            "gluten","fish", "vegan"
        ],
        "tips": [ "tip1", "tip2", "tip3"] 
    },
    {
        "recipe_id": "16-132323434", 
        "publisher": "Not Simple", 
        "source_url": "http://food2fork.com/view/37859", 
        "title": "Cokoladna torta", 
        "image_url": "http://static.food2fork.com/chickenturnover2_300e6667e66.jpg", 
        "tags": [
            "gluten","fish", "vegan"
        ],
        "tips": [ "tip1", "tip2", "tip3"] 
    },
    {
        "recipe_id": "17-777859", 
        "publisher": "Real Simple", 
        "source_url": "http://food2fork.com/view/37859", 
        "title": "Svininja pecena", 
        "image_url": "http://static.food2fork.com/chickenturnover2_300e6667e66.jpg", 
        "tags": [
            "gluten","fish", "vegan"
        ],
        "tips": [ "tip1", "tip2", "tip3"] 
    },
    {
        "recipe_id": "18-4d3252453899", 
        "publisher": "Not Simple", 
        "source_url": "http://food2fork.com/view/37859", 
        "title": "Posebna torta z keksi", 
        "image_url": "http://static.food2fork.com/chickenturnover2_300e6667e66.jpg", 
        "tags": [
            "gluten","fish", "vegan"
        ],
        "tips": [ "tip1", "tip2", "tip3"] 
    },
    { 
        "recipe_id": "19-ergtggsfdgfdsdf", 
        "publisher": "Real Simple", 
        "source_url": "http://food2fork.com/view/37859", 
        "title": "Keksi njami njami", 
        "image_url": "http://static.food2fork.com/chickenturnover2_300e6667e66.jpg", 
        "tags": [
            "gluten","fish", "vegan"
        ],
        "tips": [ "tip1", "tip2", "tip3"] 
    },
    {
        "recipe_id": "20-132323434", 
        "publisher": "Not Simple", 
        "source_url": "http://food2fork.com/view/37859", 
        "title": "Cokoladna torta z lesniki", 
        "image_url": "http://static.food2fork.com/chickenturnover2_300e6667e66.jpg", 
        "tags": [
            "gluten","fish", "vegan"
        ],
        "tips": [ "tip1", "tip2", "tip3"] 
    },
    {
        "recipe_id": "21-732323434", 
        "publisher": "Not Simple", 
        "source_url": "http://food2fork.com/view/37859", 
        "title": "Alergija zadnja", 
        "image_url": "http://static.food2fork.com/chickenturnover2_300e6667e66.jpg", 
        "tags": [
            "gluten","fish", "vegan"
        ],
        "tips": [ "tip1", "tip2", "tip3"] 
    }     
    ]`
  return JSON.parse(jsonStr)
}
var apiModule = require('../assets/api.js') // XXX test CommonJS
export default {
  name: 'recipe_editor_list',
  data () {
    var recipes = getRecipes()

    console.log('!!!!', apiModule.doSomething(), somethingElse())
    return {
      recipes: recipes
    }
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
