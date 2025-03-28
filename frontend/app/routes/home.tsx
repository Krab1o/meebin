
import type { Route } from "./+types/home"
import { Link } from "react-router-dom";
import CleanupRequestsPage from "~/components/trashJournal/trashJounral"
 
export function meta({}: Route.MetaArgs) {
  return [
    { title: "Главная страница" },
    { name: "description", content: "Welcome to React Router!" },
  ]
}
 
export default function Home() {
  return (
      <div className="flex flex-col items-center justify-center min-h-svh">
        <CleanupRequestsPage></CleanupRequestsPage>
      </div>
  )
}