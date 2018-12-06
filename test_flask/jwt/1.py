from flask import Flask, request, jsonify, make_response, render_template, Response, redirect
import jwt
import json
from flask_cors import CORS
import datetime
from functools import wraps

# https://www.youtube.com/watch?v=J5bIPtEbS0Q
# https://www.youtube.com/watch?v=WxGBoY5iNXY

# global variable
app = Flask(__name__)
CORS(app)

# our own decorator
def token_required(f):
    @wraps(f)
    def decorated(*args, **kwargs):
        token = request.args.get('token')
        print(token)
        if not token:
            return jsonify({'message': 'Token is missing!'}), 403
        try:
            data = jwt.decode(token, "mysecret")
        except:
            return jsonify({'message': 'token is invalid'}), 403
        return f(*args, **kwargs)
    return decorated


@app.route('/unprotected', methods=["GET"])
def unprotected():
    return jsonify({"message": "Anyone can view this"})
    

#http://localhost:5000/protected?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyIjoiYWJjIiwiZXhwIjoxNTQzNDIyNDI4fQ.gYUDvxSviOtSZqZiAhB2BAJ9BziNtqktDO1Zt7Dqgd0
@app.route('/protected', methods=["GET"])
@token_required
def protected():
    return jsonify({"message": "This is only available for people with valid tokens"})
    
@app.route('/login', methods=["GET"])
def login():
    auth = request.authorization
    print ("auth:", auth)

    if auth and auth.password == 'password':
        print(datetime.datetime.utcnow())       # 2018-11-28 16:01:37.433323
        print(datetime.timedelta(minutes=5))    # 0:05:00
        print(datetime.datetime.utcnow() + datetime.timedelta(minutes=5))
        token = jwt.encode({'user': auth.username, 
                            'exp': datetime.datetime.utcnow() + datetime.timedelta(minutes=5)
                           }, 'mysecret' )
        print(token)
        print(token.decode('UTF-8'))
        return jsonify({"token": token.decode('UTF-8')})

    return make_response("Could not verify! ", 401, {'WWW-Authenticate': 'Basic realm="login Required"'})

if __name__ == '__main__':
    app.run(debug=True, port=5000)
    
