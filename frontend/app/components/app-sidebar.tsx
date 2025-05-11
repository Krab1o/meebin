import { Home, FilePlus2, User, Info } from "lucide-react"
import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarFooter,
  SidebarTrigger
} from "~/components/ui/sidebar"
import { Link } from "react-router"
import { ModeToggle } from "~/components/mode-toggle"

export function AppSidebar() {
  const navItems = [
    {
      title: "Главная",
      path: "/home",
      icon: Home,
    },
    {
      title: "Создать заявку",
      path: "/create",
      icon: FilePlus2,
    },
    {
      title: "Профиль",
      path: "/profile",
      icon: User,
    },
    {
      title: "О нас",
      path: "/about",
      icon: Info,
    },
  ]

  return (
    <Sidebar className="w-64 border-r"> 
        <SidebarTrigger className="absolute left-64" /> 
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel>Меню</SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu>
              {navItems.map((item) => (
                <SidebarMenuItem key={item.path}>
                  <SidebarMenuButton asChild>
                    <Link 
                      to={item.path}
                      className="flex items-center gap-3 p-2 hover:bg-accent transition-colors"
                    >
                      <item.icon className="w-5 h-5" />
                      <span>{item.title}</span>
                    </Link>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
      
      <SidebarFooter className="p-4 border-t">
        <ModeToggle />
      </SidebarFooter>
    </Sidebar>
  )
}