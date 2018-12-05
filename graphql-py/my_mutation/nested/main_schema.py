import graphene

from model_user import User, UserList

from mutation_user import CreateUser1, UpdateUser1, DeleteUser1

class Query(graphene.ObjectType):
    users = graphene.List(User)

    def resolve_users(self, info):
        print("resolve_users")
        return UserList

class Mutations(graphene.ObjectType):
    CreateUserMutation = CreateUser1.Field(description="Create User")
    # CreateUserMutation matches the graphql user input #36.  CreateUser1 matches the Mutation Action class CreateUser1 #15
    UpdateUserMutaton = UpdateUser1.Field(description="Update User")
    DeleteUserMutaton = DeleteUser1.Field(description="Delete User")


mySchema = graphene.Schema(query=Query, mutation=Mutations)