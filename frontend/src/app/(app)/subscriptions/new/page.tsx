import { SubscriptionForm } from "@/features/subscriptions/components/SubscriptionForm";

export default function NewSubscriptionPage() {
  return (
    <div className="py-8">
      <h1 className="text-3xl font-bold mb-8 text-center">New Subscription</h1>
      <SubscriptionForm />
    </div>
  );
}
