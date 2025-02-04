import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import { ManifestOptions, VitePWA } from 'vite-plugin-pwa'

// manifest поменять pwa manifest generator
const manifest: false | Partial<ManifestOptions> | undefined = {"theme_color":"#8936FF","background_color":"#2EC6FE","icons":[{"purpose":"maskable","sizes":"512x512","src":"icon512_maskable.png","type":"image/png"},{"purpose":"any","sizes":"512x512","src":"icon512_rounded.png","type":"image/png"}],"orientation":"any","display":"standalone","lang":"ru"}
export default defineConfig({
  plugins: [react(), VitePWA({registerType: 'autoUpdate', workbox: {
    globPatterns: ['**/*.{js,css,html,ico,png,svg}'],
  },
  manifest: manifest
})],
})
