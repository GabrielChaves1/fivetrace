import useBreakpoint from "@/hooks/useBreakpoint";
import { cn } from "@/lib/utils";
import { Bell } from "lucide-react";
import { Button } from "./ui/button";
import { DropdownMenu, DropdownMenuContent, DropdownMenuSeparator, DropdownMenuTrigger } from "./ui/dropdown-menu";

export default function NotificationWidget() {
  const isDesktop = useBreakpoint()

  return (
    <>
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button size={"icon"} variant={"outline"} className="relative">
            <div className="w-2.5 h-2.5 bg-primary rounded-full absolute top-0 right-0">
              <div className="w-full h-full bg-primary animate-ping rounded-full absolute"></div>
            </div>
            <Bell size={16} />
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent className={cn("w-72 sm:w-96 mr-5 desktop:mr-0")} align={isDesktop ? "end" : "center"}>
          <p className="text-sm p-2 text-foreground">Notificações</p>
          <DropdownMenuSeparator />
          <div className="w-full h-36 flex flex-col space-y-2 items-center justify-center">
            <p className="text-sm text-foreground">Nenhuma notificação nova</p>
            <span className="text-xs text-muted-foreground text-center w-full max-w-64">Você será notificado aqui sobre quaisquer avisos sobre sua organização</span>
          </div>
        </DropdownMenuContent>
      </DropdownMenu>
    </>
  )
}