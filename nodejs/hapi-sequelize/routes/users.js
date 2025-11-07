const User = require('../models/User');
const Joi = require('@hapi/joi');

const userRoutes = [
  // CREATE - 사용자 생성
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
      try {
        const { email, username, password } = request.payload;
        const user = await User.create({ email, username, password });

        return h.response({
          success: true,
          data: user
        }).code(201);
      } catch (error) {
        return h.response({
          success: false,
          error: error.message
        }).code(400);
      }
    }
  },

  // READ - 모든 사용자 조회
  {
    method: 'GET',
    path: '/api/users',
    handler: async (request, h) => {
      try {
        const users = await User.findAll({
          order: [['id', 'ASC']]
        });

        return {
          success: true,
          data: users
        };
      } catch (error) {
        return h.response({
          success: false,
          error: error.message
        }).code(500);
      }
    }
  },

  // READ - 특정 사용자 조회
  {
    method: 'GET',
    path: '/api/users/{id}',
    options: {
      validate: {
        params: Joi.object({
          id: Joi.number().integer().required()
        })
      }
    },
    handler: async (request, h) => {
      try {
        const user = await User.findByPk(request.params.id);

        if (!user) {
          return h.response({
            success: false,
            error: 'User not found'
          }).code(404);
        }

        return {
          success: true,
          data: user
        };
      } catch (error) {
        return h.response({
          success: false,
          error: error.message
        }).code(500);
      }
    }
  },

  // UPDATE - 사용자 수정
  {
    method: 'PUT',
    path: '/api/users/{id}',
    options: {
      validate: {
        params: Joi.object({
          id: Joi.number().integer().required()
        }),
        payload: Joi.object({
          email: Joi.string().email(),
          username: Joi.string().min(3).max(50),
          password: Joi.string().min(6)
        })
      }
    },
    handler: async (request, h) => {
      try {
        const { email, username, password } = request.payload;
        const user = await User.findByPk(request.params.id);

        if (!user) {
          return h.response({
            success: false,
            error: 'User not found'
          }).code(404);
        }

        await user.update({ email, username, password });

        return {
          success: true,
          data: user
        };
      } catch (error) {
        return h.response({
          success: false,
          error: error.message
        }).code(400);
      }
    }
  },

  // DELETE - 사용자 삭제
  {
    method: 'DELETE',
    path: '/api/users/{id}',
    options: {
      validate: {
        params: Joi.object({
          id: Joi.number().integer().required()
        })
      }
    },
    handler: async (request, h) => {
      try {
        const user = await User.findByPk(request.params.id);

        if (!user) {
          return h.response({
            success: false,
            error: 'User not found'
          }).code(404);
        }

        await user.destroy();

        return {
          success: true,
          message: 'User deleted successfully'
        };
      } catch (error) {
        return h.response({
          success: false,
          error: error.message
        }).code(500);
      }
    }
  }
];

module.exports = userRoutes;
