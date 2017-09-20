<template>
  <div class="recipe_viewer">
    <div id="back-next">
      <button v-if="position > 0"v-on:click="prevPage">Back</button>

      <button v-if="this.scroll != -1" v-on:click="nextPage">Up</button>
      <button v-if="this.scroll != -1" v-on:click="prevPage">Down</button>

      <button v-if="position < instructions.length" v-on:click="nextPage">Next</button>
    </div>
    <h1>{{title}}</h1>
    <div v-if="position === 0" id="ingredient-list">
      <h2>Ingredients</h2>
      <div ref="content" id="content">
        <ul>
          <li v-for="ingredient in ingredients">
          {{ ingredient }}
          </li>
        </ul>
      </div>
    </div>
      <div v-if="position > 0" id="instructions-view">
        <h2>Step {{position}} / {{instructions.length}}</h2>
        <div ref="content" id="content">
          <div v-for="line in instructions[position - 1].instruction.split('\n')">
            {{line}}<br>
          </div>
        </div>
    </div>  
  </div>
</template>


<script>
// https://stackoverflow.com/questions/40730116/scroll-to-bottom-of-div-with-vue-js
// https://forum.vuejs.org/t/cant-change-dom-property/13397/2
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
      "instruction": "2 lorem ipsum sdfsdfgf \\n sdf\\nga\\nsd\\nga\\nsd\\nfg\\ndd\\ndd\\ndddd\\nddd\\ndd\\ndd\\ndd\\nddd\\ndd\\nddddd\\ndddddd\\nddddd\\nmiha" 
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
      instructions: recipe.instructions,
      scroll: -1
    }
  },
  mounted: function () {
    this.$nextTick(function () {
      // Code that will run only after the
      // entire view has been rendered
      // XXX this doesnt work
      this.$refs.content.style.height = (document.body.clientHeight - this.$refs.content.offsetTop).toString()
    })
  },
  updated () { // bug use computed of watcher instead. updated doc says so
    console.log('updated offsetTop', this.$refs.content.offsetTop)
    console.log('updated offsetHeight', this.$refs.content.offsetHeight)
    console.log('updated scrollHeight', this.$refs.content.scrollHeight)
    console.log('updated scrollTop', this.$refs.content.scrollTop)
    console.log('updated clientHeight', this.$refs.content.clientHeight)
    console.log('window.inerHeight', window.innerHeight)
    console.log('document.body.clientHeight', document.body.clientHeight)
    if (this.$refs.content.scrollHeight <= this.$refs.content.offsetHeight) {
      this.scrool = -1
      return
    }
    if (this.scroll === -1) {
      this.scroll = 0
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
body {
  width: 100%;
  height: 100%;
}
#content {
  -position: fixed;
  text-align: left;
  overflow: auto;
  -height: 50vh;
  -display:table;
};
#back-next {
}
</style>

