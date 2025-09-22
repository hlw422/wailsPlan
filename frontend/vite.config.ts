import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import Components from 'unplugin-vue-components/vite'
import path from 'path'  // 引入path模块

// 进行naive ui 的自动引入只需要下面这句
import { NaiveUiResolver } from 'unplugin-vue-components/resolvers'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    Components({
      resolvers: [
        // naive ui 的自动引入，只需要这一句
        NaiveUiResolver()
        // AntDesign vue 的自动引入，只需要这一句
        // AntDesignVueResolver()
      ]
    }),
  ],
  // 增加路径别名配置
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src')  // 将@指向src目录
    }
  },
  
})