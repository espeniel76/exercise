# Fastify CRUD with Sequelize ORM

Fastify와 Sequelize ORM을 사용한 Node.js CRUD API

## 기술 스택
- Fastify (v4.25.2) - 고성능 웹 프레임워크
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
fastify-sequelize/
├── config/
│   └── database.js          # Sequelize 데이터베이스 설정
├── models/
│   └── User.js              # User 모델
├── routes/
│   └── users.js             # Users 라우트
├── server.js                # Fastify 서버 진입점
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

## Fastify의 특징

### 1. 고성능
- Express보다 약 2배 빠른 성능
- JSON 파싱 최적화
- 효율적인 라우팅 시스템

### 2. Schema 기반 검증
- JSON Schema를 사용한 요청/응답 검증
- 자동 문서화 지원
- 타입 안정성

### 3. 플러그인 아키텍처
- 캡슐화된 플러그인 시스템
- 의존성 주입
- 비동기 부트스트랩

### 4. TypeScript 지원
- 훌륭한 TypeScript 타입 정의
- 타입 안전성

## Fastify vs Express

| 특징 | Fastify | Express |
|------|---------|---------|
| 성능 | 매우 빠름 (2배 이상) | 보통 |
| JSON 처리 | 최적화됨 | 기본 |
| 검증 | 내장 Schema 검증 | 별도 라이브러리 필요 |
| 플러그인 | 강력한 플러그인 시스템 | 미들웨어 |
| 학습 곡선 | 조금 가파름 | 완만함 |
| 커뮤니티 | 성장 중 | 매우 큼 |

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

## 성능 비교

Fastify는 다음과 같은 상황에서 특히 유용합니다:
- 높은 처리량이 필요한 API
- JSON을 많이 다루는 서비스
- 마이크로서비스 아키텍처
- 실시간 데이터 처리
