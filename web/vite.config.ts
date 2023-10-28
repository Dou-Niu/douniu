import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
// 配置路径别名
import { resolve } from 'path'
const pathResolve=(dir: string) =>{
  return resolve(__dirname,".",dir)
}
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve:{
    alias:{
      "@":pathResolve("src")
    }
  }
})
