<template>
  <div class="recipe_view_browser">
      <div id="title">
        <h1>Recipe Browswer</h1>
      </div>

      <div id="position-info-container">
       <div v-on:click="scrollUp" class="position-info">Previus</div>
       <div class="position-info">page {{pagePosition}}</div>
       <div v-on:click="scrollDown" class="position-info">Next</div>
      </div>

      <div ref="scroll_me" id="scroll_me">
            <router-link v-for="recipe in recipes" class="kolom" v-bind:key="recipe.recipe_id" v-bind:to="{name: 'RecipeViewer', params: { userId: recipe.recipe_id}}">{{recipe.title}}</router-link>
      </div>
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

export default {
  name: 'recipe_view_browser',
  data () {
    var recipes = getRecipes()
    if (!recipes || recipes === undefined) {
      recipes = []
    }
    return {
      recipes: recipes,
      pagePosition: '0%'
    }
  },
  computed: {

  },
  mounted: function () {
    console.log('mounted')
    this.$nextTick(function () {
      console.log('mounted - $nextTick')
      this.renderScrollPosition()
    })
  },
  updated () { // bug use computed of watcher instead. updated doc says so
    console.log('updated')
    this.$nextTick(function () {
      this.renderScrollPosition()
    })
  },
  methods: {
    renderScrollPosition: function () {
      var sm = this.$refs.scroll_me
      var allSubpages = sm.scrollHeight / sm.clientHeight
      var subpage = ((sm.scrollTop + sm.clientHeight) / sm.scrollHeight) * allSubpages
      this.pagePosition = Math.floor((subpage / allSubpages) * 100).toString() + '%'
    },
    scrollUp: function (event) {
      console.log('scrollUp')
      var sm = this.$refs.scroll_me

      if (sm.scrollTop > 0) {
        var newScroll = sm.scrollTop - sm.clientHeight

        if (newScroll < 0) {
          newScroll = 0
        }
        sm.scrollTop = newScroll
        this.renderScrollPosition()
      }
    },
    scrollDown: function (event) {
      var sm = this.$refs.scroll_me

      if (sm.scrollHeight - sm.scrollTop > sm.clientHeight) {
        var newScroll = sm.scrollTop + sm.clientHeight

        if (newScroll > (sm.scrollHeight - sm.clientHeight)) {
          newScroll = sm.scrollHeight - sm.clientHeight
        }
        sm.scrollTop = newScroll
        this.renderScrollPosition()
      }
    }
  }
}
</script>
<style>

 * {
    margin: 0px;
    padding: 0px;
    border-width: 0px;
 };
html {
	height:100%;
}
body {
    color: #000;
    overflow: hidden;
    font-family: 'Sanchez', serif;
    font-size: 120%;
    margin: 0px;
    padding: 0px;
    border-width: 0px;
}
#title {
    width: 100vw;
    height: 10vh;

};
#position-info-container {
  border: solid;
  border-width: 1px 1px 1px 0px;
  border-color:black;
  width: calc(100vw - 1px);
  height: calc(10vh - 2px);
};
.position-info {
  float:left;
  border: solid;
  border-width: 0px 0px 0px 1px;
  border-color:black;
  width: calc(33% - 1px); /* subtract 1px border width */
  height: 100%;
};
#scroll_me {
    margin: 0 auto;
    width: calc(100vw - 2*2px);
    height: calc(80vh - 2*2px);
    overflow: hidden;
    border: solid;
    border-width: 2px 2px 2px 2px;
    border-color:black;
}
.kolom {
    float: left;
    width: 100%;
    height: calc(20% - 1px);
    text-align: center;
    border: solid;
    border-width: 0px 0px 1px 0px;
    border-color:black;
    overflow: hidden;
}
</style>

