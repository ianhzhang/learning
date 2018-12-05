import graphene
import json

class Query(graphene.ObjectType):
    isStaff = graphene.Boolean()

    def resolve_isStaff(self, info):
        return True


def main():
    schema = graphene.Schema(query=Query)
    result = schema.execute(
        '''
            {
                isStaff
            }
        '''
    )

    print result.data.items()
    print result.data['isStaff']
    print "----------------------------"
    items = dict(result.data.items())
    print (json.dumps(items, indent=4))
if __name__ == '__main__':
    main()