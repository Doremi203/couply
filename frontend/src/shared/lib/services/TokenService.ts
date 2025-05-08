// // Function to check if token is expired or about to expire (within 5 minutes)
// export const isTokenExpired = (): boolean => {
//   const expiresAtStr = localStorage.getItem('tokenExpiresAt');

//   if (!expiresAtStr) {
//     return true;
//   }

//   const expiresAt = parseInt(expiresAtStr, 10);
//   // Check if token expires in less than 5 minutes (300000 ms)
//   return Date.now() > expiresAt - 300000;
// };

// // Function to get the current token
// export const getToken = (): string | null => {
//   return localStorage.getItem('token');
// };

// // Function to set token and expiration after login/refresh
// export const setToken = (token: string, expiresIn: number): void => {
//   localStorage.setItem('token', token);
//   const expiresAt = Date.now() + expiresIn * 1000;
//   localStorage.setItem('tokenExpiresAt', expiresAt.toString());
// };

// // Function to clear token on logout
// export const clearToken = (): void => {
//   localStorage.removeItem('token');
//   localStorage.removeItem('tokenExpiresAt');
// };

//TODO
