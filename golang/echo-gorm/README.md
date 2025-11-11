# Echo + GORM CRUD API

Echo 프레임워크와 GORM ORM을 사용한 RESTful CRUD API 예제입니다.

## 기술 스택

- **Web Framework**: [Echo](https://echo.labstack.com/) - 고성능, 확장 가능한 미니멀 웹 프레임워크
- **ORM**: [GORM](https://gorm.io/) - Go에서 가장 대표적인 ORM 라이브러리
- **Database**: MariaDB/MySQL
- **Environment**: godotenv

## 프로젝트 구조

```
echo-gorm/
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

## Echo 프레임워크 특징

### 핵심 기능
- **고성능**: 빠른 HTTP 라우터와 최적화된 성능
- **미니멀**: 간결하고 직관적인 API
- **확장 가능**: 풍부한 미들웨어 생태계
- **HTTP/2 지원**: 자동 TLS 및 HTTP/2 지원
- **WebSocket**: 내장 WebSocket 지원
- **템플릿 렌더링**: 다양한 템플릿 엔진 지원

### 내장 미들웨어
- **Logger**: 요청 로깅
- **Recover**: 패닉 복구
- **CORS**: Cross-Origin Resource Sharing
- **JWT**: JSON Web Token 인증
- **Rate Limit**: 요청 속도 제한
- **Gzip**: 압축 지원

### Context 기반 설계
Echo는 강력한 Context 객체를 제공합니다:

```go
// Request 바인딩
c.Bind(&user)

// JSON 응답
c.JSON(200, data)

// 경로 파라미터
c.Param("id")

// 쿼리 파라미터
c.QueryParam("name")

// Form 값
c.FormValue("email")
```

## GORM 기능

- **Auto Migration**: 자동 스키마 마이그레이션
- **Connection Pooling**: 효율적인 DB 커넥션 관리
- **Soft Delete**: DeletedAt 필드를 통한 소프트 삭제 지원
- **Logging**: 쿼리 로깅 기능

## 다른 프레임워크와 비교

| 프레임워크 | 성능 | 학습 곡선 | 미들웨어 | 특징 |
|----------|------|---------|---------|------|
| **Echo** | ⭐⭐⭐⭐⭐ | 쉬움 | 풍부 | 미니멀하고 빠름 |
| Gin | ⭐⭐⭐⭐⭐ | 쉬움 | 풍부 | 가장 빠름 |
| Gorilla | ⭐⭐⭐⭐ | 쉬움 | 수동 | 표준 라이브러리 기반 |
| Fiber | ⭐⭐⭐⭐⭐ | 중간 | 풍부 | Express.js와 유사 |

## Echo vs Gin

### Echo의 장점
- 더 풍부한 내장 기능 (WebSocket, HTTP/2 등)
- 더 우아한 에러 핸들링
- 자동 TLS 지원
- 더 직관적인 Context API

### Gin의 장점
- 약간 더 빠른 성능
- 더 큰 커뮤니티
- JSON 자동 바인딩/검증

## 에러 핸들링

Echo는 에러를 반환하면 자동으로 처리합니다:

```go
func handler(c echo.Context) error {
    // 에러를 반환하면 Echo가 처리
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return c.JSON(200, data)
}
```

## 커스텀 미들웨어 예제

```go
// 커스텀 로거
e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        start := time.Now()
        err := next(c)
        log.Printf("%s %s %v", c.Request().Method, c.Path(), time.Since(start))
        return err
    }
})
```

## 라이선스

ISC
