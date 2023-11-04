import { Protocol } from './Protocol.ts';

type Connection = {
  id: string;
  protocol: Protocol;
  address: string;
  port: number;
};

export type { Connection };
