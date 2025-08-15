# Frontend Applications

This folder contains all frontend applications for the **f25** project.  
We are following a **micro-frontend architecture** within a monorepo, using [Next.js](https://nextjs.org/) for each app.

## Structure
Each application lives in its own folder:

```
apps/
    app-name/
    package.json
    README.md
    ...
```

## Development Guidelines
- Each app is self-contained but may import from the `/shared` package.
- Keep app-specific code inside its own directory.
- Use TypeScript for type safety.
- Follow the project’s coding style and linting rules.

## Current Apps
- _To be added_

## Getting Started
To start developing a specific app:
```bash
cd apps/<app-name>
npm install
npm run dev
