import { ActionIcon, Group, Modal } from '@mantine/core';
import MonacoEditor from 'react-monaco-editor/lib/editor';
import { IconX } from '@tabler/icons-react';
import { useViewportSize } from '@mantine/hooks';

const CodeModal = (props: { opened: boolean; close: () => void }) => {
  const { height, width } = useViewportSize();
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
          theme="vs-dark"
          options={options}
          onChange={onChange}
        />
      </Modal>
    </>
  );
};

export { CodeModal };
