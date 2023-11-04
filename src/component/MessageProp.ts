import IntrinsicAttributes = React.JSX.IntrinsicAttributes;

type MessageProp = {
  content: string;
} & Partial<IntrinsicAttributes>;

export type { MessageProp };
