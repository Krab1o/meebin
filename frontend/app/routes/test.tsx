import { Button } from "~/components/ui/button";
import { Input } from "~/components/ui/input";
import { Textarea } from "~/components/ui/textarea";
import { Avatar, AvatarFallback, AvatarImage } from "~/components/ui/avatar";
import { Link } from "react-router-dom";

export default function EventPage() {
  // Моковые данные события
  const eventData = {
    title: "Уборка мусора в парке",
    photos: [
      "/trash-images/trash1.jpg",
      "/trash-images/trash6.jpg",
    ],
    address: "г. Саратов, ул. Московская, 9",
    description: "Уже бегают крысы размером с котов. Пожалуйста уберите. Я не могу это больше выносить.",
    createdAt: "15.05.2023 в 14:30",
    respondedAt: "16.05.2023 в 10:15",
    creator: {
      name: "Иван Петров",
      avatar: "https://example.com/avatar1.jpg"
    },
    responder: {
      name: "Анна Сидорова",
      avatar: "https://example.com/avatar2.jpg"
    },
    comments: [
      {
        id: 1,
        author: "Мария Иванова",
        text: "Я тоже хочу помочь! Можно присоединиться?",
        date: "15.05.2023 в 16:45"
      },
      {
        id: 2,
        author: "Сергей Васильев",
        text: "Отличная инициатива! Я с вами.",
        date: "15.05.2023 в 18:20"
      }
    ]
  };

  const hasResponder = eventData.responder.name !== "";

  return (
    <div className="container mx-auto px-4 py-8 w-[100vw] h-[90vh]">
      <div className="flex-row flex justify-between">
        {/* Заголовок */}
        <h1 className="text-3xl font-bold mb-6">{eventData.title}</h1>
        <Button>Отозваться</Button>
      </div>
      {/* Галерея фото */}
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mb-6">
        {eventData.photos.map((photo, index) => (
          <div key={index} className="rounded-lg overflow-hidden bg-gray-100">
            <img 
              src={photo} 
              alt={`Фото события ${index + 1}`} 
              className="w-full h-48 object-cover"
            />
          </div>
        ))}
      </div>

      {/* Информация о событии */}
      <div className="bg-white rounded-lg shadow p-6 mb-6">
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
          {/* Первая колонка */}
          <div>
            <div className="mb-4">
              <h2 className="text-lg font-semibold mb-2">Адрес</h2>
              <p>{eventData.address}</p>
            </div>

            <div className="mb-4">
              <h2 className="text-lg font-semibold mb-2">Описание</h2>
              <p className="whitespace-pre-line">{eventData.description}</p>
            </div>
          </div>

          {/* Вторая колонка */}
          <div>
            <div className="mb-4">
              <h2 className="text-lg font-semibold mb-2">Дата создания</h2>
              <p>{eventData.createdAt}</p>
            </div>

            <div className="mb-4">
              <h2 className="text-lg font-semibold mb-2">Дата отклика</h2>
              <p>{hasResponder ? eventData.respondedAt : "—"}</p>
            </div>
          </div>

          {/* Третья колонка - пользователи */}
          <div>
            <div className="mb-4">
              <h2 className="text-lg font-semibold mb-2">Организатор</h2>
              <div className="flex items-center gap-2">
                <Avatar>
                  <AvatarImage src={eventData.creator.avatar} />
                  <AvatarFallback>{eventData.creator.name.charAt(0)}</AvatarFallback>
                </Avatar>
                <span>{eventData.creator.name}</span>
              </div>
            </div>

            <div className="mb-4">
              <h2 className="text-lg font-semibold mb-2">Откликнувшийся</h2>
              {hasResponder ? (
                <div className="flex items-center gap-2">
                  <Avatar>
                    <AvatarImage src={eventData.responder.avatar} />
                    <AvatarFallback>{eventData.responder.name.charAt(0)}</AvatarFallback>
                  </Avatar>
                  <span>{eventData.responder.name}</span>
                </div>
              ) : (
                <p>—</p>
              )}
            </div>

            {/* Кнопка отклика */}
            {!hasResponder && (
              <Button className="w-full bg-green-600 hover:bg-green-700">
                Откликнуться
              </Button>
            )}
          </div>
        </div>
      </div>

      {/* Комментарии */}
      <div className="bg-white rounded-lg shadow p-6">
        <h2 className="text-xl font-semibold mb-4">Комментарии ({eventData.comments.length})</h2>

        {/* Список комментариев */}
        <div className="space-y-4 mb-6">
          {eventData.comments.map(comment => (
            <div key={comment.id} className="border-b pb-4">
              <div className="flex justify-between items-start mb-2">
                <div className="flex items-center gap-2">
                  <Avatar className="h-8 w-8">
                    <AvatarFallback>{comment.author.charAt(0)}</AvatarFallback>
                  </Avatar>
                  <span className="font-medium">{comment.author}</span>
                </div>
                <span className="text-sm text-gray-500">{comment.date}</span>
              </div>
              <p>{comment.text}</p>
            </div>
          ))}
        </div>

        {/* Форма добавления комментария */}
        <div>
          <h3 className="text-lg font-medium mb-2">Оставить комментарий</h3>
          <Textarea 
            placeholder="Ваш комментарий..." 
            className="mb-3"
            rows={3}
          />
          <Button>Отправить</Button>
        </div>
      </div>
    </div>
  );
}