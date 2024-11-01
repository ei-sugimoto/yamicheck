import { createConnectTransport } from '@connectrpc/connect-node';
const baseUrl = process.env.API_BASE_URL || 'http://api:8000';
export const baseTransport = createConnectTransport({
  baseUrl: baseUrl,
  httpVersion: '2',
  interceptors: [],
});
