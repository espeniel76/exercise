# Fiber + GORM CRUD API

Fiber 프레임워크와 GORM ORM을 사용한 RESTful CRUD API 예제입니다.

## 기술 스택

- **Web Framework**: [Fiber](https://gofiber.io/) - Express.js에서 영감을 받은 초고속 웹 프레임워크
- **ORM**: [GORM](https://gorm.io/) - Go에서 가장 대표적인 ORM 라이브러리
- **Database**: MariaDB/MySQL
- **Environment**: godotenv

## 프로젝트 구조

```
fiber-gorm/
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

## Fiber 프레임워크 특징

### 핵심 장점
- **초고속 성능**: fasthttp 기반으로 Go 프레임워크 중 가장 빠름 (벤치마크 1위)
- **Express.js와 유사**: Node.js 개발자에게 매우 친숙한 API
- **제로 메모리 할당**: 극도로 최적화된 메모리 사용
- **낮은 CPU 사용량**: 효율적인 리소스 활용
- **풍부한 미들웨어**: 40개 이상의 내장 미들웨어
- **빠른 라우팅**: 최적화된 라우터

### Express.js와의 유사성

**Express.js (Node.js)**:
```javascript
app.get('/api/users/:id', (req, res) => {
  const id = req.params.id;
  res.json({ data: user });
});
```

**Fiber (Go)**:
```go
app.Get("/api/users/:id", func(c *fiber.Ctx) error {
  id := c.Params("id")
  return c.JSON(fiber.Map{"data": user})
})
```

### 내장 미들웨어
- **Logger**: 요청 로깅
- **Recover**: 패닉 복구
- **CORS**: Cross-Origin Resource Sharing
- **Compress**: Gzip/Brotli 압축
- **Cache**: 응답 캐싱
- **Limiter**: Rate limiting
- **Monitor**: 성능 모니터링
- **Helmet**: 보안 헤더

### Fiber Context API

```go
// Request
c.Params("id")           // URL 파라미터
c.Query("name")          // 쿼리 파라미터
c.BodyParser(&user)      // Body 파싱
c.FormValue("email")     // Form 값
c.Cookies("token")       // 쿠키

// Response
c.JSON(data)             // JSON 응답
c.Status(200)            // 상태 코드
c.SendString("Hello")    // 문자열 응답
c.Redirect("/path")      // 리다이렉트
```

## 성능 비교

### 벤치마크 결과 (요청/초)

```
Fiber (fasthttp):    6,162,556  ⭐⭐⭐⭐⭐
Gin:                 5,976,157  ⭐⭐⭐⭐⭐
Echo:                5,664,787  ⭐⭐⭐⭐
Chi:                 4,397,590  ⭐⭐⭐⭐
Gorilla Mux:         3,416,123  ⭐⭐⭐
```

Fiber는 fasthttp를 기반으로 하여 표준 net/http보다 **10배 빠른** 성능을 제공합니다.

## 다른 프레임워크와 비교

| 프레임워크 | 기반 | 성능 | 메모리 | 학습 곡선 | Express 유사성 |
|----------|------|------|--------|---------|---------------|
| **Fiber** | fasthttp | ⭐⭐⭐⭐⭐ | 매우 낮음 | 쉬움 | 매우 높음 |
| Gin | net/http | ⭐⭐⭐⭐⭐ | 낮음 | 쉬움 | 중간 |
| Echo | net/http | ⭐⭐⭐⭐ | 낮음 | 쉬움 | 중간 |
| Gorilla | net/http | ⭐⭐⭐⭐ | 낮음 | 쉬움 | 낮음 |

## Fiber를 선택해야 하는 경우

1. **최고의 성능이 필요한 경우** - 대규모 트래픽 처리
2. **Node.js에서 전환하는 경우** - Express.js와 매우 유사한 API
3. **마이크로서비스** - 낮은 메모리와 빠른 응답 속도
4. **실시간 애플리케이션** - WebSocket, SSE 등

## Fiber의 독특한 기능

### 1. 체이닝 가능한 응답

```go
return c.Status(200).JSON(fiber.Map{
    "success": true,
    "data": user,
})
```

### 2. fiber.Map 헬퍼

```go
// 간결한 JSON 응답
return c.JSON(fiber.Map{
    "message": "Hello",
    "status": "success",
})
```

### 3. 그룹 라우팅

```go
api := app.Group("/api")
v1 := api.Group("/v1")
v1.Get("/users", getUsers)
```

## GORM 기능

- **Auto Migration**: 자동 스키마 마이그레이션
- **Connection Pooling**: 효율적인 DB 커넥션 관리
- **Soft Delete**: DeletedAt 필드를 통한 소프트 삭제 지원
- **Logging**: 쿼리 로깅 기능

## 프로덕션 최적화

```go
app := fiber.New(fiber.Config{
    Prefork:       true,  // 멀티 프로세스 모드
    CaseSensitive: true,
    StrictRouting: true,
    ServerHeader:  "Fiber",
    AppName:       "Your App v1.0.0",
})
```

## 라이선스

ISC
