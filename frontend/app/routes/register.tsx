import { Button } from "~/components/ui/button";
import { Input } from "~/components/ui/input";
import { Calendar } from "~/components/ui/calendar";
import { CalendarIcon } from "lucide-react";
import { Popover, PopoverContent, PopoverTrigger } from "~/components/ui/popover";
import { Link } from "react-router-dom";

export default function RegisterPage() {
  return (
    <div className="flex h-full w-full items-center justify-center bg-gray-50 p-4 absolute">
      <div className="w-full max-w-md space-y-6 rounded-lg bg-white p-8 shadow-md">
        <h1 className="text-center text-2xl font-bold">Регистрация</h1>
        
        <div className="space-y-4">
          <div>
            <label className="block text-sm font-medium text-gray-700">Email</label>
            <Input type="email" className="mt-1" />
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700">Дата рождения</label>
            <Popover>
              <PopoverTrigger asChild>
                <Button
                  variant={"outline"}
                  className="w-full justify-start text-left font-normal"
                >
                  <CalendarIcon className="mr-2 h-4 w-4" />
                  Выберите дату
                </Button>
              </PopoverTrigger>
              <PopoverContent className="w-auto p-0">
                <Calendar mode="single" className="rounded-md border" />
              </PopoverContent>
            </Popover>
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700">Город</label>
            <Input type="text" className="mt-1" />
          </div>

          <div className="grid grid-cols-2 gap-4">
            <div>
              <label className="block text-sm font-medium text-gray-700">Имя</label>
              <Input type="text" className="mt-1" />
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700">Фамилия</label>
              <Input type="text" className="mt-1" />
            </div>
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700">Отчество (необязательно)</label>
            <Input type="text" className="mt-1" />
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700">Пароль</label>
            <Input type="password" className="mt-1" />
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700">Подтверждение пароля</label>
            <Input type="password" className="mt-1" />
          </div>

          <Link
            to="/" 
            className="text-sm font-medium text-blue-600 hover:text-blue-500 hover:underline">  
                <Button className="w-full">Зарегистрироваться</Button>
          </Link>
        </div>
      </div>
    </div>
  );
}