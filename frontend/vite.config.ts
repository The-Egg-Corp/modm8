import { resolve } from 'path'
import { defineConfig } from 'vite'

import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve:{
    alias: {
      '@assets': resolve(__dirname, './src/assets'),
      '@components': resolve(__dirname, './src/components'),
      '@store': resolve(__dirname, './src/store'),
      '@types': resolve(__dirname, './src/types'),
      '@i18n': resolve(__dirname, './src/i18n'),
      '@backend': resolve(__dirname, './wailsjs/go'),
      '@runtime': resolve(__dirname, './wailsjs/runtime/runtime')
    }
  }
})
