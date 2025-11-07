const Router = require('koa-router');
const User = require('../models/User');

const router = new Router();

// CREATE - 사용자 생성
router.post('/api/users', async (ctx) => {
  try {
    const { email, username, password } = ctx.request.body;
    const user = await User.create({ email, username, password });

    ctx.status = 201;
    ctx.body = {
      success: true,
      data: user
    };
  } catch (error) {
    ctx.status = 400;
    ctx.body = {
      success: false,
      error: error.message
    };
  }
});

// READ - 모든 사용자 조회
router.get('/api/users', async (ctx) => {
  try {
    const users = await User.findAll({
      order: [['id', 'ASC']]
    });

    ctx.body = {
      success: true,
      data: users
    };
  } catch (error) {
    ctx.status = 500;
    ctx.body = {
      success: false,
      error: error.message
    };
  }
});

// READ - 특정 사용자 조회
router.get('/api/users/:id', async (ctx) => {
  try {
    const user = await User.findByPk(ctx.params.id);

    if (!user) {
      ctx.status = 404;
      ctx.body = {
        success: false,
        error: 'User not found'
      };
      return;
    }

    ctx.body = {
      success: true,
      data: user
    };
  } catch (error) {
    ctx.status = 500;
    ctx.body = {
      success: false,
      error: error.message
    };
  }
});

// UPDATE - 사용자 수정
router.put('/api/users/:id', async (ctx) => {
  try {
    const { email, username, password } = ctx.request.body;
    const user = await User.findByPk(ctx.params.id);

    if (!user) {
      ctx.status = 404;
      ctx.body = {
        success: false,
        error: 'User not found'
      };
      return;
    }

    await user.update({ email, username, password });

    ctx.body = {
      success: true,
      data: user
    };
  } catch (error) {
    ctx.status = 400;
    ctx.body = {
      success: false,
      error: error.message
    };
  }
});

// DELETE - 사용자 삭제
router.delete('/api/users/:id', async (ctx) => {
  try {
    const user = await User.findByPk(ctx.params.id);

    if (!user) {
      ctx.status = 404;
      ctx.body = {
        success: false,
        error: 'User not found'
      };
      return;
    }

    await user.destroy();

    ctx.body = {
      success: true,
      message: 'User deleted successfully'
    };
  } catch (error) {
    ctx.status = 500;
    ctx.body = {
      success: false,
      error: error.message
    };
  }
});

module.exports = router;
