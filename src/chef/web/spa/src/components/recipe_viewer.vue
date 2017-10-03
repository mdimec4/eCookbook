<template>
  <div class="recipe_viewer">
    <!-- header buttons -->
    <div id="back-next-home">
      <span><icon scale=2 name="home"></icon></span>
      <span v-if="this.rPagePrevShow" v-on:click="prevPage"><icon scale=2 name="arrow-left"></icon></span>
      <span v-if="this.scrollUpShow" v-on:click="scrollUp"><icon scale=2 name="arrow-up"></icon></span>
      <span v-if="this.scrollDownShow" v-on:click="scrollDown"><icon scale=2 name="arrow-down"></icon></span>
      <span v-if="this.rPageNextShow" v-on:click="nextPage"><icon scale=2 name="arrow-right"></icon></span>
    </div>
  
    <!--title-->
    <div id="title">
      <h1>{{title}}</h1>
    </div>

    <!--print content title-->
    <div v-if="rwhat==='ingredients'" id="subtitle">
      <h2>Ingredients</h2>
    </div>
    <div v-else-if="rwhat==='instructions'" id="subtitle">
      <h2>Step {{instPage+1}} / {{instructions.length}}</h2>
    </div>
    <div v-else-if="rwhat==='tips'" id="subtitle">
      <h2>Tips</h2>
    </div>

    <!--print content-->
    <div ref="content" id="content"> 
      <div v-if="rwhat==='ingredients'">
        <ul ref="textSizeHelp">
          <li v-for="ingredient in ingredients">
            {{ ingredient }}
          </li>
        </ul>
      </div>
      <div v-else-if="rwhat==='instructions'" ref="textSizeHelp">
        <div v-for="line in instructions[instPage].instruction.split('\n')">
          {{line}}
        </div>
      </div>
      <div v-else-if="rwhat==='tips'">
        <ul ref="textSizeHelp">
          <li v-for="tip in tips">
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
      "instruction": "1-3 lorem ipsum sdfsdfgf lorem ipsum sdfsdfgf\\n2-sdf\\n3-ga\\n4-sd\\n5-gasd\\n-6fg\\n7-dd\\n-8dddddd\\n-9dddlorem ipsum sdfsdfgf \\n10- sdf\\n11-ga\\n12-sd\\n13-gasd\\n14-fg\\n15-dd\\n16-dddddd\\n17-ddd\\n18- sdf\\n19-ga\\n20-sd\\n21-gasd\\n22-fg\\n23-dd\\n24-dddddd\\n25-ddd\\n26-dd\\n27-dd\\n28-dd\\n29-dd\\n30-dddddddd\\n31-dddddd\\n32-dd\\n33-d\\n34-d\\n35-d\\n36-12\\n37-23\\n38-34\\n39-4\\n40-566\\n41-778\\n42-89\\n843-90-\\n44-0-\\n45-646\\n46-364\\n47-6346\\n48-3436\\n49-46\\n50-3456534\\n51-6544\\n52-4444\\n53-44\\n54-4\\n55-4\\n56-4\\n57-4\\n58-4\\n59-44444\\n60-44444444444\\n61-4444444\\n62-444444\\n63-444444\\n64-miha" 
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
    var tmpNextShow = (recipe.instructions != null && recipe.instructions.length > 0) || (recipe.tips != null && recipe.tips.length > 0)
    return {
      rwhat: 'ingredients',
      instPage: 0,
      rPageNextShow: tmpNextShow,
      rPagePrevShow: false,
      title: recipe.title,
      ingredients: recipe.ingredients,
      instructions: recipe.instructions,
      tips: recipe.tips,
      scrollUpShow: false,
      scrollDownShow: false,
      subpagePosition: '0%'
    }
  },
  computed: {
    allPages: function () {
      var count = 0
      if (this.ingredients) {
        count++
      }
      if (this.instructions) {
        count += this.instructions.length
      }
      if (this.tips) {
        count++
      }
      return count
    },
    curentPage: function () {
      var baseIng = 0
      if (this.ingredients) {
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
      this.renderScroll()
      this.renderScrollPosition()
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
      this.renderScroll() // TODO bug use computed of watcher instead. updated() doc says so
      this.renderScrollPosition()
    })
  },
  methods: {
    prevPage: function (event) {
      switch (this.rwhat) {
        case 'ingredients':
          break
        case 'instructions':
          this.$refs.content.scrollTop = 0
          if (this.instPage > 0) {
            this.instPage--
            this.rPageNextShow = true
            this.rPagePrevShow = true
          } else {
            this.rwhat = 'ingredients'
            this.rPagePrevShow = false
            this.rPageNextShow = true
          }
          break
        case 'tips':
          this.$refs.content.scrollTop = 0
          this.rwhat = 'instructions'
          this.rPagePrevShow = true
          this.rPageNextShow = true
          break
      }
    },
    nextPage: function (event) {
      switch (this.rwhat) {
        case 'ingredients':
          this.$refs.content.scrollTop = 0
          this.rwhat = 'instructions'
          this.rPagePrevShow = true
          this.rPageNextShow = (this.instructions != null && this.instructions.length > 0) || (this.tips != null && this.tips.length > 0)
          break
        case 'instructions':
          this.$refs.content.scrollTop = 0
          if (this.instPage + 1 < (this.instructions.length)) {
            this.instPage++
            this.rPageNextShow = (this.instPage + 1 < ((this.instructions.length))) || (this.tips != null && this.tips.length > 0)
            this.rPagePrevShow = true
          } else {
            if (this.tips != null && this.tips.length > 0) {
              this.rwhat = 'tips'
              this.rPagePrevShow = true
              this.rPageNextShow = false
            }
          }
          break
        case 'tips':
          break
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
      console.log(this.$refs.textSizeHelp.childNodes)
      if (this.$refs.textSizeHelp.hasChildNodes()) {
        lineHeight = this.$refs.textSizeHelp.childNodes[0].clientHeight
      }
      if (ce.scrollTop > 0) {
        var newScroll = ce.scrollTop - (Math.floor(ce.clientHeight / lineHeight) * lineHeight)

        if (newScroll < 0) {
          newScroll = 0
        }
        ce.scrollTop = newScroll
        this.renderScroll()
      }
    },
    scrollDown: function (event) {
      var ce = this.$refs.content
      var lineHeight = 1
      // read text line height
      console.log(this.$refs.textSizeHelp.childNodes)
      if (this.$refs.textSizeHelp.hasChildNodes()) {
        lineHeight = this.$refs.textSizeHelp.childNodes[0].clientHeight
      }
      console.log('1 fs ', lineHeight)
      if (ce.scrollHeight - ce.scrollTop > ce.clientHeight) {
        var newScroll = ce.scrollTop + ce.clientHeight
        // make scroll multiple of lineHeight
        newScroll = Math.floor(newScroll / lineHeight) * lineHeight
        if (newScroll > (ce.scrollHeight - ce.clientHeight)) {
          newScroll = ce.scrollHeight - ce.clientHeight
        }
        ce.scrollTop = newScroll
        this.renderScroll()
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
 }
html {
	height:100%;
}
body {
color: #000;
overflow: hidden;
font-family: 'Sanchez', serif;
font-size: 120%;
}
#back-next-home {
    width: 100vw;
    height: 5vh
}
#title {
    width: 100vw;
    height: 15vh
}
#subtitle {
    width: 100vw;
    height: 10vh
}
#content {
    text-align: left;
    display: block;
    overflow: hidden;
    width: 100vw;
    height: 65vh
};
#position-info-container {
  border: solid;
  border-width: 1px 1px 1px 0px;
  border-color:black;
  position: absolute;
  left: 0px;
  bottom: 0px;
  width: 100vw;
  height: 5vh
};
.position-info {
  float:left;
  border: solid;
  border-width: 0px 0px 0px 1px;
  border-color:black;
  width: calc(50% - 1px); /* subtract 1px border width */
};
</style>

