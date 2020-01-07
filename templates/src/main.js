import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from 'axios'
import sha256 from 'js-sha256'
import vuetify from './plugins/vuetify'
import f from './func.vue'

Vue.prototype.$axios = axios
Vue.prototype.$sha256 = sha256
Vue.config.productionTip = false

Vue.prototype.$f = f.methods

new Vue({
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount('#app')
