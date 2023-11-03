import { AppShell } from '@mantine/core';
import { Navbar } from './Navbar.tsx';
import { Content } from './Content.tsx';
import { WorkspaceTab } from './WorkspaceTab.tsx';

const Layout = () => {
  // const [opened, { toggle }] = useDisclosure();
  return (
    <>
      <AppShell
        layout={'alt'}
        navbar={{ width: 250, breakpoint: 'sm', collapsed: { mobile: true } }}
      >
        <AppShell.Navbar>
          <Navbar />
        </AppShell.Navbar>
        <AppShell.Main>
          <WorkspaceTab />
          <Content />
        </AppShell.Main>
      </AppShell>
    </>
  );
};

export { Layout };
