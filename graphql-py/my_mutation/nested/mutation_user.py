import graphene

from model_user import User, UserList, Name

class NameInput(graphene.InputObjectType):
    fn = graphene.String()
    ln = graphene.String()

class PersonInput(graphene.InputObjectType):
    id = graphene.Int()
    name = graphene.InputField(NameInput)
    

##############################################
class CreateUser1(graphene.Mutation):                   # I call it mutation action class
    class Arguments:
        personData = PersonInput()

    person = graphene.Field(User)

    def mutate(self, info, personData):
        print("create mutate")
        print(personData)
        id = personData.id
        name = personData.name
        print("id:",id)
        print("name:",name)       
        person = User(id = personData.id, name = name)
        UserList.append(person)
        return CreateUser1(person)

class UpdateUser1(graphene.Mutation):                   # I call it mutation action class
    class Arguments:
        personData = PersonInput()

    person = graphene.Field(User)

    def mutate(self, info, personData):
        print("update mutate")
        print(personData)

        global UserList
        id = personData.id
        name = personData.name
        for user in UserList:
            if user.id == id :
                print("found")
                user.name = name
        person = User(id = personData.id, name = name)
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
