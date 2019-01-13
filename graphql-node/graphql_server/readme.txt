egghead.io
Build a graphql server

yarn init -y
yarn add graphql


Lesson 1
mkdir 1_getting_started
touch 1_getting_started/index.js
node 1_getting_started/index.js

Lesson 2

Lesson 5:
 yarn express express-graphql

 node 5_express/index.js

 In GraphiQL browser:
 http://localhost:300/qraphiql
   {
       videos {
           title
       }
   }

Lesson 6:  Passing the argument
In GraphIQL browser:
    {
        video(id: 'a') {
            title
        }
    }

Lesson 10: mutation:
In GraphiQL browser:
    mutation M {
        createVideo(title: "Foo", duration: 300, released: false) {
            id,
            title
        }
    }

Lesson 11:
In GraphiQL browser:
    mutation M {
        createVideo(video: {
            title: "Foo", 
            duration: 300, 
            released: false
        }) {
            id,
            title
        }
    }



Other resources:
https://malcoded.com/posts/graphql-with-angular-apollo
https://qiita.com/alokrawat050/items/7fc8a240e047ea6c2e06
https://stackoverflow.com/questions/49618392/when-to-use-watchquery-or-query-in-apollo-angular

python client:
https://github.com/graphql-python/gql
https://stackoverflow.com/questions/45957784/how-to-consume-the-github-graphql-api-using-python
https://medium.com/@hugo__df/python-graphql-client-requests-example-using-gql-1b5ea3506f9b
https://codewithhugo.com/python-graphql-client-requests-example-using-gql/