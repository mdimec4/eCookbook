<template>
  <div class="recipe_editor">
    <h1>Editor</h1>

    <div style="color: gray;" v-if="recipe.recipe_id">
        <p>Editing recipe with unique id: {{recipe.recipe_id}}</p>
    </div>

    <!-- title editor-->
    <div>
        <h2>Title:</h2>
        <input v-model="recipe.title">
    </div>

    <!-- INGREDIENTS EDITOR-->
    <div>
        <h2>Ingredients:</h2>
        <div v-for="(ingredient, index) in recipe.ingredients">
            <input v-model="recipe.ingredients[index]" type="text" :name="'ingredient'+index">
            <span v-on:click="removeIngredient(index)"><icon style="color: red;"scale=1 name="remove"></icon></span>
        </div>
        <div v-on:click="addIngredient()"><icon style="color: red;"scale=1.5 name="plus"></icon></div>
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
        <h2>Tips:</h2>
        <div v-for="(tip, index) in recipe.tips">
            <input v-model="recipe.tips[index]" type="text" :name="'tip'+index">
            <span v-on:click="removeTip(index)"><icon style="color: red;"scale=1 name="remove"></icon></span>
        </div>
        <div v-on:click="addTip()"><icon style="color: red;"scale=1.5 name="plus"></icon></div>
    </div>


    <button v-on:click="submit()">Submit</button>
  </div>
</template>

<script>
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
      "image_url": "http://static.food2fork.com/chickenturnover2_300e6667e66.jpg",
      "instruction": "lorem ipsum sdfgasdgasdfg" 
    },
    {
      "image_url": "http://static.food2fork.com/chickenturnover2_300e6667e66.jpg", 
      "instruction": "2 lorem ipsum sdfs<br>dfgf \\n sdf\\nga\\nsd\\ngasd\\nfg\\ndd\\ndddddd\\nddd\\ndd\\ndd\\ndd\\ndd\\ndddddddd\\ndddddd\\nddddd\\n1223344\\n566778\\n89890-\\n0-\\n646\\n364\\n6346\\n343646\\n3456534\\n6544\\n4444\\n4444444\\n44444\\n44444444444\\n4444444\\n444444\\n444444\\nmiha" 
    },
    { 
      "image_url": "http://static.food2fork.com/chickenturnover2_300e6667e66.jpg", 
      "instruction": "1-3 lorem ipsum sdfsdfgf lorem ipsum sdfsdfgf\\n2-sdf\\n3-ga\\n4-sd\\n5-gasd\\n-6fg\\n7-dd\\n-8dddddd\\n-9dddlorem ipsum sdfsdfgf \\n10- sdf\\n11-ga\\n12-sd\\n13-gasd\\n14-fg\\n15-dd\\n16-dddddd\\n17-ddd\\n18- sdf\\n19-ga\\n20-sd\\n21-gasd\\n22-fg\\n23-dd\\n24-dddddd\\n25-ddd\\n26-dd\\n27-dd\\n28-dd\\n29-dd\\n30-dddddddd\\n31-dddddd\\n32-dd\\n33-d\\n34-d\\n35-d\\n36-12\\n37-23\\n38-34\\n39-4\\n40-566\\n41-778\\n42-89\\n843-90-\\n44-0-\\n45-646\\n46-364\\n47-6346\\n48-3436\\n49-46\\n50-3456534\\n51-6544\\n52-4444\\n53-44\\n54-4\\n55-4\\n56-4\\n57-4\\n58-4\\n59-44444\\n60-44444444444\\n61-4444444\\n62-444444\\n63-444444\\n64-miha" 
    }
    ], 
    "tips": [ "tip1", "tip2", "tip3"] 
    }`
  return JSON.parse(jsonStr)
}

export default {
  name: 'recipe_editor',
  data () {
    var id = this.$route.params.id
    if (id && (typeof id === 'string' || id instanceof String) && id !== '') {
      var recipe = getRecipe(id)
      return {
        recipe: recipe
      }
    } else {
      return {
        recipe: {
          title: '',
          ingredients: [''],
          instructions: [{number: 1, instruction: ''}],
          tips: ['']
        }
      }
    }
  },
  methods: {
    removeIngredient: function (index) {
      console.log('rem', index)
      this.recipe.ingredients.splice(index, 1)
    },
    addIngredient: function () {
      this.recipe.ingredients.push('')
    },
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

      var jsonRecipe = JSON.stringify(this.recipe)
      console.log(jsonRecipe)
      // TODO redirect back to recipe menu
      this.$router.push({name: 'RecipeViewer', params: { userId: this.recipe.recipe_id }})
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
