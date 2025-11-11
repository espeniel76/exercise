"""
User model definition
"""
from django.db import models


class User(models.Model):
    """User model for CRUD operations"""

    id = models.AutoField(primary_key=True)
    email = models.CharField(max_length=255, unique=True)
    username = models.CharField(max_length=50, unique=True)
    password = models.CharField(max_length=255)
    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)

    class Meta:
        db_table = 'users'
        ordering = ['id']

    def __str__(self):
        return self.username

    def to_dict(self):
        """Convert model instance to dictionary"""
        return {
            'id': self.id,
            'email': self.email,
            'username': self.username,
            'password': self.password,
            'created_at': self.created_at.isoformat() if self.created_at else None,
            'updated_at': self.updated_at.isoformat() if self.updated_at else None,
        }
