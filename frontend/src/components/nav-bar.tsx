import { NAVIGATION_ITEMS, NavigationItem } from "@/config/navigation";
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
    <div className="w-full hidden desktop:flex sticky inset-0 flex-col justify-between bg-container/50 max-w-xs border-r border-foreground/10">
      <nav className="flex flex-col">
        <ul className="flex flex-1 flex-col">
          {Object.keys(groupedItems).map(category => (
            <div className="space-y-3 p-6 px-8 border-b border-foreground/10" key={category}>
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
        </ul>
      </nav>
    </div>
  )
}