import Header from "@/components/header";
import MaxWidthWrapper from "@/components/max-width-wrapper";
import { Plus } from "lucide-react";
import Server from "../components/server";

export default function Dashboard() {
  return (
    <MaxWidthWrapper>
      <Header />
      <div className="container mx-auto border-foreground/10 flex-1 space-y-6 py-10">
        <div className="leading-none">
          <h2 className="text-lg font-bold">Seus servidores</h2>
          <p className="text-sm text-muted-foreground">Lista de todos servidores cadastrados por você na plataforma.</p>
        </div>
        <div className="grid grid-cols-1 md:grid-cols-2 desktop:grid-cols-3 gap-6 w-full content-start">
          <Server />
          <div className="border p-4 group border-dashed border-foreground/10 hover:border-primary/50 transition cursor-pointer rounded-md grid place-items-center">
            <div className="p-2 group-hover:border-primary/50 border-foreground/10 transition grid place-items-center border rounded-full">
              <Plus />
            </div>
          </div>
        </div>
      </div>
    </MaxWidthWrapper>
  )
}