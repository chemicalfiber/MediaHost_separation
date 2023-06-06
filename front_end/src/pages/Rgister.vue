<template>
  <v-form ref="registerForm" @submit="submitForm">

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
<!--TODO：异步校验用户名是否可用-->
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
                  :value="passwordConfirm"
                  v-model="passwordConfirm"
                  :error-messages="errMessage.passwordConfirm"
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

  </v-form>
</template>

<script>
export default {
  name: "register",
  data() {
    return {
      username: "",
      // nickname: "",
      password: "",
      // passwordConfirm: "",
      errMessage: {
        username: "",
        // nickname: "",
        password: "",
        // passwordConfirm: "",
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
            "/user/register",
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
          // TODO：将后端传递过来的错误信息重新渲染到该组件上
          console.log(error);
        })
      }
    }
  }
}
</script>

<style scoped>

</style>