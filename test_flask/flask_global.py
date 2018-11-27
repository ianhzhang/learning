from flask import Flask,request, jsonify
import json
from flask_cors import CORS
import time


# global variable
app = Flask(__name__)
CORS(app)
a = "hello g"
    
@app.route('/hello1', methods=["GET"])
def hello1():
    global a
    old_a = a 
    print old_a
    a = "hello 1"
    print request.remote_addr
    print old_a
    return jsonify({"rslt": old_a})
    
@app.route('/hello2', methods=["GET"])
def hello2():
    global a
    old_a = a 
    a = "hello 2"
    print request.remote_addr
    print "hello2"
    return jsonify({"rslt": old_a})
    


if __name__ == '__main__':
    app.run(debug=True, port=5000)

'''
It seems to me that flask can keep its global, at least in this example
'''