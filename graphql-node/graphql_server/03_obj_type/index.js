'use strict'
console.log("Started")
const { graphql, buildSchema } = require('graphql')

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
    }


    type Schema {
        query: Query
    }
`);

// 2. What value to return is through resolver.
const resolvers = {
    video: () => ({
        id: '1',
        title: 'Foo',
        duration: 180,
        watched: true
    })
};


// 3. Define query
const my_query = `
    query myFirstQuery {
        video {
            id,
            title,
            duration,
            watched
        }
    }
`;

// 4. execute query
graphql(schema, my_query, resolvers)
    .then( (result) => console.log(result) )    // { data: { foo: 'bar' } }
    .catch( (error) => console.log(error));


/*
{ data: 
   { video: { id: '1', title: 'Foo', duration: 180, watched: true } } }

*/
