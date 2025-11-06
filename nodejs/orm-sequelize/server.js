const express = require('express');
const bodyParser = require('body-parser');
require('dotenv').config();

const sequelize = require('./config/database');
const userRoutes = require('./routes/users');

const app = express();
const PORT = process.env.PORT || 3000;

// Middleware
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

// Routes
app.use('/api/users', userRoutes);

// Health check
app.get('/', (req, res) => {
  res.json({
    message: 'Node.js + Sequelize CRUD API',
    status: 'running'
  });
});

// Database connection and server start
const startServer = async () => {
  try {
    // 데이터베이스 연결 테스트
    await sequelize.authenticate();
    console.log('Database connection established successfully.');

    // 테이블 동기화 (개발 환경에서만 사용)
    await sequelize.sync({ alter: true });
    console.log('Database synchronized.');

    // 서버 시작
    app.listen(PORT, () => {
      console.log(`Server is running on port ${PORT}`);
    });
  } catch (error) {
    console.error('Unable to connect to the database:', error);
    process.exit(1);
  }
};

startServer();
