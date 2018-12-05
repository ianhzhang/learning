import graphene

from model_user import User, UserList, UserList1

class PersonInput(graphene.InputObjectType):
    fn = graphene.String()
    ln = graphene.String()
    id = graphene.Int()

##############################################
class CreateUser1(graphene.Mutation):                   # I call it mutation action class
    class Arguments:
        personData = PersonInput()

    person = graphene.Field(User)

    def mutate(self, info, personData):
        print("create mutate")
        print(personData)       
        person = User(id = personData.id, fn = personData.fn, ln= personData.ln)
        UserList.append(person)
        return CreateUser1(person=person)

class UpdateUser1(graphene.Mutation):                   # I call it mutation action class
    class Arguments:
        personData = PersonInput()

    person = graphene.Field(User)

    def mutate(self, info, personData):
        print("update mutate")
        print(personData)

        global UserList
        id = personData.id
        for user in UserList:
            if user.id == id :
                print("found")
                user.fn = personData.fn
                user.ln = personData.ln   
        person = User(id = personData.id, fn = personData.fn, ln = personData.ln)
        return UpdateUser1(person=person)

class DeleteUser1(graphene.Mutation):                   # I call it mutation action class
    class Arguments:
        personData = PersonInput()

    person = graphene.Field(User)

    def mutate(self, info, personData):
        print("delete mutate")
        print(personData)

        global UserList
        id = personData.id

        index = -1
        for idx, user in enumerate(UserList):
            if user.id == id :
                del UserList[idx]
                break
         
        person = User(id = personData.id)
        return DeleteUser1(person=person)
