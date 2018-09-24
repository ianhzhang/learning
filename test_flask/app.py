from flask import Flask,request, jsonify, make_response, render_template,Response,redirect
import json
from flask_cors import CORS

from xhtml2pdf import pisa
from cStringIO import StringIO
from functools import wraps 


# global variable
app = Flask(__name__)
CORS(app)

    
@app.route('/hello', methods=["GET"])
def hello():
    print request.remote_addr
    return jsonify({"rslt": "hello World"})
    


if __name__ == '__main__':
    app.run(debug=True, port=5001)
    
