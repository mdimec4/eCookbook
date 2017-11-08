<template>
  <div class="recipe_viewer">
    <!-- header buttons -->
    <div id="menu">
      <router-link v-bind:to="{name: 'RecipeViewBrowser'}" style="color: black;"><icon scale=2 name="home"></icon></router-link>
      <!--
        <span v-if="this.rPagePrevShow" v-on:click="buttonBack"><icon scale=2 name="arrow-left"></icon></span>
      <span v-if="this.rPageNextShow" v-on:click="buttonNext"><icon scale=2 name="arrow-right"></icon></span>
      -->
    </div>

    <!-- transperend overlayng almost whole page next and prev buttons-->
    <span class="transperentPrevNext"  id="transperentPrev" v-on:click="buttonBack"></span>
    <span class="transperentPrevNext" id="transperentNext" v-on:click="buttonNext"></span>


    <!--title-->
    <div id="title">
      <h1>{{recipe.title}}</h1>
    </div>

    <!--print content title-->
    <div v-if="rwhat==='ingredients'" id="subtitle">
      <h2>Ingredients</h2>
    </div>
    <div v-else-if="rwhat==='instructions'" id="subtitle">
      <h2>Step {{instPage+1}} / {{recipe.instructions.length}}</h2>
    </div>
    <div v-else-if="rwhat==='tips'" id="subtitle">
      <h2>Tips</h2>
    </div>

    <!--print content-->
    <div ref="content" id="content">
      <div v-if="rwhat==='ingredients'">
        <ul ref="textSizeHelp">
          <li v-for="ingredient in recipe.ingredients">
            {{ ingredient }}
          </li>
        </ul>
      </div>
      <div v-else-if="rwhat==='instructions'" ref="textSizeHelp">
        <div v-for="line in recipe.instructions[instPage].instruction.split('\n')">
          {{line}}
        </div>
      </div>
      <div v-else-if="rwhat==='tips'">
        <ul ref="textSizeHelp">
          <li v-for="tip in recipe.tips">
            {{ tip }}
          </li>
        </ul>
      </div>
    </div>

    <div id="position-info-container">
      <div class="position-info">page {{curentPage}}/{{ allPages}}</div>
      <div class="position-info">supbapge position {{subpagePosition}}</div>
    </div>
  </div>
</template>


<script>
import { getRecipes } from './api'
// https://stackoverflow.com/questions/40730116/scroll-to-bottom-of-div-with-vue-js
// https://forum.vuejs.org/t/cant-change-dom-property/13397/2
// https://stackoverflow.com/questions/4106538/difference-between-offsetheight-and-clientheight
// http://vanseodesign.com/css/css-positioning/ // understand pusitioning
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
  lastPageChange: '',
  data () {
    var recipe = {
      recipe_id: '',
      title: '',
      source_url: '',
      ingredients: [],
      instructions: [],
      tips: []
    }
    return {
      rwhat: 'ingredients',
      instPage: 0,
      // rPageNextShow: tmpNextShow,
      // rPagePrevShow: false,
      recipe: recipe,
      subpagePosition: '0%'
    }
  },
  computed: {
    allPages: function () {
      var count = 0
      if (this.recipe.ingredients) {
        count++
      }
      if (this.recipe.instructions) {
        count += this.recipe.instructions.length
      }
      if (this.recipe.tips) {
        count++
      }
      return count
    },
    curentPage: function () {
      var baseIng = 0
      if (this.recipe.ingredients) {
        baseIng = 1
      }
      switch (this.rwhat) {
        case 'ingredients':
          return 1
        case 'instructions':
          return baseIng + (this.instPage + 1)
        case 'tips':
          return baseIng + (this.instPage + 1) + 1
      }
      return -1
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
      this.renderScrollPosition()
    })

    var idParam = this.$route.params.id
    if (idParam && (typeof idParam === 'string' || idParam instanceof String) && idParam !== '') {
      getRecipes(idParam).then((recipe) => {
        if (recipe === undefined || recipe === null || typeof recipe !== 'object') {
          console.error('recipe ' + idParam + ' not found')
          return
        }
        if (recipe.recipe_id === undefined || recipe.recipe_id == null) {
          recipe.recipe_id = ''
        }
        if (recipe.title === undefined || recipe.title == null) {
          recipe.title = ''
        }
        if (recipe.source_url === undefined || recipe.source_url == null) {
          recipe.source_url = ''
        }
        if (recipe.ingredients === undefined || recipe.ingredients == null) {
          recipe.ingredients = []
        }
        if (recipe.instructions === undefined || recipe.instructions == null) {
          recipe.instructions = []
        }
        if (recipe.tips === undefined || recipe.tips == null) {
          recipe.tips = []
        }
        this.recipe = recipe
      }, (err) => {
        console.error('getRecipes promise: ', err)
      })
    } else {
      console.error('recipe id parameter missing')
    }
  },
  updated () { // bug use computed of watcher instead. updated doc says so
    console.log('updated')
    this.$nextTick(function () {
      console.log('updated - $nextTick')
      var ce = this.$refs.content
      if (this.lastPageChange === 'next') {
        this.lastPageChange = ''
        ce.scrollTop = 0
      } else if (this.lastPageChange === 'prev') {
        this.lastPageChange = ''
        ce.scrollTop = ce.scrollHeight - ce.clientHeight
      }
      this.renderScrollPosition()
    })
  },
  methods: {
    buttonNext: function () {
      var ce = this.$refs.content
      if (ce.scrollHeight - ce.scrollTop > ce.clientHeight) {
        this.scrollDown()
        return
      }
      this.nextPage()
    },
    buttonBack: function () {
      var ce = this.$refs.content
      if (ce.scrollTop > 0) {
        this.scrollUpShow = true
        this.scrollUp()
        return
      }
      this.prevPage()
    },
    prevPage: function (event) {
      switch (this.rwhat) {
        case 'ingredients':
          break
        case 'instructions':
          this.lastPageChange = 'prev'
          if (this.instPage > 0) {
            this.instPage--
            // this.rPageNextShow = true
            // this.rPagePrevShow = true
          } else {
            this.rwhat = 'ingredients'
            // this.rPagePrevShow = false
            // this.rPageNextShow = true
          }
          break
        case 'tips':
          this.lastPageChange = 'prev'
          this.rwhat = 'instructions'
          // this.rPagePrevShow = true
          // this.rPageNextShow = true
          break
      }
    },
    nextPage: function (event) {
      switch (this.rwhat) {
        case 'ingredients':
          this.lastPageChange = 'next'
          this.rwhat = 'instructions'
          // this.rPagePrevShow = true
          // this.rPageNextShow = (this.recipe.instructions != null && this.recipe.instructions.length > 0) || (this.recipe.tips != null && this.recipe.tips.length > 0)
          break
        case 'instructions':
          if (this.instPage + 1 < (this.recipe.instructions.length)) {
            this.lastPageChange = 'next'
            this.instPage++
            // this.rPageNextShow = (this.instPage + 1 < ((this.recipe.instructions.length))) || (this.tips != null && this.tips.length > 0)
            // this.rPagePrevShow = true
          } else {
            if (this.recipe.tips != null && this.recipe.tips.length > 0) {
              this.lastPageChange = 'next'
              this.rwhat = 'tips'
              // this.rPagePrevShow = true
              // this.rPageNextShow = false
            }
          }
          break
        case 'tips':
          break
      }
    },
    renderScrollPosition: function () {
      var ce = this.$refs.content
      var allSubpages = ce.scrollHeight / ce.clientHeight
      var subpage = ((ce.scrollTop + ce.clientHeight) / ce.scrollHeight) * allSubpages
      this.subpagePosition = Math.floor((subpage / allSubpages) * 100).toString() + '%'
    },
    scrollUp: function (event) {
      var ce = this.$refs.content
      var lineHeight = 1
      // read text line height
      if (this.$refs.textSizeHelp.hasChildNodes()) {
        lineHeight = this.$refs.textSizeHelp.childNodes[0].clientHeight
      }
      if (ce.scrollTop > 0) {
        var newScroll = ce.scrollTop - (Math.floor(ce.clientHeight / lineHeight) * lineHeight)

        if (newScroll < 0) {
          newScroll = 0
        }
        ce.scrollTop = newScroll
        this.renderScrollPosition()
      }
    },
    scrollDown: function (event) {
      var ce = this.$refs.content
      var lineHeight = 1
      // read text line height
      if (this.$refs.textSizeHelp.hasChildNodes()) {
        lineHeight = this.$refs.textSizeHelp.childNodes[0].clientHeight
      }
      if (ce.scrollHeight - ce.scrollTop > ce.clientHeight) {
        var newScroll = ce.scrollTop + ce.clientHeight
        // make scroll multiple of lineHeight
        newScroll = Math.floor(newScroll / lineHeight) * lineHeight
        if (newScroll > (ce.scrollHeight - ce.clientHeight)) {
          newScroll = ce.scrollHeight - ce.clientHeight
        }
        ce.scrollTop = newScroll
        this.renderScrollPosition()
      }
    }
  }
}
</script>
<style scoped>
 * {
    margin: 0px;
    padding: 0px;
    border-width: 0px;
 };
html {
	height:100%;
};

body {
  color: #000;
  overflow: hidden;
  font-family: 'Sanchez', serif;
  font-size: 120%;

  margin: 0px;
  padding: 0px;
  border-width: 0px;
};

.transperentPrevNext {
  position: absolute;
  background: transparent;
  height: 100%;
  width: 50vw;
  z-index: 1;
};
#transperentPrev {
  left: 0;
};
#transperentNext {
   left: 50vw;
};


#menu {
    width: 100vw;
    height: 5vh;
}
#title {
    width: 100vw;
    height: 10vh;
}
#subtitle {
    width: 100vw;
    height: 5vh;
}
#content {
    padding-left: 5vw;
    text-align: left;
    display: block;
    overflow: hidden;
    width: 95vw;
    height: 75vh;
};
#position-info-container {
  border: solid;
  border-width: 1px 1px 1px 0px;
  border-color:black;
  position: fixed;
  left: 0px;
  bottom: 0px;
  width: calc(100vw - 1px);
  height: 5vh;
};
.position-info {
  float:left;
  border: solid;
  border-width: 0px 0px 0px 1px;
  border-color:black;
  width: calc(50% - 1px); /* subtract 1px border width */
  height: 100%;
};
</style>

