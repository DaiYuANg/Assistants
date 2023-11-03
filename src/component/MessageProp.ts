import { JSX } from 'react/jsx-runtime';
import IntrinsicAttributes = JSX.IntrinsicAttributes;

type MessageProp = {
  content: string;
} & Partial<IntrinsicAttributes>;

export type { MessageProp };
