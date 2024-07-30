import Grid from "@/components/grid";
import MaxWidthWrapper from "@/components/max-width-wrapper";
import { Button } from "@/components/ui/button";

export default function Confirm() {
  return (
    <MaxWidthWrapper className="flex-row">
      <section className="flex-1 h-screen container mx-auto hidden lg:flex sticky top-0 items-center justify-center">
        <Grid className="absolute z-0 w-3/4 object-cover bg-auto h-full"/>
        <div className="flex flex-col space-y-4 z-10 w-fit max-w-xl">
          <h1 className="text-7xl">Plataforma <br />de <b className="text-primary tracking-tighter">FiveM</b></h1>
          <p className="text-md opacity-60">Não perca um único evento. A <b>Fivetrace</b> te ajuda nisso.</p>
        </div>
      </section>
      <section className="container mx-auto w-full max-w-xl py-12 flex flex-col items-center justify-center gap-12 bg-container border-x border-foreground/10">
        <div className="w-full max-w-sm text-center space-y-8">
          <div className="space-y-2">
            <h1 className="text-3xl">Parabéns!</h1>
            <p className="text-muted-foreground">Seu cadastro foi concluído com <span className="text-primary font-semibold">sucesso</span>, para acessar a plataforma basta clicar no botão abaixo.</p>
          </div>
          <Button className="w-full h-12">Entrar na plataforma</Button>
        </div>
      </section>
    </MaxWidthWrapper>
  )
}