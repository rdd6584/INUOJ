import Vue from 'vue'
import VueRouter from 'vue-router'
import ff from '../func.vue'

var f = ff.methods
Vue.use(VueRouter)

const forBdmin = (to, from, next) => {
  f.getUserValid().then(
    res => {
    if (res === null) next('/login')
    else if (f.admin == '' || parseInt(f.admin) < 1) next('/wrongAccess')
    else next()
  })
}

const forUsers = (to, from, next) => {
  f.getUserValid().then(
    res => {
    if (res === null) next('/login')
    else next()
  })
}

const forNotUsers = (to, from, next) => {
  f.getUserValid().then(
    res => {
    if(res === null) next()
    else next('/wrongAccess')
  })
}

const routes = [
  { path: '/', component: () => import('../views/home.vue') },

  { path: '/register', beforeEnter: forNotUsers, component: () => import('../views/register.vue') },
  { path: '/login', beforeEnter: forNotUsers, component: () => import('../views/login.vue') },
  { path: '/auth', beforeEnter: forNotUsers, component: () => import('../views/auth.vue') },

  { path: '/profile/:id', beforeEnter: forUsers, component: () => import('../views/profile.vue') },
  { path: '/modify/:id', beforeEnter: forUsers, component: () => import('../views/modifyPass.vue') },
  { path: '/status', beforeEnter: forUsers, component: () => import('../views/status.vue') },
  { path: '/submit/:num', beforeEnter: forUsers, component: () => import('../views/submit.vue') },
  { path: '/list', beforeEnter: forUsers, component: () => import('../views/list.vue') },
  { path: '/problem/:num', beforeEnter: forUsers, component: () => import('../views/problem.vue') },
  { path: '/source/:subm_no', beforeEnter: forUsers, component: () => import('../views/source.vue') },

  { path: '/writepost', component: () => import('../views/writePost.vue') },
  { path: '/post/:num', component: () => import('../views/post.vue') },
  { path: '/board', component: () => import('../views/board.vue') },

  { path: '/sup/list', beforeEnter: forBdmin, component: () => import('../views/sup/list.vue') },
  { path: '/sup/detail/:ori_no', beforeEnter: forBdmin, component: () => import('../views/sup/detail.vue') },

  { path: '/wrongAccess', component: () => import('../views/404page.vue') },
  { path: '/*', component: () => import('../views/404page.vue') },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
