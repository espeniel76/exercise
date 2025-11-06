# Node.js CRUD with Sequelize ORM

Sequelize ORM을 사용한 Node.js CRUD API

## 기술 스택
- Node.js + Express
- Sequelize ORM (v6.35.2)
- MariaDB

## 설치 및 실행

```bash
# 의존성 설치
npm install

# 서버 실행
npm start

# 개발 모드 (nodemon)
npm run dev
```

## API 엔드포인트

### 사용자 생성 (CREATE)
```bash
POST http://localhost:3000/api/users
Content-Type: application/json

{
  "email": "user@example.com",
  "username": "testuser",
  "password": "password123"
}
```

### 모든 사용자 조회 (READ)
```bash
GET http://localhost:3000/api/users
```

### 특정 사용자 조회 (READ)
```bash
GET http://localhost:3000/api/users/:id
```

### 사용자 수정 (UPDATE)
```bash
PUT http://localhost:3000/api/users/:id
Content-Type: application/json

{
  "email": "updated@example.com",
  "username": "updateduser",
  "password": "newpassword123"
}
```

### 사용자 삭제 (DELETE)
```bash
DELETE http://localhost:3000/api/users/:id
```

## 데이터베이스 스키마

```sql
CREATE TABLE `users` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `email` VARCHAR(255) NOT NULL,
  `username` VARCHAR(50) NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_email` (`email`),
  UNIQUE KEY `uk_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```
