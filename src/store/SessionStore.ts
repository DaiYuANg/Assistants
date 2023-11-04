import { Connection } from '../type/Connection.ts';
import { create } from 'zustand';
import { devtools } from 'zustand/middleware';

type GlobalConnectionStoreType = {
  connections: Array<Connection>;
  addConnection: (connection: Connection) => void;
};

const useSessionStore = create<GlobalConnectionStoreType>()(
  devtools(
    (set) => ({
      connections: [],
      addConnection: (connection: Connection) => {
        set(() => {
          return {
            connections: [connection],
          };
        });
      },
    }),
    { name: 'layoutStore' },
  ),
);

export { useSessionStore };

export type { GlobalConnectionStoreType };
