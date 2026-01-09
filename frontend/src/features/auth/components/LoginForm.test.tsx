import { describe, it, expect, vi } from 'vitest'
import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import { axe } from 'vitest-axe'
import { LoginForm } from './LoginForm'

// Mock useRouter
const pushMock = vi.fn()
vi.mock('next/navigation', () => ({
  useRouter: () => ({
    push: pushMock,
  }),
}))

// Mock fetch
global.fetch = vi.fn()

describe('LoginForm Feature', () => {
  it('renders login form with accessible fields', async () => {
    const { container } = render(<LoginForm />)
    
    expect(screen.getByLabelText(/email/i)).toBeInTheDocument()
    expect(screen.getByLabelText(/password/i)).toBeInTheDocument()
    expect(screen.getByRole('button', { name: /sign in/i })).toBeInTheDocument()

    const results = await axe(container)
    expect(results).toHaveNoViolations()
  })

  it('submits the form successfully', async () => {
    (global.fetch as any).mockResolvedValueOnce({ ok: true })

    render(<LoginForm />)
    
    fireEvent.change(screen.getByLabelText(/email/i), { target: { value: 'test@example.com' } })
    fireEvent.change(screen.getByLabelText(/password/i), { target: { value: 'password123' } })
    
    const button = screen.getByRole('button', { name: /sign in/i })
    fireEvent.click(button)

    expect(button).toBeDisabled()
    expect(button).toHaveTextContent(/signing in/i)

    await waitFor(() => {
        expect(pushMock).toHaveBeenCalledWith('/dashboard')
    })
  })
})
