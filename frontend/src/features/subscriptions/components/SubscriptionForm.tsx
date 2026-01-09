'use client'

import React, { useState } from 'react';
import { useRouter } from 'next/navigation';
import { Card, CardContent, CardHeader, CardTitle, Button } from "@/components/ui";
import { Input, Label, Select } from "@/components/ui/form";
import { Currency, BillingCycle } from "@/types";

export function SubscriptionForm() {
  const router = useRouter();
  const [loading, setLoading] = useState(false);
  const [formData, setFormData] = useState({
    name: '',
    price: '',
    currency: 'BRL' as Currency,
    category: '',
    billingCycle: 'monthly' as BillingCycle,
    nextPaymentDate: '',
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setFormData(prev => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    try {
        const payload = {
            ...formData,
            price: parseFloat(formData.price),
        };

        const res = await fetch('/api/subscriptions', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(payload)
        });

        if (!res.ok) throw new Error('Failed to create');

        router.push('/dashboard');
        router.refresh(); // Refresh server components if any (mostly client for now but good practice)
    } catch (error) {
        console.error(error);
        alert('Error creating subscription');
    } finally {
        setLoading(false);
    }
  };

  return (
    <Card className="max-w-2xl mx-auto">
        <CardHeader>
            <CardTitle>Add New Subscription</CardTitle>
        </CardHeader>
        <CardContent>
            <form onSubmit={handleSubmit} className="space-y-6">
                <div className="space-y-2">
                    <Label htmlFor="name">Service Name</Label>
                    <Input 
                        id="name" 
                        name="name" 
                        placeholder="e.g. Netflix, Gym, Adobe" 
                        value={formData.name} 
                        onChange={handleChange}
                        required 
                    />
                </div>

                <div className="grid grid-cols-2 gap-4">
                    <div className="space-y-2">
                        <Label htmlFor="price">Price</Label>
                        <Input 
                            id="price" 
                            name="price" 
                            type="number" 
                            step="0.01" 
                            placeholder="0.00" 
                            value={formData.price} 
                            onChange={handleChange}
                            required 
                        />
                    </div>
                    <div className="space-y-2">
                        <Label htmlFor="currency">Currency</Label>
                        <Select 
                            id="currency" 
                            name="currency" 
                            value={formData.currency} 
                            onChange={handleChange}
                        >
                            <option value="BRL">BRL (R$)</option>
                            <option value="USD">USD ($)</option>
                            <option value="EUR">EUR (€)</option>
                        </Select>
                    </div>
                </div>

                <div className="grid grid-cols-2 gap-4">
                    <div className="space-y-2">
                        <Label htmlFor="category">Category</Label>
                        <Select 
                            id="category" 
                            name="category" 
                            value={formData.category} 
                            onChange={handleChange}
                            required
                        >
                            <option value="">Select a category</option>
                            <option value="Entertainment">Entertainment</option>
                            <option value="Software">Software</option>
                            <option value="Infrastructure">Infrastructure</option>
                            <option value="Health">Health</option>
                            <option value="Education">Education</option>
                            <option value="Other">Other</option>
                        </Select>
                    </div>
                    <div className="space-y-2">
                        <Label htmlFor="billingCycle">Billing Cycle</Label>
                        <Select 
                            id="billingCycle" 
                            name="billingCycle" 
                            value={formData.billingCycle} 
                            onChange={handleChange}
                        >
                            <option value="monthly">Monthly</option>
                            <option value="yearly">Yearly</option>
                        </Select>
                    </div>
                </div>

                <div className="space-y-2">
                    <Label htmlFor="nextPaymentDate">Next Payment Date</Label>
                    <Input 
                        id="nextPaymentDate" 
                        name="nextPaymentDate" 
                        type="date" 
                        value={formData.nextPaymentDate} 
                        onChange={handleChange}
                        required 
                    />
                </div>

                <div className="flex justify-end gap-4 pt-4">
                    <Button type="button" className="bg-muted text-muted-foreground hover:bg-muted/80" onClick={() => router.back()}>
                        Cancel
                    </Button>
                    <Button type="submit" disabled={loading}>
                        {loading ? 'Creating...' : 'Create Subscription'}
                    </Button>
                </div>
            </form>
        </CardContent>
    </Card>
  );
}
