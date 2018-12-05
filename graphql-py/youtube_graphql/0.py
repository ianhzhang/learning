import graphene

class Query(graphene.ObjectType):
    myQuery = graphene.String(name=graphene.String(default_value="stranger"))

    def resolve_myQuery(self, info, name):
        return 'Hello ' + name

def main():
    schema = graphene.Schema(query=Query)
    result = schema.execute(
        '''
        { 
            myQuery 
            }
        '''
    )
    print result.data.items()
    print result.data['myQuery']

if __name__ == '__main__':
    main()