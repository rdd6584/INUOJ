import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from 'axios'
import sha256 from 'js-sha256'
import vuetify from './plugins/vuetify';

Vue.prototype.$axios = axios
Vue.prototype.$sha256 = sha256
Vue.config.productionTip = false

const f = new Vue({
  methods: {
    getUserValid() {
      return this.getAuth()
    },
    getToken() {
      return localStorage.getItem('token')
    },
    getAuth() {
      return localStorage.getItem('auth')
    },
  },
})

Vue.prototype.$f = f

new Vue({
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount('#app')
