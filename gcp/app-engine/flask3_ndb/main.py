from flask import Flask, render_template, url_for, redirect, request,send_from_directory
from google.appengine.ext import ndb

app = Flask(__name__)

class Music(ndb.Model):
    song = ndb.StringProperty()
    singer = ndb.StringProperty()


@app.route('/')
def index():
    print "Index"
    
    results = Music.query()
    print results

    for each in results:
        print "singer:", each.singer
        print "song:", each.song
        print "\n"
    
    return render_template('index.html', all_items=results)
    #return send_from_directory('static','index.html')


@app.route('/save', methods=['POST'])
def save():
    print "Save"
    song = request.form['song']
    singer = request.form['singer']
    print "Song:", song
    print "Singer:", singer
    music = Music()
    music.song = song
    music.singer = singer
    music.put()

    return redirect(url_for('index'))   # this is the function name, not url path

@app.route('/search', methods=['GET'])
def search():
    print "Search"
    # https://cloud.google.com/appengine/docs/standard/python/ndb/queries
    results = Music.query()
    print results

    for each in results:
        print "singer:", each.singer
        print "song:", each.song
        print "\n"
    

    return redirect(url_for('index'))   # this is the function name, not url path

# Other files.
@app.route('/<path:path>')
def file(path):
    return send_from_directory('static',path)