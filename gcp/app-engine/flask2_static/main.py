from flask import Flask, render_template, request,send_from_directory

app = Flask(__name__)

@app.route('/')
def index():
    print "Hello 1"
    return send_from_directory('static','index.html')

@app.route('/<path:path>')
def file(path):
    return send_from_directory('static',path)