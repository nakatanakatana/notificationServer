import { createRootRoute, Link, Outlet } from "@tanstack/solid-router";

export const Route = createRootRoute({
  component: () => (
    <div>
      <nav style={{ padding: "1rem", border: "1px solid #ccc" }}>
        <Link to="/">Home</Link>
      </nav>
      <main style={{ padding: "1rem" }}>
        <Outlet />
      </main>
    </div>
  ),
});
