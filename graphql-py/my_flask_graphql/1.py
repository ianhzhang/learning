from flask import Flask
from flask_graphql import GraphQLView
import graphene
from datetime import datetime

#from models import db_session
#from schema import schema

class User(graphene.ObjectType):
    id = graphene.ID()
    username = graphene.String()
    last_login = graphene.DateTime()

class Query(graphene.ObjectType):
    users = graphene.List(User)

    def resolve_users(self, info):
        return [
            User(username="Ian", last_login=datetime.now()),
            User(username="Kevin", last_login=datetime.now()),
            User(username="Lucas", last_login=datetime.now())
        ]

schema = graphene.Schema(query=Query)


app = Flask(__name__)

app.add_url_rule('/graphiql',  # the endpoint

                 # setting the view function
                 view_func=GraphQLView.as_view(
                     'graphql',

                     # we set the schema that we generate
                     schema=schema,

                     # We say that we *do* want a graphiql interface
                     graphiql=True
                 )
)


if __name__ == '__main__':
    app.run(host='0.0.0.0', debug=True)


'''
python3 1.py
http://localhost:5000/graphiql

{
  users {
    username
    lastLogin
  } 
}
'''