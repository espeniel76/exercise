const Koa = require('koa');
const bodyParser = require('koa-bodyparser');
const cors = require('@koa/cors');
require('dotenv').config();

const sequelize = require('./config/database');
const userRoutes = require('./routes/users');

const app = new Koa();
const PORT = process.env.PORT || 3001;

// Middleware
app.use(cors());
app.use(bodyParser());

// Error handling middleware
app.use(async (ctx, next) => {
  try {
    await next();
  } catch (err) {
    ctx.status = err.status || 500;
    ctx.body = {
      success: false,
      error: err.message
    };
    console.error('Error:', err);
  }
});

// Routes
app.use(userRoutes.routes());
app.use(userRoutes.allowedMethods());

// Health check
const Router = require('koa-router');
const healthRouter = new Router();

healthRouter.get('/', async (ctx) => {
  ctx.body = {
    message: 'Koa + Sequelize CRUD API',
    status: 'running'
  };
});

app.use(healthRouter.routes());

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
