"""
User routes - CRUD operations
"""
from flask import Blueprint, request, jsonify
from models.user import User
from config.database import db

users_bp = Blueprint('users', __name__)

@users_bp.route('/api/users', methods=['POST'])
def create_user():
    """
    CREATE - Create a new user
    """
    try:
        data = request.get_json()

        # Validate required fields
        if not data or not data.get('email') or not data.get('username') or not data.get('password'):
            return jsonify({
                'success': False,
                'error': 'Email, username, and password are required'
            }), 400

        # Create new user
        user = User(
            email=data['email'],
            username=data['username'],
            password=data['password']
        )

        db.session.add(user)
        db.session.commit()

        return jsonify({
            'success': True,
            'data': user.to_dict()
        }), 201

    except Exception as e:
        db.session.rollback()
        return jsonify({
            'success': False,
            'error': str(e)
        }), 500


@users_bp.route('/api/users', methods=['GET'])
def get_users():
    """
    READ - Get all users
    """
    try:
        users = User.query.all()
        return jsonify({
            'success': True,
            'data': [user.to_dict() for user in users]
        }), 200

    except Exception as e:
        return jsonify({
            'success': False,
            'error': str(e)
        }), 500


@users_bp.route('/api/users/<int:user_id>', methods=['GET'])
def get_user(user_id):
    """
    READ - Get a single user by ID
    """
    try:
        user = User.query.get(user_id)

        if not user:
            return jsonify({
                'success': False,
                'error': 'User not found'
            }), 404

        return jsonify({
            'success': True,
            'data': user.to_dict()
        }), 200

    except Exception as e:
        return jsonify({
            'success': False,
            'error': str(e)
        }), 500


@users_bp.route('/api/users/<int:user_id>', methods=['PUT'])
def update_user(user_id):
    """
    UPDATE - Update a user
    """
    try:
        user = User.query.get(user_id)

        if not user:
            return jsonify({
                'success': False,
                'error': 'User not found'
            }), 404

        data = request.get_json()

        # Update fields if provided
        if data.get('email'):
            user.email = data['email']
        if data.get('username'):
            user.username = data['username']
        if data.get('password'):
            user.password = data['password']

        db.session.commit()

        return jsonify({
            'success': True,
            'data': user.to_dict()
        }), 200

    except Exception as e:
        db.session.rollback()
        return jsonify({
            'success': False,
            'error': str(e)
        }), 500


@users_bp.route('/api/users/<int:user_id>', methods=['DELETE'])
def delete_user(user_id):
    """
    DELETE - Delete a user
    """
    try:
        user = User.query.get(user_id)

        if not user:
            return jsonify({
                'success': False,
                'error': 'User not found'
            }), 404

        db.session.delete(user)
        db.session.commit()

        return jsonify({
            'success': True,
            'message': 'User deleted successfully'
        }), 200

    except Exception as e:
        db.session.rollback()
        return jsonify({
            'success': False,
            'error': str(e)
        }), 500
