const User = require('../models/User');

async function userRoutes(fastify, options) {
  // CREATE - 사용자 생성
  fastify.post('/api/users', async (request, reply) => {
    try {
      const { email, username, password } = request.body;
      const user = await User.create({ email, username, password });

      reply.code(201).send({
        success: true,
        data: user
      });
    } catch (error) {
      reply.code(400).send({
        success: false,
        error: error.message
      });
    }
  });

  // READ - 모든 사용자 조회
  fastify.get('/api/users', async (request, reply) => {
    try {
      const users = await User.findAll({
        order: [['id', 'ASC']]
      });

      reply.send({
        success: true,
        data: users
      });
    } catch (error) {
      reply.code(500).send({
        success: false,
        error: error.message
      });
    }
  });

  // READ - 특정 사용자 조회
  fastify.get('/api/users/:id', async (request, reply) => {
    try {
      const user = await User.findByPk(request.params.id);

      if (!user) {
        return reply.code(404).send({
          success: false,
          error: 'User not found'
        });
      }

      reply.send({
        success: true,
        data: user
      });
    } catch (error) {
      reply.code(500).send({
        success: false,
        error: error.message
      });
    }
  });

  // UPDATE - 사용자 수정
  fastify.put('/api/users/:id', async (request, reply) => {
    try {
      const { email, username, password } = request.body;
      const user = await User.findByPk(request.params.id);

      if (!user) {
        return reply.code(404).send({
          success: false,
          error: 'User not found'
        });
      }

      await user.update({ email, username, password });

      reply.send({
        success: true,
        data: user
      });
    } catch (error) {
      reply.code(400).send({
        success: false,
        error: error.message
      });
    }
  });

  // DELETE - 사용자 삭제
  fastify.delete('/api/users/:id', async (request, reply) => {
    try {
      const user = await User.findByPk(request.params.id);

      if (!user) {
        return reply.code(404).send({
          success: false,
          error: 'User not found'
        });
      }

      await user.destroy();

      reply.send({
        success: true,
        message: 'User deleted successfully'
      });
    } catch (error) {
      reply.code(500).send({
        success: false,
        error: error.message
      });
    }
  });
}

module.exports = userRoutes;
