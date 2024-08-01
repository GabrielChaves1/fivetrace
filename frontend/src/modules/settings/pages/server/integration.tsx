import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { ArrowUpRight, Copy } from "lucide-react";

export default function ServerSettingsIntegration() {
  return (
    <>
      <Card>
        <CardHeader className="border-b">
          <CardTitle>Chaves de API</CardTitle>
          <CardDescription>Configure as chaves de API para integrações com outros serviços.</CardDescription>
          <Button variant={"outline"} size={"sm"} className="w-max px-2">
            <ArrowUpRight size={14} />
            API Docs
          </Button>
        </CardHeader>
        <CardContent className="p-6 space-y-6">
          <div className="grid grid-cols-1 lg:grid-cols-2 gap-2 w-full">
            <h3 className="text-sm opacity-90">Chave do cliente</h3>
            <div className="flex flex-col space-y-1">
              <div className="flex items-center space-x-2">
                <Input className="h-8 text-xs px-1" placeholder="Nome do servidor" value={"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6IndoZ29"} disabled>
                  <Button size={"sm"} className="h-6" variant={"outline"}>
                    <Copy size={14} />
                  </Button>
                </Input>
              </div>
              <span className="text-xs text-muted-foreground/50">Essa chave é utilizada para autenticação</span>
            </div>
          </div>
          <div className="grid grid-cols-1 lg:grid-cols-2 gap-2 w-full">
            <h3 className="text-sm opacity-90">Chave secreta</h3>
            <div className="flex flex-col space-y-1">
              <div className="flex items-center space-x-2">
                <Input className="h-8 text-xs px-1" placeholder="Nome do servidor" type="password" value={"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6IndoZ29"} disabled>
                  <Button size={"sm"} className="h-6" variant={"outline"}>
                    <Copy size={14} />
                  </Button>
                </Input>
              </div>
              <span className="text-xs text-muted-foreground/50">Não compartilhe sua chave secreta com ninguem!</span>
            </div>
          </div>
        </CardContent>
      </Card>
    </>
  )
}