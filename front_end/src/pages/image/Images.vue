<template>
  <div>
    <v-container>
      <v-row>
        <v-col cols="4">
          <h1>我的图片</h1>
        </v-col>
      </v-row>
      <v-row>
        <v-col
            cols="4"
            v-for="image in imageList"
            :key="image._id"
        >
          <v-card max-width="600" max-height="340" :to='{path:"/images/"+image._id}'>
            <v-img style="display: block" :src="image.link" :alt="image.name" max-width="500" max-height="250"/>
            <div>{{image.name}}</div>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "Images",
  data() {
    return {
      imageList: undefined,
    }
  },
  beforeMount() {
    axios.get(
        "/user/img/" + localStorage.getItem("user_id")
    ).then(response => {
      console.log(response)
      this.imageList = response.data.data;
    }).catch(error => {
      console.log(error)
    })
  }
}
</script>

<style scoped>

</style>