import { QueryClient } from "@tanstack/solid-query";
import { createConnectTransport } from "@connectrpc/connect-web";

export const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 5000,
      refetchOnWindowFocus: false,
    },
  },
});

export const transport = createConnectTransport({
  baseUrl: "/api",
});
