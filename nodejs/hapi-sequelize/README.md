# Hapi.js CRUD with Sequelize ORM

Hapi.js와 Sequelize ORM을 사용한 Node.js CRUD API

## 기술 스택
- Hapi.js (v21.3.2) - 엔터프라이즈급 웹 프레임워크
- Joi (v17.1.1) - 강력한 데이터 검증
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
hapi-sequelize/
├── config/
│   └── database.js          # Sequelize 데이터베이스 설정
├── models/
│   └── User.js              # User 모델
├── routes/
│   └── users.js             # Users 라우트 (Joi 검증 포함)
├── server.js                # Hapi.js 서버 진입점
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

## Hapi.js의 특징

### 1. 설정 기반 (Configuration-Centric)
- 라우트를 객체로 정의
- 명시적이고 구조화된 설정
- 가독성 높은 코드

### 2. 내장 입력 검증 (Joi)
- 강력한 스키마 검증
- 자동 데이터 검증
- 타입 안정성
- 에러 메시지 자동 생성

### 3. 플러그인 아키텍처
- 모듈화된 구조
- 재사용 가능한 컴포넌트
- 확장성

### 4. 보안 기능
- 내장 보안 헤더
- CSRF 보호
- XSS 방어

### 5. 엔터프라이즈 지원
- Walmart에서 개발 및 사용
- 대규모 트래픽 처리
- 안정성과 신뢰성

## Hapi.js vs Express vs Fastify

| 특징 | Hapi.js | Express | Fastify |
|------|---------|---------|---------|
| 설정 방식 | Configuration | Code | Code |
| 검증 | Joi 내장 | 별도 라이브러리 | JSON Schema |
| 플러그인 | 강력함 | 미들웨어 | 강력함 |
| 학습 곡선 | 가파름 | 완만함 | 보통 |
| 성능 | 보통 | 보통 | 매우 빠름 |
| 보안 | 매우 강함 | 기본 | 강함 |
| 엔터프라이즈 | 최적 | 보통 | 좋음 |

## Joi 검증 예시

```javascript
// 라우트에서 자동으로 검증됨
{
  method: 'POST',
  path: '/api/users',
  options: {
    validate: {
      payload: Joi.object({
        email: Joi.string().email().required(),
        username: Joi.string().min(3).max(50).required(),
        password: Joi.string().min(6).required()
      })
    }
  },
  handler: async (request, h) => {
    // payload가 이미 검증됨
    // 실패 시 자동으로 400 에러 반환
  }
}
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

## Hapi.js를 선택해야 하는 경우

1. **엔터프라이즈급 애플리케이션**
   - 대규모 팀 프로젝트
   - 높은 안정성 요구
   - 장기 유지보수

2. **강력한 검증이 필요한 경우**
   - 복잡한 데이터 구조
   - API 입력 검증
   - 타입 안정성

3. **보안이 중요한 경우**
   - 금융 서비스
   - 의료 시스템
   - 민감한 데이터 처리

4. **명확한 구조가 필요한 경우**
   - 팀 협업
   - 코드 가독성
   - 명시적 설정 선호

## Hapi.js의 장점

✅ **강력한 검증**: Joi를 통한 완벽한 데이터 검증
✅ **보안**: 엔터프라이즈급 보안 기능 내장
✅ **구조화**: 명확하고 일관된 코드 구조
✅ **플러그인**: 재사용 가능한 모듈화
✅ **안정성**: 대규모 프로덕션 환경에서 검증됨
✅ **문서화**: 우수한 공식 문서

## Hapi.js의 실제 사용 사례

- **Walmart**: 블랙 프라이데이 트래픽 처리
- **Auth0**: 인증 서비스
- **npm**: 패키지 레지스트리
- **Mozilla**: 다양한 서비스

## 성능 vs 안정성

Hapi.js는 최고 성능보다는 **안정성, 보안, 유지보수성**을 우선시합니다.
- Express보다 약간 느림
- Fastify보다 느림
- 하지만 엔터프라이즈 환경에서 더 안정적
