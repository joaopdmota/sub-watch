export type Currency = 'BRL' | 'USD' | 'EUR';
export type BillingCycle = 'monthly' | 'yearly';
export type SubscriptionStatus = 'ACTIVE' | 'CANCELLED' | 'PAUSED';

export interface Subscription {
  id: string;
  name: string;
  price: number;
  currency: Currency;
  category: string;
  billingCycle: BillingCycle;
  status: SubscriptionStatus;
  nextPaymentDate: string;
  logoUrl?: string;
}

export interface DashboardSummary {
  totalMonthly: number;
  totalYearly: number;
  activeSubscriptions: number;
  recentActivity: Array<{
    id: string;
    action: 'payment' | 'created' | 'cancelled';
    title: string;
    date: string;
    amount: number;
  }>;
}

export interface User {
  id: string;
  name: string;
  email: string;
  avatarUrl?: string;
}
