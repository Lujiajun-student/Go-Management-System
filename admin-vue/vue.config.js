const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  lintOnSave: false, // 关闭校验
  productionSourceMap: false, // 选择是否生成source map
  publicPath: '/', // 部署应用时的基本url
  outputDir: 'dist', // build输出的文件目录
  assetsDir: 'assets', // 放置静态文件夹目录
  devServer: {
    port: 8081,
    host: '0.0.0.0', // 运行域名
    https: false, // 不需要https
    open: false, // 是否直接打开浏览器
    proxy: {
      "/api": {
        target:"http://localhost:8080", // 配置后端服务地址
        changeOrigin: true,
      }
    },
    client: {
      overlay: false // 关闭全屏报错
    }
  },
})
