# NestJS CRUD with TypeORM

TypeORM을 사용한 NestJS CRUD API

## 기술 스택
- NestJS (v10.3.0)
- TypeORM (v0.3.19)
- MariaDB
- TypeScript
- Class Validator

## 설치 및 실행

```bash
# 의존성 설치
npm install

# 빌드
npm run build

# 개발 모드 실행 (hot-reload)
npm run start:dev

# 프로덕션 모드 실행
npm run start:prod
```

## 프로젝트 구조

```
src/
├── users/
│   ├── dto/
│   │   ├── create-user.dto.ts     # 사용자 생성 DTO
│   │   └── update-user.dto.ts     # 사용자 수정 DTO
│   ├── entities/
│   │   └── user.entity.ts         # User 엔티티
│   ├── users.controller.ts        # Users 컨트롤러
│   ├── users.service.ts           # Users 서비스
│   └── users.module.ts            # Users 모듈
├── app.module.ts                  # App 모듈
└── main.ts                        # 진입점
```

## API 엔드포인트

### 사용자 생성 (CREATE)
```bash
POST http://localhost:3002/api/users
Content-Type: application/json

{
  "email": "user@example.com",
  "username": "testuser",
  "password": "password123"
}
```

### 모든 사용자 조회 (READ)
```bash
GET http://localhost:3002/api/users
```

### 특정 사용자 조회 (READ)
```bash
GET http://localhost:3002/api/users/:id
```

### 사용자 수정 (UPDATE)
```bash
PUT http://localhost:3002/api/users/:id
Content-Type: application/json

{
  "email": "updated@example.com",
  "username": "updateduser",
  "password": "newpassword123"
}
```

### 사용자 삭제 (DELETE)
```bash
DELETE http://localhost:3002/api/users/:id
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

## 주요 특징

### 1. TypeORM
- 엔티티 기반 데이터 모델링
- 자동 마이그레이션 (synchronize: true)
- Repository 패턴
- 데코레이터 기반 설정

### 2. Validation
- class-validator를 사용한 요청 검증
- DTO 기반 데이터 전송
- 자동 타입 변환

### 3. Dependency Injection
- NestJS의 강력한 DI 시스템
- 모듈화된 구조
- 테스트 용이성

### 4. 에러 처리
- 내장 예외 처리
- NotFoundException 자동 처리
- 일관된 에러 응답

## NestJS vs Express

### NestJS 장점
- TypeScript 기본 지원
- 의존성 주입 (DI)
- 모듈화된 아키텍처
- 데코레이터 기반 라우팅
- 내장 Validation
- 확장성과 유지보수성
