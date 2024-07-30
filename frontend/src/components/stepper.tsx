import { cn } from "@/lib/utils";
import { ComponentProps, useEffect, useState } from "react";

interface StepperProps extends ComponentProps<'div'> {
  currentStep: number;
  steps: number;
}

export default function Stepper({ currentStep, steps, className, ...props }: StepperProps) {
  const [percentage, setPercentage] = useState(0);

  useEffect(() => {
    setPercentage(((currentStep - 1) / (steps - 1)) * 100);
  }, [currentStep, steps]);

  return (
    <div className={cn("w-full max-w-sm flex justify-between items-center relative", className)} {...props}>
      <div className="w-full h-0.5 bg-foreground/10 absolute z-0">
        <div className="h-full bg-primary transition-all duration-300" style={{ width: `${percentage}%` }} />
      </div>
      {Array.from({ length: steps }).map((_, index) => {
        return (
          <div
            key={index}
            className={cn("bg-background z-[1] text-md border-2 grid place-items-center", {
              "w-10 h-10 rounded-full": true,
              "border-primary": index < currentStep,
              "border-foreground/10": index >= currentStep,
              "text-foreground": index < currentStep,
            })}>
            {index + 1}
          </div>
        )
      })}
    </div>
  );
}