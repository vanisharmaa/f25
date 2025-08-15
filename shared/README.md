### **f25/shared/README.md**

# Shared Packages & Utilities

This folder contains **shared code** used across multiple apps and services in the **f25** project.

## Purpose
- Avoid code duplication.
- Provide common utilities, types, constants, and helpers.
- Maintain a single source of truth for shared logic.

## Structure
```
shared/
    utils/ # Generic utility functions
    types/ # Shared TypeScript/Go types
    constants/ # Global constants
    README.md
```

## Guidelines
- Only include **truly reusable** code here.
- Do not include app/service-specific logic.
- Keep shared code framework-agnostic when possible.

## Importing Shared Code
From an app:
```ts
import { myFunction } from '@f25/shared/utils';
