from flask import Flask,request
import graphene
import json

class Query(graphene.ObjectType):
    myQuery = graphene.String(name=graphene.String(default_value="stranger"))

    def resolve_myQuery(self, info, name):
        print "resolve name: ", name
        return {"rslt": 'Hello ' + name}

schema = graphene.Schema(query=Query)


app = Flask(__name__)

@app.route('/graphql', methods=['POST'])
def graphql():
    print "1=================================="
    print request.data
    print "2------------------------------------"
    data = json.loads(request.data)
    print data['query']
    print "3---------------------------------------"
    data1 = schema.execute(data['query']).data
    #print data1['myQuery']
    print "4-----------------------------------------"
    return data1['myQuery']

if __name__ == '__main__':
   app.run(debug=True)

# curl -H "Content-Type: application/json" -d '{"query":"{myQuery}"}' http://localhost:5000/graphql


'''
http://localhost:5000/graphql

{
	"query":"{myQuery}"
}

1==================================
{
	"query":"{myQuery}"
}
2------------------------------------
{myQuery}
3---------------------------------------
{'rslt': 'Hello stranger'}
4-----------------------------------------



========================================================
http://localhost:5000/graphql

{
	"query":"{myQuery (name: \"ian\")}"
}
cannot be
	"query":"{myQuery (name: 'ian')}"
'''