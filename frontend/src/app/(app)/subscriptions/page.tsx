import { SubscriptionList } from "@/features/subscriptions/components/SubscriptionList";
import Link from 'next/link';
import { Button } from "@/components/ui";

export default function SubscriptionsPage() {
  return (
    <div className="space-y-8">
      <div className="flex justify-between items-center">
        <div>
            <h1 className="text-3xl font-bold tracking-tight">Subscriptions</h1>
            <p className="text-muted-foreground">Manage your recurring expenses.</p>
        </div>
        <Link href="/subscriptions/new">
            <Button>+ Add Subscription</Button>
        </Link>
      </div>
      
      <SubscriptionList />
    </div>
  );
}
