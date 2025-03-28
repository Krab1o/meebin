import * as React from "react"
import { Moon, Sun } from "lucide-react"
import { useTheme } from "~/components/theme-provider"
import { Button } from "~/components/ui/button"

export function ModeToggle() {
  const { theme, setTheme } = useTheme()
  const [mounted, setMounted] = React.useState(false)

  // Эффект для определения, что компонент смонтирован (работает в браузере)
  React.useEffect(() => {
    setMounted(true)
  }, [])

  if (!mounted) {
    // Рендерим заглушку на сервере
    return (
      <Button variant="ghost" size="sm" className="w-full justify-start gap-3">
        <Sun className="h-4 w-4 opacity-0" />
        <span className="opacity-0">Загрузка...</span>
      </Button>
    )
  }

  return (
    <Button
      variant="ghost"
      size="sm"
      onClick={() => setTheme(theme === "dark" ? "light" : "dark")}
      className="w-full justify-start gap-3"
    >
      {theme === "dark" ? (
        <Sun className="h-4 w-4" />
      ) : (
        <Moon className="h-4 w-4" />
      )}
      <span>{theme === "dark" ? "Светлая тема" : "Тёмная тема"}</span>
    </Button>
  )
}