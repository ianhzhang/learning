from flask import Flask,request
import graphene
import json

class Query(graphene.ObjectType):
    myQuery = graphene.String(name=graphene.String(default_value="stranger"))

    def resolve_myQuery(self, info, name):
        return {"rslt": 'Hello ' + name}


app = Flask(__name__)
schema = graphene.Schema(query=Query)
@app.route('/graphql', methods=['POST'])
def graphql():
    print "abcde3"
    print request.data
    data = json.loads(request.data)
    print data['query']
    data1 = schema.execute(data['query']).data
    print data1['myQuery']
    return data1['myQuery']

if __name__ == '__main__':
   app.run(debug=True)

# curl -H "Content-Type: application/json" -d '{"query":"{myQuery}"}' http://localhost:5000/graphql