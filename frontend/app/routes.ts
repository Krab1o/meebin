import { type RouteConfig, index, route } from "@react-router/dev/routes";

export default 
[
    index("routes/login.tsx"),
    route("*", "routes/not-found.tsx"),
    route("about", "routes/about.tsx"),
    route("profile", "routes/profile.tsx"),
    route("home", "routes/home.tsx"),
    route("register", "routes/register.tsx"),
    route("create", "routes/create.tsx"),
    route("test", "routes/test.tsx"),
] satisfies RouteConfig;
