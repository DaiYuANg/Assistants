import { Box, Menu, rem, Text } from '@mantine/core';
import {
  IconArrowsLeftRight,
  IconCopy,
  IconMessageCircle,
  IconPhoto,
  IconSearch,
  IconTrash,
} from '@tabler/icons-react';
import { MessageProp } from './MessageProp.ts';
import { useCallback } from 'react';

const Message = ({ content, key }: MessageProp) => {
  const copyToClipboard = useCallback(async () => {
    navigator.clipboard.writeText(content).then(() => {
      console.log('copied success');
    });
  }, []);

  return (
    <>
      <Menu
        shadow="md"
        key={key}
        transitionProps={{ transition: 'pop', duration: 100 }}
      >
        <Menu.Target key={key}>
          <Box
            className={[
              'block',
              'w-full',
              'bg-sky-200',
              'text-black',
              'hover:bg-sky-700',
              'p-1',
              'rounded',
            ].join(' ')}
            key={key}
          >
            <Text size={'md'}>{content}</Text>
          </Box>
        </Menu.Target>

        <Menu.Dropdown>
          <Menu.Label>Actions</Menu.Label>
          <Menu.Item
            leftSection={
              <IconCopy style={{ width: rem(14), height: rem(14) }} />
            }
            onClick={copyToClipboard}
          >
            Copy
          </Menu.Item>
          <Menu.Item
            leftSection={
              <IconMessageCircle style={{ width: rem(14), height: rem(14) }} />
            }
          >
            Messages
          </Menu.Item>
          <Menu.Item
            leftSection={
              <IconPhoto style={{ width: rem(14), height: rem(14) }} />
            }
          >
            Gallery
          </Menu.Item>
          <Menu.Item
            leftSection={
              <IconSearch style={{ width: rem(14), height: rem(14) }} />
            }
            rightSection={
              <Text size="xs" c="dimmed">
                ⌘K
              </Text>
            }
          >
            Search
          </Menu.Item>

          <Menu.Divider />

          <Menu.Label>Danger zone</Menu.Label>
          <Menu.Item
            leftSection={
              <IconArrowsLeftRight
                style={{ width: rem(14), height: rem(14) }}
              />
            }
          >
            Transfer my data
          </Menu.Item>
          <Menu.Item
            color="red"
            leftSection={
              <IconTrash style={{ width: rem(14), height: rem(14) }} />
            }
          >
            Delete my account
          </Menu.Item>
        </Menu.Dropdown>
      </Menu>
    </>
  );
};

export { Message };
