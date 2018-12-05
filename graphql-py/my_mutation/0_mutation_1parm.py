import graphene
import json
from datetime import datetime

class User(graphene.ObjectType):
    id = graphene.ID()
    username = graphene.String()


##############################################
class CreateUser1(graphene.Mutation):
    class Arguments:
        #print("1---------1")                # at init time
        username = graphene.String()

    #print("2-------------2")                # at init time
    user = graphene.Field(User)

    def mutate(self, info, username):       
        #print("4--------4", username)        # username is Bob   

        user = User(username = username)
        return CreateUser1(user=user)

class Mutations(graphene.ObjectType):
    #print("3---------3")
    CreateUser2 = CreateUser1.Field()


def main():
    schema = graphene.Schema(mutation=Mutations)
    result = schema.execute(
        '''
            mutation {
                CreateUser2(username: "Bob") {
                    user {
                        username

                    }
                }
            }
        '''
    )

    print "Result ----------------------------"
    items = dict(result.data.items())
    print (json.dumps(items, indent=4))

if __name__ == '__main__':
    main()