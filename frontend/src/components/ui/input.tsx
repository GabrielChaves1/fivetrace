import { cn } from "@/lib/utils"
import { forwardRef } from "react"

export interface InputProps
  extends React.InputHTMLAttributes<HTMLInputElement> {
    children?: React.ReactNode 
  }

const Input = forwardRef<HTMLInputElement, InputProps>(
  ({ className, type, children, ...props }, ref) => {
    return (
      <div className={cn("flex items-center h-14 relative w-full rounded-md border border-input bg-background", className)}>
        <input
          type={type}
          className={cn("px-4 py-1 h-full flex-1 text-sm transition-colors file:border-0 bg-transparent file:bg-transparent file:text-sm file:font-medium placeholder:opacity-70 focus-visible:outline-none focus-visible:ring-ring disabled:opacity-50",
            "[aria-error]:border-destructive/60",
            className
          )}
          ref={ref}
          {...props}
        />
        {children}
      </div>
    )
  }
)
Input.displayName = "Input"

export { Input }

