const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  // configureWebpack: {
  //   watch: true,
  // },

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
