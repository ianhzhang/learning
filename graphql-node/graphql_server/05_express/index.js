'use strict'
console.log("Started")
const express = require('express')
const graphqlHTTP = require('express-graphql')
const { graphql, buildSchema } = require('graphql')

const PORT = process.env.PORT || 3000;
const server = express();

// 1. Define schema
// 1.1 type field:  create the description of the capabilities of graphql server
// 1.2 scema field: query/mutation.  The query use type Query.  The allows query the schema for foo.
const schema = buildSchema( `
    type Video {
        id: ID,
        title: String,
        duration: Int,
        watched: Boolean
    }

    type Query {
        video: Video
        videos: [Video]
    }


    type Schema {
        query: Query
    }
`);

// 2. What value to return is through resolver.
const videoA = {
    id: 'a',
    title: 'Create a GraphQL Schema',
    duration: 120,
    watched: true
}
const videoB = {
    id: 'b',
    title: 'Ember.js CLI',
    duration: 120,
    watched: false
}
const videos = [
    videoA,
    videoB
]

const resolvers = {
    video: () => ({
        id: '1',
        title: 'Foo',
        duration: 180,
        watched: true
    }),
    videos: () => videos
};

/*
// 3. Define query
const my_query = `
    query myFirstQuery {
        videos {
            id,
            title,
            duration,
            watched
        }
    }
`;
*/


/*
// 4. execute query
graphql(schema, my_query, resolvers)
    .then( (result) => console.log(result) )    // { data: { foo: 'bar' } }
    .catch( (error) => console.log(error));

*/

server.use('graphql', graphqlHTTP({
    schema, 
    graphiql: true,
    rootValue: resolvers
}))

server.listen( PORT, () => {
    console.log(`Listening on http://localhost:${PORT}`)
})





/*
{ data: { videos: [ [Object], [Object] ] } }

*/
