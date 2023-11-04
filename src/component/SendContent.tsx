import classes from './Content.module.css';
import { Button, Center, Grid, TextInput } from '@mantine/core';
import { IconCode } from '@tabler/icons-react';
import { useHotkeys } from '@mantine/hooks';

const SendContent = () => {
  useHotkeys([['enter', () => send()]]);
  const send = () => {
    console.log('send');
  };
  return (
    <>
      <Grid
        className={[
          classes.content,
          // 'fixed',
          // 'bottom-0',
          'w-full',
        ].join(' ')}
        p={'xs'}
      >
        <Grid.Col span={11}>
          <TextInput rightSection={<IconCode />} />
        </Grid.Col>
        <Grid.Col span={1}>
          <Center>
            <Button>send</Button>
          </Center>
        </Grid.Col>
      </Grid>
    </>
  );
};

export { SendContent };
