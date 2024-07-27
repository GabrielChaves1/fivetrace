import { NAVIGATION_ITEMS, NavigationItem } from "@/config/navigation";
import { LogOut } from "lucide-react";
import { Link } from "react-router-dom";

export default function Navbar() {
  const groupedItems = NAVIGATION_ITEMS.reduce((acc: any, item) => {
    const { category } = item;
    if (!acc[category]) {
      acc[category] = [];
    }
    acc[category].push(item);
    return acc;
  }, [])

  return (
    <div className="w-full hidden desktop:flex flex-col justify-between bg-container/50 max-w-xs border-r border-foreground/10">
      <nav className="flex flex-col">
        <div className="h-20 flex items-center px-6 border-b border-foreground/10">
          <h1 className="text-3xl font-bold tracking-tighter">five.<span className="text-primary">trace</span></h1>
        </div>
        <ul className="flex flex-1 flex-col">
          {Object.keys(groupedItems).map(category => (
            <div className="space-y-3 p-6 border-b border-foreground/10" key={category}>
              <h3 className="font-bold text-sm">{category}</h3>
              <ul className="flex flex-col gap-1">
                {groupedItems[category]?.map(({icon: Icon, title, to}: NavigationItem) => (
                  <Link to={to} key={to}>
                    <li className="flex items-center gap-2 text-sm text-foreground/60 hover:text-foreground/80 transition">
                      {Icon && <Icon size={14}/>}
                      {title}
                    </li>
                  </Link>
                ))}
              </ul>
            </div>
          ))}
          <div className="p-6 border-b border-foreground/10">
            <ul className="flex flex-col gap-1">
              <Link to="/logout">
                <li className="flex items-center gap-2 text-sm text-foreground/60 hover:text-foreground/80 transition">
                  <LogOut size={14}/>
                  Logout
                </li>
              </Link>
            </ul>
          </div>
        </ul>
      </nav>
    </div>
  )
}