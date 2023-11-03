import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App.tsx';
import '@mantine/notifications/styles.css';
import '@mantine/spotlight/styles.css';
import { MantineProvider } from '@mantine/core';
import theme from './theme.ts';
import '@mantine/nprogress/styles.css';
import { Notifications } from '@mantine/notifications';
import { GlobalSpotlight } from './component/GlobalSpotlight.tsx';
import { NavigationProgress } from '@mantine/nprogress';
import { ModalsProvider } from '@mantine/modals';
ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <MantineProvider theme={theme} defaultColorScheme={'auto'}>
      <Notifications position={'top-right'} limit={3} />
      <NavigationProgress />
      <GlobalSpotlight />
      <ModalsProvider>
        <App />
      </ModalsProvider>
    </MantineProvider>
  </React.StrictMode>,
);
