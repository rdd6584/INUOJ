import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

// route guard 작성하기. beforeEnter
const routes = [
  {
    path: '/',
    name: 'home',
    component: () => import('../views/home.vue')
  },
  {
    path: '/register',
    name: 'register',
    component: () => import('../views/register.vue')
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('../views/login.vue')
  },
  {
    path: '/users/:userId',
    name: 'users',
    component: () => import('../views/users.vue'),
  },
  { // auth인증 methods만 구현
    path: '/auth',
    name: 'auth',
    component: () => import('../views/auth.vue')
  },
  {
    path: '/status',
    name: 'status',
    component: () => import('../views/status.vue')
  },
  {
    path: '/test',
    name: 'test',
    component: () => import('../views/status.vue')
  },
  {
    path: '/*',
    component: () => import('../views/404page.vue')
  },
  {
    path: '/'
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
