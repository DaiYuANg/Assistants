import MonacoEditor from 'react-monaco-editor/lib/editor';

const message = Array.from({ length: 100 }, (_, index) => index++);

const MessageContent = () => {
  // const items = message.map((item) => (
  //   <>
  //     <Message content={item.toString()} key={item + 1} />
  //     <Divider key={item + 2} />
  //   </>
  // ));
  const onChange = () => {};
  const options = {
    selectOnLineNumbers: true,
  };
  return (
    <>
      {/*<ScrollArea.Autosize h={'90vh'}>*/}
      <MonacoEditor
        width="800"
        height="600"
        language="javascript"
        theme="vs-dark"
        options={options}
        onChange={onChange}
      />
      {/*<Container fluid pt={"sm"}>*/}
      {/*  {items}*/}
      {/*</Container>*/}
      {/*</ScrollArea.Autosize>*/}
    </>
  );
};

export { MessageContent };
