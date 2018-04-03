<template>
  <div class="recipe_editor">
    <h1>Editor</h1>

    <!-- recipe unuque id -->
    <div style="color: gray;" v-if="recipe.recipe_id">
        <p>Editing recipe with unique id: {{recipe.recipe_id}}</p>
    </div>

    <p id="error_msg" ref="error_msg" style="color: red;">{{errorMsg}}</p>

    <!-- title editor-->
    <div>
        <h2>Title:</h2>
        <input v-model="recipe.title">
    </div>

        <!-- RECIPE SOURCE URL -->
    <div>
      <h2>Recipe source URL (optional):</h2>
      <input v-model="recipe.source_url" type="text" name="source_url">
    </div>

    <!-- INGREDIENTS EDITOR-->
    <div>
        <h2>Ingredients:</h2>
        <p>Insert ingredients, one ingredient per line:</p>
        <textarea v-model="ingredientsText"></textarea>
    </div>


    <!-- INSTRUCTIONS EDITOR -->
    <div>
      <h2>Instructions:</h2>
      <div v-for="(instruction, index) in recipe.instructions">
        <h3>step {{index + 1}}</h3>
        <textarea v-model="recipe.instructions[index].instruction"></textarea>
        <span v-on:click="removeInstruction(index)"><icon style="color: red;"scale=1 name="remove"></icon></span>
      </div>
      <div v-on:click="addInstruction()"><icon style="color: red;"scale=1.5 name="plus"></icon></div>
    </div>

    <!-- TIPS EDITOR-->
    <div>
        <h2>Tips: (optional)</h2>
        <div v-for="(tip, index) in recipe.tips">
            <input v-model="recipe.tips[index]" type="text" :name="'tip'+index">
            <span v-on:click="removeTip(index)"><icon style="color: red;"scale=1 name="remove"></icon></span>
        </div>
        <div v-on:click="addTip()"><icon style="color: red;"scale=1.5 name="plus"></icon></div>
    </div>


    <router-link  v-bind:to="{name: 'RecipeEditorList'}"><button>Cancel</button></router-link>
    <button v-on:click="submit()">Submit</button>
  </div>
</template>

<script>
import { getRecipes, postOrPutRecipe } from './api'

export default {
  name: 'recipe_editor',
  data () {
    var recipe = {
      recipe_id: '',
      backend: 'manual-entry',
      title: '',
      source_url: '',
      ingredients: [],
      instructions: [],
      tips: []
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
      errorMsg: '',
      ingredientsText: '',
      recipe: recipe
    }
  },
  mounted: function () {
    // there are basicly two options. Eather we are editing new recipe, or we are editing exiting recipe.
    // If we are editing exiting recipe, then 'id' param is set and we need to feth exiting recipe from server
    // and populate editor fields with existing data
    var idParam = this.$route.params.id
    if (idParam && (typeof idParam === 'string' || idParam instanceof String) && idParam !== '') {
      getRecipes(idParam).then((recipe) => {
        if (recipe === undefined || recipe === null || typeof recipe !== 'object') {
          this.errorMsg = 'recipe ' + idParam + ' not found'
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
        // transform ingredients into textarea format
        var ingText = recipe.ingredients.reduce((text, line) => {
          console.log('->', text, line)
          return text + line + '\n'
        }, '')
        this.recipe = recipe
        this.ingredientsText = ingText
      }, (err) => {
        console.error('getRecipes promise: ', err)
        this.errorMsg = err
      })
    }
  },
  methods: {
    removeTip: function (index) {
      console.log('rem', index)
      this.recipe.tips.splice(index, 1)
    },
    addTip: function () {
      this.recipe.tips.push('')
    },
    removeInstruction: function (index) {
      console.log('rem', index)
      this.recipe.instructions.splice(index, 1)
    },
    addInstruction: function () {
      this.recipe.instructions.push({instruction: ''})
    },
    submit: function () {
      // transform textArea back to ingredients array
      this.recipe.ingredients = this.ingredientsText.split('\n')
      // filter out empty ingredients
      this.recipe.ingredients = this.recipe.ingredients.filter((ingredient) => {
        return ingredient.length > 0
      })
      // filter out empty instructions
      this.recipe.instructions = this.recipe.instructions.filter((inst) => {
        console.log('filter instruction ', inst.instruction, ' ', inst.instruction.length)
        return inst.instruction.length > 0
      })
      // replace
      // filter out empty tips
      this.recipe.tips = this.recipe.tips.filter((tip) => {
        return tip.length > 0
      })

      if (this.recipe.title === '') {
        this.errorMsg = 'You need to write recipe title!'
        this.$refs.error_msg.scrollIntoView()
        return
      }
      if (this.recipe.ingredients.length === 0) {
        this.errorMsg = 'You need to write ingrediend list!'
        this.$refs.error_msg.scrollIntoView()
        return
      }
      if (this.recipe.instructions.length === 0) {
        this.errorMsg = 'You need to write instructions!'
        this.$refs.error_msg.scrollIntoView()
        return
      }
      var idParam = this.$route.params.id
      var method = 'POST'
      // New manual recipes do not have ID set
      // In oposite case, we are doing recipe update. (HTTP PUT)
      if (typeof idParam === 'string' && idParam !== '') {
        method = 'PUT'
      }
      postOrPutRecipe(method, this.recipe).then((location) => {
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
