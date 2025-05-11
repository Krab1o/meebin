import { Button } from "~/components/ui/button";
import { Input } from "~/components/ui/input";
import { Link } from "react-router-dom";

export default function LoginPage() {
  return (
    <div className="flex h-full w-full items-center justify-center bg-gray-50 p-4 absolute">
      <div className="w-full max-w-md space-y-6 rounded-lg bg-white p-8 shadow-md">
        <h1 className="text-center text-2xl font-bold">Вход</h1>
        
        <div className="space-y-4">
          <div>
            <label className="block text-sm font-medium text-gray-700">Email</label>
            <Input type="email" className="mt-1" />
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700">Пароль</label>
            <Input type="password" className="mt-1" />
          </div>

          <Link to="/home">
            <Button className="w-full">Войти</Button>
          </Link>

          <div className="text-center">
            <Link 
              to="/register" 
              className="text-sm font-medium text-blue-600 hover:text-blue-500 hover:underline"
            >
              Еще нет аккаунта? Зарегистрироваться
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
}