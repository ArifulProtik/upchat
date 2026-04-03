# UpChat

> A free AI chat application that gives you access to multiple free LLM models in one place. Chat with any AI model, share your conversations publicly, or keep them private.

## Project Overview

UpChat is a full-stack AI chat application that aggregates free LLM models into a single interface. Users can select from various AI models (like ChatGPT alternatives, DeepSeek, etc.) and start chatting. Every conversation can be set to **public** or **private** -- public chats are visible to all users and can be searched, discovered, and explored by the community.

**Current State:** Early-stage scaffold -- authentication UI and database schema are in place. Backend has a health check endpoint with Gin + ent ORM + PostgreSQL. Frontend features signin/signup pages with Zod validation, shadcn-svelte component library, and light/dark theming.

**Planned Features:**

- Multi-model AI chat with streaming responses
- Community chat feed -- browse all public conversations
- Search and filter public chats.
- User authentication and session management
- Conversation history and management

## Tech Stack

**Backend:**

- Go 1.26 + Gin framework
- ent ORM (Entity Framework)
- PostgreSQL (lib/pq)
- air (hot reload)

**Frontend:**

- SvelteKit 5 (runes mode)
- TypeScript (strict mode)
- TailwindCSS v4 + shadcn-svelte (maia style, zinc base, tabler icons)
- Zod v4 (validation)
- bits-ui, formsnap, sveltekit-superforms

## Prerequisites

- Go 1.26+
- Bun (or npm/yarn/pnpm)
- PostgreSQL 14+
- `make`

## How to Run

### 1. Setup Environment

```bash
cp .env.example .env
```

Edit `.env`:

```env
GIN_MODE=debug
DATABASE_URL=postgresql://localhost:5432/upchat?sslmode=disable
DATABASE_DRIVER=postgres
PORT=8081
```

### 2. Create Database

```bash
createdb upchat
```

### 3. Start Everything

```bash
make dev          # Runs backend (air) + frontend (bun dev) together
```

Or run separately:

```bash
# Terminal 1 - Backend
make run

# Terminal 2 - Frontend
cd ui && bun install && bun run dev
```

## Development Commands

| Command                        | Description                            |
| ------------------------------ | -------------------------------------- |
| `make dev`                     | Run backend + frontend with hot reload |
| `make build`                   | Compile Go binary to `bin/myapp`       |
| `make run`                     | Build and run backend + frontend       |
| `make generate`                | Generate ent code from schema          |
| `make schema schema_name=User` | Create new ent schema                  |
| `make test`                    | Run tests with race detection          |
| `make lint`                    | Run golangci-lint                      |
| `make clean`                   | Remove build artifacts                 |
| `cd ui && bun run check`       | Type check frontend                    |
| `cd ui && bun run lint`        | ESLint + Prettier                      |
| `cd ui && bun run format`      | Auto-format code                       |

## Documentation References

- [Go Documentation](https://go.dev/doc/)
- [Gin Framework](https://gin-gonic.com/docs/)
- [ent ORM](https://entgo.io/docs/getting-started)
- [SvelteKit 5](https://svelte.dev/docs/kit)
- [Svelte 5 Runes](https://svelte.dev/docs/svelte/runes)
- [TailwindCSS v4](https://tailwindcss.com/docs)
- [shadcn-svelte](https://www.shadcn-svelte.com/)
- [Zod](https://zod.dev/)
- [bits-ui](https://www.bits-ui.com/)
