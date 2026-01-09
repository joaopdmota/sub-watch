import { SettingsForm } from "@/features/settings/components/SettingsForm";

export default function SettingsPage() {
  return (
    <div className="space-y-8 max-w-4xl">
      <div>
        <h1 className="text-3xl font-bold tracking-tight">Settings</h1>
        <p className="text-muted-foreground">Manage your account preferences and settings.</p>
      </div>
      <SettingsForm />
    </div>
  );
}
