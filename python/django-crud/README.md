# Django CRUD API

Django와 내장 Django ORM을 사용한 Python CRUD API

## 기술 스택
- **Framework**: Django (v5.0.0) - Python의 풀스택 웹 프레임워크
- **ORM**: Django ORM (Built-in) - Django의 강력한 내장 ORM
- **Database**: MariaDB/MySQL
- **Database Driver**: mysqlclient
- **Environment**: python-dotenv
- **CORS**: django-cors-headers

## 프로젝트 구조

```
django-crud/
├── config/                  # Django 프로젝트 설정
│   ├── __init__.py
│   ├── settings.py         # 전역 설정
│   ├── urls.py             # URL 라우팅
│   ├── wsgi.py             # WSGI 설정
│   └── asgi.py             # ASGI 설정
├── users/                   # Users 앱
│   ├── __init__.py
│   ├── models.py           # User 모델 정의
│   ├── views.py            # CRUD 뷰 함수
│   ├── urls.py             # 앱 URL 라우팅
│   ├── admin.py            # Django Admin 설정
│   └── apps.py             # 앱 설정
├── manage.py               # Django 관리 명령어
├── requirements.txt        # Python 의존성
├── .env                    # 환경변수
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

**Note**: `mysqlclient`는 C 라이브러리가 필요합니다:
- **Windows**: [MySQL C API](https://dev.mysql.com/downloads/c-api/) 설치 필요
- **Linux**: `sudo apt-get install python3-dev default-libmysqlclient-dev build-essential`
- **Mac**: `brew install mysql`

대안으로 `PyMySQL`을 사용하려면:
```bash
pip install PyMySQL
# settings.py에 추가: import pymysql; pymysql.install_as_MySQLdb()
```

### 3. 환경변수 설정

`.env` 파일을 편집하여 데이터베이스 연결 정보를 입력하세요.

### 4. 데이터베이스 마이그레이션

```bash
# 마이그레이션 파일 생성
python manage.py makemigrations

# 마이그레이션 실행
python manage.py migrate
```

### 5. (선택사항) Admin 계정 생성

```bash
python manage.py createsuperuser
```

### 6. 서버 실행

```bash
# 개발 모드
python manage.py runserver 0.0.0.0:3001

# 기본 포트 (8000)
python manage.py runserver
```

서버는 기본적으로 `http://localhost:3001`에서 실행됩니다.

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
PUT /api/users/:id/
Content-Type: application/json

{
  "email": "updated@example.com",
  "username": "updateduser",
  "password": "newpassword123"
}
```

#### 5. DELETE - 사용자 삭제
```bash
DELETE /api/users/:id/delete
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

## Django의 주요 특징

### 1. Django ORM (Built-in)
- **Model 정의**: Python 클래스로 데이터베이스 테이블 정의
- **자동 마이그레이션**: `makemigrations`와 `migrate` 명령어
- **QuerySet API**: 강력하고 직관적인 쿼리 인터페이스
- **관계 정의**: ForeignKey, ManyToMany 등 쉬운 관계 설정
- **트랜잭션**: 자동 트랜잭션 관리
- **데이터베이스 추상화**: 여러 DB 엔진 지원

### 2. Django Admin
- **자동 관리 패널**: `/admin`에서 모델 관리 가능
- **커스터마이징**: `admin.py`에서 리스트, 필터, 검색 설정
- **즉시 사용 가능**: 별도 구현 없이 CRUD 인터페이스 제공

```bash
# Admin 패널 접속
http://localhost:3001/admin
```

### 3. Django 프레임워크
- **Batteries Included**: ORM, Admin, 인증, 세션 등 모든 기능 내장
- **보안**: CSRF, XSS, SQL Injection 기본 방어
- **확장성**: 앱 기반 모듈화 아키텍처
- **MTV 패턴**: Model-Template-View (MVC와 유사)

### 4. 마이그레이션 시스템
```bash
# 모델 변경사항 감지 및 마이그레이션 파일 생성
python manage.py makemigrations

# 데이터베이스에 적용
python manage.py migrate

# 마이그레이션 상태 확인
python manage.py showmigrations

# 특정 마이그레이션으로 롤백
python manage.py migrate users 0001
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
SECRET_KEY=django-insecure-dev-key-please-change-in-production
DEBUG=True
```

## Django vs Flask vs FastAPI

| 특징 | Django | Flask | FastAPI |
|------|--------|-------|---------|
| 철학 | Batteries Included | Micro | Modern, Fast |
| ORM | Built-in (Django ORM) | 별도 (SQLAlchemy) | 별도 (SQLAlchemy) |
| Admin 패널 | 내장 | 별도 구현 필요 | 별도 구현 필요 |
| 인증/권한 | 내장 | 별도 패키지 | 별도 패키지 |
| 학습 곡선 | 가파름 | 완만함 | 중간 |
| 적합한 프로젝트 | 대규모 풀스택 | 작은 API, 유연성 필요 | 고성능 API |
| 성능 | 보통 | 보통 | 매우 빠름 |
| 비동기 | 3.0+부터 지원 | 미지원 (기본) | 네이티브 지원 |

## Django의 장점

1. **빠른 개발**: 모든 기능이 내장되어 있어 설정이 최소화
2. **Admin 패널**: 자동 생성되는 강력한 관리 인터페이스
3. **ORM**: 매우 강력하고 직관적인 쿼리 API
4. **보안**: 기본적으로 안전한 설정
5. **문서화**: 매우 상세한 공식 문서
6. **커뮤니티**: 거대한 생태계와 풍부한 패키지

## 프로덕션 배포

프로덕션 환경에서는 WSGI 서버 사용을 권장합니다:

```bash
# Gunicorn 설치
pip install gunicorn

# Gunicorn으로 실행
gunicorn config.wsgi:application --bind 0.0.0.0:3001

# 워커 프로세스 지정
gunicorn config.wsgi:application --bind 0.0.0.0:3001 --workers 4
```

## 유용한 Django 명령어

```bash
# 프로젝트 구조 확인
python manage.py check

# 데이터베이스 셸
python manage.py dbshell

# Python 셸 (Django 환경)
python manage.py shell

# 정적 파일 수집
python manage.py collectstatic

# 테스트 실행
python manage.py test

# 개발 서버 실행 (다른 포트)
python manage.py runserver 8080
```

## 라이선스

ISC
