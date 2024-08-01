import NotificationWidget from "./notification-widget";
import UserDropdown from "./user-dropdown";

export default function Header() {
  return (
    <div className="w-full z-50 sticky top-0 bg-card border-b">
      <div className="h-16 container mx-auto flex items-center border-l border-r justify-between">
        <h1 className="text-3xl font-bold tracking-tighter">five.<span className="text-primary">trace</span></h1>
        <div className="flex items-center space-x-3">
          <NotificationWidget />
          <UserDropdown />
        </div>
      </div>
    </div>
  )
}