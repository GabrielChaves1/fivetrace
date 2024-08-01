import { z } from "zod"

export const schema = z.object({
  organization: z.string({ message: "O nome da organização é obrigatório" }).min(1, { message: "O nome da organização é obrigatório" }).max(20, { message: "A organização deve ter no máximo 20 caracteres" }),
  email: z.string({ message: "O e-mail é obrigatório" }).min(1, { message: "O e-mail é obrigatório" }).email({ message: "O e-mail é inválido" }),
  password: z.string({ message: "A senha é obrigatória" }).min(8, { message: "A senha deve ter no mínimo 8 caracteres" }),
  confirmPassword: z.string({ message: "A confirmação da senha é obrigatória" }).min(8, { message: "A senha deve ter no mínimo 8 caracteres" }),
  country: z.string()
}).required().refine(data => data.password === data.confirmPassword, {
  message: "As senhas não coincidem",
  path: ['confirmPassword']
})