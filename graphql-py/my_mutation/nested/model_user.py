import graphene

class Name(graphene.ObjectType):
    fn = graphene.String()
    ln = graphene.String()

class User(graphene.ObjectType):
    id = graphene.ID()
    name = graphene.Field(Name)



UserList = [
    User(id=1, name=Name(fn="Ian", ln="Zhang")),
    User(id=2, name=Name(fn="Kevin", ln="Zhang"))
]

