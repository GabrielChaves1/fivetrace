import { ThemeProvider } from "@/components/theme-provider"
import { Outlet } from "react-router-dom"

function App() {
  return (
    <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
      <div className="w-screen h-full bg-background overflow-x-hidden font-sans antialiased">
        <Outlet />
      </div>
    </ThemeProvider>
  )
}

export default App
