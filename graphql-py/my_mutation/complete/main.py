
import graphene
import json

from main_schema import Query, Mutations, mySchema
from model_user import User, UserList

def main():
    print("0. len:=", len(UserList))
    create()
    print("1. len:=", len(UserList))
    update()
    print("2. len:=", len(UserList))
    delete()
    print("3. len:=", len(UserList))
    queryAll()

def create():
    print("create ---------------------------------------------------------")
    result = mySchema.execute(
        '''
            mutation {
                CreateUserMutation(personData: {id: 3, fn:"Bob", ln:"Zhang"}) {
                    person {
                        id
                        fn
                        ln
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
                UpdateUserMutaton(personData: {id:3, fn:"Hong", ln:"Sun"}) {
                    person {
                        id
                        fn
                        ln
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
                DeleteUserMutaton(personData: {id:3, fn:"Hong", ln:"Sun"}) {
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
                    id
                    fn
                    ln
                }
            }
        '''
    )
    printResult(result)



def printResult(result):
    print("Result ----------------------------")
    items = dict(result.data.items())
    print (json.dumps(items, indent=4))

if __name__ == '__main__':
    main()