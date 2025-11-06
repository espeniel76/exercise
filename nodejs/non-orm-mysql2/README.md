# Node.js CRUD with mysql2 (Non-ORM)

mysql2를 사용한 Node.js CRUD API (ORM 없이 순수 SQL 사용)

## 기술 스택
- Node.js + Express
- mysql2 (v3.6.5) - Promise 기반
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
POST http://localhost:3001/api/users
Content-Type: application/json

{
  "email": "user@example.com",
  "username": "testuser",
  "password": "password123"
}
```

### 모든 사용자 조회 (READ)
```bash
GET http://localhost:3001/api/users
```

### 특정 사용자 조회 (READ)
```bash
GET http://localhost:3001/api/users/:id
```

### 사용자 수정 (UPDATE)
```bash
PUT http://localhost:3001/api/users/:id
Content-Type: application/json

{
  "email": "updated@example.com",
  "username": "updateduser",
  "password": "newpassword123"
}
```

### 사용자 삭제 (DELETE)
```bash
DELETE http://localhost:3001/api/users/:id
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

## ORM vs Non-ORM 차이점

### ORM (Sequelize)
- 객체 지향적 데이터 조작
- 자동 마이그레이션 지원
- 타입 안정성
- 복잡한 쿼리는 학습 곡선 존재

### Non-ORM (mysql2)
- 순수 SQL 사용
- 직접적인 데이터베이스 제어
- 성능 최적화 용이
- SQL 지식 필수
