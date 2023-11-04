import {
  ActionIcon,
  Button,
  Card,
  Combobox,
  Divider,
  Grid,
  Group,
  NumberInput,
  ScrollAreaAutosize,
  Select,
  Stack,
  TextInput,
} from '@mantine/core';
import { ProtocolSessions } from './ProtocolSessions.tsx';
import { AddSessionModal } from './AddSessionModal.tsx';

const Navbar = () => {
  return (
    <>
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
    </>
  );
};

export { Navbar };
