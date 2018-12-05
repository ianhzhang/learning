import graphene


class User(graphene.ObjectType):
    id = graphene.ID()
    fn = graphene.String()
    ln = graphene.String()


UserList = [
    User(id=1, fn="Ian", ln="Zhang"),
    User(id=2, fn="Kevin", ln="Zhang")
]

UserList1 = [
    User(id=1, fn="Ian", ln="Zhang"),
    User(id=2, fn="Kevin", ln="Zhang")
]