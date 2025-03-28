import { TabsDemo } from "~/components/userPage/userPage";
import type { Route } from "./+types/profile";

export function meta({}: Route.MetaArgs) {
  return [{ title: "Page Not Found" }];
}

export default function NotFound() {
  return (
    <TabsDemo></TabsDemo>
  );
}