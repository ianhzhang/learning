from flask import Flask, request, jsonify, make_response, render_template, Response, redirect
import json
from flask_cors import CORS




# global variable
app = Flask(__name__)
CORS(app)

    
@app.route('/set', methods=["GET"])
def setcookie():
    resp = make_response('Setting cookie!')
    resp.set_cookie('framework', 'flask')
    return resp

@app.route('/get', methods=["GET"])
def getcookie():
    framework = request.cookies.get('framework')

    return 'The framework is ' + framework

if __name__ == '__main__':
    app.run(debug=True, port=5000)