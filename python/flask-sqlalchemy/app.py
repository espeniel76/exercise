"""
Flask application entry point
"""
import os
from flask import Flask, jsonify
from flask_cors import CORS
from dotenv import load_dotenv
from config.database import init_db
from routes.users import users_bp

# Load environment variables
load_dotenv()

# Create Flask application
app = Flask(__name__)

# Enable CORS
CORS(app)

# Initialize database
init_db(app)

# Register blueprints
app.register_blueprint(users_bp)

# Health check endpoint
@app.route('/', methods=['GET'])
def health_check():
    return jsonify({
        'message': 'Flask + SQLAlchemy CRUD API',
        'status': 'running'
    }), 200

if __name__ == '__main__':
    port = int(os.getenv('PORT', 3001))
    app.run(host='0.0.0.0', port=port, debug=True)
