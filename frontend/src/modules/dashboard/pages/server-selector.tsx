import Header from "@/components/header";
import MaxWidthWrapper from "@/components/max-width-wrapper";
import { Plus } from "lucide-react";
import { useLoaderData } from "react-router-dom";
import IServer from "../../../types/server";
import ServerBox from "../components/server-box";

export default function Dashboard() {
  const loader: any = useLoaderData()
  const servers: IServer[] = loader.data

  return (
    <MaxWidthWrapper>
      <Header />
      <div className="container mx-auto flex-1 space-y-6 py-10 border-l border-r">
        <div className="leading-none">
          <h2 className="text-lg font-bold">Seus servidores</h2>
          <p className="text-sm text-muted-foreground">Lista de todos servidores cadastrados por você na plataforma.</p>
        </div>
        <div className="grid grid-cols-1 md:grid-cols-2 desktop:grid-cols-3 gap-6 w-full content-start">
          {servers.map((server) => (
            <ServerBox key={server.id} {...server}/>
          ))}
          <div className="border p-4 group border-dashed  hover:border-primary/50 cursor-pointer rounded-md grid place-items-center">
            <div className="p-2 group-hover:border-primary/50 grid place-items-center border rounded-full">
              <Plus />
            </div>
          </div>
        </div>
      </div>
    </MaxWidthWrapper>
  )
}