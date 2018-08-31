from flask import Flask,request, jsonify, make_response
import json

#pip install pyjwt==1.4.2
import jwt

# sudo pip install passlib
from passlib.apps import custom_app_context as pwd_context


# global variable
app = Flask(__name__)

@app.route('/register1', methods=["POST"])
def register1():
    print("register1")
    post_data = request.form
    user_id = post_data['user']
    password = post_data['pass']

    password_hass = pwd_context.encrypt(password)
    print "password_hash:"
    print(password_hass)

    ## Store user_id and password_hash to database

    return make_response(jsonify({"result": "success"}))


# curl -v -X POST -d 'user=izhang&pass=admin' http://localhost:5000/login1
@app.route('/login1', methods=["POST"])
def login1():
    print("login1")
    post_data = request.form
    user_id = post_data['user']
    password = post_data['pass']
    print('user:', user_id)
    print('pass:', password)

    # check user id and password
    # assume we get the password_hash from database according to user id
    password_hash = "$6$rounds=656000$24aa93hP1zm2sJOM$cHPM9/jDaYsZ.7aMXMZitXiHxyrRC/G.5ObcV5MlWS1Bb3.zTamaN7QkQl6/AoVB/puBAur/uPfb.l4IaYJc61"
    
    verify_rslt = pwd_context.verify(password, password_hash)
    print("verify_rslt:", verify_rslt)

    if verify_rslt is not True:
        return make_response(jsonify({"result": "Fail"}))

    # create access token
    payload = {
        "sub": user_id
    }

    auth_token = jwt.encode(payload, "abcd", algorithm='HS256')
    print(auth_token)   # eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJpemhhbmcifQ.INRzfZ4YzSP-vAK1NYbyCkX8E2ddVDlhLcV5SkIOz7s
    print(auth_token.decode())

    respObj = {
        'status': 'success',
        'auth_token': auth_token.decode()
    }
    # we should put the auth_token in cookie

    return make_response(jsonify(respObj))

@app.route('/getInfo', methods=["GET"])
def getInfo():
    token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJpemhhbmcifQ.INRzfZ4YzSP-vAK1NYbyCkX8E2ddVDlhLcV5SkIOz7s"
    rslt = jwt.decode(token,"abcd")
    print(rslt)
    return "success"

if __name__ == '__main__':
    app.run(debug=True)
    