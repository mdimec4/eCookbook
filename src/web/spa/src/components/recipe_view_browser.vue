<template>
  <div class="recipe_view_browser">
      <h1>Recipe Browswer</h1>
      <div class="scroll-me">
        <a href="https://google.com" style="float: left;">
          <div class="kolom">Go To Google</div>
        </a>
        <div class="kolom" style="float: left;">a</div>
        <div class="kolom"></div>
        <div class="kolom"></div>
        <div class="kolom"></div>
        <div class="kolom"></div>
        <div class="kolom"></div>
        <div class="kolom"></div>
        <div class="kolom"></div>
        <div class="kolom"></div>
        <div class="kolom"></div>
        <div class="kolom"></div>
        <div class="kolom"></div>
        <div class="kolom"></div>
        <div class="kolom"></div>
        <div class="kolom"></div>
        <div class="clearfix"></div>
      </div>
</vue-scrollbar>
  </div>
</template>


<script>

function getRecipes (id) {
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
    }      
    ]`
  return JSON.parse(jsonStr)
}

export default {
  name: 'recipe_view_browser',
  data () {
    var recepies = getRecipes()
    if (!recepies || recepies === undefined) {
      recepies = []
    }
    return {
      recepies: recepies
    }
  },
  computed: {
  },
  mounted: function () {
    console.log('mounted')
    this.$nextTick(function () {
      console.log('mounted - $nextTick')
    })
  },
  updated () { // bug use computed of watcher instead. updated doc says so
    console.log('updated')
    this.$nextTick(function () {
    })
  },
  methods: {
  }
}
</script>
<style>
/*
 * {
    margin: 0px;
    padding: 0px;
    border-width: 0px;
 }
 */
html {
	height:100%;
}
body {
color: #000;
overflow: hidden;
font-family: 'Sanchez', serif;
font-size: 120%;
}
</style>

