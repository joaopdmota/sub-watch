'use client'

import { useEffect, useState } from 'react'
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui"

interface Subscription {
  id: number
  name: string
  price: number
  currency: string
  billingCycle: string
  nextPayment: string
  category: string
}

export function SubscriptionList() {
  const [subscriptions, setSubscriptions] = useState<Subscription[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    fetch('/api/subscriptions')
      .then((res) => res.json())
      .then((data) => {
        setSubscriptions(data)
        setLoading(false)
      })
      .catch((err) => {
          console.error(err)
          setLoading(false)
      })
  }, [])

  if (loading) return <div>Carregando assinaturas...</div>

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      {subscriptions.map((sub) => (
        <Card key={sub.id} className="hover:border-primary/50 transition-colors cursor-pointer">
          <CardHeader className="pb-2">
            <div className="flex justify-between items-start">
                <CardTitle>{sub.name}</CardTitle>
                <span className="text-xs bg-secondary px-2 py-1 rounded text-secondary-foreground">{sub.category}</span>
            </div>
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold mb-1">
                {sub.currency} {sub.price.toFixed(2)}
            </div>
            <div className="text-xs text-muted-foreground">
                {sub.billingCycle} • Próx: {new Date(sub.nextPayment).toLocaleDateString('pt-BR')}
            </div>
          </CardContent>
        </Card>
      ))}
    </div>
  )
}
