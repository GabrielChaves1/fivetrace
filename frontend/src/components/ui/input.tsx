import { cn } from "@/lib/utils"
import { forwardRef } from "react"
import { useFormField } from "./form"

export interface InputProps
  extends React.InputHTMLAttributes<HTMLInputElement> { }

const Input = forwardRef<HTMLInputElement, InputProps>(
  ({ className, type, ...props }, ref) => {
    const { error } = useFormField()

    return (
      <input
        type={type}
        className={cn("flex h-14 w-full rounded-md border border-input bg-background px-4 py-1 text-sm transition-colors file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:opacity-70 focus-visible:outline-none focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50",
          error && "border-destructive/60",
          className
        )}
        ref={ref}
        {...props}
      />
    )
  }
)
Input.displayName = "Input"

export { Input }

