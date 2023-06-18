const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: [
    'vuetify'
  ],
  chainWebpack: config =>{
    config.plugin('html')
        .tap(args => {
          args[0].title = "图床——传你想传";
          return args;
        })
  },
    devServer:{
        host:"0.0.0.0",
        port:8080,
        allowedHosts:[
            ".dynv6.net",
            "10.35.48.32",
            "192.168.2.10",
            "127.0.0.1",
        ]
    }
})
