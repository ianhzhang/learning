import graphene
import json
from datetime import datetime

class User(graphene.ObjectType):
    id = graphene.ID()
    username = graphene.String()
    last_login = graphene.DateTime(required=False)

class Query(graphene.ObjectType):
    users = graphene.List(User, first=graphene.Int())

    def resolve_users(self, info, first):
        return [
            User(username="Ian", last_login=datetime.now()),
            User(username="Kevin", last_login=datetime.now()),
            User(username="Lucas", last_login=datetime.now())
        ][:first]



##############################################
class CreateUser(graphene.Mutation):
    class Arguments:
        username = graphene.String()

    user = graphene.Field(User)

    def mutate(self, info, username):
        user = User(username = username)
        return CreateUser(user=user)

class Mutations(graphene.ObjectType):
    create_user = CreateUser.Field()


def main():
    schema = graphene.Schema(query=Query, mutation=Mutations)
    result = schema.execute(
        '''
            mutation createUser {
                createUser(username: "Bob") {
                    user {
                        username
                    }
                }
            }
        '''
    )

    print result.data.items()
    print "----------------------------"
    items = dict(result.data.items())
    print (json.dumps(items, indent=4))
if __name__ == '__main__':
    main()