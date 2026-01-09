
import { DashboardOverview } from "@/features/dashboard/components/DashboardOverview";
import { SubscriptionList } from "@/features/subscriptions/components/SubscriptionList";

export default function DashboardPage() {
  return (
    <div className="space-y-8">
      <div>
        <h1 className="text-3xl font-bold tracking-tight">Dashboard</h1>
        <p className="text-muted-foreground">Overview of your recurring expenses.</p>
      </div>
      
      <DashboardOverview />
      
      <section>
        <h2 className="text-xl font-semibold mb-4 text-muted-foreground">Your Subscriptions</h2>
        <SubscriptionList />
      </section>
    </div>
  );
}
