

export interface GetSubscriptionResponse {
    subscriptionId: string;
    status: Status;
    autoRenew: boolean;
    startDate: string;
    endDate: string;
    paymentIds: string[];
    plan: Plan;
}


export enum Plan {
   unspecified ='SUBSCRIPTION_PLAN_UNSPECIFIED',
   monthly = 'SUBSCRIPTION_PLAN_MONTHLY',
   annual = 'SUBSCRIPTION_PLAN_ANNUAL',
   semiAnnual = 'SUBSCRIPTION_PLAN_SEMI_ANNUAL',
}


export enum Status {
    unspecified = 'SUBSCRIPTION_STATUS_UNSPECIFIED',
    active = 'SUBSCRIPTION_STATUS_ACTIVE',
    expired = 'SUBSCRIPTION_STATUS_EXPIRED',
    cancelled = 'SUBSCRIPTION_STATUS_CANCELED',
    pending = 'SUBSCRIPTION_STATUS_PENDING_PAYMENT'
}


export interface CancelRequest {
    subscriptionId: string;
}



export interface CreateSubRequst {
    plan: Plan;
    autoRenew: boolean;
}

