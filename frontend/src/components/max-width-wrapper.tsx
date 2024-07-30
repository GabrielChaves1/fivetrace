import { cn } from '@/lib/utils';
import { ComponentProps, ReactNode } from 'react';

type MaxWidthWrapperProps = ComponentProps<'main'> & {
  children: ReactNode;
}

export default function MaxWidthWrapper({ children, className, ...props }: MaxWidthWrapperProps) {
  return (
    <main className={cn('w-screen h-screen flex flex-col', className)} {...props}>
      {children}
    </main>
  )
}