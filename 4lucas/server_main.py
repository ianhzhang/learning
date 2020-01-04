from flask import Flask,request, jsonify,send_from_directory,send_file,render_template
from flask_cors import CORS
from _DataLoader import DataLoader

# global variable
app = Flask(__name__, static_url_path='', template_folder='static' , static_folder='static')
# enable cross domain
CORS(app)

# User can get index.html without .html
@app.route("/")
def index():
    return render_template('index.html')

# Test
@app.route("/test")
def test():
    return "This is test"

@app.route('/imageInfoAll', methods=["GET"])
def imageInfoAll():
    return jsonify(imageList)

@app.route('/images/<path:filename>')
def get_file(filename):
    #end_from_directory('data', path,  mimetype='image/jpeg')
    send_file('data/myimg.jpg', mimetype='image/jpeg')
    return ""


if __name__ == '__main__':
    # Load imageList (search all jpeg, analysis image file)
    dataLoader = DataLoader()
    imageList = dataLoader.loadImageInfo("static/assets/tree_for_candidate")
    app.run(host= '0.0.0.0', port="8085", debug=True)
