-- ========================================
-- Node.js CRUD 프로젝트 데이터베이스 설정
-- 데이터베이스: espeniel
-- ========================================

-- 데이터베이스 생성 (이미 존재한다면 스킵)
CREATE DATABASE IF NOT EXISTS espeniel
  CHARACTER SET utf8mb4
  COLLATE utf8mb4_unicode_ci;

-- espeniel 데이터베이스 사용
USE espeniel;

-- ========================================
-- users 테이블 생성
-- ========================================

DROP TABLE IF EXISTS users;

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

-- ========================================
-- 샘플 데이터 삽입 (선택사항)
-- ========================================

INSERT INTO users (email, username, password) VALUES
('user1@example.com', 'user1', 'password123'),
('user2@example.com', 'user2', 'password456'),
('user3@example.com', 'user3', 'password789');

-- ========================================
-- 확인 쿼리
-- ========================================

SELECT * FROM users;
