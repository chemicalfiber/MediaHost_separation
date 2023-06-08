import Vue from 'vue'
import App from './App.vue'
// 引入路由插件
import VueRouter from "vue-router";
import router from "@/router";

import vuetify from './plugins/vuetify'

import axios from 'axios'
Vue.prototype.$axios = axios
axios.defaults.baseURL = 'http://10.35.48.32:8823'  // 设置axios请求基本路径
// Vue.config.productionTip = false

Vue.use(VueRouter)

new Vue({
  vuetify,
  router, // 使用路由
  render: h => h(App)
}).$mount('#app')
