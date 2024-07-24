import CountrySelector from "@/components/country-selector";
import Grid from "@/components/grid";
import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import api from "@/services/axios";
import { zodResolver } from "@hookform/resolvers/zod";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { schema } from "./schemas";

export default function Signup() {
  const [emailSent, setEmailSent] = useState(false);

  const form = useForm<z.infer<typeof schema>>({
    resolver: zodResolver(schema),
    defaultValues: {
      country: "br"
    }
  })

  const onSubmit = async (values: z.infer<typeof schema>) => {
    console.log("Cadastro realizado com sucesso!");
    const res = await api.post("/signup", {
      ...values
    })
    console.log(res);

    setEmailSent(true);
  }

  return (
    <main className="w-screen min-h-screen flex">
      <section className="flex-1 h-screen container mx-auto hidden lg:flex sticky top-0 items-center justify-center">
        <Grid className="absolute z-0 w-3/4 object-cover bg-auto h-full"/>
        <div className="flex flex-col space-y-4 z-10 w-fit max-w-xl">
          <h1 className="text-7xl">Plataforma <br />de <b className="text-primary tracking-tighter">FiveM</b></h1>
          <p className="text-md opacity-50">Não perca um único evento. A <b>Luminog</b> te ajuda nisso.</p>
        </div>
      </section>
      {emailSent ? (
        <div className="container mx-auto w-full max-w-xl py-12 flex flex-col items-center justify-center gap-8 bg-container border-x border-foreground/10">
          {/* <img src="/logo.svg" /> */}
          <div className="space-y-10 w-fit max-w-xs text-center">
            <p className="text-md text-foreground/75">Estamos felizes em informar que sua conta foi cadastrada com <span className="text-primary font-semibold">sucesso</span> em nossa plataforma!</p>
            <hr className="border-foreground/10" />
            <p className="text-md text-foreground/75">Para completar o processo de cadastro e ativar sua conta, por favor verifique sua caixa de entrada de email. Enviamos um <span className="text-primary font-semibold">link de confirmação</span> para o endereço de email fornecido durante o cadastro.</p>
          </div>
        </div>
      ) : (
        <section className="container mx-auto w-full max-w-xl py-12 flex flex-col items-center justify-center gap-12 bg-container border-x border-foreground/10">
          {/* <img src="/logo.svg" /> */}
          <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-10 w-full max-w-sm">
              <div className="space-y-4">
                <FormField
                  control={form.control}
                  name="organization"
                  render={({ field }) => (
                    <FormItem className="mb-8">
                      <FormLabel>Organização</FormLabel>
                      <FormControl>
                        <Input placeholder="Nome da organização" {...field} />
                      </FormControl>
                      <FormDescription>
                        O nome da sua organização pode ser alterado.
                      </FormDescription>
                      <FormMessage />
                    </FormItem>
                  )}
                />
                <FormField
                  control={form.control}
                  name="email"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>E-mail</FormLabel>
                      <FormControl>
                        <Input type="email" placeholder="Seu e-mail" {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                <FormField
                  control={form.control}
                  name="password"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Senha</FormLabel>
                      <FormControl>
                        <Input type="password" placeholder="Sua senha" {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
                <FormField
                  control={form.control}
                  name="confirmPassword"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Confirme sua senha</FormLabel>
                      <FormControl>
                        <Input type="password" placeholder="Sua senha" {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                <FormField
                  control={form.control}
                  name="country"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>País</FormLabel>
                      <FormControl>
                        <CountrySelector />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </div>
              <div className="flex flex-col items-center space-y-4">
                <Button type="submit" className="h-12 w-full font-medium">
                  {/* <Loader /> */}
                  Cadastrar-se
                </Button>
                <span className="text-foreground/60 text-sm">
                  Já possui uma conta?{" "}
                  <a className="text-primary underline cursor-pointer hover:text-primary/80 transition-colors">Entre agora</a>
                </span>
              </div>
            </form>
          </Form>
          {/* <form className="flex flex-col w-full px-20 space-y-10 container mx-auto" onSubmit={handleSubmit(onSubmit)}>
          <fieldset>
            <div className="space-y-2 flex flex-col">
              <label htmlFor="organization" className="opacity-75 text-sm">Nome da organização</label>
              <Input type="text" id="organization" {...register("organization")} placeholder="Nome da sua organização" />
              <p className="text-destructive/80 text-sm">{errors.organization?.message}</p>
            </div>
          </fieldset>
          <fieldset className="space-y-4">
            <div className="space-y-2 flex flex-col">
              <label htmlFor="email" className="opacity-75 text-sm">E-mail</label>
              <Input type="email" id="email" {...register("email")} placeholder="Seu e-mail" />
              <p className="text-destructive/80 text-sm">{errors.organization?.message}</p>
            </div>
            <div className="space-y-2 flex flex-col">
              <label htmlFor="password" className="opacity-75 text-sm">Senha</label>
              <Input type="password" id="password" {...register("password")} placeholder="Sua senha" />
              <p className="text-destructive/80 text-sm">{errors.organization?.message}</p>
            </div>
            <div className="space-y-2 flex flex-col">
              <label htmlFor="confirm-password" className="opacity-75 text-sm">Confirmar Senha</label>
              <Input type="password" id="confirm-password" {...register("confirmPassword")} placeholder="Sua senha" />
              <p className="text-destructive/80 text-sm">{errors.organization?.message}</p>
            </div>
          </fieldset>
          <div className="flex flex-col items-center space-y-4">
            <button type="submit" className="h-12 w-full bg-primary/45 border border-primary/50 font-medium rounded-md transition-colors hover:bg-primary/50 hover:border-primary">
              Cadastrar-se
            </button>
            <span className="text-foreground/60 text-sm">
              Já possui uma conta?{" "}
              <a className="text-primary underline cursor-pointer hover:text-primary/80 transition-colors">Entre agora</a>
            </span>
          </div>
        </form> */}
        </section>
      )}
    </main>
  )
}