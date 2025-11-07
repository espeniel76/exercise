const fastify = require('fastify')({ logger: true });
const cors = require('@fastify/cors');
require('dotenv').config();

const sequelize = require('./config/database');
const userRoutes = require('./routes/users');

const PORT = process.env.PORT || 3001;

// CORS 설정
fastify.register(cors);

// Routes 등록
fastify.register(userRoutes);

// Health check
fastify.get('/', async (request, reply) => {
  return {
    message: 'Fastify + Sequelize CRUD API',
    status: 'running'
  };
});

// 서버 시작 함수
const start = async () => {
  try {
    // 데이터베이스 연결 테스트
    await sequelize.authenticate();
    console.log('Database connection established successfully.');

    // 테이블 동기화 (개발 환경에서만 사용)
    await sequelize.sync({ alter: true });
    console.log('Database synchronized.');

    // Fastify 서버 시작
    await fastify.listen({ port: PORT, host: '0.0.0.0' });
    console.log(`Server is running on port ${PORT}`);
  } catch (error) {
    fastify.log.error(error);
    process.exit(1);
  }
};

start();
