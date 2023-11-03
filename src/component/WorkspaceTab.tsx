import { rem, Tabs } from '@mantine/core';
import {
  IconMessageCircle,
  IconPhoto,
  IconSettings,
} from '@tabler/icons-react';
import classes from './Content.module.css';

const WorkspaceTab = () => {
  const iconStyle = { width: rem(12), height: rem(12) };
  return (
    <>
      <Tabs
        className={['absolute', classes.content, 'opacity-0'].join(' ')}
        defaultValue="gallery"
        pt={'xs'}
      >
        <Tabs.List>
          <Tabs.Tab
            value="gallery"
            leftSection={<IconPhoto style={iconStyle} />}
          >
            Gallery
          </Tabs.Tab>
          <Tabs.Tab
            value="messages"
            leftSection={<IconMessageCircle style={iconStyle} />}
          >
            Messages
          </Tabs.Tab>
          <Tabs.Tab
            value="settings"
            leftSection={<IconSettings style={iconStyle} />}
          >
            Settings
          </Tabs.Tab>
        </Tabs.List>

        <Tabs.Panel value="gallery"> </Tabs.Panel>

        <Tabs.Panel value="messages"> </Tabs.Panel>

        <Tabs.Panel value="settings"> </Tabs.Panel>
      </Tabs>
    </>
  );
};

export { WorkspaceTab };
