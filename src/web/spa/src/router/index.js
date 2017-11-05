import Vue from 'vue'
import Router from 'vue-router'
// import RecipeEditorList component
import RecipeEditorList from '@/components/recipe_editor_list'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'RecipeEditorList',
      component: RecipeEditorList
    }
  ]
})
