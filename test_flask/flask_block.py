from flask import Flask,request, jsonify
import json
from flask_cors import CORS
import time


# global variable
app = Flask(__name__)
CORS(app)

    
@app.route('/hello1', methods=["GET"])
def hello1():
    print request.remote_addr
    print "hello1"
    return jsonify({"rslt": "hello World1"})
    
@app.route('/hello2', methods=["GET"])
def hello2():
    print request.remote_addr
    print "hello2"
    cnt = 0
    while cnt<99999999:
        cnt = cnt + 1
        #print cnt
    #time.sleep(10)
    print "\ntimeout 2"
    return jsonify({"rslt": "hello World2"})
    


if __name__ == '__main__':
    app.run(debug=True, port=5000)
    
'''

localhost:5000/hello2 first
then localhsot:5000/hello1 second

hello1 returns without problem
hello2 tied with thread.

If hello1 and hello2 send from same brower tab: it will have:
timeout 2
127.0.0.1 - - [21/Nov/2018 11:08:30] "GET /hello2 HTTP/1.1" 200 -
Error on request:
Traceback (most recent call last):
  File "/usr/local/lib/python2.7/site-packages/werkzeug/serving.py", line 270, in run_wsgi
    execute(self.server.app)
  File "/usr/local/lib/python2.7/site-packages/werkzeug/serving.py", line 261, in execute
    write(data)
  File "/usr/local/lib/python2.7/site-packages/werkzeug/serving.py", line 236, in write
    self.send_header('Server', self.version_string())
  File "/usr/local/Cellar/python@2/2.7.15_1/Frameworks/Python.framework/Versions/2.7/lib/python2.7/BaseHTTPServer.py", line 412, in send_header
    self.wfile.write("%s: %s\r\n" % (keyword, value))
IOError: [Errno 32] Broken pipe

If it is not same tab.  Every thing will be fine.
'''