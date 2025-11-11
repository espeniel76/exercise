"""
FastAPI application entry point
"""
import os
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from dotenv import load_dotenv

from config.database import init_db
from routes.users import router as users_router

# Load environment variables
load_dotenv()

# Create FastAPI application
app = FastAPI(
    title="FastAPI CRUD API",
    description="FastAPI with SQLAlchemy ORM CRUD API",
    version="1.0.0",
    docs_url="/docs",
    redoc_url="/redoc"
)

# Configure CORS
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Initialize database
init_db()

# Include routers
app.include_router(users_router)


# Health check endpoint
@app.get("/", tags=["health"])
def health_check():
    """Health check endpoint"""
    return {
        "message": "FastAPI + SQLAlchemy CRUD API",
        "status": "running",
        "docs": "/docs",
        "redoc": "/redoc"
    }


if __name__ == "__main__":
    import uvicorn

    port = int(os.getenv("PORT", 3001))
    uvicorn.run(
        "main:app",
        host="0.0.0.0",
        port=port,
        reload=True
    )
