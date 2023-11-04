import { invoke } from '@tauri-apps/api';
import { Methods } from './methods.ts';
import { Protocol } from '../type/Protocol.ts';

type TestConnectionArg = {
  protocol: Protocol;
  addr: string;
};
type TestConnectionResult = {
  result: string;
};
const TestProtocol = (
  arg: TestConnectionArg,
): Promise<TestConnectionResult> => {
  return invoke<TestConnectionResult>(Methods.TEST_CONNECTION, arg);
};

export { TestProtocol };
export type { TestConnectionArg };
