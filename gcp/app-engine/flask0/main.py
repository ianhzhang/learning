import sys
import os

sys.path.insert(1, os.path.join(os.path.abspath('.'), 'lib'))

# pip install -t lib flask

from flask import Flask

app = Flask(__name__)

@app.route('/')
def form():
    return 'Hello World'