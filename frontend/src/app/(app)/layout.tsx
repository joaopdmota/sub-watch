import { AppSidebar } from "@/components/layout/AppSidebar";

export default function AppLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="flex min-h-screen bg-background text-foreground">
      <AppSidebar />
      <main className="flex-1 overflow-y-auto max-h-screen">
          <div className="container mx-auto p-6 md:p-8 max-w-6xl">
            {children}
          </div>
      </main>
    </div>
  );
}
