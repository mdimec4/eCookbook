<template>
  <div class="allrecipes">
    <h1>allrecipes.com entry</h1>

    <p id="error_msg" ref="error_msg" style="color: red;">{{errorMsg}}</p>

    <!-- RECIPE SOURCE URL OR ID-->
    <div>
      <h2>insert allrecipes.com recipe URL or recipe ID:</h2>
      <p style="color: gray">Example allrecipes.com recipe URL would be: "https://www.allrecipes.com/recipe/246866/rigatoni-alla-genovese/"</p>
      <p style="color: gray">recipe ID of the same recipe would be "246866"</p>
      <input v-model="indentifier" type="text" name="source_url">
    </div>


    <router-link  v-bind:to="{name: 'RecipeEditorList'}"><button>Cancel</button></router-link>
    <button v-on:click="submit()">Submit</button>
  </div>
</template>

<script>
import { postOrPutRecipe } from './api'

export default {
  name: 'allrecipes',
  data () {
    var recipe = {
      recipe_id: '',
      backend: 'allrecipes.com',
      source_url: ''
    }
    return {
      indentifier: '',
      errorMsg: '',
      recipe: recipe
    }
  },
  watch: {
    indentifier: function (val) {
      var myRe = /^https:\/\/www\.allrecipes\.com\/recipe\/[0-9]+.*/g
      var myArray = myRe.exec(val)
      if (myArray !== null && myArray.length > 0) {
        this.recipe.source_url = val
      } else {
        this.recipe.source_url = ''
      }
      myRe = /^[0-9]*$/g
      myArray = myRe.exec(val)
      if (myArray !== null && myArray.length > 0) {
        this.recipe.recipe_id = 'allrecipes.com--' + val
      } else {
        this.recipe.recipe_id = ''
      }
    }
  },
  methods: {
    submit: function () {
      console.log(this.recipe)
      if (this.recipe.recipe_id === '' && this.recipe.source_url === '') {
        this.errorMsg = 'You need to write recipe ID or URL!'
        this.$refs.error_msg.scrollIntoView()
        return
      }
      postOrPutRecipe('POST', this.recipe).then((location) => {
        this.errorMsg = ''
        console.log('recipe url: ', location)
        // redirect back to recipe menu
        this.$router.push({name: 'RecipeEditorList'})
      }, (err) => {
        console.error('postOrPutRecipe promise: ', err)
        this.errorMsg = err
      })
    }
  }
}

</script>


<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
};
textarea {
  width: 90vw;
  height: 25vh;
  overflow:scroll;
};
input {
  width: 90vw;
};
</style>
