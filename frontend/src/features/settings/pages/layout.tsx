import Header from "@/components/header";
import MaxWidthWrapper from "@/components/max-width-wrapper";
import Navbar from "@/components/nav-bar";
import { Outlet } from "react-router-dom";

export default function Settings() {
  return (
    <MaxWidthWrapper>
      <Header />
      <div className="container mx-auto p-0 border-l border-r border-foreground/10 flex flex-1">
        <Navbar />
        <div className="w-full h-full p-5 px-8">
          <Outlet />
        </div>
      </div>
    </MaxWidthWrapper>
  )
}