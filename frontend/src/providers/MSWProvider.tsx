'use client'

import { useEffect, useState } from 'react'

export function MSWProvider({
  children,
}: {
  children: React.ReactNode
}) {
  const [mswReady, setMswReady] = useState(false)

  useEffect(() => {
    async function init() {
      // Typically only run MSW in development
      if (process.env.NODE_ENV === 'development' && typeof window !== 'undefined') {
        const { worker } = await import('@/mocks/browser')
        await worker.start({
            onUnhandledRequest: 'bypass',
        })
        setMswReady(true)
      } else {
         setMswReady(true)
      }
    }

    init()
  }, [])

  if (!mswReady) {
      // You can return a loading spinner here if you want to block the app until MSW is ready
      return (
        <div style={{ display: 'flex', height: '100vh', justifyContent: 'center', alignItems: 'center', background: '#0a0a0a', color: '#fff' }}>
            Loading SubWatch Environment...
        </div>
      )
  }

  return <>{children}</>
}
