export interface CreatePaymentRequest {
  subscriptionId: string;
  amount: string;
  currency: string;
}

export interface CreatePaymentResponse {
  paymentId: string;
  status: string;
  updatedAt: string;
}

export interface GetPaymentRequest {
  paymentId: string;
}

export interface GetPaymentResponse {
  paymentId: string;
  status: string;
  updatedAt: string;
}
