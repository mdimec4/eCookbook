<template>
  <div class="recipe_viewer">
    <div id="back-next">
      <button v-if="rpage > 0"v-on:click="prevPage">Back</button>

      <button v-if="this.scrollUpShow" v-on:click="scrollUp">Up</button>
      <button v-if="this.scrollDownShow" v-on:click="scrollDown">Down</button>

      <button v-if="rpage < instructions.length" v-on:click="nextPage">Next</button>
    </div>
    <h1>{{title}}</h1>
    <div v-if="rpage === 0" id="ingredient-list">
      <h2>Ingredients</h2>
      <div ref="content" id="content">
        <ul>
          <li v-for="ingredient in ingredients">
          {{ ingredient }}
          </li>
        </ul>
      </div>
    </div>
      <div v-if="rpage > 0" id="instructions-view">
        <h2>Step {{rpage}} / {{instructions.length}}</h2>
        <div ref="content" id="content">
          <div v-for="line in instructions[rpage - 1].instruction.split('\n')">
            {{line}}<br>
          </div>
        </div>
    </div>  
  </div>
</template>


<script>
// https://stackoverflow.com/questions/40730116/scroll-to-bottom-of-div-with-vue-js
// https://forum.vuejs.org/t/cant-change-dom-property/13397/2
// https://stackoverflow.com/questions/4106538/difference-between-offsetheight-and-clientheight
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
      "instruction": "2 lorem ipsum sdfsdfgf \\n sdf\\nga\\nsd\\ngasd\\nfg\\ndd\\ndddddd\\nddd\\ndd\\ndd\\ndd\\ndd\\ndddddddd\\ndddddd\\nddddd\\n1223344\\n566778\\n89890-\\n0-\\n646\\n364\\n6346\\n343646\\n3456534\\n6544\\n4444\\n4444444\\n44444\\n44444444444\\n4444444\\n444444\\n444444\\nmiha" 
    },
    {
      "number": 3, 
      "image_url": "http://static.food2fork.com/chickenturnover2_300e6667e66.jpg", 
      "instruction": "3 lorem ipsum sdfsdfgf \\n sdf\\nga\\nsd\\ngasd\\nfg\\ndd\\ndddddd\\nddd\\ndd\\ndd\\ndd\\ndd\\ndddddddd\\ndddddd\\nddddd\\n1223344\\n566778\\n89890-\\n0-\\n646\\n364\\n6346\\n343646\\n3456534\\n6544\\n4444\\n4444444\\n44444\\n44444444444\\n4444444\\n444444\\n444444\\nmiha" 
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
      rpage: 0,
      title: recipe.title,
      ingredients: recipe.ingredients,
      instructions: recipe.instructions,
      scrollUpShow: false,
      scrollDownShow: false
    }
  },
  mounted: function () {
    console.log('mounted')
    this.$nextTick(function () {
      console.log('mounted - $nextTick')
      // Code that will run only after the
      // entire view has been rendered
      // XXX this doesnt work
      // var ce = this.$refs.content
      // console.log(ce.style.height)
      // ce.style.height = window.inerHeight - (ce.getBoundingClientRect().top + ce.clientTop + window.getComputedStyle(ce).padingTop)
      // console.log(ce.style.height)
      // ce.scrollTop = ce.scrollHeight
      this.renderScroll()
    })
  },
  updated () { // bug use computed of watcher instead. updated doc says so
    console.log('updated')
    console.log('updated offsetTop', this.$refs.content.offsetTop)
    console.log('updated offsetHeight', this.$refs.content.offsetHeight)
    console.log('updated scrollHeight', this.$refs.content.scrollHeight)
    console.log('updated scrollTop', this.$refs.content.scrollTop)
    console.log('updated clientHeight', this.$refs.content.clientHeight)
    console.log('window.inerHeight', window.innerHeight)
    console.log('document.body.clientHeight', document.body.clientHeight)
    this.$nextTick(function () {
      console.log('updated - $nextTick')
      this.renderScroll()
    })
  },
  methods: {
    prevPage: function (event) {
      if (this.rpage > 0) {
        this.$refs.content.scrollTop = 0
        this.rpage--
      }
    },
    nextPage: function (event) {
      if (this.rpage < this.instructions.length) {
        this.$refs.content.scrollTop = 0
        this.rpage++
      }
    },
    renderScroll: function () {
      var ce = this.$refs.content
      if (ce.clientHeight === ce.scrollHeight) {
        this.scrollUpShow = false
        this.scrollDownShow = false
        return
      }
      if (ce.scrollHeight - ce.scrollTop > ce.clientHeight) {
        this.scrollDownShow = true
      } else {
        this.scrollDownShow = false
      }
      if (ce.scrollTop > 0) {
        this.scrollUpShow = true
      } else {
        this.scrollUpShow = false
      }
    },
    scrollUp: function (event) {
      var ce = this.$refs.content
      if (ce.scrollTop > 0) {
        var newScroll = ce.scrollTop - ce.clientHeight
        if (newScroll < 0) {
          newScroll = 0
        }
        ce.scrollTop = newScroll
        this.renderScroll()
      }
    },
    scrollDown: function (event) {
      var ce = this.$refs.content
      if (ce.scrollHeight - ce.scrollTop > ce.clientHeight) {
        var newScroll = ce.scrollTop + ce.clientHeight
        if (newScroll > ce.scrollHeight) {
          newScroll = ce.scrollHeight
        }
        ce.scrollTop = newScroll
        this.renderScroll()
      }
    }
  }
}
</script>
<style>
html {
	height:100%;
}
body {
color: #000;
overflow: hidden;
font-family: 'Sanchez', serif;
font-size: 120% !important

}
#content {
    margin: 0 auto !important;
    width: auto !important;
    display: block;
    overflow: hidden;
    height: calc(100vh - 354px);

};
#back-next {
}
</style>

