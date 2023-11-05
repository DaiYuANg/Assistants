import { ActionIcon, Group, Modal } from '@mantine/core';
import { IconX } from '@tabler/icons-react';
import { useColorScheme, useViewportSize } from '@mantine/hooks';
import { lazy, useState } from 'react';
const MonacoEditor = lazy(() =>
  import('react-monaco-editor/lib/editor').then((module) => ({
    default: module.default,
  })),
);
const CodeModal = (props: { opened: boolean; close: () => void }) => {
  const { height, width } = useViewportSize();
  const color = useColorScheme();
  const [moncaoTheme] = useState<string>(
    color === 'dark' ? 'vs-dark' : 'vs-light',
  );
  const options = {
    selectOnLineNumbers: true,
    fontSize: 16,
    fontFamily: 'Jetbrains Mono',
  };
  const onChange = (v: any) => {
    console.log(v);
  };
  return (
    <>
      <Modal
        padding={0}
        opened={props.opened}
        onClose={props.close}
        fullScreen
        withCloseButton={false}
        radius={0}
        transitionProps={{ transition: 'fade', duration: 200 }}
        pos={'relative'}
        className={['p-0', 'm-0'].join(' ')}
      >
        <Group
          pos={'fixed'}
          className={['z-50', 'top-2', 'right-2'].join(' ')}
          justify={'end'}
        >
          <ActionIcon onClick={props.close}>
            <IconX />
          </ActionIcon>
        </Group>
        <MonacoEditor
          width={width}
          height={height}
          language="text"
          theme={moncaoTheme}
          options={options}
          onChange={onChange}
        />
      </Modal>
    </>
  );
};

export { CodeModal };
