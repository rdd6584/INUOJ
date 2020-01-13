import Vue from 'vue'
import VueRouter from 'vue-router'
import ff from '../func.vue'

var f = ff.methods
Vue.use(VueRouter)

var forNotUsers = function(next) {
  f.getUserValid().then(
    res => {
    if(res === null) next()
    else next('/wrongAccess')
  })
}

var forUsers = function(next) {
  f.getUserValid().then(
    res => {
    if(res === null) next('/login')
    else next()
  })
}

const routes = [
  {
    path: '/',
    component: () => import('../views/home.vue')
  },
  {
    path: '/register',
    beforeEnter: (to, from, next) => { forNotUsers(next) },
    component: () => import('../views/register.vue')
  },
  {
    path: '/login',
    beforeEnter: (to, from, next) => { forNotUsers(next) },
    component: () => import('../views/login.vue')
  },
  { // auth인증 methods만 구현
    path: '/auth',
    beforeEnter: (to, from, next) => { forNotUsers(next) },
    component: () => import('../views/auth.vue')
  },
  {
    path: '/status',
    beforeEnter: (to, from, next) => { forUsers(next) },
    component: () => import('../views/status.vue')
  },
  {
    path: '/submit/:num',
    beforeEnter: (to, from, next) => { forUsers(next) },
    component: () => import('../views/submit.vue')
  },
  {
    path: '/list',
    // beforeEnter: (to, from, next) => { forUsers(next) },
    component: () => import('../views/list.vue')
  },
  {
    path: '/problem/:num',
    // beforeEnter: (to, from, next) => { forUsers(next) },
    component: () => import('../views/problem.vue')
  },
  {
    path: '/sup/list',
    beforeEnter: (to, from, next) => { forUsers(next) },
    component: () => import('../views/sup/list.vue'),
  },
  {
    path: '/sup/detail/:ori_no',
    beforeEnter: (to, from, next) => { forUsers(next) },
    component: () => import('../views/sup/detail.vue'),
  },
  {
    path: '/status',
    name: 'test',
    component: () => import('../views/status.vue')
  },
  {
    path: '/test2',
    component: () => import('../semiViews/textEditor.vue')
  },
  {
    path: '/wrongAccess',
    component: () => import('../views/404page.vue'),
  },
  {
    path: '/*',
    component: () => import('../views/404page.vue'),
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
