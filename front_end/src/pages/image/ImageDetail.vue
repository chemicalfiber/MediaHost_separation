<template>
  <div>
    <v-container>
      <h1>详情</h1>
      <v-card
          max-height="800"
          max-width="1270"
      >
        <v-img :src="fileInfo.file_link" max-width="640px" max-height="480px" alt="图片预览"/>
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

        <v-row >
          <v-col cols="6">
            <v-btn color="error" @click="deleteThis(fileInfo.info._id)">删除这个文件</v-btn>
          </v-col>
        </v-row>

      </v-form>

    </v-container>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "ImageDetail",
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
        this.$router.push("/images")
      }).catch(error => console.log(error))
    }
  }
}
</script>

<style scoped>

</style>