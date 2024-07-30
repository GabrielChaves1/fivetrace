import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from "@/components/ui/tooltip";
import { Info } from "lucide-react";

export default function Server() {
  return (
    <div className="p-4 flex gap-4 bg-container h-max border border-foreground/10 rounded-md">
      <img src="/example.png" className="w-24 rounded-xl object-cover p-1 border bg-background border-foreground/10" alt="" />
      <div className="flex flex-col justify-between flex-1">
        <div className="flex justify-between">
          <div className="leading-none">
            <p className="font-bold text-md">Copacabana Roleplay</p>
            <span className="text-sm text-muted-foreground">192.0.0.1</span>
          </div>
          <TooltipProvider>
            <Tooltip>
              <TooltipTrigger asChild>
                <Info size={16} className="cursor-pointer text-muted-foreground" />
              </TooltipTrigger>
              <TooltipContent side="left" sideOffset={10}>
                <div className="space-y-2 p-1 flex flex-col">
                  <div className="flex gap-1 items-center">
                    <p className="text-muted-foreground">Status:</p>
                    <span className="text-primary flex gap-1 items-center font-semibold">Operacional</span>
                  </div>
                  <div className="flex gap-2 items-center">
                    <p className="text-muted-foreground">Plano:</p>
                    <Badge variant={"outline"}>Teste Gratuito</Badge>
                  </div>
                </div>
              </TooltipContent>
            </Tooltip>
          </TooltipProvider>
        </div>
        <Button size="sm">Acessar painel</Button>
      </div>
    </div>
  )
}