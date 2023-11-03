import {
  Button,
  Card,
  Divider,
  NumberInput,
  ScrollAreaAutosize,
  Select,
  Stack,
  TextInput,
} from '@mantine/core';
import { HeightLevelSetting } from './HeightLevelSetting.tsx';
import { IconDeviceDesktop } from '@tabler/icons-react';
import { invoke } from '@tauri-apps/api';
import { useCallback } from 'react';
import { useForm } from '@mantine/form';
import { notifications } from '@mantine/notifications';
import { nprogress } from '@mantine/nprogress';

const Navbar = () => {
  const connectionForm = useForm({
    initialValues: {
      email: '',
      termsOfService: false,
    },

    validate: {
      // email: (value) => (/^\S+@\S+$/.test(value) ? null : "Invalid email"),
    },
  });

  const testConnection = useCallback(() => {
    console.log(connectionForm.values);
    nprogress.start();
    invoke('test_connection', { invokeMessage: 'test' })
      .then((result) => {
        console.log(result);
      })
      .catch((e) => {
        notifications.show({
          color: 'red',
          title: 'Connection result',
          message: e,
        });
        console.log(e);
      });
  }, [connectionForm]);

  return (
    <>
      <form onSubmit={testConnection}>
        <ScrollAreaAutosize
          style={{
            minHeight: '90vh',
          }}
        >
          <Stack p={'md'}>
            <Card withBorder shadow="sm" padding="lg" radius="md">
              <Card.Section p={'xs'}>
                <Select
                  label="Select Protocol"
                  description="Select Connection Protocol"
                  placeholder="Pick value"
                  rightSection={<IconDeviceDesktop />}
                  data={['React', 'Angular', 'Vue', 'Svelte']}
                  {...connectionForm.getInputProps('protocol')}
                />
              </Card.Section>
              <Divider />
              <Card.Section p={'xs'}>
                <TextInput
                  label="Input Ip Address"
                  description="Input Connection Ip address"
                  placeholder="Input placeholder"
                  {...connectionForm.getInputProps('address')}
                />
              </Card.Section>
              <Divider />
              <Card.Section p={'xs'}>
                <NumberInput
                  label="Input Port"
                  description="Input Connection Port"
                  placeholder="Input placeholder"
                  {...connectionForm.getInputProps('port')}
                />
              </Card.Section>
            </Card>
            <Button onClick={testConnection}>Test Connection</Button>
            <HeightLevelSetting />
          </Stack>
        </ScrollAreaAutosize>
      </form>
    </>
  );
};

export { Navbar };
