import react from '@vitejs/plugin-react';
import { visualizer } from 'rollup-plugin-visualizer';
import { defineConfig, splitVendorChunkPlugin, type PluginOption } from 'vite';
import { compression } from 'vite-plugin-compression2';
import TurboConsole from 'vite-plugin-turbo-console';
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    react(),
    visualizer() as PluginOption,
    TurboConsole(),
    splitVendorChunkPlugin(),
    compression(),
  ],
  build: {},
  envPrefix: ['VITE_', 'TAURI_'],
});
