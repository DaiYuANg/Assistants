import {
  Affix,
  Button,
  Container,
  Divider,
  rem,
  ScrollArea,
  Transition,
} from '@mantine/core';
import { useVirtualizer } from '@tanstack/react-virtual';
import { IconArrowUp } from '@tabler/icons-react';
import { useWindowScroll } from '@mantine/hooks';
import { useRef } from 'react';
import { Message } from './Message.tsx';
const message = Array.from({ length: 100 }, (_, index) => index++);

const MessageContent = () => {
  const containerRef = useRef<HTMLDivElement | null>(null);
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  useVirtualizer({
    count: 10000,
    getScrollElement: () => containerRef.current!,
    estimateSize: () => 35,
  });
  const [scroll, scrollTo] = useWindowScroll();
  const items = message.map((item) => (
    <>
      <Message content={item.toString()} key={item + 1} />
      <Divider key={item + 2} />
    </>
  ));

  return (
    <>
      <ScrollArea.Autosize h={'92vh'}>
        <div ref={containerRef}></div>
        <Container fluid pt={'sm'} pb={'sm'}>
          {items}
        </Container>
        <Affix position={{ bottom: 20, right: 20 }}>
          <Transition transition="slide-up" mounted={scroll.y > 0}>
            {(transitionStyles) => (
              <Button
                leftSection={
                  <IconArrowUp style={{ width: rem(16), height: rem(16) }} />
                }
                style={transitionStyles}
                onClick={() => scrollTo({ y: 0 })}
              >
                Scroll to top
              </Button>
            )}
          </Transition>
        </Affix>
      </ScrollArea.Autosize>
    </>
  );
};

export { MessageContent };
