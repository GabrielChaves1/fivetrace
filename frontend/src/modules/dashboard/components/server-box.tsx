import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from "@/components/ui/tooltip";
import { Check, Info } from "lucide-react";
import IServer from "../../../types/server";
import { Link } from "react-router-dom";

export default function ServerBox({ id, ip, name }: IServer) {
  return (
    <div className="p-4 flex gap-4 bg-card h-max border rounded-md">
      <Avatar className="w-24 h-24 rounded-md *:rounded-md">
        <AvatarFallback>{name[0]}</AvatarFallback>
        <AvatarImage src="/example.png" alt={name} />
      </Avatar>
      <div className="flex flex-col justify-between flex-1">
        <div className="flex justify-between">
          <div className="leading-none">
            <div className="flex space-x-2 items-center">
              <p className="font-bold text-md">{name}</p>
              <TooltipProvider>
                <Tooltip>
                  <TooltipTrigger asChild>
                    <Check size={16} className="cursor-pointer text-primary" />
                  </TooltipTrigger>
                  <TooltipContent>
                    <div className="space-y-2 p-1 items-center flex flex-col">
                      <div className="flex gap-1 items-center">
                        <p className="text-muted-foreground">Status:</p>
                        <span className="text-primary flex gap-1 items-center font-semibold">Operacional</span>
                      </div>
                      <Button size="sm" variant={"outline"} className="w-full">Healthcheck</Button>
                    </div>
                  </TooltipContent>
                </Tooltip>
              </TooltipProvider>
            </div>
            <span className="text-sm text-muted-foreground">{ip}</span>
          </div>
          <TooltipProvider>
            <Tooltip>
              <TooltipTrigger asChild>
                <Info size={16} className="cursor-pointer text-muted-foreground" />
              </TooltipTrigger>
              <TooltipContent side="left" sideOffset={10}>
                <div className="space-y-2 p-1 flex flex-col">
                  <div className="flex gap-2 items-center">
                    <p className="text-muted-foreground">Plano:</p>
                    <Badge variant={"outline"}>Teste Gratuito</Badge>
                  </div>
                  <div className="flex gap-2 items-center">
                    <p className="text-muted-foreground">Expira em:</p>
                    <span>8 dias</span>
                  </div>
                </div>
              </TooltipContent>
            </Tooltip>
          </TooltipProvider>
        </div>
        <Button asChild size="sm">
          <Link to={`/panel/${id}/logs`}>Acessar painel</Link>
        </Button>
      </div>
    </div>
  )
}