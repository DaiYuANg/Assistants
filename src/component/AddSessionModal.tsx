import {
  ActionIcon,
  Button,
  Divider,
  Modal,
  NumberInput,
  Select,
  Stack,
  TextInput,
} from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import { IconDeviceDesktop, IconPlus } from '@tabler/icons-react';
import { Protocol } from '../type/Protocol.ts';
import { useForm } from '@mantine/form';
import { useCallback, useState } from 'react';
import { nprogress } from '@mantine/nprogress';
import { TestProtocol } from '../calls';
import { notifications } from '@mantine/notifications';
import { catchError, from, of, tap } from 'rxjs';
import { debounce } from 'lodash';

type connectionFormType = {
  protocol: Protocol;
  address: string;
  port: number;
};

const SessionModal = (props: { opened: boolean; close: () => void }) => {
  const connectionForm = useForm<connectionFormType>({
    initialValues: {
      protocol: Protocol.TCP_CLIENT,
      address: '127.0.0.1',
      port: 0,
    },

    validate: {
      // email: (value) => (/^\S+@\S+$/.test(value) ? null : "Invalid email"),
    },
  });
  const [buttonText, setButtonText] = useState<string>('Test Connection');

  const testConnection = useCallback(
    debounce(() => {
      nprogress.start();
      switch (connectionForm.values.protocol) {
        case Protocol.TCP_CLIENT:
          break;
        case Protocol.TCP_SERVER:
          break;
        case Protocol.UDP_CLIENT:
          break;
        case Protocol.UDP_SERVER:
          break;
      }
      const connectionFullAddress =
        connectionForm.values.address +
        ':' +
        String(connectionForm.values.port);
      from(
        TestProtocol({
          protocol: connectionForm.values.protocol,
          addr: connectionFullAddress,
        }),
      )
        .pipe(
          tap((r) => {
            console.log(r);
          }),
          catchError((e) => {
            notifications.show({
              color: 'red',
              title: 'Connection',
              message: e,
            });
            return of(null); // 返回一个包含null值的Observable，以继续执行
          }),
        )
        .subscribe();
    }, 200),
    [connectionForm],
  );
  return (
    <>
      <Modal
        opened={props.opened}
        overlayProps={{
          backgroundOpacity: 0.55,
          blur: 3,
        }}
        onClose={props.close}
        title="Add Session"
      >
        <form onSubmit={testConnection}>
          {buttonText === 'Save Session' && (
            <TextInput label="Session Name" placeholder="Input Session Name" />
          )}
          <Select
            label="Select Protocol"
            description="Select Connection Protocol"
            placeholder="Pick value"
            rightSection={<IconDeviceDesktop />}
            data={Object.values(Protocol)}
            {...connectionForm.getInputProps('protocol')}
          />
          <TextInput
            label="Input Ip Address"
            description="Input Connection Ip address"
            placeholder="Input placeholder"
            {...connectionForm.getInputProps('address')}
          />
          <NumberInput
            label="Input Port"
            description="Input Connection Port"
            placeholder="Input placeholder"
            {...connectionForm.getInputProps('port')}
          />
          <Divider />
          <Stack pt={'md'}>
            <Button onClick={testConnection}>{buttonText}</Button>
          </Stack>
        </form>
      </Modal>
    </>
  );
};

const AddSessionModal = () => {
  const [opened, { open, close }] = useDisclosure(false);

  return (
    <>
      <SessionModal opened={opened} close={close} />

      <ActionIcon size={'sm'} onClick={open}>
        <IconPlus />
      </ActionIcon>
    </>
  );
};

export { AddSessionModal };
