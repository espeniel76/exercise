# Flask CRUD with SQLAlchemy ORM

Flask와 SQLAlchemy ORM을 사용한 Python CRUD API

## 기술 스택
- **Framework**: Flask (v3.0.0) - Python의 경량 웹 프레임워크
- **ORM**: SQLAlchemy (v3.1.1) - Python에서 가장 널리 사용되는 ORM
- **Database**: MariaDB/MySQL
- **Database Driver**: PyMySQL
- **Environment**: python-dotenv

## 프로젝트 구조

```
flask-sqlalchemy/
├── config/
│   └── database.py      # 데이터베이스 연결 설정
├── models/
│   └── user.py          # User 모델 정의
├── routes/
│   └── users.py         # User 라우트 핸들러 (Blueprint)
├── app.py               # 애플리케이션 진입점
├── requirements.txt     # Python 의존성
├── .env                 # 환경변수
└── README.md
```

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
# 개발 모드 (디버그 활성화)
python app.py

# 프로덕션 모드
flask run --host=0.0.0.0 --port=3001
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
    "created_at": "2024-01-01T00:00:00",
    "updated_at": "2024-01-01T00:00:00"
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

### SQLAlchemy 기능
- **ORM Mapping**: 클래스 기반 데이터 모델링
- **Auto Migration**: create_all()을 통한 자동 테이블 생성
- **Session Management**: 트랜잭션 관리 및 롤백 지원
- **Connection Pooling**: 효율적인 DB 커넥션 관리
- **Query API**: 강력한 쿼리 빌더
- **Relationship**: Foreign Key 및 관계 정의 지원

### Flask 프레임워크
- **경량**: 미니멀한 코어, 필요한 기능만 추가
- **Blueprint**: 모듈화된 라우트 관리
- **WSGI**: 표준 Python 웹 서버 인터페이스
- **Jinja2**: 내장 템플릿 엔진 (API에서는 미사용)
- **Flask-CORS**: CORS 지원

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

## Node.js/Go 버전과의 비교

이 프로젝트는 다른 언어 구현과 동일한 구조와 기능을 제공합니다:

| 기능 | Node.js (Sequelize) | Go (GORM) | Python (SQLAlchemy) |
|------|---------------------|-----------|---------------------|
| Framework | Express/Fastify | Gin/Echo | Flask |
| ORM | Sequelize | GORM | SQLAlchemy |
| 라우팅 | Middleware/Plugin | Function | Blueprint |
| 타입 | JavaScript | 정적 타입 | 동적 타입 (Type hints 지원) |
| 성능 | 빠름 | 매우 빠름 | 보통 |
| 개발 속도 | 빠름 | 보통 | 매우 빠름 |

## Python의 장점

- **읽기 쉬운 코드**: 간결하고 명확한 문법
- **빠른 개발**: 적은 코드로 많은 기능 구현
- **풍부한 생태계**: 데이터 분석, ML/AI 통합 용이
- **강력한 ORM**: SQLAlchemy의 유연성과 기능성
- **REPL 지원**: 대화형 개발 및 디버깅

## 프로덕션 배포

프로덕션 환경에서는 WSGI 서버 사용을 권장합니다:

```bash
# Gunicorn 설치
pip install gunicorn

# Gunicorn으로 실행
gunicorn -w 4 -b 0.0.0.0:3001 app:app
```

## 라이선스

ISC
