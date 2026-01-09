'use client'

import { Card, CardContent } from "@/components/ui"

interface Props {
  title: string
  value: string | number
  subtext?: string
  trend?: 'up' | 'down' | 'neutral'
}

export function KPICard({ title, value, subtext }: Props) {
  return (
    <Card>
      <CardContent className="p-6">
        <div className="text-sm font-medium text-muted-foreground">{title}</div>
        <div className="text-3xl font-bold mt-2">{value}</div>
        {subtext && <p className="text-xs text-muted-foreground mt-1">{subtext}</p>}
      </CardContent>
    </Card>
  )
}
