"""
URL configuration for Django CRUD project.
"""
from django.contrib import admin
from django.urls import path, include
from django.http import JsonResponse

def health_check(request):
    """Health check endpoint"""
    return JsonResponse({
        'message': 'Django CRUD API',
        'status': 'running'
    })

urlpatterns = [
    path('admin/', admin.site.urls),
    path('', health_check),
    path('api/', include('users.urls')),
]
