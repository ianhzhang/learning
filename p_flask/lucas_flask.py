from flask import Flask,request, jsonify
from flask_cors import CORS

# global variable
app = Flask(__name__)
CORS(app)


@app.route('/hello', methods=["GET"])
def hello():
    returnDic = {"rslt": "hello World1"}
    print(returnDic)
    return jsonify(returnDic)

@app.route('/filelist', methods=["GET"])
def filelist_func():
    filelist = ["File1", "File2", "File3"]
    return jsonify(filelist)

@app.route('/param', methods=["GET"])
def param():
    params = {
        "param1": 23.2,
        "param2":45,
    }

    return jsonify(params)

@app.route('/curveData', methods=["GET"])
def curveData():
    curveData = [
        { "curve1": [1.0,2.0,3.0,4,0 ] },
        { "curve2": [31.30,32.0,43.0,54,60 ] },
        { "curve3": [61.0,62.0,63.0,74,0 ] }


    ]
    return jsonify(curveData)

if __name__ == '__main__':
    app.run(port="8085", debug=True)
