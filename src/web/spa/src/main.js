// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
// import the vue instance
import Vue from 'vue'
// import the App component
import App from './App'
// import the vue router
import VueRouter from 'vue-router'
// import RecipeViewBrowser component
import RecipeViewBrowser from './components/recipe_view_browser'
// import RecipeViewer component
import RecipeViewer from './components/recipe_viewer'
// import RecipeEditorList component
import RecipeEditorList from './components/recipe_editor_list'
// import RecipeEditor component
import RecipeEditor from './components/recipe_editor'

import 'vue-awesome/icons'
import Icon from 'vue-awesome/components/Icon'

// tell vue to use the router
Vue.component('icon', Icon)

Vue.use(VueRouter)
const routes = [
// route for the RecipeEditor passing in params
{path: '/', component: RecipeEditorList, name: 'RecipeEditorList'},
// route for the RecipeEditor passing in params
{path: '/recipe_editor', component: RecipeEditor, name: 'RecipeEditorNew'},
// route for the RecipeEditor passing in params
{path: '/recipe_editor/:id', component: RecipeEditor, name: 'RecipeEditor'},
// route for the RecipeViewBrowser passing in params
{ path: '/device/recipe_view_browser', component: RecipeViewBrowser, name: 'RecipeViewBrowser' },
// route for the RecipeViewer passing in params
{ path: '/device/recipe_viewer/:id', component: RecipeViewer, name: 'RecipeViewer' }
]

// Create the router instance and pass the `routes` option
// You can pass in additional options here, but let's
// keep it simple for now.
const router = new VueRouter({
  routes, // short for routes: routes
  mode: 'history' /* 'hash' */
})

// place the router guard
router.beforeEach((to, from, next) => {
  // check if the path user is going to is our param path
  if (to.path === '/param') {
    // check if the user item is already set
    if (localStorage.getItem('user') === undefined) {
      // prompt for username
      var user = prompt('please enter your username')
      // prompt for password
      var pass = prompt('please enter your password')
      // check if th username and password given equals our preset details
      if (user === 'username' && pass === 'password') {
        // set the user item
        localStorage.setItem('user', user)
        // move to the route
        next()
      } else {
        // alert the username and pass is wrong
        alert('Wrong username and password, you do not have permission to access that route')
        // return, do not move to the route
        return
      }
    }
  }
  next()
})

// instatinat the vue instance
new Vue({
// define the selector for the root component
  el: '#app',
  // pass the template to the root component
  template: '<App/>',
  // declare components that the root component can access
  components: { App },
  // pass in the router to the Vue instance
  router
}).$mount('#app')// mount the router on the app
