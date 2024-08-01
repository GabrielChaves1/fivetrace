import Header from "@/components/header";
import MaxWidthWrapper from "@/components/max-width-wrapper";
import { Avatar, AvatarImage } from "@/components/ui/avatar";
import { Tabs, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Outlet } from "react-router-dom";

export default function Panel() {
  return (
    <MaxWidthWrapper>
      <Header />
      <div className="container mx-auto p-0 border-l border-r flex flex-1">
        <div className="flex flex-col w-full h-full pt-6">
          <div className="space-y-4">
            <div className="flex gap-2 items-center px-8">
              <Avatar className="w-14 h-14 rounded-md">
                <AvatarImage className="rounded-md" src="/example.png" />
              </Avatar>
              <div>
                <h1 className="text-lg font-semibold text-foreground/90">Painel de monitoramento</h1>
                <p className="text-sm text-muted-foreground">Monitore todos os dados em tempo real do servidor <b>Copacabana Roleplay</b></p>
              </div>
            </div>
            <Tabs value="logs">
              <TabsList className="px-8">
                <TabsTrigger value="logs">
                  Monitoramento de Logs
                </TabsTrigger>
                <TabsTrigger disabled value="statistics">
                  Estatísticas
                </TabsTrigger>
                <TabsTrigger disabled value="statistics">
                  Reinicialização
                </TabsTrigger>
              </TabsList>
            </Tabs>
          </div>
          <div className="p-8 h-full">
            <Outlet />
          </div>
        </div>
      </div>
    </MaxWidthWrapper>
  )
}