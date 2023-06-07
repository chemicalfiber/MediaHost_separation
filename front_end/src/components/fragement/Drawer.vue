<template>
  <!--此处的app属性为了配合App.vue中的v-app进行渲染，否则会跑版-->
<!--  expand-on-hover为"悬浮时展开"-->
    <v-navigation-drawer
        app
        permanent
    >
      <v-list>

        <v-list-item link v-if="isLogin" @click="logout">
          <v-list-item-content>
            <v-list-item-title class="text-h6">
              {{ nickname }}
            </v-list-item-title>
            <v-list-item-subtitle>{{ username }}</v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>

        <v-list-item link v-else to="/login">
          <v-list-item-content>
            <v-list-item-title class="text-h6">
              登录
            </v-list-item-title>
            <v-list-item-subtitle>点此登录</v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
      </v-list>

      <v-divider></v-divider>

      <v-list
          nav
          dense
      >
        <v-list-item link to="/">
          <v-list-item-icon>
            <v-icon color="blue darken-2">mdi-cloud-upload</v-icon>
          </v-list-item-icon>
          <v-list-item-title>传文件</v-list-item-title>
        </v-list-item>
        <v-list-item link>
          <v-list-item-icon>
            <v-icon color="brown">mdi-folder-play</v-icon>
          </v-list-item-icon>
          <v-list-item-title>我的媒体</v-list-item-title>
        </v-list-item>
        <v-list-item link>
          <v-list-item-icon>
            <v-icon color="green">mdi-account-circle</v-icon>
          </v-list-item-icon>
          <v-list-item-title>账户</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

</template>

<script>
export default {
  name: "Drawer",
  data() {
    return {
      isLogin: false,
      username:localStorage.getItem('username'),
      nickname:localStorage.getItem('nickname'),
    }
  },
  mounted() {
    // TODO：使用vuex对isLogin进行管理
    this.isLogin = localStorage.getItem("x-token");
  },
  methods:{
    logout(){
      localStorage.clear();
      location.reload();
      // this.$router.push("/login");
    }
  }
}
</script>

<style scoped>

</style>