const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  // configureWebpack: {
  //   watch: true,
  // },
  devServer: {
    proxy: "http://localhost:8890"
  },

  transpileDependencies: [
    'quasar'
  ],

  pluginOptions: {
    quasar: {
      importStrategy: 'kebab',
      rtlSupport: false
    }
  }
})
