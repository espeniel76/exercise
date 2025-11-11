# Gorilla Mux + GORM CRUD API

Gorilla Mux 라우터와 GORM ORM을 사용한 RESTful CRUD API 예제입니다.

## 기술 스택

- **Router**: [Gorilla Mux](https://github.com/gorilla/mux) - Go에서 가장 오래되고 안정적인 HTTP 라우터
- **ORM**: [GORM](https://gorm.io/) - Go에서 가장 대표적인 ORM 라이브러리
- **Database**: MariaDB/MySQL
- **Environment**: godotenv

## 프로젝트 구조

```
gorilla-gorm/
├── config/
│   └── database.go      # 데이터베이스 연결 설정
├── models/
│   └── user.go          # User 모델 정의
├── routes/
│   └── users.go         # User 라우트 핸들러
├── main.go              # 애플리케이션 진입점
├── go.mod               # Go 모듈 정의
├── .env.example         # 환경변수 예제
├── .gitignore           # Git 제외 파일
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

### Gorilla Mux 특징
- **강력한 라우팅**: 정규표현식, 호스트, 메서드, 스키마 매칭
- **표준 라이브러리 기반**: net/http와 완벽 호환
- **서브 라우터**: 모듈화된 라우팅 구조
- **미들웨어 지원**: 체인 가능한 미들웨어
- **안정성**: 2012년부터 운영, 매우 안정적
- **URL 변수**: 편리한 경로 파라미터 추출

### GORM 기능
- **Auto Migration**: 자동 스키마 마이그레이션
- **Connection Pooling**: 효율적인 DB 커넥션 관리
- **Soft Delete**: DeletedAt 필드를 통한 소프트 삭제 지원
- **Logging**: 쿼리 로깅 기능

### 커스텀 미들웨어
- **CORS**: Cross-Origin Resource Sharing 지원
- **Logging**: 모든 요청 로깅

## Gorilla Mux의 장점

1. **표준 라이브러리 기반**: net/http Handler 인터페이스와 완전 호환
2. **안정성**: 10년 이상 검증된 프로덕션 레벨 라이브러리
3. **유연성**: 매우 강력한 라우팅 기능
4. **학습 곡선**: 표준 라이브러리를 알면 쉽게 배울 수 있음
5. **커뮤니티**: 방대한 문서와 예제

## 다른 프레임워크와 비교

| 프레임워크 | 특징 | 장점 |
|----------|------|------|
| **Gorilla Mux** | 표준 기반 라우터 | 안정성, 유연성, 표준 호환 |
| Gin | 고성능 프레임워크 | 빠른 속도, 미들웨어 풍부 |
| Echo | 미니멀 프레임워크 | 빠른 속도, 쉬운 사용 |
| Fiber | Express-like | 매우 빠름, Node.js와 유사 |

## 라우팅 예제

Gorilla Mux는 다양한 라우팅 옵션을 제공합니다:

```go
// 경로 파라미터
r.HandleFunc("/users/{id}", handler)

// 정규표현식
r.HandleFunc("/users/{id:[0-9]+}", handler)

// 메서드 지정
r.HandleFunc("/users", handler).Methods("GET", "POST")

// 호스트 매칭
r.HandleFunc("/", handler).Host("example.com")

// 서브 라우터
api := r.PathPrefix("/api").Subrouter()
api.HandleFunc("/users", usersHandler)
```

## 라이선스

ISC
