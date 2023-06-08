<template>
  <div>
    <v-container>
      <v-row>
        <v-col cols="4">
          <h1>我的视频</h1>
        </v-col>
      </v-row>
      <v-simple-table>
        <template v-slot:default>
          <thead>
          <tr>
            <th class="text-left">
              视频名
            </th>
            <th class="text-left">
              上传日期
            </th>
            <th>
              操作
            </th>
          </tr>
          </thead>
          <tbody>
          <tr
              v-for="video in videoList"
              :key="video._id"
          >
            <td>{{ video.name }}</td>
            <td>{{ video.upload_date }}</td>
            <td>
              <v-btn :to='{path:"/videos/"+video._id}' color="primary">查看详情</v-btn>
            </td>
          </tr>
          </tbody>
        </template>
      </v-simple-table>
    </v-container>
  </div>
</template>

<script>

import axios from "axios";

export default {
  name: "Video",
  data(){
    return {
      videoList : undefined,
    }
  },
  beforeMount() {
    axios.get(
        "/user/video/" + localStorage.getItem("user_id")
    ).then(response => {
      console.log(response)
      this.videoList = response.data.data;
    }).catch(error => {
      console.log(error)
    })
  }
}
</script>

<style scoped>

</style>