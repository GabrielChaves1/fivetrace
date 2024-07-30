import { Link } from "react-router-dom";
import { Avatar, AvatarFallback, AvatarImage } from "./ui/avatar";
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator, DropdownMenuTrigger } from "./ui/dropdown-menu";

export default function UserDropdown() {
  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Avatar>
          <AvatarFallback>GC</AvatarFallback>
          <AvatarImage src="/example.png" alt="Guilherme Carvalho" />
        </Avatar>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end" sideOffset={10}>
        <DropdownMenuLabel>Organização</DropdownMenuLabel>
        <DropdownMenuSeparator />
        <DropdownMenuItem>Meu Perfil</DropdownMenuItem>
        <DropdownMenuItem asChild>
          <Link to="/settings">
            Configurações
          </Link>
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}