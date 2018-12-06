from flask import Flask, session, jsonify, request
from flask_cors import CORS

# global variable
app = Flask(__name__)
CORS(app)

app.secret_key  = "223232www456469"

@app.route('/', methods=["GET"])
def index():

    print session.keys()
    if 'cntKey' in session:
        cnt = session['cntKey']
        cnt = cnt + 1
        session['cntKey'] = cnt
    else:
        cnt = 0
        session['cntKey'] = cnt

    return jsonify({"rslt": cnt})

@app.route('/drop', methods=["GET"])
def drop():
    session.pop('cntKey', None)
    return 'Dropped'

if __name__ == '__main__':
    app.run(debug=True, port=5000)
    
'''
Differnt brwoser will have differnt session object.
But different tab in the same browser will share  the same session
'''