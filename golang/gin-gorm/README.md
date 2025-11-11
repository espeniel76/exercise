# Gin + GORM CRUD API

Gin 프레임워크와 GORM ORM을 사용한 RESTful CRUD API 예제입니다.

## 기술 스택

- **Web Framework**: [Gin](https://github.com/gin-gonic/gin) - Go에서 가장 빠르고 인기있는 웹 프레임워크
- **ORM**: [GORM](https://gorm.io/) - Go에서 가장 대표적인 ORM 라이브러리
- **Database**: MariaDB/MySQL
- **Environment**: godotenv

## 프로젝트 구조

```
gin-gorm/
├── config/
│   └── database.go      # 데이터베이스 연결 설정
├── models/
│   └── user.go          # User 모델 정의
├── routes/
│   └── users.go         # User 라우트 핸들러
├── main.go              # 애플리케이션 진입점
├── go.mod               # Go 모듈 정의
├── .env.example         # 환경변수 예제
└── README.md
```

## 설치 및 실행

### 1. 환경변수 설정

```bash
cp .env.example .env
```

`.env` 파일을 편집하여 데이터베이스 연결 정보를 입력하세요.

### 2. 의존성 설치

```bash
go mod download
```

또는 실행 시 자동으로 다운로드됩니다.

### 3. 서버 실행

```bash
# 일반 실행
go run main.go

# 빌드 후 실행
go build -o server
./server
```

서버는 기본적으로 `http://localhost:3001`에서 실행됩니다.

## API 엔드포인트

### Health Check
- `GET /` - 서버 상태 확인

### User CRUD Operations

#### 1. CREATE - 사용자 생성
```bash
POST /api/users
Content-Type: application/json

{
  "email": "user@example.com",
  "username": "johndoe",
  "password": "password123"
}
```

#### 2. READ - 모든 사용자 조회
```bash
GET /api/users
```

#### 3. READ - 특정 사용자 조회
```bash
GET /api/users/:id
```

#### 4. UPDATE - 사용자 수정
```bash
PUT /api/users/:id
Content-Type: application/json

{
  "email": "updated@example.com",
  "username": "updateduser",
  "password": "newpassword123"
}
```

#### 5. DELETE - 사용자 삭제
```bash
DELETE /api/users/:id
```

## 응답 형식

### 성공 응답
```json
{
  "success": true,
  "data": {
    "id": 1,
    "email": "user@example.com",
    "username": "johndoe",
    "password": "password123",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

### 에러 응답
```json
{
  "success": false,
  "error": "User not found"
}
```

## 주요 특징

### GORM 기능
- **Auto Migration**: 자동 스키마 마이그레이션
- **Validation**: 모델 레벨 데이터 검증 (email, required 등)
- **Connection Pooling**: 효율적인 DB 커넥션 관리
- **Soft Delete**: DeletedAt 필드를 통한 소프트 삭제 지원
- **Logging**: 쿼리 로깅 기능

### Gin 프레임워크
- **빠른 성능**: httprouter 기반의 고성능 라우팅
- **Middleware**: CORS 등 미들웨어 지원
- **JSON Validation**: 자동 JSON 바인딩 및 검증
- **Error Handling**: 우아한 에러 핸들링

## 개발 모드

환경변수 `GIN_MODE`를 설정하여 모드를 변경할 수 있습니다:

```bash
# Debug mode (default)
GIN_MODE=debug

# Release mode
GIN_MODE=release
```

## 데이터베이스 설정

GORM은 MariaDB/MySQL을 자동으로 감지하고 최적화합니다:

- Connection Pool 설정
  - Max Open Connections: 5
  - Max Idle Connections: 0
  - Connection Lifetime: 1시간

## Node.js 버전과의 비교

이 프로젝트는 `fastify-sequelize` 프로젝트와 동일한 구조와 기능을 제공합니다:

| 기능 | Node.js | Go |
|------|---------|-----|
| Framework | Fastify | Gin |
| ORM | Sequelize | GORM |
| 라우팅 | Plugin 기반 | Function 기반 |
| 타입 | JavaScript | 정적 타입 (Go) |
| 성능 | 빠름 | 매우 빠름 |

## 라이선스

ISC
