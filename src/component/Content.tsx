import { Box } from '@mantine/core';
import classes from './Content.module.css';
import { SendContent } from './SendContent.tsx';
import { MessageContent } from './MessageContent.tsx';
import { Divider } from '@mantine/core';

const Content = () => {
  return (
    <>
      <Box className={[classes.content].join(' ')}>
        <MessageContent />
        <Divider />
        <SendContent />
      </Box>
    </>
  );
};

export { Content };
