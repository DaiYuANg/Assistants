import { Accordion } from '@mantine/core';
import { IconNetwork } from '@tabler/icons-react';
import { Protocol } from '../type/Protocol.ts';

const protocolSession = Object.values(Protocol).map((value, index, array) => {
  console.log(value);
  return {
    value: value,
    icon: <IconNetwork />,
  };
});

const ProtocolSessions = () => {
  const items = protocolSession.map((item) => (
    <Accordion.Item key={item.value} value={item.value}>
      <Accordion.Control icon={item.icon}>{item.value}</Accordion.Control>
      <Accordion.Panel></Accordion.Panel>
    </Accordion.Item>
  ));
  return (
    <>
      <Accordion>{items}</Accordion>
    </>
  );
};

export { ProtocolSessions };
