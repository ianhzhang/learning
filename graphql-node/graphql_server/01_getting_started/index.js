'use strict'
console.log("Started")
const { graphql, buildSchema } = require('graphql')

// 1. Define schema
// 1.1 type field:  create the description of the capabilities of graphql server
// 1.2 scema field: query/mutation.  The query use type Query.  The allows query the schema for foo.
const schema = buildSchema( `
    type Query {
        foo: String
    }


    type Schema {
        query: Query
    }
`);

// 2. What value to return is through resolver.
const resolvers = {
    foo: () => 'bar'
};


// 3. Define query
const my_query = `
    query myFirstQuery {
        foo
    }
`;

// 4. execute query
graphql(schema, my_query, resolvers)
    .then( (result) => console.log(result) )    // { data: { foo: 'bar' } }
    .catch( (error) => console.log(error));



// This is a server code.   When we use node or express to receive my_query string.
// We execute the step 4.