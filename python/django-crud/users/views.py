"""
User views - CRUD operations
"""
import json
from django.http import JsonResponse
from django.views.decorators.csrf import csrf_exempt
from django.views.decorators.http import require_http_methods
from .models import User


@csrf_exempt
@require_http_methods(["POST"])
def create_user(request):
    """
    CREATE - Create a new user
    """
    try:
        data = json.loads(request.body)

        # Validate required fields
        if not data.get('email') or not data.get('username') or not data.get('password'):
            return JsonResponse({
                'success': False,
                'error': 'Email, username, and password are required'
            }, status=400)

        # Create new user
        user = User.objects.create(
            email=data['email'],
            username=data['username'],
            password=data['password']
        )

        return JsonResponse({
            'success': True,
            'data': user.to_dict()
        }, status=201)

    except json.JSONDecodeError:
        return JsonResponse({
            'success': False,
            'error': 'Invalid JSON'
        }, status=400)
    except Exception as e:
        return JsonResponse({
            'success': False,
            'error': str(e)
        }, status=500)


@require_http_methods(["GET"])
def get_users(request):
    """
    READ - Get all users
    """
    try:
        users = User.objects.all()
        return JsonResponse({
            'success': True,
            'data': [user.to_dict() for user in users]
        }, status=200)

    except Exception as e:
        return JsonResponse({
            'success': False,
            'error': str(e)
        }, status=500)


@require_http_methods(["GET"])
def get_user(request, user_id):
    """
    READ - Get a single user by ID
    """
    try:
        user = User.objects.get(id=user_id)
        return JsonResponse({
            'success': True,
            'data': user.to_dict()
        }, status=200)

    except User.DoesNotExist:
        return JsonResponse({
            'success': False,
            'error': 'User not found'
        }, status=404)
    except Exception as e:
        return JsonResponse({
            'success': False,
            'error': str(e)
        }, status=500)


@csrf_exempt
@require_http_methods(["PUT"])
def update_user(request, user_id):
    """
    UPDATE - Update a user
    """
    try:
        user = User.objects.get(id=user_id)
        data = json.loads(request.body)

        # Update fields if provided
        if data.get('email'):
            user.email = data['email']
        if data.get('username'):
            user.username = data['username']
        if data.get('password'):
            user.password = data['password']

        user.save()

        return JsonResponse({
            'success': True,
            'data': user.to_dict()
        }, status=200)

    except User.DoesNotExist:
        return JsonResponse({
            'success': False,
            'error': 'User not found'
        }, status=404)
    except json.JSONDecodeError:
        return JsonResponse({
            'success': False,
            'error': 'Invalid JSON'
        }, status=400)
    except Exception as e:
        return JsonResponse({
            'success': False,
            'error': str(e)
        }, status=500)


@csrf_exempt
@require_http_methods(["DELETE"])
def delete_user(request, user_id):
    """
    DELETE - Delete a user
    """
    try:
        user = User.objects.get(id=user_id)
        user.delete()

        return JsonResponse({
            'success': True,
            'message': 'User deleted successfully'
        }, status=200)

    except User.DoesNotExist:
        return JsonResponse({
            'success': False,
            'error': 'User not found'
        }, status=404)
    except Exception as e:
        return JsonResponse({
            'success': False,
            'error': str(e)
        }, status=500)
