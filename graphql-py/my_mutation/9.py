import graphene
import json
from datetime import datetime
from pprint import pprint

class User(graphene.ObjectType):
    id = graphene.ID()
    fn = graphene.String()
    ln = graphene.String()

class PersonInput(graphene.InputObjectType):
    id = graphene.Int()
    fn = graphene.String()
    ln = graphene.String()


class Query(graphene.ObjectType):
    users = graphene.List(User)

    def resolve_users(self, info):
        print("resolve_users")
        return usersList

usersList = [
    User(id=1, fn="Ian", ln="Zhang"),
    User(id=2, fn="Kevin", ln="Zhang")
]

##############################################

class CreateUser(graphene.Mutation):
    class Arguments:
        personData = PersonInput()

    person = graphene.Field(User)

    def mutate(self, info, personData):
        print personData       
        person = User(id = personData.id, fn = personData.fn, ln= personData.ln)
        usersList.append(person)
        return CreateUser(person=person)

class Mutations(graphene.ObjectType):
    CreateUserMutation = CreateUser.Field()


def main():
    schema = graphene.Schema(query=Query, mutation=Mutations)
    
    print("---------1.  Query -------------------------------")
    result = schema.execute(
        '''
            {
                users {
                    id
                    fn
                    ln
                }
            }
        '''
    )

    items = dict(result.data.items())
    print(json.dumps(items, indent=4))
    

    print("---------2.  Mutation -------------------------------")
    result = schema.execute(
        '''
            mutation {
                CreateUserMutation(personData: {id:3, fn:"Lucas", ln:"Zhang"}) {
                    person {
                        id
                        fn
                        ln
                    }
                }
            }
        '''
    )

    items = dict(result.data.items())
    print(json.dumps(items, indent=4))

    print("---------3.  Query after insert -------------------------------")
    result = schema.execute(
        '''
            {
                users {
                    id
                    fn
                    ln
                }
            }
        '''
    )

    print(result.data.items())
    items = dict(result.data.items())
    print(json.dumps(items, indent=4))

if __name__ == '__main__':
    main()