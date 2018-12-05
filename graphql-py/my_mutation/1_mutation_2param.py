# https://docs.graphene-python.org/en/latest/types/mutations/

import graphene
import json

class User(graphene.ObjectType):
    fn = graphene.String()
    ln = graphene.String()

class PersonInput(graphene.InputObjectType):
    fn = graphene.String()
    ln = graphene.String()

##############################################
class CreateUser1(graphene.Mutation):                   # I call it mutation action class
    class Arguments:
        personData = PersonInput()

    person = graphene.Field(User)

    def mutate(self, info, personData):
        print personData       
        person = User(fn = personData.fn, ln= personData.ln)
        return CreateUser1(person=person)

class Mutations(graphene.ObjectType):
    CreateUserMutation = CreateUser1.Field()            
    # CreateUserMutation matches the graphql user input #36.  CreateUser1 matches the Mutation Action class CreateUser1 #15


def main():
    schema = graphene.Schema(mutation=Mutations)
    result = schema.execute(
        '''
            mutation {
                CreateUserMutation(personData: {fn:"Bob", ln:"Zhang"}) {
                    person {
                        fn
                        ln
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