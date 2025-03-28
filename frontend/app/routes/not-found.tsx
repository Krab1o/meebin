import type { Route } from "./+types/not-found";

export function meta({}: Route.MetaArgs) {
  return [{ title: "Page Not Found" }];
}

export default function NotFound() {
  return (
    <div className="p-8 text-center">
      <h1 className="text-2xl font-bold">404</h1>
      <p>Page not found.</p>
    </div>
  );
}