<template>
  <v-form ref="uploadForm" @submit.prevent="upload" class="form">
    <v-container>
      <v-row class="justify-center">
        <v-col cols="4">
          <v-file-input
              name="file"
              show-size
              truncate-length="25"
              placeholder="选择一个图片或视频..."
              :rules="[rules.required]"
              accept="image/*,video/*"
          ></v-file-input>
        </v-col>
      </v-row>
    </v-container>

    <v-container>
      <v-row class="justify-center">
        <v-col cols="4">
          <v-btn
              color="primary"
              :type="'submit'">
            <v-icon>mdi-cloud-upload</v-icon>
            上传
          </v-btn>
        </v-col>
      </v-row>
    </v-container>
  </v-form>
</template>

<script>
import axios from "axios";

export default {
  name: "Upload",
  data(){
    return{
      rules:{required: value => !!value || '此项为必填项！'},
    }
  },
  methods: {
    upload(event) {
      // console.log(event)
      // console.log(event.target);
      // console.log(event.target['file'])
      console.log(event.target['file'].files[0])
      if (this.$refs.uploadForm.validate()){
        let formData = new FormData();
        formData.append("file",event.target['file'].files[0]);
        formData.append('x-token',localStorage.getItem('x-token'))
        axios.post(
            "/file/upload",
            formData,
            {
              headers:{
                'Content-Type':"application/form-data"
              }
            }
        ).then(resp => {
          console.log(resp)
        }).catch(error => {
          // console.log(error)
          console.log(error.response.data.message)
        })
      }
    }
  }
}
</script>

<style scoped>
.form{
  position: relative;
  top: 300px;
}
</style>