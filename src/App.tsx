import '@mantine/core/styles.css';
import './App.scss';
import { Layout } from './component';
import { MouseEvent } from 'react';
import { useWindowEvent } from '@mantine/hooks';
import './i18n';

function App() {
  useWindowEvent('contextmenu', (event) => {
    event.preventDefault();
  });
  const onContextmenu = (e: MouseEvent<HTMLDivElement>) => {
    return () => {
      e.preventDefault();
    };
  };
  return (
    <>
      <div onContextMenu={onContextmenu}>
        <Layout />
      </div>
    </>
  );
}

export default App;
