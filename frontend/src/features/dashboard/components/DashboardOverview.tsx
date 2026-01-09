'use client'

import React, { useEffect, useState } from 'react';
import { KPICard } from './KPICard';
import { Card, CardHeader, CardTitle, CardContent } from "@/components/ui";
import { DashboardSummary } from "@/types";

export function DashboardOverview() {
    const [summary, setSummary] = useState<DashboardSummary | null>(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        fetch('/api/dashboard/summary')
            .then(res => res.json())
            .then(data => {
                setSummary(data);
                setLoading(false);
            })
            .catch(err => {
                console.error(err);
                setLoading(false);
            });
    }, []);

    if (loading) return <div>Loading dashboard...</div>;
    if (!summary) return <div>Failed to load data.</div>;

    return (
        <div className="space-y-8">
            <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
                <KPICard 
                    title="Total Monthly" 
                    value={`R$ ${summary.totalMonthly.toFixed(2)}`} 
                    subtext="Forecast for this month"
                />
                <KPICard 
                    title="Total Yearly" 
                    value={`R$ ${summary.totalYearly.toFixed(2)}`}
                    subtext="Projected annual cost"
                />
                <KPICard 
                    title="Active Subscriptions" 
                    value={summary.activeSubscriptions} 
                />
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                <Card className="col-span-1">
                    <CardHeader>
                        <CardTitle>Recent Activity</CardTitle>
                    </CardHeader>
                    <CardContent>
                        <div className="space-y-4">
                            {summary.recentActivity.map((activity) => (
                                <div key={activity.id} className="flex justify-between items-center border-b pb-2 last:border-0 last:pb-0">
                                    <div>
                                        <div className="font-medium text-sm">{activity.title}</div>
                                        <div className="text-xs text-muted-foreground">{new Date(activity.date).toLocaleDateString()}</div>
                                    </div>
                                    <div className="text-sm font-semibold">
                                        {activity.action === 'payment' ? '-' : '+'} R$ {activity.amount.toFixed(2)}
                                    </div>
                                </div>
                            ))}
                        </div>
                    </CardContent>
                </Card>
                
                {/* Placeholder for Chart */}
                <Card className="col-span-1 border-dashed border-2 flex items-center justify-center min-h-[300px] text-muted-foreground bg-accent/20">
                    Category Distribution Chart Needed (Recharts)
                </Card>
            </div>
        </div>
    );
}
