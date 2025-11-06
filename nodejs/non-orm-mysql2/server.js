const express = require('express');
const bodyParser = require('body-parser');
require('dotenv').config();

const db = require('./config/database');
const userRoutes = require('./routes/users');

const app = express();
const PORT = process.env.PORT || 3001;

// Middleware
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

// Routes
app.use('/api/users', userRoutes);

// Health check
app.get('/', (req, res) => {
  res.json({
    message: 'Node.js + mysql2 CRUD API (Non-ORM)',
    status: 'running'
  });
});

// Database connection test and server start
const startServer = async () => {
  try {
    // 데이터베이스 연결 테스트
    await db.query('SELECT 1');
    console.log('Database connection established successfully.');

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
