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
import { AlertCircle } from "lucide-react";

export default function OrganizationSettingsGeneral() {
  return (
    <div>
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
              <Input className="h-10" placeholder="Nome do servidor" />
            </div>
            <div className="space-y-2">
              <Label className="text-sm text-muted-foreground">IP do servidor</Label>
              <Input className="h-10" placeholder="IP do servidor" />
            </div>
          </div>
        </CardContent>
        <CardFooter className="border-t flex items-center justify-end space-x-2  p-4 px-6">
          <Button variant={"outline"} size={"sm"}>Cancelar</Button>
          <Button size={"sm"}>Salvar alterações</Button>
        </CardFooter>
      </Card>
      <Card>
        <CardHeader className="border-b ">
          <CardTitle>Zona de Perigo</CardTitle>
        </CardHeader>
        <CardContent className="p-6">
          <Alert variant={"warn"}>
            <AlertCircle className="h-4 w-4" />
            <AlertTitle>Desativar servidor</AlertTitle>
            <AlertDescription>
              Ao desativar esse servidor, todas operações serão automaticamente canceladas.
            </AlertDescription>
            <div className="mt-2">
              <AlertDialog>
                <AlertDialogTrigger asChild>
                  <Button variant="warn" size={"sm"}>Desativar servidor</Button>
                </AlertDialogTrigger>
                <AlertDialogContent>
                  <AlertDialogHeader>
                    <AlertDialogTitle>Tem certeza absoluta?</AlertDialogTitle>
                    <AlertDialogDescription>
                      Ao desativar esse servidor, todas operações serão automaticamente canceladas.
                    </AlertDialogDescription>
                  </AlertDialogHeader>
                  <div className="space-y-1">
                    <p className="text-sm text-foreground">Para confirmar digite <span className="font-semibold text-foreground">"Copacabana Roleplay"</span> no campo abaixo.</p>
                    <Input className="h-10" />
                  </div>
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