import { ReactNode } from "react";

// Input Component
interface InputProps extends React.InputHTMLAttributes<HTMLInputElement> {}
export function Input({ className = "", ...props }: InputProps) {
  return (
    <input
      className={`flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 ${className}`}
      {...props}
    />
  )
}

// Label Component
interface LabelProps extends React.LabelHTMLAttributes<HTMLLabelElement> {}
export function Label({ className = "", ...props }: LabelProps) {
  return (
    <label
      className={`text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70 ${className}`}
      {...props}
    />
  )
}

// Select (Simplified for now, native select)
interface SelectProps extends React.SelectHTMLAttributes<HTMLSelectElement> {}
export function Select({ className = "", children, ...props }: SelectProps) {
    return (
        <div className="relative">
            <select className={`flex h-10 w-full items-center justify-between rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background placeholder:text-muted-foreground focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 appearance-none ${className}`} {...props}>
                {children}
            </select>
            <div className="absolute right-3 top-3 pointer-events-none opacity-50">
               ▼
            </div>
        </div>
    )
}
