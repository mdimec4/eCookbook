<template>
    <div class="recipe_viewer">
        <h1>{{title}}</h1>
        <div v-if="position === 0" id="ingredient-list">
          <h2>Ingredients</h2>
          <ul>
            <li v-for="ingredient in ingredients">
              {{ ingredient }}
            </li>
          </ul>
        </div>
        <div v-if="position > 0" id="instructions-view">
          <h2>Step {{position}} / {{instructions.length}}</h2>
            <div>
              {{instructions[position - 1].instruction}}
            </div>
        </div>
        <button v-if="position > 0"v-on:click="prevPage">Back</button>
        <button v-if="position < instructions.length" v-on:click="nextPage">Next</button>
    </div>
</template>


<script>
function getRecipe (id) {
  var jsonStr = `{ 
    "recipe_id": "37859", 
    "publisher": "Real Simple", 
    "source_url": "http://food2fork.com/view/37859", 
    "title": "Chicken and Gruyre Turnovers", 
    "image_url": "http://static.food2fork.com/chickenturnover2_300e6667e66.jpg", 
    "tags": [
        "gluten","fish", "vegan"
    ],
    "ingredients": [
      "1 1/2 cups shredded rotisserie chicken",
      "1 1/2 cups grated Gruyre",
      "1 cup frozen peas",
      "2 sheets (one 17.25-ounce package) frozen puff pastry, thawed",
      "1 large egg, beaten",
      "1/4 cup Dijon mustard"
    ],
    "instructions": [
    {
      "number": 1,
      "image_url": "http://static.food2fork.com/chickenturnover2_300e6667e66.jpg",
      "instruction": "lorem ipsum sdfgasdgasdfg" 
    },
    {
      "number": 2, 
      "image_url": "http://static.food2fork.com/chickenturnover2_300e6667e66.jpg", 
      "instruction": "2 lorem ipsum \\n sdfgasdgasdfgdddddddddddddddddddddddddddddddddddddd miha" 
    }
    ], 
    "tips": [ "tip1", "tip2", "tip3"] 
    }`
  return JSON.parse(jsonStr)
}
/*
function sliceArrOfStingsToSets (maxChar, arr) {
  var joinPartialCollection = function (collection, partialCollection) {
    if (partialCollection.length > 0) {
      collection.push(partialCollection)
    }
    return collection
  }
  var joinSingleString = function (collection, str) {
    collection.push([str])
    return collection
  }

  var rec = function (arr, index, maxChar, partialCollection, partialLen, collection) {
    if (index >= arr.length) {
      return joinPartialCollection(collection, partialCollection)
    }

    // if next string is to long. add it to new category
    if (arr[index].length > maxChar) {
      return rec(arr, index + 1, maxChar, [], 0,
        joinSingleString(joinPartialCollection(collection, partialCollection), arr[index]))
    }

    if (partialLen + arr[index].length > maxChar) {
      return rec(arr, index + 1, maxChar, [arr[index]], arr[index].length,
         joinPartialCollection(collection, partialCollection))
    }

    partialCollection.push(arr[index])
    partialLen += arr[index].length
    return rec(arr, index + 1, maxChar, partialCollection, partialLen, collection)
  }
  return rec(arr, 0, maxChar, [], 0, [])
}
*/

export default {
  name: 'recipe_viewer',
  data () {
    var recipe = getRecipe(this.$route.params.id)
    return {
      position: 0,
      title: recipe.title,
      ingredients: recipe.ingredients,
      instructions: recipe.instructions
    }
  },
  methods: {
    prevPage: function (event) {
      if (this.position > 0) {}
      this.position--
    },
    nextPage: function (event) {
      if (this.position < this.instructions.length) {
        this.position++
      }
    }
  }
}
</script>
<style>
.recipe_viewer {
  text-align: left;
};
</style>

