import Vue from 'vue'
import VueRouter from 'vue-router'
import UserLogin from '../components/UserLogin.vue'
import UserRegister from '../components/UserRegister.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: UserLogin
  },
  {
    path: '/register',
    name: 'Register',
    component: UserRegister
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
