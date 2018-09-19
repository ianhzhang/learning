from flask import Flask,request, jsonify, make_response, render_template,Response,redirect
import json
from flask_cors import CORS

from xhtml2pdf import pisa
from cStringIO import StringIO
from functools import wraps 


# global variable
app = Flask(__name__)
CORS(app)



@app.before_request
def limit_remote_addr():
    print request.remote_addr
    if request.remote_addr != '0.0.0.0' and request.remote_addr != '127.0.0.1':
        return "Cannot find the page.", 404

    
@app.route('/hello', methods=["GET"])
def hello():
    print request.remote_addr
    return jsonify({"rslt": "hello World"})
    
@app.route('/getInfo', methods=["GET"])
def getInfo():
    user = {
        'username': "Ian Zhang"
    }
    htmlContent = render_template('t.html', user=user)
    pdfBuff = StringIO()

    #pisa.CreatePDF(StringIO(htmlContent.encode('utf-8')), pdfBuff)
    resultFile = open("abc.pdf", "w+b")
    pisa.CreatePDF(htmlContent.encode('utf-8'), dest= resultFile)

    return htmlContent

@app.route("/getPdf/<id>", methods=["GET"])
def getPdf(id):

    user = {
        'username': "Ian Zhang"
    }
    htmlContent = render_template('r1.html', user=user)
    print htmlContent
    pdfBuff = StringIO()

    pisa.CreatePDF(StringIO(htmlContent.encode('utf-8')), pdfBuff)
    #print pdfBuff.getvalue()
    response = make_response(pdfBuff.getvalue())
    response.headers['Content-Type'] = 'application/pdf'
    return response



if __name__ == '__main__':
    app.run(debug=True)
    
