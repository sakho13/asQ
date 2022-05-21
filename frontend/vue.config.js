const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  configureWebpack: {
    watchOptions: {
      aggregateTimeout: 300,
      poll: 1000,
    }
  },
  devServer: {
    proxy: "http://localhost:8890"
  },

  transpileDependencies: true,

  pluginOptions: {
    quasar: {
      importStrategy: 'kebab',
      rtlSupport: false
    }
  }
})
