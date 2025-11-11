"""
Database configuration for Flask application
"""
import os
from flask_sqlalchemy import SQLAlchemy

db = SQLAlchemy()

def init_db(app):
    """
    Initialize database connection
    """
    # Build database URI
    db_host = os.getenv('DB_HOST', 'localhost')
    db_user = os.getenv('DB_USER', 'root')
    db_password = os.getenv('DB_PASSWORD', '')
    db_name = os.getenv('DB_NAME', 'espeniel')
    db_port = os.getenv('DB_PORT', '3306')

    # MySQL/MariaDB connection string
    database_uri = f'mysql+pymysql://{db_user}:{db_password}@{db_host}:{db_port}/{db_name}?charset=utf8mb4'

    app.config['SQLALCHEMY_DATABASE_URI'] = database_uri
    app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False
    app.config['SQLALCHEMY_ENGINE_OPTIONS'] = {
        'pool_size': 5,
        'pool_recycle': 3600,
        'pool_pre_ping': True,
    }

    db.init_app(app)

    with app.app_context():
        db.create_all()
        print('Database synchronized')

    return db
