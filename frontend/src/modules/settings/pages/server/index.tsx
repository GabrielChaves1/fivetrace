import { Badge } from "@/components/ui/badge";
import { Tabs, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Link, Outlet, useLocation, useParams } from "react-router-dom";

export default function ServerSettings() {
  const location = useLocation()
  const params = useParams()

  const selectedTab = location.pathname.split("/").pop()

return (
    <div className="flex flex-col h-full gap-8">
      <div className="space-y-4 pt-8 z-10 sticky top-[-1.5rem] bg-background">
        <div className="w-full flex flex-col gap-4 px-8">
          <div className="space-y-0">
            <h1 className="text-md md:text-lg flex items-center font-semibold text-foreground/90">
              Copacabana Roleplay
              <Badge variant={"outline"} className="ml-2 font-medium text-muted-foreground">Plano Enterprise</Badge>
            </h1>
            <p className="text-xs md:text-sm text-muted-foreground">Nesta aba você poderá configurar tudo relacionado ao servidor</p>
          </div>
        </div>
        <Tabs value={selectedTab}>
          <TabsList className="px-8">
            <Link to={`/settings/servers/${params.serverId}/general`}>
              <TabsTrigger value="general">
                Geral
              </TabsTrigger>
            </Link>
            <Link to={`/settings/servers/${params.serverId}/team`}>
              <TabsTrigger value="team">
                Equipe
              </TabsTrigger>
            </Link>
            <Link to={`/settings/servers/${params.serverId}/integration`}>
              <TabsTrigger value="integration">
                Integração
              </TabsTrigger>
            </Link>
          </TabsList>
        </Tabs>
      </div>
      <div className="px-8 space-y-8">
        <Outlet />
      </div>
    </div>
  )
}