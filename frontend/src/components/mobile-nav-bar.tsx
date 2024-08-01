import { NAVIGATION_ITEMS, NavigationItem } from "@/config/navigation";
import { LogOut, MenuIcon } from "lucide-react";
import { Link } from "react-router-dom";
import NotificationWidget from "./notification-widget";
import { Avatar, AvatarFallback, AvatarImage } from "./ui/avatar";
import { Button } from "./ui/button";
import { Sheet, SheetContent, SheetTrigger } from "./ui/sheet";

export default function MobileNavbar() {
  const groupedItems = NAVIGATION_ITEMS.reduce((acc: any, item) => {
    const { category } = item;
    if (!acc[category]) {
      acc[category] = [];
    }
    acc[category].push(item);
    return acc;
  }, [])

  return (
    <div className="bg-container/50 desktop:hidden">
      <nav className="flex items-center justify-between px-10 border-b  gap-10">
        <Sheet>
          <SheetTrigger asChild>
            <Button variant={"ghost"} size={"icon"}>
              <MenuIcon size={20} />
            </Button>
          </SheetTrigger>
          <SheetContent side={"left"}>
            <nav className="w-full">
              {Object.keys(groupedItems).map(category => (
                <div className="space-y-3 p-6 border-b " key={category}>
                  <h3 className="font-bold text-sm">{category}</h3>
                  <ul className="flex flex-col gap-1">
                    {groupedItems[category]?.map(({ icon: Icon, title, to }: NavigationItem) => (
                      <Link to={to} key={to}>
                        <li className="flex items-center h-8 gap-2 text-sm text-foreground/60 hover:text-foreground/80 transition">
                          {Icon && <Icon size={14} />}
                          {title}
                        </li>
                      </Link>
                    ))}
                  </ul>
                </div>
              ))}
              <div className="p-6 border-b ">
                <ul className="flex flex-col gap-1">
                  <Link to="/logout">
                    <li className="flex items-center gap-2 text-sm text-foreground/60 hover:text-foreground/80 transition">
                      <LogOut size={14} />
                      Logout
                    </li>
                  </Link>
                </ul>
              </div>
            </nav>
          </SheetContent>
        </Sheet>
        <div className="h-20 grid place-items-center">
          <h1 className="text-2xl sm:text-3xl font-bold tracking-tighter">five.<span className="text-primary">trace</span></h1>
        </div>
        <div className="flex space-x-2 items-center">
          <NotificationWidget />
          <Avatar>
            <AvatarFallback>GC</AvatarFallback>
            <AvatarImage src="/example.png" alt="Guilherme Carvalho" />
          </Avatar>
        </div>
      </nav>
    </div>
  )
}