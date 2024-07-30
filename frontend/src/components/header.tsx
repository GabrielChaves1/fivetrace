import NotificationWidget from "./notification-widget";
import UserDropdown from "./user-dropdown";

export default function Header() {
  return (
    <div className="w-full py-6 z-50 sticky top-0 bg-container border-b border-foreground/10 shadow-md shadow-foreground/[2%]">
      <div className="h-full container mx-auto flex items-center justify-between">
        <h1 className="text-3xl font-bold tracking-tighter">five.<span className="text-primary">trace</span></h1>
        <div className="flex items-center space-x-3">
          <NotificationWidget />
          <UserDropdown />
        </div>
      </div>
    </div>
  )
}