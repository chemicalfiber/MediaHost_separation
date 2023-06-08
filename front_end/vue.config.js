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
  }
})
