from flask import Flask, request, session, redirect, jsonify, make_response, render_template, Response, redirect, g, url_for
import json
from flask_cors import CORS

app = Flask(__name__)
CORS(app)
app.secret_key  = "223232456469"


@app.route('/', methods=["GET", "POST"])
def index():
    session.pop('user', None)
    if request.method == 'POST':
        if request.form['password'] == 'password':
            session['user'] = request.form['username']
            return redirect(url_for('protected'))

    return render_template('index.html')

@app.route('/protected', methods=["GET"])
def protected():
    session['method'] = 'try'
    if g.user:
        return render_template('protected.html')
    return redirect(url_for('index'))


@app.before_request
def before_request():
    g.user = None
    if 'user' in session:
        g.user = session['user']

@app.route('/get', methods=["GET"])
def get():
    if 'user' in session:
        return session['user']
    else:
        return 'not login'

@app.route('/drop', methods=["GET"])
def drop():
    print session.keys()
    session.pop('user', None)
    print session.keys()
    return 'Dropped'

if __name__ == '__main__':
    app.run(debug=True, port=5000)