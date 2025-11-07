const Hapi = require('@hapi/hapi');
require('dotenv').config();

const sequelize = require('./config/database');
const userRoutes = require('./routes/users');

const PORT = process.env.PORT || 3001;

const init = async () => {
  // Hapi 서버 생성
  const server = Hapi.server({
    port: PORT,
    host: '0.0.0.0',
    routes: {
      cors: {
        origin: ['*']
      }
    }
  });

  // Health check route
  server.route({
    method: 'GET',
    path: '/',
    handler: (request, h) => {
      return {
        message: 'Hapi.js + Sequelize CRUD API',
        status: 'running'
      };
    }
  });

  // User routes 등록
  server.route(userRoutes);

  try {
    // 데이터베이스 연결 테스트
    await sequelize.authenticate();
    console.log('Database connection established successfully.');

    // 테이블 동기화 (개발 환경에서만 사용)
    await sequelize.sync({ alter: true });
    console.log('Database synchronized.');

    // 서버 시작
    await server.start();
    console.log(`Server is running on port ${server.info.port}`);
  } catch (error) {
    console.error('Error starting server:', error);
    process.exit(1);
  }
};

// 에러 핸들링
process.on('unhandledRejection', (err) => {
  console.error('Unhandled rejection:', err);
  process.exit(1);
});

init();
