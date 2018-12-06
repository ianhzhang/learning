from flask import Flask, session, jsonify, request
from flask_cors import CORS

# global variable
app = Flask(__name__)
CORS(app)

app.secret_key  = "223232456469"
    
@app.route('/', methods=["GET"])
def index():
    print request.remote_addr
    session['user'] = "Ian"
    return jsonify({"rslt": "hello World"})
    
@app.route('/get', methods=["GET"])
def get():
    if 'user' in session:
        return session['user']
    else:
        return 'not login'

@app.route('/drop', methods=["GET"])
def drop():
    session.pop('user', None)
    return 'Dropped'

if __name__ == '__main__':
    app.run(debug=True, port=5000)
    
