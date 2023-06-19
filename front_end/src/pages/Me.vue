<template>
  <div>
    <v-container>
      <v-row>
        <v-col cols="4">
          <h1>个人中心</h1>
          <v-btn color="secondary" @click="logout">登出</v-btn>
        </v-col>
      </v-row>
    </v-container>
    <v-form ref="userForm" @submit="submitForm" class="form">

      <v-container>
        <v-row class="offset-4" style="color: red">
          {{formErrMessage}}
        </v-row>
      </v-container>
      <v-container>
        <v-row class="justify-center">
          <v-col
              cols="4">
            <v-text-field
                prepend-icon="mdi-account-circle"
                hide-details="auto"
                label="用户API token"
                :value="user_id"
                disabled
                :error-messages="'这是您在其他系统中调用本系统接口的API token，千万不要泄漏！'"
            >
            </v-text-field>
          </v-col>
        </v-row>
      </v-container>

      <v-container>
        <v-row class="justify-center">
          <v-col
              cols="4">
            <v-text-field
                prepend-icon="mdi-account-circle-outline"
                hide-details="auto"
                label="昵称"
                :value="nickname"
                v-model="nickname"
                :error-messages="errMessage.nickname"
                :rules="[rules.max,rules.min]"
            >

            </v-text-field>
          </v-col>
        </v-row>
      </v-container>

      <v-container>
        <v-row class="justify-center">
          <v-col
              cols="4">
            <v-text-field
                prepend-icon="mdi-lock"
                hide-details="auto"
                label="密码"
                :type="'password'"
                :value="password"
                v-model="password"
                :error-messages="errMessage.password"
                :rules="[rules.passwordMin,rules.passwordMax]"
            >

            </v-text-field>
          </v-col>
        </v-row>
      </v-container>

      <v-container>
        <v-row class="justify-center">
          <v-col
              cols="4">
            <v-text-field
                prepend-icon="mdi-lock-outline"
                hide-details="auto"
                label="确认密码"
                :type="'password'"
                :value="confirmPassword"
                v-model="confirmPassword"
                :error-messages="errMessage.confirmPassword"
                :rules="[rules.max,rules.min]"
                @keyup.enter="submitForm"
            >

            </v-text-field>
          </v-col>
        </v-row>
      </v-container>


      <v-container>
        <v-row class="justify-center">
          <v-col
              cols="4">
            <v-btn
                color="secondary"
                class="mr-4"
                @click="submitForm"
            >
              <v-icon left>mdi-account-arrow-up</v-icon>
              更新个人信息
            </v-btn>
            <v-btn to="/doc"><v-icon left>mdi-progress-question</v-icon>如何使用上述的API token？</v-btn>
          </v-col>
        </v-row>

      </v-container>

    </v-form>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "Me",
  data() {
    return {
      user_id:"",
      username: "",
      nickname: "",
      password: "",
      confirmPassword: "",
      formErrMessage:"",
      errMessage: {
        username: "",
        nickname: "",
        password: "",
        confirmPassword: "",
      },
      // 表单校验规则
      rules: {
        required: value => !!value || '此项为必填项！',
        max: value => value.length <= 20 || '最多只能20个字符',
        min: value => value.length >= 3 || '至少需要3个字符',
        passwordMax: value => value.length <= 16 || '最多只能16个字符',
        passwordMin: value => value.length >= 8 || '至少需要8个字符',
        email: value => {
          const pattern = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
          return pattern.test(value) || 'Invalid e-mail.'
        }
      }
    }
  },
  methods: {
    logout() {
      localStorage.clear();
      location.reload();
      // this.$router.push("/login");
    },
    submitForm(){
      // console.log(this.$refs.userForm.validate());
      if (this.$refs.userForm.validate()){
        let formData = new FormData();
        formData.append("user_id",this.user_id);
        formData.append("username",this.username)
        formData.append("password",this.password)
        formData.append("nickname",this.nickname)
        formData.append("confirmPassword",this.confirmPassword)
        axios.post(
            "/user/update",
            formData,{
              headers:{
                'Content-Type':"application/form-data"
              }
            }
        ).then(resp => {
          console.log(resp)
          if(resp.data.message === "OK"){
            this.$router.push("/login")  // 登录成功，走你～
          }
        }).catch(error=>{
          console.log(error);
          this.formErrMessage = error.response.data.message;
        })
      }

    }
  },
  mounted() {
    this.username = localStorage.getItem("username");
    this.nickname = localStorage.getItem("nickname");
    this.user_id = localStorage.getItem("user_id");
  }
}
</script>

<style scoped>

</style>