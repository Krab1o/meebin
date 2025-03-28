import { type RouteConfig, index, route } from "@react-router/dev/routes";

export default 
[
    index("routes/home.tsx"),
    route("*", "routes/not-found.tsx"),
    route("about", "routes/about.tsx"),
    route("profile", "routes/profile.tsx"),
] satisfies RouteConfig;
