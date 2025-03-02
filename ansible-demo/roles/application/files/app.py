# Basic flask api to demonstrate ansible deployment
"""Sample Application
This is a sample application to demonstrate ansible deployment
"""
from flask import Flask
app = Flask(__name__)


@app.route('/')
def hello_world():
    """Sample API
    This is a sample API to demonstrate ansible deployment

    Returns:
        string: Plain old hello world message
    """
    return 'Hello, World!'


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
