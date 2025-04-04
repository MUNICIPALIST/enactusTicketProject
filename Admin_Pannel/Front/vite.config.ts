import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    host: '0.0.0.0',
    port: 5173,
    watch: {
      usePolling: true
    },
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, '')
      }
    },
    allowedHosts: [
      'localhost',
      '127.0.0.1',
      'vencera.tech',
      '.vencera.tech' // Разрешает все поддомены
    ]
  },
  preview: {
    host: '0.0.0.0',
    port: 5173
  },
  base: '/'
})
