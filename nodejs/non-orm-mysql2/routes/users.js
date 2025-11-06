const express = require('express');
const router = express.Router();
const db = require('../config/database');

// CREATE - 사용자 생성
router.post('/', async (req, res) => {
  try {
    const { email, username, password } = req.body;

    const [result] = await db.query(
      'INSERT INTO users (email, username, password) VALUES (?, ?, ?)',
      [email, username, password]
    );

    const [users] = await db.query(
      'SELECT * FROM users WHERE id = ?',
      [result.insertId]
    );

    res.status(201).json({
      success: true,
      data: users[0]
    });
  } catch (error) {
    res.status(400).json({
      success: false,
      error: error.message
    });
  }
});

// READ - 모든 사용자 조회
router.get('/', async (req, res) => {
  try {
    const [users] = await db.query('SELECT * FROM users ORDER BY id ASC');

    res.json({
      success: true,
      data: users
    });
  } catch (error) {
    res.status(500).json({
      success: false,
      error: error.message
    });
  }
});

// READ - 특정 사용자 조회
router.get('/:id', async (req, res) => {
  try {
    const [users] = await db.query(
      'SELECT * FROM users WHERE id = ?',
      [req.params.id]
    );

    if (users.length === 0) {
      return res.status(404).json({
        success: false,
        error: 'User not found'
      });
    }

    res.json({
      success: true,
      data: users[0]
    });
  } catch (error) {
    res.status(500).json({
      success: false,
      error: error.message
    });
  }
});

// UPDATE - 사용자 수정
router.put('/:id', async (req, res) => {
  try {
    const { email, username, password } = req.body;

    const [result] = await db.query(
      'UPDATE users SET email = ?, username = ?, password = ? WHERE id = ?',
      [email, username, password, req.params.id]
    );

    if (result.affectedRows === 0) {
      return res.status(404).json({
        success: false,
        error: 'User not found'
      });
    }

    const [users] = await db.query(
      'SELECT * FROM users WHERE id = ?',
      [req.params.id]
    );

    res.json({
      success: true,
      data: users[0]
    });
  } catch (error) {
    res.status(400).json({
      success: false,
      error: error.message
    });
  }
});

// DELETE - 사용자 삭제
router.delete('/:id', async (req, res) => {
  try {
    const [result] = await db.query(
      'DELETE FROM users WHERE id = ?',
      [req.params.id]
    );

    if (result.affectedRows === 0) {
      return res.status(404).json({
        success: false,
        error: 'User not found'
      });
    }

    res.json({
      success: true,
      message: 'User deleted successfully'
    });
  } catch (error) {
    res.status(500).json({
      success: false,
      error: error.message
    });
  }
});

module.exports = router;
