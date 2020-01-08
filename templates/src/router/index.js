import Vue from 'vue'
import VueRouter from 'vue-router'
import ff from '../func.vue'

var f = ff.methods
Vue.use(VueRouter)

function forNotUsers(next) {
  f.getUserValid().then(
    res => {
    if(res === null) next()
    else next('/wrongAccess')
  })
}

function forUsers(next) {
  f.getUserValid().then(
    res => {
    if(res === null) next('/login')
    else next()
  })
}

const routes = [
  {
    path: '/',
    name: 'home',
    component: () => import('../views/home.vue')
  },
  {
    path: '/register',
    name: 'register',
    beforeEnter: (to, from, next) => { forNotUsers(next) },
    component: () => import('../views/register.vue')
  },
  {
    path: '/login',
    name: 'login',
    beforeEnter: (to, from, next) => { forNotUsers(next) },
    component: () => import('../views/login.vue')
  },
  { // auth인증 methods만 구현
    path: '/auth',
    name: 'auth',
    beforeEnter: (to, from, next) => { forNotUsers(next) },
    component: () => import('../views/auth.vue')
  },
  {
    path: '/status',
    name: 'status',
    beforeEnter: (to, from, next) => { forUsers(next) },
    component: () => import('../views/status.vue')
  },
  {
    path: '/test',
    name: 'test',
    component: () => import('../views/createProblem.vue')
  },
  {
    path: '/test2',
    name: 'test2',
    component: () => import('../semiViews/textEditor.vue')
  },
  {
    path: '/wrongAccess',
    name: 'wrongAccess',
    component: () => import('../views/404page.vue'),
  },
  {
    path: '/*',
    redirect: '/wrongAccess',
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
