<template>
  <div class="recipe_view_browser">
      <div id="title">
        <h1>Recipe Browswer</h1>
      </div>

      <div id="selector-container">
       <div v-on:click="scrollUp" class="selector">Previus</div>
       <div class="selector">page {{pagePosition}}</div>
       <div v-on:click="scrollDown" class="selector">Next</div>
      </div>

      <div ref="scroll_me" id="scroll_me">
            <router-link v-for="recipe in recipes" class="kolom" v-bind:key="recipe.recipe_id" v-bind:to="{name: 'RecipeViewer', params: { id: recipe.recipe_id}}">{{recipe.title}}</router-link>
      </div>
  </div>
</template>

<script>
import { getRecipes } from './api'

export default {
  name: 'recipe_view_browser',
  data () {
    return {
      recipes: [],
      pagePosition: '0/0'
    }
  },
  mounted: function () {
    console.log('mounted')
    this.$nextTick(function () {
      console.log('mounted - $nextTick')
      this.renderScrollPosition()
    })
    getRecipes().then((recipes) => {
      this.recipes = recipes
    }, (err) => {
      console.error('getRecipes promise: ', err)
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
      this.pagePosition = Math.ceil(subpage).toString() + '/' + Math.ceil(allSubpages).toString()
      console.log('page ', subpage, ' / ', allSubpages)
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
<style scoped>
.recipe_view_browser {
 position: fixed;
 top: 0px;
 left: 0px;
 font-size: 150%;
};
#title {
    width: 100vw;
    height: 10vh;

};
#selector-container {
  border: solid;
  border-width: 1px 1px 1px 0px;
  border-color:black;
  width: calc(100vw - 1px);
  height: calc(10vh - 2px);
};
.selector {
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

