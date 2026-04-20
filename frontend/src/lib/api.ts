export type Provider = 'lm-studio' | 'ollama' | 'external-endpoint';

export type LLMSettings = {
  provider: Provider;
  endpoint: string;
  model: string;
  apiKey: string;
  temperature: number;
  systemPrompt: string;
};

declare global {
  interface Window {
    APP_CONFIG?: {
      apiBase?: string;
    };
  }
}

const apiBase = (window.APP_CONFIG?.apiBase || '/api').replace(/\/$/, '');

async function parseResponse<T>(response: Response): Promise<T> {
  if (!response.ok) {
    const text = await response.text();
    throw new Error(text || `Request failed with ${response.status}`);
  }
  return response.json() as Promise<T>;
}

export async function getSettings(provider: Provider): Promise<LLMSettings> {
  const response = await fetch(`${apiBase}/settings/${provider}`);
  return parseResponse<LLMSettings>(response);
}

export async function listSettings(): Promise<LLMSettings[]> {
  const response = await fetch(`${apiBase}/settings`);
  return parseResponse<LLMSettings[]>(response);
}

export async function saveSettings(provider: Provider, payload: LLMSettings): Promise<LLMSettings> {
  const response = await fetch(`${apiBase}/settings/${provider}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload),
  });
  return parseResponse<LLMSettings>(response);
}

export async function deleteSettings(provider: Provider): Promise<void> {
  const response = await fetch(`${apiBase}/settings/${provider}`, { method: 'DELETE' });
  if (!response.ok) {
    const text = await response.text();
    throw new Error(text || `Delete failed with ${response.status}`);
  }
}
