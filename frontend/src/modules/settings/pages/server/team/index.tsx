import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { faker } from "@faker-js/faker"
import { columns } from "./columns"
import { DataTable } from "./data-table"

export type Member = {
  sub: string
  avatar: string
  username: string
  role: 'manager' | 'member'
  email: string
}
 
const payments = (): Member[] => {
  return Array.from({ length: 6 }).map((_, i) => ({
    sub: faker.string.uuid(),
    avatar: faker.image.avatar(),
    username: faker.person.fullName(),
    role: 'member',
    email: faker.internet.email().toLowerCase()
  }))
}

export default function ServerSettingsTeam() {
  const data = payments()

  return (
    <div className="flex flex-col gap-2">
      <div className="flex space-x-2 justify-between">
        <Input placeholder="Pesquisar usuário" className="h-8 text-xs" />
        <Button className="h-8 text-xs">Convidar</Button>
      </div>
      <DataTable columns={columns} data={data} />
    </div>
  )
}