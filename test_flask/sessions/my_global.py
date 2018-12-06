from flask import Flask, session, jsonify, request
from flask_cors import CORS

# global variable
app = Flask(__name__)
CORS(app)

cnt = 0

@app.route('/', methods=["GET"])
def index():
    global cnt
    cnt = cnt +1
    return jsonify({"rslt": cnt})

@app.route('/drop', methods=["GET"])
def drop():
    session.pop('cnt', None)
    return 'Dropped'

if __name__ == '__main__':
    app.run(debug=True, port=5000)
    
