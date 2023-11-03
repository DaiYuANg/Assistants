import classes from './Content.module.css';
import { Button, Grid, TextInput } from '@mantine/core';
import { IconCode } from '@tabler/icons-react';

const SendContent = () => {
  return (
    <>
      <Grid
        className={[
          classes.content,
          'fixed',
          'bottom-0',
          'w-full',
          'border-b-2 border-amber-100',
        ].join(' ')}
        p={'sm'}
      >
        <Grid.Col span={10}>
          <TextInput rightSection={<IconCode />} />
        </Grid.Col>
        <Grid.Col span={2}>
          <Button>send</Button>
        </Grid.Col>
      </Grid>
    </>
  );
};

export { SendContent };
