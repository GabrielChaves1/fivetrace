import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import { Badge } from "@/components/ui/badge"
import { ColumnDef } from "@tanstack/react-table"
import { Member } from "."
 
export const columns: ColumnDef<Member>[] = [
  {
    accessorKey: "username",
    header: "Usuário",
    cell: ({ row }) => (
      <div className="flex items-center space-x-3 text-muted-foreground">
        <Avatar className="w-8 h-8">
          <AvatarFallback>{row.original.username[0]}</AvatarFallback>
          <AvatarImage src={row.original.avatar} alt={row.original.username} />
        </Avatar>
        <span>{row.original.username}</span>
      </div>
    )
  },
  {
    accessorKey: "email",
    header: "Email",
    cell: ({ row }) => (
      <span className="text-muted-foreground">{row.original.email}</span>
    )
  },
  {
    accessorKey: "role",
    header: "Cargo",
    cell: ({ row }) => (
      <Badge variant={"outline"} className="text-muted-foreground">
        {row.original.role === 'manager' ? 'Gerente' : 'Membro'}
      </Badge>
    )
  },
  
]