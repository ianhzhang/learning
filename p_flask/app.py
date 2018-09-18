from flask import Flask,request, jsonify, make_response, render_template,Response
import json
from flask_cors import CORS

from xhtml2pdf import pisa
from cStringIO import StringIO

# global variable
app = Flask(__name__)
CORS(app)

@app.route('/hello', methods=["GET"])
def hello():
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
    
