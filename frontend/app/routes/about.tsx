import { Card, CardHeader, CardTitle, CardDescription, CardContent } from "~/components/ui/card"
import { Avatar, AvatarImage, AvatarFallback } from "~/components/ui/avatar"
import { Carousel, CarouselContent, CarouselItem } from "~/components/ui/carousel"
import { Badge } from "~/components/ui/badge"
import Autoplay from "embla-carousel-autoplay"
import React from "react"



const technologies = [
  {
    name: "React",
    description: "Библиотека для построения пользовательских интерфейсов",
    icon: "react-icon.png"
  },
  {
    name: "TypeScript",
    description: "Статическая типизация для JavaScript",
    icon: "typescript-icon.png"
  },
  {
    name: "React Router 7",
    description: "Маршрутизация в React-приложениях",
    icon: "react-router-icon.png"
  },
  {
    name: "shadcn/ui",
    description: "Современные компоненты интерфейса",
    icon: "shadcn-icon.png"
  },
  {
    name: "Vite",
    description: "Быстрый сборщик проектов",
    icon: "vite-icon.png"
  }
]

export default function AboutPage() {
  const plugin = React.useRef(
    Autoplay({ delay: 3000, stopOnInteraction: false })
  )

  return (
    <div className="container mx-auto px-4 py-12">
      <Card className="max-w-3xl mx-auto">
        <CardHeader className="text-center">
          <Avatar className="w-24 h-24 mx-auto mb-4">
            <AvatarImage src="/avatar.jpg" />
            <AvatarFallback>BU</AvatarFallback>
          </Avatar>
          <CardTitle className="text-3xl">Богдан Устюшин</CardTitle>
          <CardDescription className="text-xl">
            Дипломная работа
          </CardDescription>
        </CardHeader>

        <CardContent className="space-y-8">
          <section>
            <h2 className="text-2xl font-semibold mb-4">О проекте</h2>
            <p className="text-lg">
              Это приложение для создания заявок на очистку мусора на улицах города.
              Пользователи могут отмечать проблемные места на карте, оставлять описания
              и отслеживать статус выполнения заявок.
            </p>
          </section>

          <section>
            <h2 className="text-2xl font-semibold mb-4">Используемые технологии</h2>
            <Carousel
              plugins={[plugin.current]}
              className="w-full"
              opts={{
                align: "start",
                loop: true,
              }}
            >
              <CarouselContent>
                {technologies.map((tech) => (
                  <CarouselItem key={tech.name} className="md:basis-1/2 lg:basis-1/3">
                    <Card className="h-full">
                      <CardHeader>
                        <div className="flex items-center space-x-4">
                          <Avatar>
                            <AvatarImage src={`/icons/${tech.icon}`} />
                            <AvatarFallback>{tech.name[0]}</AvatarFallback>
                          </Avatar>
                          <CardTitle>{tech.name}</CardTitle>
                        </div>
                      </CardHeader>
                      <CardContent>
                        <p>{tech.description}</p>
                      </CardContent>
                    </Card>
                  </CarouselItem>
                ))}
              </CarouselContent>
            </Carousel>
          </section>

          <section>
            <h2 className="text-2xl font-semibold mb-4">Функционал приложения</h2>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <Badge variant="outline" className="py-2 px-4 text-base w-auto">
                📍 Отметка мест с мусором
              </Badge>
              <Badge variant="outline" className="py-2 px-4 text-base w-auto">
                📷 Загрузка фотографий
              </Badge>
              <Badge variant="outline" className="py-2 px-4 text-base w-auto">
                🔔 Уведомления о статусе
              </Badge>
              <Badge variant="outline" className="py-2 px-4 text-base w-auto">
                🗺️ Интерактивная карта
              </Badge>
            </div>
          </section>
        </CardContent>
      </Card>
    </div>
  )
}