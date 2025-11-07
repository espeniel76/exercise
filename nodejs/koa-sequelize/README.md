# Koa CRUD with Sequelize ORM

Koa와 Sequelize ORM을 사용한 Node.js CRUD API

## 기술 스택
- Koa (v2.14.2) - 차세대 웹 프레임워크
- Koa Router (v12.0.1)
- Koa Bodyparser (v4.4.1)
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

## 프로젝트 구조

```
koa-sequelize/
├── config/
│   └── database.js          # Sequelize 데이터베이스 설정
├── models/
│   └── User.js              # User 모델
├── routes/
│   └── users.js             # Users 라우트
├── server.js                # Koa 서버 진입점
├── .env                     # 환경 변수
├── package.json
└── README.md
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

## Koa의 특징

### 1. Async/Await 중심 설계
- ES2017 async/await를 완전히 지원
- 콜백 지옥 없음
- 깔끔한 비동기 처리

### 2. 미들웨어 캐스케이드
- 양파 모델(Onion Model) 미들웨어
- 요청/응답 흐름 제어 용이
- 강력한 에러 핸들링

### 3. Context (ctx) 객체
- request와 response를 하나로 통합
- 더 나은 API 설계
- 편리한 사용성

### 4. 경량화
- 핵심 기능만 포함
- 필요한 기능은 플러그인으로 추가
- 작은 번들 사이즈

## Koa vs Express

| 특징 | Koa | Express |
|------|-----|---------|
| 비동기 처리 | async/await 네이티브 | 콜백/Promise |
| 미들웨어 | 양파 모델 | 선형 모델 |
| 번들 크기 | 작음 | 중간 |
| 내장 기능 | 최소 | 많음 |
| 에러 핸들링 | 강력함 | 기본 |
| 학습 곡선 | 보통 | 완만 |
| 커뮤니티 | 중간 | 매우 큼 |

## Koa 미들웨어 흐름 (양파 모델)

```
Request
   ↓
┌─────────────────┐
│  Middleware 1   │ → await next()
│  ┌───────────┐  │
│  │ MW 2      │  │ → await next()
│  │ ┌───────┐ │  │
│  │ │ MW 3  │ │  │ → await next()
│  │ │       │ │  │
│  │ │ Route │ │  │ ← Response
│  │ └───────┘ │  │
│  └───────────┘  │
└─────────────────┘
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

## 환경 변수 (.env)

```env
DB_HOST=15.165.117.90
DB_USER=root
DB_PASSWORD=anjgkrp
DB_NAME=espeniel
DB_PORT=3306
PORT=3001
```

## Koa를 선택해야 하는 경우

- 모던 JavaScript(async/await)를 선호하는 경우
- 깔끔한 에러 핸들링이 필요한 경우
- 미들웨어 흐름을 세밀하게 제어하고 싶은 경우
- 경량화된 프레임워크를 원하는 경우
- Express 팀이 만든 차세대 프레임워크를 경험하고 싶은 경우

## Koa의 장점

1. **더 나은 에러 핸들링**: try-catch로 모든 에러를 잡을 수 있음
2. **깔끔한 코드**: async/await로 가독성 향상
3. **모듈러**: 필요한 기능만 추가 가능
4. **성능**: Express보다 약간 더 빠름
5. **미래 지향적**: 최신 JavaScript 기능 활용
