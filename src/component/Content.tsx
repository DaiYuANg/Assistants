import { Box } from '@mantine/core';
import classes from './Content.module.css';
import { SendContent } from './SendContent.tsx';
import { MessageContent } from './MessageContent.tsx';

const Content = () => {
  return (
    <>
      <Box className={[classes.content].join(' ')}>
        <MessageContent />
        <SendContent />
      </Box>
    </>
  );
};

export { Content };
