<template>
  <v-form ref="loginForm" @submit="submitForm" class="form">
  <v-container>
    <v-row class="justify-center">
      <h1>登录</h1>
    </v-row>
  </v-container>
    <v-container>
      <v-row class="justify-center">
        <v-col
            cols="4">
          <v-text-field
              prepend-icon="mdi-account-circle"
              hide-details="auto"
              label="用户名"
              :value="username"
              v-model="username"
              :error-messages="errMessage.username"
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
              @keyup.enter="submitForm"
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
          <v-btn
              color="secondary"
              class="mr-4"
              @click="submitForm"
          >
            登录
          </v-btn>
        </v-col>
      </v-row>
    </v-container>

    <v-container>
      <v-row class="offset-4">
        没有账户？<router-link to="/register">注册一个</router-link>
      </v-row>
    </v-container>
  </v-form>

</template>

<script>
import {Base64} from 'js-base64';
export default {
  name: "Login",
  mounted() {
    localStorage.clear();
  },
  data() {
    return {
      username: "",
      password: "",
      errMessage: {
        username: "",
        password: "",
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
    submitForm() {
      // console.log(this.$refs.loginForm.validate());
      if (this.$refs.loginForm.validate()){
        let formData = new FormData();
        formData.append("username",this.username)
        formData.append("password",this.password)
        this.$axios.post(
            "/user/login",
            formData,{
              headers:{
                'Content-Type':"application/form-data"
              }
            }
        ).then(resp => {
          // console.log(resp)
          if(resp.data.message === "OK"){
            const token = resp.data.data.token;
            localStorage.setItem("x-token",token)  // 将token存好
            // console.log(token.indexOf('.'));
            // console.log(token.lastIndexOf('.'));
            // 解析token中的载体
            const payload = token.substring(token.indexOf('.')+1, token.lastIndexOf('.'));
            console.log(payload);
            const payloadDecodeStr = Base64.decode(payload);
            let payloadObj = JSON.parse(payloadDecodeStr);

            console.log(payloadObj)

            // 存放载体的关键信息
            localStorage.setItem('nickname',payloadObj.nickname);
            localStorage.setItem('username',payloadObj.username);
            localStorage.setItem('user_id',payloadObj.user_id);

            location.href = '/';  // 登录成功，刷新页面让显示用户信息相关组件重新渲染
          }
        }).catch(error=>{
          console.log(error);
          // console.log(this);
          this.errMessage.username = error.response.data.message;
          this.errMessage.password = error.response.data.message;
          this.password = '';
        })
      }
    }
  }
}
</script>

<style scoped>
.form{
  position: relative;
  top: 200px;
}
</style>