import Grid from "@/components/grid";
import MaxWidthWrapper from "@/components/max-width-wrapper";
import Stepper from "@/components/stepper";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";

import { Button } from "@/components/ui/button";
import {
  Form, FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { zodResolver } from "@hookform/resolvers/zod";
import { ArrowRight } from "lucide-react";
import { useForm } from "react-hook-form";
import { z } from "zod";

const formSchema = z.object({
  image: z.string({ message: "A imagem do servidor é obrigatória" }).url(),
  serverName: z.string({ message: "O nome do servidor é obrigatório" }).min(1),
  serverIp: z.string({ message: "O endereço IP do servidor é obrigatório" }).min(1),
})

export default function Welcome() {
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
  })

  const onSubmit = (data: z.infer<typeof formSchema>) => {
    console.log(data)
  }

  return (
    <MaxWidthWrapper className="flex-col items-center justify-center">
      <Grid className="absolute left-2/4 top-2/4 w-screen object-contain translate-x-[-50%] translate-y-[-50%] opacity-30" />
      <div className="z-10 flex flex-col items-center px-8 gap-5">

        <div className="space-y-1 text-center">
          <h1 className="text-2xl desktop:text-3xl font-bold">Bem-vindo a <span className="tracking-tighter">five.<span className="text-primary">trace</span></span>!</h1>
          <p className="text-sm desktop:text-md text-muted-foreground">Para começar vamos configurar seu primeiro servidor</p>
        </div>

        <Stepper currentStep={1} steps={2}/>

        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4 flex flex-col container mx-auto">
            <FormField
              control={form.control}
              name="image"
              render={({ field }) => (
                <FormItem >
                  <FormControl>
                    <div className="flex items-center space-x-4">
                      <Avatar className="w-20 h-20 rounded-md">
                        <AvatarFallback></AvatarFallback>
                        <AvatarImage className="rounded-sm" src={field.value || "https://avatars.githubusercontent.com/u/37236608?v=4"} alt="Icone do servidor" />
                      </Avatar>
                      <Button size={"sm"} variant={"outline"} type="button">Selecionar Imagem</Button>
                    </div>
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="serverName"
              render={({ field }) => (
                <FormItem >
                  <FormLabel>Nome do servidor</FormLabel>
                  <FormControl>
                    <Input type="text" placeholder="Nome do servidor" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="serverIp"
              render={({ field }) => (
                <FormItem >
                  <FormLabel>IP do servidor</FormLabel>
                  <FormControl>
                    <Input type="text" placeholder="0.0.0.0" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <Button type="submit" className="self-end">
              Próximo
              <ArrowRight size={14}/>
            </Button>
          </form>
        </Form>
      </div>
    </MaxWidthWrapper>
  )
}