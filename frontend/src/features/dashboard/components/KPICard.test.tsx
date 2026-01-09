import { describe, it, expect } from 'vitest'
import { render, screen } from '@testing-library/react'
import { axe } from 'vitest-axe'
import { KPICard } from './KPICard'

describe('KPICard Feature', () => {
  it('renders title and value', async () => {
    const { container } = render(<KPICard title="Total Spent" value="R$ 100,00" />)
    
    expect(screen.getByText("Total Spent")).toBeInTheDocument()
    expect(screen.getByText("R$ 100,00")).toBeInTheDocument()
    
    const results = await axe(container)
    expect(results).toHaveNoViolations()
  })

  it('renders subtext when provided', () => {
    render(<KPICard title="Test" value="123" subtext="vs last month" />)
    expect(screen.getByText("vs last month")).toBeInTheDocument()
  })
})
