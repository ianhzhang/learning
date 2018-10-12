from flask import Flask,request, jsonify, make_response, render_template,Response,redirect
import json
from flask_cors import CORS


# global variable
application = Flask(__name__)
CORS(application)

    
@application.route('/api/hello', methods=["GET"])
def hello():
    print request.remote_addr
    return jsonify({"rslt": "hello World"})
    


if __name__ == '__main__':
    application.run(debug=True, port=5000, host='0.0.0.0')
    
