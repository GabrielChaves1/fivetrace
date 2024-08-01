import { cva, type VariantProps } from "class-variance-authority"
import * as React from "react"

import { cn } from "@/lib/utils"

const badgeVariants = cva(
  "inline-flex items-center rounded-md border px-2 py-0.5 text-[.7rem] md:text-xs transition-colors focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2",
  {
    variants: {
      variant: {
        default:
          "border-transparent bg-primary text-primary-foreground shadow hover:bg-primary/80",
        secondary:
          "border-transparent bg-secondary text-secondary-foreground hover:bg-secondary/80",
        outline: "text-foreground",
      },
      color: {
        destructive: "border-transparent bg-destructive text-foreground",
        warn: "border-transparent bg-warn text-foreground",
        important: "border-transparent bg-important text-foreground",
        secondary: "border-transparent bg-secondary text-foreground",
      }
    },
    compoundVariants: [
      {
        variant: "outline",
        color: "destructive",
        class: "border-destructive/50 text-destructive/90 bg-transparent",
      },
      {
        variant: "outline",
        color: "warn",
        class: "border-warn/50 text-warn/90 bg-transparent"
      },  
      {
        variant: "outline",
        color: "important",
        class: "border-important/50 text-important/90 bg-transparent"
      },
      {
        variant: "outline",
        color: "secondary",
        class: "border-secondary/50 text-secondary/90 bg-transparent"
      }
    ],
    defaultVariants: {
      variant: "default",
    },
  }
)

export interface BadgeProps
  extends Omit<React.HTMLAttributes<HTMLDivElement>, "color">,
    VariantProps<typeof badgeVariants> {}

function Badge({ className, variant, color, ...props }: BadgeProps) {
  return (
    <div className={cn(badgeVariants({ variant, color }), className)} {...props} />
  )
}

export { Badge, badgeVariants }

