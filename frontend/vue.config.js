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
    vuetify: {
			// https://github.com/vuetifyjs/vuetify-loader/tree/next/packages/vuetify-loader
		}
  }
})
