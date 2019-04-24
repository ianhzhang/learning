
from flask import Flask

app = Flask(__name__)

@app.route('/')
def form():
    return 'Hello World 1'