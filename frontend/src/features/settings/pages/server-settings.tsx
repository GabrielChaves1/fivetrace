import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog";
import { Avatar, AvatarImage } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Tabs, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { AlertCircle } from "lucide-react";

export default function ServerSettings() {
  return (
    <div className="flex flex-col h-full gap-8">
      <div className="space-y-4">
        <div>
          <h1 className="text-lg font-semibold text-foreground/90">Copacabana Roleplay</h1>
          <p className="text-sm text-muted-foreground">Nesta aba você poderá configurar tudo relacionado ao servidor</p>
        </div>
        <Tabs defaultValue="general">
          <TabsList>
            <TabsTrigger value="general">Geral</TabsTrigger>
            <TabsTrigger value="account">Chaves de acesso</TabsTrigger>
            <TabsTrigger value="password">Integração</TabsTrigger>
          </TabsList>
        </Tabs>
      </div>
      <Card>
        <CardContent className="p-6 grid grid-cols-1 gap-6 md:grid-cols-2">
          <h3 className="font-medium">Configurações gerais</h3>
          <div className="flex flex-col gap-4">
            <div className="flex space-x-3">
              <Avatar className="w-20 h-20 rounded-md">
                <AvatarImage className="rounded-md" src="/example.png" alt="Ícone do servidor" />
              </Avatar>
              <div className="flex flex-col gap-2">
                <Label className="text-sm text-muted-foreground">Logo do servidor</Label>
                <Button size={"sm"} variant={"outline"} className="w-max">Alterar logo</Button>
              </div>
            </div>
            <div className="space-y-2">
              <Label className="text-sm text-muted-foreground">Nome do servidor</Label>
              <Input className="h-10" placeholder="Nome do servidor" value={"Copacabana Roleplay"} />
            </div>
            <div className="space-y-2">
              <Label className="text-sm text-muted-foreground">IP do servidor</Label>
              <Input className="h-10" placeholder="IP do servidor" value={"127.0.0.1"} />
            </div>
          </div>
        </CardContent>
        <CardFooter className="border-t flex items-center justify-end space-x-2 border-foreground/10 p-4 px-6">
          <Button variant={"outline"} size={"sm"}>Cancelar</Button>
          <Button size={"sm"}>Salvar alterações</Button>
        </CardFooter>
      </Card>
      <Card>
        <CardHeader className="border-b border-foreground/10">
          <CardTitle>Zona de Perigo</CardTitle>
        </CardHeader>
        <CardContent className="p-6">
          <Alert variant={"destructive"}>
            <AlertCircle className="h-4 w-4" />
            <AlertTitle>Deletar servidor</AlertTitle>
            <AlertDescription>
              Ao deletar esse servidor, todos os dados serão perdidos e não poderão ser recuperados.
            </AlertDescription>
            <div className="mt-2">
              <AlertDialog>
                <AlertDialogTrigger asChild>
                  <Button variant="destructive" size={"sm"}>Deletar servidor</Button>
                </AlertDialogTrigger>
                <AlertDialogContent>
                  <AlertDialogHeader>
                    <AlertDialogTitle>Tem certeza absoluta?</AlertDialogTitle>
                    <AlertDialogDescription>
                      Essa ação não pode ser desfeita. Isso irá deletar permanentemente seu servidor
                      e remover seus dados dos nossos servidores.
                    </AlertDialogDescription>
                  </AlertDialogHeader>
                  <AlertDialogFooter>
                    <AlertDialogCancel>Cancelar</AlertDialogCancel>
                    <AlertDialogAction>Continuar</AlertDialogAction>
                  </AlertDialogFooter>
                </AlertDialogContent>
              </AlertDialog>
            </div>
          </Alert>
        </CardContent>
      </Card>
    </div>
  )
}