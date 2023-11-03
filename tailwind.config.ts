import type { Config } from 'tailwindcss';

export default {
  content: ['./src/**/*.{js,jsx,ts,tsx}'],
  theme: {
    extend: {},
  },
  darkMode: ['class', '[data-mode="dark"]'],
  plugins: [],
} satisfies Config;
