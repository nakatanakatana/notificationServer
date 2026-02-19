import type { Transport } from "@connectrpc/connect";
import { createContext, type ParentComponent, useContext } from "solid-js";

const TransportContext = createContext<Transport>();

export const TransportProvider: ParentComponent<{ transport: Transport }> = (
  props,
) => {
  return (
    <TransportContext.Provider value={props.transport}>
      {props.children}
    </TransportContext.Provider>
  );
};

export function useTransport(): Transport {
  const transport = useContext(TransportContext);
  if (!transport) {
    throw new Error("useTransport must be used within a TransportProvider");
  }
  return transport;
}
