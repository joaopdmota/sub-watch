import { describe, it, expect } from 'vitest'
import { render, screen } from '@testing-library/react'
import { axe } from 'vitest-axe'
import { Hero } from './Hero'
import { vi } from 'vitest'

// Mock next/link to avoid issues outside Next app context
vi.mock('next/link', () => {
    return {
        __esModule: true,
        default: ({ children }: { children: React.ReactNode }) => {
            return <a>{children}</a>;
        }
    }
})

describe('Hero Feature', () => {
  it('renders hero content', async () => {
    const { container } = render(<Hero />)
    
    expect(screen.getByText(/domine suas assinaturas/i)).toBeInTheDocument()
    expect(screen.getByRole('button', { name: /começar agora/i })).toBeInTheDocument()
    
    const results = await axe(container)
    expect(results).toHaveNoViolations()
  })
})
