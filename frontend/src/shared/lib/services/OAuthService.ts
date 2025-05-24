export function generateCodeVerifier(length = 64): string {
  return generateRandomString(length);
}

export function generateState(length = 32): string {
  return generateRandomString(length);
}

export async function generateCodeChallenge(codeVerifier: string): Promise<string> {
  // SHA-256 хеширование
  const encoder = new TextEncoder();
  const data = encoder.encode(codeVerifier);
  return await crypto.subtle.digest('SHA-256', data).then(hash => {
    return btoa(String.fromCharCode(...new Uint8Array(hash)))
      .replace(/\+/g, '-')
      .replace(/\//g, '_')
      .replace(/=+$/, '');
  });
}

function generateRandomString(length = 32): string {
  const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-._~';
  const array = new Uint8Array(length);
  crypto.getRandomValues(array);

  return Array.from(array)
    .map(byte => chars[byte % chars.length])
    .join('');
}

export function generateDeviceId(): string {
  return crypto.randomUUID();
}
