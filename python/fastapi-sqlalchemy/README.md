# FastAPI CRUD with SQLAlchemy ORM

FastAPI와 SQLAlchemy ORM을 사용한 현대적인 Python CRUD API

## 기술 스택
- **Framework**: FastAPI (v0.109.0) - 현대적이고 빠른 Python 웹 프레임워크
- **ORM**: SQLAlchemy (v2.0.25) - Python의 가장 강력한 ORM
- **Server**: Uvicorn - Lightning-fast ASGI 서버
- **Validation**: Pydantic v2 - 타입 기반 데이터 검증
- **Database**: MariaDB/MySQL
- **Database Driver**: PyMySQL
- **Environment**: python-dotenv

## 프로젝트 구조

```
fastapi-sqlalchemy/
├── config/
│   └── database.py         # SQLAlchemy 데이터베이스 설정
├── models/
│   └── user.py            # SQLAlchemy User 모델
├── schemas/
│   └── user.py            # Pydantic 스키마 (검증 및 직렬화)
├── routes/
│   └── users.py           # User CRUD 라우트
├── main.py                # FastAPI 애플리케이션 진입점
├── requirements.txt       # Python 의존성
├── .env                   # 환경변수
└── README.md
```

## FastAPI의 주요 특징

### 1. **자동 API 문서화** 🎯
- **Swagger UI**: http://localhost:3001/docs
- **ReDoc**: http://localhost:3001/redoc
- 코드 작성만으로 자동 생성!

### 2. **Pydantic 기반 검증**
- 타입 힌트로 자동 검증
- Email 검증 자동화
- 명확한 에러 메시지

### 3. **비동기 지원**
- async/await 네이티브 지원
- 높은 성능 (Node.js/Go 수준)

### 4. **타입 안전성**
- Python 타입 힌트 활용
- IDE 자동완성 지원
- 런타임 타입 검증

## 설치 및 실행

### 1. 가상 환경 생성 (권장)

```bash
# Windows
python -m venv venv
venv\Scripts\activate

# Linux/Mac
python3 -m venv venv
source venv/bin/activate
```

### 2. 의존성 설치

```bash
pip install -r requirements.txt
```

### 3. 환경변수 설정

`.env` 파일을 편집하여 데이터베이스 연결 정보를 입력하세요.

### 4. 서버 실행

```bash
# 개발 모드 (hot-reload)
python main.py

# 또는 uvicorn 직접 실행
uvicorn main:app --reload --host 0.0.0.0 --port 3001

# 프로덕션 모드
uvicorn main:app --host 0.0.0.0 --port 3001 --workers 4
```

서버는 기본적으로 `http://localhost:3001`에서 실행됩니다.

## API 문서

서버 실행 후 자동 생성되는 문서:

- **Swagger UI**: http://localhost:3001/docs
  - 대화형 API 테스트 가능
  - "Try it out" 버튼으로 바로 테스트

- **ReDoc**: http://localhost:3001/redoc
  - 깔끔한 문서 형식
  - 스키마 상세 설명

## API 엔드포인트

### Health Check
- `GET /` - 서버 상태 확인

### User CRUD Operations

#### 1. CREATE - 사용자 생성
```bash
POST /api/users/
Content-Type: application/json

{
  "email": "user@example.com",
  "username": "johndoe",
  "password": "password123"
}
```

**검증 규칙**:
- email: 유효한 이메일 형식 필수
- username: 1-50자
- password: 1-255자

#### 2. READ - 모든 사용자 조회
```bash
GET /api/users/
```

#### 3. READ - 특정 사용자 조회
```bash
GET /api/users/{user_id}
```

#### 4. UPDATE - 사용자 수정
```bash
PUT /api/users/{user_id}
Content-Type: application/json

{
  "email": "updated@example.com",
  "username": "updateduser",
  "password": "newpassword123"
}
```

**참고**: 모든 필드가 선택적 (Optional)

#### 5. DELETE - 사용자 삭제
```bash
DELETE /api/users/{user_id}
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
    "created_at": "2024-01-01T00:00:00",
    "updated_at": "2024-01-01T00:00:00"
  }
}
```

### 에러 응답
```json
{
  "detail": "User not found"
}
```

FastAPI는 HTTP 상태 코드도 자동으로 설정합니다:
- 200: OK
- 201: Created
- 404: Not Found
- 400: Bad Request
- 422: Validation Error

## FastAPI의 특별한 기능

### 1. **자동 데이터 검증**
```python
class UserCreate(BaseModel):
    email: EmailStr  # 이메일 형식 자동 검증
    username: str = Field(..., min_length=1, max_length=50)
    password: str = Field(..., min_length=1)
```

### 2. **의존성 주입 (Dependency Injection)**
```python
def get_user(user_id: int, db: Session = Depends(get_db)):
    # db 세션이 자동으로 주입됨
    # 함수 종료 시 자동으로 정리됨
```

### 3. **타입 힌트 기반**
```python
@router.get("/{user_id}", response_model=SuccessResponse)
def get_user(user_id: int, ...):
    # user_id는 자동으로 int로 변환
    # 실패 시 422 에러 자동 반환
```

### 4. **자동 문서화**
- 코드 주석이 API 문서에 자동 포함
- 스키마가 자동으로 문서화
- Request/Response 예제 자동 생성

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

## FastAPI vs Flask vs Django

| 특징 | FastAPI | Flask | Django |
|------|---------|-------|--------|
| **성능** | 매우 빠름 ⚡ | 보통 | 보통 |
| **비동기** | 네이티브 지원 | 미지원 | 3.0+부터 지원 |
| **타입 검증** | 자동 (Pydantic) | 수동 | 수동/Form |
| **API 문서** | 자동 생성 🎯 | 수동 | 수동 |
| **학습 곡선** | 중간 | 쉬움 | 어려움 |
| **ORM** | 선택 (SQLAlchemy) | 선택 | 내장 (Django ORM) |
| **적합한 경우** | 현대적 API | 작은 프로젝트 | 풀스택 앱 |

## FastAPI의 장점

1. **자동 문서화**: Swagger UI와 ReDoc 자동 생성
2. **빠른 성능**: Starlette 기반, 비동기 지원
3. **타입 안전성**: Pydantic으로 자동 검증
4. **현대적**: Python 3.7+ 타입 힌트 활용
5. **개발자 경험**: 뛰어난 IDE 지원
6. **에러 처리**: 명확한 검증 에러 메시지

## 프로덕션 배포

### Uvicorn with Workers
```bash
# 멀티 워커로 실행
uvicorn main:app --host 0.0.0.0 --port 3001 --workers 4
```

### Gunicorn + Uvicorn Workers
```bash
pip install gunicorn

gunicorn main:app \
  --workers 4 \
  --worker-class uvicorn.workers.UvicornWorker \
  --bind 0.0.0.0:3001
```

### Docker
```dockerfile
FROM python:3.11-slim

WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "3001"]
```

## 테스트

FastAPI는 테스트가 매우 쉽습니다:

```python
from fastapi.testclient import TestClient
from main import app

client = TestClient(app)

def test_create_user():
    response = client.post("/api/users/", json={
        "email": "test@example.com",
        "username": "testuser",
        "password": "test123"
    })
    assert response.status_code == 201
```

## 유용한 명령어

```bash
# 개발 서버 실행 (auto-reload)
uvicorn main:app --reload

# 다른 포트로 실행
uvicorn main:app --port 8000

# 로그 레벨 설정
uvicorn main:app --log-level debug

# 프로덕션 모드
uvicorn main:app --host 0.0.0.0 --port 3001 --workers 4
```

## 왜 FastAPI를 선택해야 하나?

- ✅ **최고의 성능**: Python 웹 프레임워크 중 가장 빠름
- ✅ **자동 검증**: Pydantic으로 타입 안전성 보장
- ✅ **자동 문서화**: 별도 작업 없이 Swagger/ReDoc 생성
- ✅ **현대적**: async/await, 타입 힌트 등 최신 Python 기능 활용
- ✅ **생산성**: 적은 코드로 많은 기능 구현
- ✅ **개발자 경험**: IDE 자동완성, 명확한 에러 메시지

FastAPI는 특히 **RESTful API 개발**에 최적화되어 있습니다!

## 라이선스

ISC
