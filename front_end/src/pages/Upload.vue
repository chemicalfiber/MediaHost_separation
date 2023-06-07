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

    <v-container>
      <v-row class="justify-center">
        <v-col cols="4">
<!--          <div class="text-center mb-4">-->
<!--            <v-btn-->
<!--                color="primary"-->
<!--                @click="alert = !alert"-->
<!--            >-->
<!--              Toggle-->
<!--            </v-btn>-->
<!--          </div>-->
          <v-alert
              :value="alert"
              type="success"
              dark
              border="top"
              icon="check-circle"
              transition="scale-transition"
          >
            上传完成<br>
            文件访问链接：{{fileLink}}
          </v-alert>
        </v-col>
      </v-row>
    </v-container>

  </v-form>
</template>

<script>
import axios from "axios";

export default {
  name: "Upload",
  data() {
    return {
      rules: {required: value => !!value || '此项为必填项！'},
      fileLink:"",
      // fileInfoKey: "",
      // fileKey: "",
      alert: false,
    }
  },
  methods: {
    upload(event) {
      // console.log(event)
      // console.log(event.target);
      // console.log(event.target['file'])
      this.alert = false;
      console.log(event.target['file'].files[0])
      if (this.$refs.uploadForm.validate()) {
        let formData = new FormData();
        formData.append("file", event.target['file'].files[0]);
        formData.append('x-token', localStorage.getItem('x-token'))  // 传递token到后端，由后端解析token中的用户名
        axios.post(
            "/file/upload",
            formData,
            {
              headers: {
                'Content-Type': "application/form-data"
              }
            }
        ).then(resp => {
          console.log(resp)
          this.alert = true;
          this.fileLink = resp.data.data.fileLink;
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
.form {
  position: relative;
  top: 300px;
}
</style>