<template>
  <v-app>
<!--    左侧抽屉-->
    <Drawer></Drawer>
    <v-main>
      <v-container>
              <router-view/>

<!--        <Login></Login>-->
      </v-container>
<!--      <HelloWorld/>-->

    </v-main>
  </v-app>

</template>

<script>
// import HelloWorld from './components/HelloWorld';
import axios from "axios";
import Drawer from "@/components/fragement/Drawer";
import Login from "@/pages/Login";

export default {
  name: 'App',

  components: {
    Login,
    // HelloWorld,
    Drawer,
  },

  data: () => ({
    //
  }),
  mounted() {
    axios.get("/api/checkToken",{params:{"x-token":localStorage.getItem("x-token")}}).catch(error => {
      this.$router.push("/login")
    })
  }
};
</script>
