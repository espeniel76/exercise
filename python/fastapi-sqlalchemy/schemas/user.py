"""
Pydantic schemas for User model
"""
from datetime import datetime
from typing import Optional
from pydantic import BaseModel, EmailStr, Field


class UserBase(BaseModel):
    """Base User schema"""
    email: EmailStr = Field(..., description="User email address")
    username: str = Field(..., min_length=1, max_length=50, description="Username")


class UserCreate(UserBase):
    """Schema for creating a user"""
    password: str = Field(..., min_length=1, max_length=255, description="User password")


class UserUpdate(BaseModel):
    """Schema for updating a user"""
    email: Optional[EmailStr] = Field(None, description="User email address")
    username: Optional[str] = Field(None, min_length=1, max_length=50, description="Username")
    password: Optional[str] = Field(None, min_length=1, max_length=255, description="User password")


class UserResponse(UserBase):
    """Schema for user response"""
    id: int
    password: str
    created_at: datetime
    updated_at: datetime

    class Config:
        from_attributes = True


class SuccessResponse(BaseModel):
    """Generic success response"""
    success: bool = True
    data: UserResponse


class SuccessListResponse(BaseModel):
    """Success response for list of users"""
    success: bool = True
    data: list[UserResponse]


class ErrorResponse(BaseModel):
    """Error response"""
    success: bool = False
    error: str


class DeleteResponse(BaseModel):
    """Response for delete operation"""
    success: bool = True
    message: str
