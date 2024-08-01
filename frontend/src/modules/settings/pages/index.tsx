import Header from "@/components/header";
import MaxWidthWrapper from "@/components/max-width-wrapper";
import { Button } from "@/components/ui/button";
import { cn } from "@/lib/utils";
import IServer from "@/types/server";
import { ArrowLeft } from "lucide-react";
import { Link, Outlet, useLoaderData, useLocation } from "react-router-dom";

export default function Settings() {
  const loader = useLoaderData()
  const servers: IServer[] = (loader as any).data

  const location = useLocation()

  return (
    <MaxWidthWrapper>
      <Header />
      <div className="container mx-auto p-0 border-l border-r flex flex-1">
        <nav className="max-w-xs py-8 hidden lg:flex flex-col w-full bg-card border-r content-start">
          <ul className="flex flex-1 fixed px-8 flex-col space-y-4">
            <Button className="w-max" size={"sm"} asChild variant={"outline"}>
              <Link to={"/dashboard"}>
                <ArrowLeft size={16} />
                Voltar
              </Link>
            </Button>
            <div>
              <h3 className="font-bold text-sm">Geral</h3>
              <ul className="flex flex-col gap-1 px-2">
                <Link to={"/settings/organization/general"}>
                  <li className="flex py-2 items-center gap-2 text-xs text-muted-foreground hover:text-foreground transition">
                    Configurações da conta
                  </li>
                </Link>
              </ul>
            </div>
            <div>
              <h3 className="font-bold text-sm">Servidores</h3>
              <ul className="flex flex-col px-2">
                {servers.map(server => (
                  <Link key={server.id} to={`/settings/servers/${server.id}/general`}>
                    <li className={cn("flex py-2 items-center gap-2 text-xs text-muted-foreground transition", {
                      "text-primary": server.id === location.pathname.split("/")[3],
                      "hover:text-foreground": server.id !== location.pathname.split("/")[3]
                    })}>
                      {server.name}
                    </li>
                  </Link>
                ))}
              </ul>
            </div>
          </ul>
        </nav>
        <div className="w-full h-full">
          <Outlet />
        </div>
      </div>
    </MaxWidthWrapper>
  )
}