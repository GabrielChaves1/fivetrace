import Grid from "@/components/grid";

export default function Confirm() {
  return (
    <main className="w-screen min-h-screen flex">
      <section className="flex-1 h-screen container mx-auto hidden lg:flex sticky top-0 items-center justify-center">
        <Grid className="absolute z-0 w-3/4 object-cover bg-auto h-full"/>
        <div className="flex flex-col space-y-4 z-10 w-fit max-w-xl">
          <h1 className="text-7xl">Plataforma <br />de <b className="text-primary tracking-tighter">FiveM</b></h1>
          <p className="text-md opacity-60">Não perca um único evento. A <b>Luminog</b> te ajuda nisso.</p>
        </div>
      </section>
      <section className="container mx-auto w-full max-w-xl py-12 flex flex-col items-center justify-center gap-12 bg-container border-x border-foreground/10">
        
      </section>
    </main>
  )
}