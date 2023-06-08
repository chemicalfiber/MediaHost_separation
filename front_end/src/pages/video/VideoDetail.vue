<template>
<div>
  <v-container>
    <h1>详情</h1>
    <v-card
        max-height="800"
        max-width="1270"
    >
      <video :src="fileInfo.file_link" controls width="850" height="480"/>
    </v-card>
    <v-form>

      <v-row >
        <v-col cols="6">
          上传者：{{fileInfo.upload_user}}
        </v-col>
      </v-row>

      <v-row >
        <v-col cols="6">
          文件名：{{fileInfo.info.name}}
        </v-col>
      </v-row>

      <v-row >
        <v-col cols="6">
          文件类型：{{fileInfo.info.type}}
        </v-col>
      </v-row>

      <v-row >
        <v-col cols="6">
          上传时间：{{fileInfo.info.upload_date}}
        </v-col>
      </v-row>

      <v-row >
        <v-col cols="6">
          文件访问链接：{{fileInfo.file_link}}
        </v-col>
      </v-row>
    </v-form>
    <v-row >
      <v-col cols="6">
        <v-btn color="error" @click="deleteThis(fileInfo.info._id)">删除这个文件</v-btn>
      </v-col>
    </v-row>
  </v-container>
</div>
</template>

<script>
import axios from "axios";

export default {
  name: "VideoDetail",
  data() {
    return {
      fileInfo: undefined,
    }
  },
  beforeMount() {
    let params = this.$route.params;
    console.log(params.id)
    axios.get(
        "/file/" + params.id
    ).then(response => {
      console.log(response)
      this.fileInfo = response.data.data;
    }).catch(error => {
      console.log(error)
    })
  },
  methods:{
    deleteThis(id){
      console.log(id)
      axios.get("/file/delete/"+id).then(response => {
        console.log(response);
        this.$router.push("/videos")
      }).catch(error => console.log(error))
    }
  }
}
</script>

<style scoped>

</style>