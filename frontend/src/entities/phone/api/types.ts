export interface PhoneParams {
  phone: string;
}

export interface ConfirmParams {
  phone: string;
  code: string;
}

export interface CodeResponse {
  sendAgainIn: number;
}
