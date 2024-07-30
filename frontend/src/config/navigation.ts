import { ElementType } from "react";

export interface NavigationItem {
  title: string
  to: string
  category: string
  icon: ElementType
}

export const NAVIGATION_ITEMS = [
  {
    title: "Configurações da organização",
    to: "/general",
    category: "Geral",
  },
  {
    title: "Copacabana Roleplay",
    to: "/settings/cidade-alta",
    category: "Servidores",
  },
];