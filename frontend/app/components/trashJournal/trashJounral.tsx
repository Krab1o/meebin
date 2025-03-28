import { Card, CardHeader, CardContent, CardFooter } from "~/components/ui/card";
import { Button } from "~/components/ui/button";
import { ScrollArea } from "~/components/ui/scroll-area";
import { useEffect, useState } from "react";
import { motion, AnimatePresence } from "framer-motion";

interface CleanupRequest {
  id: number;
  address: string;
  imageUrl: string;
  accepted: boolean;
  acceptedBy?: string;
}

const randomNames = [
  "Алексей Петров",
  "Мария Иванова",
  "Дмитрий Смирнов",
  "Елена Кузнецова",
  "Иван Попов",
  "Ольга Васильева",
  "Сергей Соколов",
  "Анна Михайлова",
  "Андрей Новиков",
  "Наталья Федорова"
];

const generateRandomAddress = () => {
  const streets = [
    "Московская", "Рахова", "Кирова", "Чернышевского", "Горького",
    "Астраханская", "Соколовая", "Лермонтова", "Чапаева", "Б. Садовая"
  ];
  return `ул. ${streets[Math.floor(Math.random() * streets.length)]}, ${
    Math.floor(Math.random() * 100) + 1
  }, Саратов`;
};

const getRandomLocalImage = () => {
  const imageCount = 5;
  const randomIndex = Math.floor(Math.random() * imageCount) + 1;
  return `/trash-images/trash${randomIndex}.jpg`;
};

const CleanupRequestsPage = () => {
  const [requests, setRequests] = useState<CleanupRequest[]>(() =>
    Array.from({ length: 5 }, (_, i) => ({
      id: i + 1,
      address: generateRandomAddress(),
      imageUrl: getRandomLocalImage(),
      accepted: false
    }))
  );
  const [nextId, setNextId] = useState(8);

  const handleAcceptRequest = (id: number) => {
    // Сначала помечаем как принятую
    setRequests(requests.map(request => 
      request.id === id 
        ? { ...request, accepted: true, acceptedBy: "Вы" }
        : request
    ));

    // Через 1 секунду заменяем на новую
    setTimeout(() => {
      setRequests(prev => [
        ...prev.filter(r => r.id !== id),
        {
          id: nextId,
          address: generateRandomAddress(),
          imageUrl: getRandomLocalImage(),
          accepted: false
        }
      ]);
      setNextId(prev => prev + 1);
    }, 1000);
  };

  useEffect(() => {
    const interval = setInterval(() => {
      setRequests(prev => {
        const availableRequests = prev.filter(r => !r.accepted);
        if (availableRequests.length === 0) return prev;

        const randomRequestIndex = Math.floor(Math.random() * availableRequests.length);
        const randomRequest = availableRequests[randomRequestIndex];
        const randomName = randomNames[Math.floor(Math.random() * randomNames.length)];

        // Помечаем как принятую
        const updated = prev.map(request =>
          request.id === randomRequest.id
            ? { ...request, accepted: true, acceptedBy: randomName }
            : request
        );

        // Через 1 секунду заменяем на новую
        setTimeout(() => {
          setRequests(prevRequests => [
            ...prevRequests.filter(r => r.id !== randomRequest.id),
            {
              id: nextId,
              address: generateRandomAddress(),
              imageUrl: getRandomLocalImage(),
              accepted: false
            }
          ]);
          setNextId(prev => prev + 1);
        }, 1000);

        return updated;
      });
    }, Math.random() * 15000 + 10000);

    return () => clearInterval(interval);
  }, [nextId]);

  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-3xl font-bold mb-6">Заявки на уборку мусора</h1>
      <ScrollArea className="h-[600px] rounded-md border p-4">
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          {requests.map((request) => (
            <motion.div
              key={request.id}
              layout
              initial={{ opacity: 0 }}
              animate={{ opacity: 1 }}
              exit={{ opacity: 0 }}
              transition={{ duration: 0.3 }}
              className={request.accepted ? "opacity-50" : ""}
            >
              <Card className="hover:shadow-lg transition-shadow">
                <CardHeader>
                  <h3 className="text-xl font-semibold">{request.address}</h3>
                </CardHeader>
                <CardContent>
                  <img
                    src={request.imageUrl}
                    alt={`Мусор по адресу ${request.address}`}
                    className="w-full h-48 object-cover rounded-md"
                  />
                </CardContent>
                <CardFooter className="flex flex-col gap-2">
                  {request.accepted ? (
                    <div className="w-full p-2 bg-green-50 text-green-800 rounded text-center">
                      {request.acceptedBy === "Вы"
                        ? "Вы приняли эту заявку"
                        : `Заявка принята: ${request.acceptedBy}`}
                    </div>
                  ) : (
                    <Button
                      onClick={() => handleAcceptRequest(request.id)}
                      className="w-full"
                    >
                      Принять заявку
                    </Button>
                  )}
                </CardFooter>
              </Card>
            </motion.div>
          ))}
        </div>
      </ScrollArea>
    </div>
  );
};

export default CleanupRequestsPage;