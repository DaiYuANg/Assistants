import { createTheme, rem } from '@mantine/core';
import '@fontsource/jetbrains-mono/700-italic.css';
import '@fontsource/jetbrains-mono/700.css'; // Specify weight
const theme = createTheme({
  fontFamily: 'Jetbrains Mono',
  fontSmoothing: true,
  shadows: {
    md: '1px 1px 3px rgba(0, 0, 0, .25)',
    xl: '5px 5px 3px rgba(0, 0, 0, .25)',
  },

  headings: {
    fontFamily: 'Roboto, sans-serif',
    sizes: {
      h1: { fontSize: rem(36) },
    },
  },
  fontSizes: {},
});

export default theme;
