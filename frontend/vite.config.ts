import react from '@vitejs/plugin-react';
import { defineConfig } from 'vite';
import { ManifestOptions, VitePWA } from 'vite-plugin-pwa';

// manifest поменять pwa manifest generator
const manifest: false | Partial<ManifestOptions> | undefined = {
  theme_color: '#8936FF',
  background_color: '#2EC6FE',
  icons: [
    { purpose: 'maskable', sizes: '512x512', src: 'LOGO.png', type: 'image/png' },
    { purpose: 'any', sizes: '512x512', src: 'LOGO.png', type: 'image/png' },
  ],
  orientation: 'any',
  display: 'standalone',
  lang: 'ru',
};
export default defineConfig({
  plugins: [
    react(),
    VitePWA({
      strategies: 'injectManifest',
      srcDir: 'src',
      filename: 'service-worker.ts',
      registerType: 'autoUpdate',
      workbox: {
        navigateFallback: '/index.html',
      },
      injectManifest: {
        globPatterns: ['**/*.{js,css,html,ico,png,svg}'],
      },
      devOptions: {
        enabled: true,
        type: 'module',
      },
      injectRegister: 'auto',
      manifest: manifest,
    }),
  ],
});
