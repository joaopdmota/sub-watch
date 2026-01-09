'use client'

import React from 'react';
import { Card, CardContent, CardHeader, CardTitle, Button } from "@/components/ui";
import { Input, Label, Select } from "@/components/ui/form";

export function SettingsForm() {
  return (
    <div className="space-y-6">
        <Card>
            <CardHeader>
                <CardTitle>Preferences</CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
                <div className="grid gap-2">
                    <Label htmlFor="currency">Default Currency</Label>
                    <Select id="currency" defaultValue="BRL">
                        <option value="BRL">BRL (R$)</option>
                        <option value="USD">USD ($)</option>
                        <option value="EUR">EUR (€)</option>
                    </Select>
                </div>
                <div className="grid gap-2">
                    <Label htmlFor="theme">Theme</Label>
                    <Select id="theme" defaultValue="dark">
                        <option value="dark">Dark</option>
                        <option value="light">Light</option>
                        <option value="system">System</option>
                    </Select>
                </div>
            </CardContent>
        </Card>

        <Card>
            <CardHeader>
                <CardTitle>Profile</CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
                <div className="grid gap-2">
                    <Label htmlFor="name">Display Name</Label>
                    <Input id="name" defaultValue="João Paulo" />
                </div>
                <div className="grid gap-2">
                    <Label htmlFor="email">Email</Label>
                    <Input id="email" defaultValue="joao@example.com" disabled />
                </div>
            </CardContent>
        </Card>

        <div className="flex justify-end">
            <Button>Save Changes</Button>
        </div>
    </div>
  );
}
