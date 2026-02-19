import { createFileRoute } from "@tanstack/solid-router";

export const Route = createFileRoute("/")({
  component: Home,
});

function Home() {
  return (
    <div>
      <h1>Welcome to {{PROJECT_NAME}}</h1>
      <p>Your app is running!</p>
    </div>
  );
}
