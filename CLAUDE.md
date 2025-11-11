# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

This is a polyglot CRUD practice repository demonstrating identical User management APIs across multiple language/framework/ORM combinations. All implementations connect to a shared MariaDB database (`espeniel`) and implement the same RESTful API contract.

### Shared Database Configuration
- **Database**: MariaDB on 15.165.117.90:3306
- **Database Name**: espeniel
- **Schema**: Users table with fields: id, email, username, password, created_at, updated_at
- **Setup Script**: [database-setup.sql](database-setup.sql)

All projects use environment variables for database configuration (see .env files in each project directory).

## Project Structure

The repository is organized by language with subdirectories for each framework+ORM combination:

### Node.js Implementations (nodejs/)
- **nestjs-typeorm/** - NestJS with TypeORM (Port 3002)
- **fastify-sequelize/** - Fastify with Sequelize ORM (Port 3001)
- **koa-sequelize/** - Koa with Sequelize ORM
- **hapi-sequelize/** - Hapi with Sequelize ORM
- **orm-sequelize/** - Express with Sequelize ORM (Port 3000)
- **non-orm-mysql2/** - Express with raw mysql2 (no ORM) (Port 3001)

### Go Implementations (golang/)
- **gin-gorm/** - Gin with GORM (Port 3001)
- **echo-gorm/** - Echo with GORM
- **fiber-gorm/** - Fiber with GORM
- **gorilla-gorm/** - Gorilla Mux with GORM

### Python Implementations (python/)
- Empty directory (planned implementations)

## Common Commands

### Node.js Projects
```bash
# Install dependencies
npm install

# Run in development mode (with hot-reload)
npm run dev

# Run in production mode
npm start

# Build (for TypeScript projects like NestJS)
npm run build
npm run start:prod
```

### Go Projects
```bash
# Install dependencies
go mod download

# Run directly
go run main.go

# Build and run
go build -o server
./server  # or .\server.exe on Windows
```

## API Contract

All implementations expose identical REST endpoints:

- `GET /` - Health check
- `POST /api/users` - Create user (requires: email, username, password)
- `GET /api/users` - List all users
- `GET /api/users/:id` - Get user by ID
- `PUT /api/users/:id` - Update user (accepts: email, username, password)
- `DELETE /api/users/:id` - Delete user

Response format:
```json
{
  "success": true,
  "data": { /* user object */ }
}
```

Error response:
```json
{
  "success": false,
  "error": "Error message"
}
```

## Architecture Patterns

### Node.js Projects (Express-like)
Common structure across Express, Fastify, Koa, Hapi:
```
project/
├── config/
│   └── database.js      # Database connection setup
├── models/
│   └── User.js          # User model definition
├── routes/
│   └── users.js         # Route handlers
└── server.js            # Application entry point
```

### NestJS (Module-based)
```
src/
├── users/
│   ├── dto/             # Data Transfer Objects
│   ├── entities/        # TypeORM entities
│   ├── users.controller.ts
│   ├── users.service.ts
│   └── users.module.ts
├── app.module.ts
└── main.ts
```

### Go Projects
```
project/
├── config/
│   └── database.go      # Database connection (GORM)
├── models/
│   └── user.go          # User struct with GORM tags
├── routes/
│   └── users.go         # Route handlers
└── main.go              # Entry point with framework setup
```

## Key Implementation Differences

### ORM vs Non-ORM
- **ORM projects** (Sequelize, TypeORM, GORM): Use model definitions with automatic migrations
- **Non-ORM** (mysql2): Direct SQL queries with connection pooling

### Framework Characteristics
- **Gin**: Fastest Go framework, minimal middleware
- **Echo/Fiber**: Modern, Express-like API in Go
- **Gorilla Mux**: Standard library-style routing
- **NestJS**: Angular-inspired DI and modules, TypeScript-first
- **Fastify**: High-performance, schema-based validation
- **Express**: Most common, middleware-heavy

## Environment Configuration

All projects expect a `.env` file with these variables:
```env
DB_HOST=db host
DB_USER=db user
DB_PASSWORD=db password
DB_NAME=db name
DB_PORT=3306
PORT=3001  # or 3000, 3002 depending on project
```

## Testing

API test collection is available in [apidog-collection.json](apidog-collection.json) (OpenAPI 3.0 format) with examples for all endpoints.

## Adding New Implementations

When adding a new framework/language:
1. Create directory under appropriate language folder
2. Implement all 5 CRUD operations following the API contract
3. Use the shared database schema (run database-setup.sql if needed)
4. Configure environment variables matching the pattern above
5. Add README.md documenting setup and framework-specific features
6. Ensure consistent port configuration to avoid conflicts
