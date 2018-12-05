
import graphene
import json

from main_schema import Query, Mutations, mySchema
from model_user import User, UserList

def main():
    create()
    update()
    queryAll()
    delete()
    queryAll()
    pass

def create():
    print("create ---------------------------------------------------------")
    result = mySchema.execute(
        '''
            mutation {
                CreateUserMutation(personData: {id: 3, name: { fn: "Lucas", ln: "Zhang" } }) {
                    person {
                        id
                        name {
                            fn
                            ln
                        }
                    }
                }
            }
        '''
    )
    printResult(result)

def update():
    print("update ---------------------------------------------------------")
    result = mySchema.execute(
        '''
            mutation {
                UpdateUserMutaton(personData: {id:3, name:{fn:"Hong", ln:"Sun" } }) {
                    person {
                        id
                        name {
                            fn
                            ln
                        }
                    }
                }
            }
        '''
    )
    printResult(result)

def delete():
    print("delete ---------------------------------------------------------")
    result = mySchema.execute(
        '''
            mutation {
                DeleteUserMutaton(personData: {id:3, name:{fn:"Hong", ln:"Sun" } }) {
                    person {
                        id
                    }
                }
            }
        '''
    )
    printResult(result)


def queryAll():
    print("queryAll ---------------------------------------------------------")
    result = mySchema.execute(
        '''
            {
                users {
                    name {
                        fn
                        ln
                    }
                }
            }
        '''
    )
    printResult(result)



def printResult(result):
    print("Result ----------------------------")
    if result and result.data:
        items = dict(result.data.items())
        print (json.dumps(items, indent=4))
    else:
        print("None")

if __name__ == '__main__':
    main()