# f25 Monorepo

This repository contains the complete source code for the **f25** project, including frontend applications, backend microservices, and shared packages.

We follow a **monorepo** architecture to keep all related code in a single place, making development, testing, and deployment more efficient.

---

## 📂 Repository Structure
```
f25/
│
├── apps/ # Frontend applications (Next.js)
├── services/ # Backend microservices (Go)
├── shared/ # Shared utilities, types, and constants
├── package.json
└── README.md
```

---

## 🛠 Tech Stack
### Frontend
- [Next.js](https://nextjs.org/) for React-based apps
- [TypeScript](https://www.typescriptlang.org/) for type safety
- [Tailwind CSS](https://tailwindcss.com/) for styling

### Backend
- [Go](https://go.dev/) for microservices
- REST APIs
- PostgreSQL (planned)

### DevOps & Tooling
- Git + GitHub for version control
- npm / pnpm for package management
- Docker (planned)
- GitHub Actions (planned)

---

## 📁 Folder Details
- **[`apps/`](./apps/README.md)** → Contains all frontend applications. Each app is self-contained.
- **[`services/`](./services/README.md)** → Contains backend microservices. Each service runs independently.
- **[`shared/`](./shared/README.md)** → Contains code that is shared between apps and services.

---

## 🚀 Getting Started

### 1️⃣ Clone the repository
```bash
git clone https://github.com/<your-username>/f25.git
cd f25
npm install
cd apps/<app-name>
npm install
cd apps/<app-name>
npm run dev
cd services/<service-name>
go run main.go
