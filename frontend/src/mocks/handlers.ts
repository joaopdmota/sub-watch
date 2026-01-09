import { http, HttpResponse } from 'msw'
import { DashboardSummary, Subscription } from '@/types'

const subscriptions: Subscription[] = [
  { 
    id: '1', 
    name: 'Netflix', 
    price: 55.90, 
    currency: 'BRL', 
    billingCycle: 'monthly', 
    nextPaymentDate: '2024-02-15T00:00:00Z', 
    category: 'Entertainment',
    status: 'ACTIVE',
    logoUrl: 'https://upload.wikimedia.org/wikipedia/commons/0/08/Netflix_2015_logo.svg'
  },
  { 
    id: '2', 
    name: 'Spotify', 
    price: 21.90, 
    currency: 'BRL', 
    billingCycle: 'monthly', 
    nextPaymentDate: '2024-02-10T00:00:00Z', 
    category: 'Music',
    status: 'ACTIVE'
  },
  { 
    id: '3', 
    name: 'AWS', 
    price: 15.00, 
    currency: 'USD', 
    billingCycle: 'monthly', 
    nextPaymentDate: '2024-02-01T00:00:00Z', 
    category: 'Infrastructure',
    status: 'ACTIVE'
  },
  { 
    id: '4', 
    name: 'Adobe Creative Cloud', 
    price: 224.00, 
    currency: 'BRL', 
    billingCycle: 'monthly', 
    nextPaymentDate: '2024-02-22T00:00:00Z', 
    category: 'Software',
    status: 'PAUSED'
  },
]

export const handlers = [
  // User Auth Mock
  http.post('/api/auth/login', () => {
    return HttpResponse.json({
        user: { 
            id: 'u1', 
            name: 'Joao Paulo', 
            email: 'joao@example.com',
            avatarUrl: 'https://github.com/joaopdmota.png'
        },
        token: 'fake-jwt-token'
    })
  }),

  http.get('/api/user/me', () => {
    return HttpResponse.json({
        id: 'u1', 
        name: 'Joao Paulo', 
        email: 'joao@example.com',
        avatarUrl: 'https://github.com/joaopdmota.png'
    })
  }),

  // Subscriptions CRUD
  http.get('/api/subscriptions', () => {
    return HttpResponse.json(subscriptions)
  }),

  http.get('/api/subscriptions/:id', ({ params }) => {
    const sub = subscriptions.find(s => s.id === params.id)
    if (!sub) return new HttpResponse(null, { status: 404 })
    return HttpResponse.json(sub)
  }),

  http.post('/api/subscriptions', async ({ request }) => {
    const newSub = await request.json() as Partial<Subscription>
    const created: Subscription = {
        ...newSub,
        id: Math.random().toString(36).substr(2, 9),
        status: 'ACTIVE',
    } as Subscription
    
    // In memory push (won't persist reload but good for demo)
    subscriptions.push(created)
    return HttpResponse.json(created, { status: 201 })
  }),

  // Dashboard Aggregates
  http.get('/api/dashboard/summary', () => {
    const totalMonthly = subscriptions
        .filter(s => s.status === 'ACTIVE' && s.billingCycle === 'monthly')
        .reduce((acc, s) => acc + (s.currency === 'USD' ? s.price * 5 : s.price), 0)

    const summary: DashboardSummary = {
        totalMonthly,
        totalYearly: totalMonthly * 12,
        activeSubscriptions: subscriptions.filter(s => s.status === 'ACTIVE').length,
        recentActivity: [
            { id: 'ev1', action: 'payment', title: 'Netflix charged', date: '2024-01-15T10:00:00Z', amount: 55.90 },
            { id: 'ev2', action: 'created', title: 'New AWS Subscription', date: '2024-01-10T14:30:00Z', amount: 15.00 },
        ]
    }
    return HttpResponse.json(summary)
  })
]
