from flask import Flask, jsonify, request
from flask_cors import CORS
import json

print __name__
app = Flask(__name__)
CORS(app)

myFamily = [
    {
        "ln": "Zhang",
        "fn": "Kevin"
    },
    {
        "ln": "Zhang",
        "fn": "Lucas 1"
    },
    {
        "ln": "Zhang",
        "fn": "Ian"
    },
    {
        "ln": "Sun",
        "fn": "Hong"
    }
]

@app.route('/family/getAll', methods=["GET"])           # This is another url route
def getAll():
    print("GET get all")
    return jsonify(myFamily)

@app.route('/family/getPerson/<int:index>', methods=["GET"])           # This is another url route
def getPerson(index):
    print("GET get one")
    person = myFamily[index]
    return jsonify(person)

@app.route('/family/addPerson', methods=["POST"])           # This is another url route
def addPerson():
    print("POST: add a Person")
    payload_data = request.get_json(force=True)
    print payload_data
    print type(payload_data)
    myFamily.append(payload_data)
    return jsonify(payload_data)

@app.route('/family/deletePerson/<int:index>', methods=["DELETE"])           # This is another url route
def deletePerson(index):
    print ("Delete:", index)
    del myFamily[index]
    return jsonify({"result": "Success Kevin"})


if __name__ == "__main__":
    app.run(debug=True, port=5000)