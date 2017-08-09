// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
// import the vue instance
import Vue from 'vue'
// import the App component
import App from './App'
// import the vue router
import VueRouter from 'vue-router'
// tell vue to use the router
// import the hello component
import Hello from './components/Hello'
// import the about component
import About from './components/About'
// import the param component
import Param from './components/Param'
// import paramdetails component
import paramdetails from './components/paramdetails'

Vue.use(VueRouter)
const routes = [
// define the root url of the application.
{ path: '/', component: Hello },
// route for the about route of the web page
{ path: '/about', component: About },
// oute for the param route of the webpage
{ path: '/param', component: Param },
// route for the paramdetails passing in params
{ path: '/Paramdetails/:id', component: paramdetails, name: 'Paramdetails' }
]

// Create the router instance and pass the `routes` option
// You can pass in additional options here, but let's
// keep it simple for now.
const router = new VueRouter({
  routes, // short for routes: routes
  mode: 'history'
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
