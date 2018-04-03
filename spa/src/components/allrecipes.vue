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
    // there are basicly two options. Eather we are editing new recipe, or we are editing exiting recipe.
    // If we are editing nre recipe, then 'id' param is not set and we will nt fetch exiting recipe from server.
    var idParam = this.$route.params.id
    if (!idParam || !(typeof idParam === 'string' || idParam instanceof String) || idParam === '') {
      // we want to make UI user frendly so we will offer one
      // empty ingredient/instruction/tip in advance
      // if user will click subbmit and he/she didn't insert nthing into emty field
      // created here, this wont be a problem. since submit will filter out
      // empty fields anyway, before further processing
      if (recipe.instructions.length === 0) {
        recipe.instructions.push({instruction: ''})
      }
      if (recipe.tips.length === 0) {
        recipe.tips.push('')
      }
    }
    return {
      indentifier: '',
      errorMsg: '',
      recipe: recipe
    }
  },
  watch: {
    indentifier: function (val) {
          var myRe = /^https:\/\/www\.allrecipes\.com\/recipe\/.*/g
          var myArray = myRe.exec(val)
          if (myArray,length > 0) {
               this.recipe.source_url = val
               return
          }
          myRe = /^[0-9]*$/g
          myArray = myRe.exec(val)
          if (myArray,length > 0) {
               this.recipe.recipe_id = val
               return
          }
    }
  },
  methods: {
        submit: function () { 
                if (this.recipe.title === '') {
                         this.errorMsg = 'You need to write recipe title!'
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
