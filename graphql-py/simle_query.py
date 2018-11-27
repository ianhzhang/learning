import graphene

class Query(graphene.ObjectType):
    print "1--------------------"
    myQuery = graphene.String(name=graphene.String(default_value="stranger"))

    def resolve_myQuery(self, info, name):
        print "2----------------------------"
        print "2 info:", info
        print "2 name:", name
        return 'Hello ' + name

def main():
    print "a-----------------------------"
    schema = graphene.Schema(query=Query)
    print "b-------------------------------"
    result = schema.execute('{ myQuery }')
    print "c-----------------------------"
    print result.data['myQuery']

if __name__ == '__main__':
    main()


'''
1--------------------
a-----------------------------
b-------------------------------
2----------------------------
2 info: <graphql.execution.base.ResolveInfo object at 0x1057f25a0>
2 name: stranger
c-----------------------------
Hello stranger
'''