"""
URL routing for users app
"""
from django.urls import path
from . import views

urlpatterns = [
    path('users', views.get_users, name='get_users'),
    path('users/', views.create_user, name='create_user'),
    path('users/<int:user_id>', views.get_user, name='get_user'),
    path('users/<int:user_id>/', views.update_user, name='update_user'),
    path('users/<int:user_id>/delete', views.delete_user, name='delete_user'),
]
