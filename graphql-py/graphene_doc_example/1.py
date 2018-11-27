import graphene

class Query(graphene.ObjectType):
    hello = graphene.String(argument=graphene.String(default_value="stranger"))

    def resolve_hello(self, info, argument):
        return 'Hello ' + argument

schema = graphene.Schema(query=Query)



# ==================================================

result = schema.execute('{ hello }')
print(result.data['hello']) # "Hello stranger"

# or passing the argument in the query
result = schema.execute('{ hello (argument: "ian") }')
print(result.data['hello']) # "Hello graph"