'use strict'
console.log("Started")
const express = require('express')
const graphqlHTTP = require('express-graphql')

const { 
    GraphQLSchema,
    GraphQLObjectType,
    GraphQLID,
    GraphQLString,
    GraphQLInt,
    GraphQLBoolean
 } = require('graphql')

const PORT = process.env.PORT || 3000;
const server = express();

// Data
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

getVideoById = (id) = new Promise((resolve) => {
    const [video] = videos.filter( (video) => {
        return video.id == id;
    });
    resolve(video);
})



// 1. Define schema

// 1.1  Our end model in graphql Type.
videType = new GraphQLObjectType( {
    name: 'Video',
    description: 'A video on Egghead.io',
    fields: {
        id: {
            type: GraphQLID,
            description: 'The id of the video'
        },
        title: {
            type: GraphQLString,
            description: 'The title of the video'
        },
        duration: {
            type: GraphQLInt,
            description: 'The duration of the video'
        },
        watched: {
            type: GraphQLBoolean,
            description: 'Whether or not the viewer has watch the video'
        }
    }

})

// 1.2 What can we query
queryType - new GraphQLObjectType ({
    name: 'QueryType',
    description: 'The root query type.',
    fields: {
        video: {
            type: videType,
            args : {                                        // Argument
                id: {
                    type: GraphQLID,
                    description: 'The id of the video.'
                }
            },
            resolve: (_, args) => {
                return getVideoById(args.id)                // defined earlier.
            }
        }
    }
})
const schema = new GraphQLSchema({
    query: queryType,
    //mutation,
    //subscription
})


// 2. What value to return is through resolver.



server.use('graphql', graphqlHTTP({
    schema, 
    graphiql: true
}))

server.listen( PORT, () => {
    console.log(`Listening on http://localhost:${PORT}`)
})


/*
{ data: { videos: [ [Object], [Object] ] } }

*/
