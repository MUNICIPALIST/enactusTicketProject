import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  base: '/aidana/',
  server: {
    host: '0.0.0.0',
    port: 5173,
    watch: {
      usePolling: true
    },
    proxy: {
      '/api': {
        target: 'http://172.17.0.1:8080',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, '')
      }
    },
    cors: true,
    allowedHosts: [
      'localhost',
      '127.0.0.1',
      'vencera.tech',
      '.vencera.tech',
      '176.123.178.135'
    ]
  },
  preview: {
    host: '0.0.0.0',
    port: 5173
  },
  build: {
    outDir: 'dist',
    assetsDir: 'assets',
    sourcemap: false,
    manifest: true,
    rollupOptions: {
      output: {
        manualChunks: undefined
      }
    }
  }
})
