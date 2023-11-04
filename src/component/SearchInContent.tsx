import { Button, Dialog, Group, Text, TextInput } from '@mantine/core';
import { useDisclosure, useHotkeys } from '@mantine/hooks';

const SearchInContent = () => {
  useHotkeys([
    ['mod+F', () => toggle()],
    // ['ctrl+K', () => console.log('Trigger search')],
    // ['alt+mod+shift+X', () => console.log('Rick roll')],
  ]);

  const [opened, { toggle, close }] = useDisclosure(false);
  return (
    <>
      <Dialog
        opened={opened}
        withCloseButton
        onClose={close}
        size="lg"
        radius="md"
        position={{ top: 20, right: 20 }}
      >
        <Text size="sm" mb="xs" fw={500}>
          Search In Content
        </Text>

        <Group align="flex-end">
          <TextInput
            size={'sm'}
            placeholder="hello@gluesticker.com"
            style={{ flex: 1 }}
          />
          <Button size={'sm'} onClick={close}>
            Subscribe
          </Button>
        </Group>
      </Dialog>
    </>
  );
};

export { SearchInContent };
