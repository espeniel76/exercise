"""
User routes - CRUD operations
"""
from fastapi import APIRouter, Depends, HTTPException, status
from sqlalchemy.orm import Session
from typing import List

from config.database import get_db
from models.user import User
from schemas.user import (
    UserCreate,
    UserUpdate,
    UserResponse,
    SuccessResponse,
    SuccessListResponse,
    DeleteResponse
)

router = APIRouter(prefix="/api/users", tags=["users"])


@router.post("/", response_model=SuccessResponse, status_code=status.HTTP_201_CREATED)
def create_user(user: UserCreate, db: Session = Depends(get_db)):
    """
    CREATE - Create a new user
    """
    # Check if email already exists
    if db.query(User).filter(User.email == user.email).first():
        raise HTTPException(
            status_code=status.HTTP_400_BAD_REQUEST,
            detail="Email already registered"
        )

    # Check if username already exists
    if db.query(User).filter(User.username == user.username).first():
        raise HTTPException(
            status_code=status.HTTP_400_BAD_REQUEST,
            detail="Username already taken"
        )

    # Create new user
    db_user = User(
        email=user.email,
        username=user.username,
        password=user.password
    )
    db.add(db_user)
    db.commit()
    db.refresh(db_user)

    return {"success": True, "data": db_user}


@router.get("/", response_model=SuccessListResponse)
def get_users(db: Session = Depends(get_db)):
    """
    READ - Get all users
    """
    users = db.query(User).all()
    return {"success": True, "data": users}


@router.get("/{user_id}", response_model=SuccessResponse)
def get_user(user_id: int, db: Session = Depends(get_db)):
    """
    READ - Get a single user by ID
    """
    user = db.query(User).filter(User.id == user_id).first()

    if not user:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND,
            detail="User not found"
        )

    return {"success": True, "data": user}


@router.put("/{user_id}", response_model=SuccessResponse)
def update_user(user_id: int, user_update: UserUpdate, db: Session = Depends(get_db)):
    """
    UPDATE - Update a user
    """
    user = db.query(User).filter(User.id == user_id).first()

    if not user:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND,
            detail="User not found"
        )

    # Update fields if provided
    if user_update.email is not None:
        # Check if email is already taken by another user
        existing = db.query(User).filter(
            User.email == user_update.email,
            User.id != user_id
        ).first()
        if existing:
            raise HTTPException(
                status_code=status.HTTP_400_BAD_REQUEST,
                detail="Email already registered"
            )
        user.email = user_update.email

    if user_update.username is not None:
        # Check if username is already taken by another user
        existing = db.query(User).filter(
            User.username == user_update.username,
            User.id != user_id
        ).first()
        if existing:
            raise HTTPException(
                status_code=status.HTTP_400_BAD_REQUEST,
                detail="Username already taken"
            )
        user.username = user_update.username

    if user_update.password is not None:
        user.password = user_update.password

    db.commit()
    db.refresh(user)

    return {"success": True, "data": user}


@router.delete("/{user_id}", response_model=DeleteResponse)
def delete_user(user_id: int, db: Session = Depends(get_db)):
    """
    DELETE - Delete a user
    """
    user = db.query(User).filter(User.id == user_id).first()

    if not user:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND,
            detail="User not found"
        )

    db.delete(user)
    db.commit()

    return {"success": True, "message": "User deleted successfully"}
