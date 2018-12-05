import graphene
import json
from datetime import datetime

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


def main():
    schema = graphene.Schema(query=Query)
    result = schema.execute(
        '''
            {
                users {
                    username
                }
            }
        '''
    )

    print result.data.items()
    print result.data['users']
    print "----------------------------"
    items = dict(result.data.items())
    print (json.dumps(items, indent=4))
if __name__ == '__main__':
    main()