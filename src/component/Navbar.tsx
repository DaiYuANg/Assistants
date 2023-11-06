import { Divider, Grid, Group, ScrollAreaAutosize, Stack } from '@mantine/core';
import { ProtocolSessions } from './ProtocolSessions.tsx';
import { AddSessionModal } from './AddSessionModal.tsx';

const Navbar = () => {
  return (
    <>
      <div className={['bg-transparent', 'relative'].join(' ')}>
        <ScrollAreaAutosize>
          <Stack p={'md'}>
            <Grid>
              <Grid.Col span={10}>Protocol Assistant</Grid.Col>
              <Grid.Col span={2}>
                <Group justify={'end'} align={'center'}>
                  <AddSessionModal />
                </Group>
              </Grid.Col>
            </Grid>
            <Divider />
            <ProtocolSessions />
          </Stack>
        </ScrollAreaAutosize>
      </div>
    </>
  );
};

export { Navbar };
