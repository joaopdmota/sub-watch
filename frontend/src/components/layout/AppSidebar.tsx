'use client'

import React from 'react';
import Link from 'next/link';
import { usePathname } from 'next/navigation';

const navItems = [
  { href: '/dashboard', label: 'Dashboard', icon: '📊' },
  { href: '/subscriptions', label: 'Subscriptions', icon: '💳' },
  { href: '/settings', label: 'Settings', icon: '⚙️' },
];

export function AppSidebar() {
  const pathname = usePathname();

  return (
    <aside className="w-64 border-r bg-card h-screen hidden md:flex flex-col">
      <div className="p-6 border-b">
        <h1 className="text-xl font-bold tracking-tight text-primary">SubWatch</h1>
      </div>
      <nav className="flex-1 p-4 space-y-1">
        {navItems.map((item) => {
          const isActive = pathname.startsWith(item.href);
          return (
            <Link
              key={item.href}
              href={item.href}
              className={`flex items-center gap-3 px-4 py-3 text-sm font-medium rounded-md transition-colors ${
                isActive
                  ? 'bg-primary text-primary-foreground'
                  : 'text-muted-foreground hover:bg-accent hover:text-accent-foreground'
              }`}
            >
              <span>{item.icon}</span>
              {item.label}
            </Link>
          );
        })}
      </nav>
      <div className="p-4 border-t">
        <div className="flex items-center gap-3">
             <div className="w-8 h-8 rounded-full bg-gray-700"></div>
             <div className="text-sm">
                <p className="font-medium">João Paulo</p>
                <p className="text-xs text-muted-foreground">Pro Plan</p>
             </div>
        </div>
      </div>
    </aside>
  );
}
