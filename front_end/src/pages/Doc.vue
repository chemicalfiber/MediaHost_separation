<template>
  <div>
    <v-container>
      <v-row>
        <v-col cols="4">
          <h1>使用说明</h1>
        </v-col>
      </v-row>
    </v-container>
    <v-container>
        <v-simple-table>
          <template v-slot:default>
            <thead>
            <tr>
              <th class="text-left">
                API
              </th>
              <th class="text-left">
                请求方法
              </th>
              <th class="text-left">
                请求参数
              </th>
              <th class="text-left">
                返回数据
              </th>
              <th class="text-left">
               简介
              </th>
            </tr>
            </thead>
            <tbody>
            <tr
                v-for="item in apis"
                :key="item.path"
            >
              <td>{{ item.path }}</td>
              <td>{{item.method}}</td>
              <td>{{item.params}}</td>
              <td>{{item.return_type}}</td>
              <td>{{ item.description }}</td>
            </tr>
            </tbody>
          </template>
        </v-simple-table>
    </v-container>
  </div>
</template>

<script>
export default {
name: "Doc",
  data () {
    return {
      apis: [
        {
          path: '/file/upload',
          method:"POST",
          params:"一个form-data，请求body中至少需要包含：「userId」 ： 一个包含24个字符的用户API token，「file」 ： 一个文件对象",
          return_type:"json字符串",
          description: "上传文件到本系统。注意：本系统只接收图片和视频。",
        },
        {
          path: '/f/:id',
          method: "GET",
          params:"id：上传文件成功后，从本系统返回的文件访问ID，通常是一个62进制表示数字。",
          return_type:"文件",
          description: "使用本系统返回的文件ID取得一个文件，使用REST ful请求，将文件访问ID放在请求路径中。即可得到对应的文件，你可以将文件访问链接放在任何需要使用该文件的地方，例如html img标签的src属性中。",
        }
      ],
    }
  },
}
</script>

<style scoped>

</style>