import { AppShell, Divider } from '@mantine/core';
import { Navbar } from './Navbar.tsx';
import { Content } from './Content.tsx';
import { WorkspaceTab } from './WorkspaceTab.tsx';
import { SearchInContent } from './SearchInContent.tsx';

const defaultWidth: number = 250;
const Layout = () => {
  return (
    <>
      <AppShell
        withBorder={false}
        layout={'alt'}
        navbar={{
          width: defaultWidth,
          breakpoint: 'sm',
          collapsed: { mobile: true },
        }}
      >
        <AppShell.Navbar>
          <Navbar />
          <Divider
            orientation={'vertical'}
            className={[
              'absolute',
              'right-0',
              'w-1',
              'top-0',
              'cursor-col-resize',
              'select-none',
              'h-full-view',
            ].join(' ')}
          />
        </AppShell.Navbar>
        <AppShell.Main>
          <SearchInContent />
          <WorkspaceTab />
          <Content />
        </AppShell.Main>
      </AppShell>
    </>
  );
};

export { Layout };
