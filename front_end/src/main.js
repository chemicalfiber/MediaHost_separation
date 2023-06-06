import Vue from 'vue'
import App from './App.vue'
// 引入路由插件
import VueRouter from "vue-router";
import router from "@/router";
// 引入Vuex插件（实际上在store的js中引入了）
import store from "@/store";  // 引入store
import vuetify from './plugins/vuetify'

import axios from 'axios'
Vue.prototype.$axios = axios
axios.defaults.baseURL = 'http://127.0.0.1:8823'  // 设置axios请求基本路径
// Vue.config.productionTip = false

Vue.use(VueRouter)

new Vue({
  vuetify,
  router, // 使用路由
  store,  // 使用Vuex的store对象
  render: h => h(App)
}).$mount('#app')