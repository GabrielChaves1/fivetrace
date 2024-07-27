import { ArrowTopRightIcon } from "@radix-ui/react-icons";
import { ElementType } from "react";

export interface NavigationItem {
  title: string
  to: string
  category: string
  icon: ElementType
}

export const NAVIGATION_ITEMS = [
  {
    title: "Página Inicial",
    to: "/home",
    category: "Geral",
  },
  {
    title: "Configurações",
    to: "/settings",
    category: "Geral",
  },
  {
    title: "Guia",
    to: "/settings",
    category: "Documentação",
    icon: ArrowTopRightIcon,
  },
  {
    title: "API Reference",
    to: "/settings",
    category: "Documentação",
    icon: ArrowTopRightIcon,
  },
];