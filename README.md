# 简介
MediaHost图床，一个存储图片与视频的图床程序。
- 前端：VUE(v2.6) + [Vuetify(v2)](https://v2.vuetifyjs.com/zh-Hans/)
- 后端：Golang(v1.16.4) + Gin(v1.7.4)
- 数据库：MongoDB(v4.2.24)
### 兄弟项目
SpringBoot(v2.7.6)版后端实现在这里：[https://github.com/chemicalfiber/media_host_springboot_backend](https://github.com/chemicalfiber/media_host_springboot_backend)

# 部署运行
首先克隆此存储库。
## 前端部署
下载并安装node(v18.15.*)。

### 正确配置下列文件中的属性：
- `front_end/vue.config.js`
```javascript
devServer:{
        host:"0.0.0.0", // 想要服务的地址
        port:8080,
        allowedHosts:[
            ".dynv6.net",   // 允许的地址，请务必填写你的域名/IP，避免无法访问
            "10.35.48.32",
            "192.168.2.10",
            "127.0.0.1",
        ]
    }
```

- `front_end/src/main.js`

```javascript
// 设置axios请求基本路径，请将它指向你后端的地址
axios.defaults.baseURL = 'http://127.0.0.1:8823'
```


### 配置完成后，运行项目：

```shell
cd front_end # 进入front_end文件夹
npm install # 拉取npm依赖
npm run serve
```

## 后端部署
下载并安装Go SDK(v1.16.4)。

### 正确配置下列文件中的属性：
- `back_end/config.json`

配置完成后，务必删除JSON中的注释，因为JSON不支持注释。
```json
{
  "port": "8823", // 后端工作的端口
  "address": "0.0.0.0", // 想要服务的地址
  "self_domain": "http://10.35.48.32", // 本图床运行的域名，在返回文件访问链接的时候需要使用到
  "file_storage_path": "/Users/chemicalfiber",  // 已废弃的配置项
  "mongodb_config": { // 数据库相关配置
    "host": "192.168.2.6",
    "port": "27017",
    "database": "media_host"
  }
}
```

### 配置完成后，运行项目：
```shell
cd back_end
go run .
```
