import Grid from "@/components/grid";
import MaxWidthWrapper from "@/components/max-width-wrapper";
import MobileNavbar from "@/components/mobile-nav-bar";
import Navbar from "@/components/nav-bar";
import NotificationWidget from "@/components/notification-widget";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbList,
  BreadcrumbPage
} from "@/components/ui/breadcrumb";
import useBreakpoint from "@/hooks/useBreakpoint";

export default function Dashboard() {
  const isDesktop = useBreakpoint()

  return (
    <MaxWidthWrapper className="flex-col desktop:flex-row">
      {isDesktop ? <Navbar /> : <MobileNavbar />}
      <div className="w-full flex flex-col">
        <div className="h-20 hidden px-10 desktop:flex items-center justify-between bg-container/50 border-b border-foreground/10">
          <Breadcrumb>
            <BreadcrumbList>
              <BreadcrumbItem>
                <BreadcrumbPage>Página Inicial</BreadcrumbPage>
              </BreadcrumbItem>
            </BreadcrumbList>
          </Breadcrumb>
          <div className="flex items-center space-x-2">
            <NotificationWidget />
            <Avatar>
              <AvatarFallback>GC</AvatarFallback>
              <AvatarImage src="https://avatars.githubusercontent.com/u/37236608?v=4" alt="Guilherme Carvalho" />
            </Avatar>
          </div>
        </div>
        <div className="flex-1 w-full">
          
        </div>
      </div>
    </MaxWidthWrapper>
  )
}