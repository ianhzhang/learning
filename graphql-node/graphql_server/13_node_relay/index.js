'use strict'
console.log("Started")
const express = require('express')
const graphqlHTTP = require('express-graphql')

const { 
    GraphQLSchema,
    GraphQLObjectType,
    GraphQLNonNull,                         // Lesson 8.
    GraphQLList,                            // Lesson 9.
    GraphQLInputObjectType,                 // Lesson 11.
    GraphQLInterfaceType,                   // Lesson 12
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

const getVideoById = (id) = new Promise((resolve) => {
    const [video] = videos.filter( (video) => {
        return video.id == id;
    });
    resolve(video);
});

const getVideos = () => new Promise( (resolve) => resolve(videos) )

const createVideo = ({title, duration, released }) => {
    const video = {
        id: (new Buffer(title,'utf8')).toString('base64'),
        title,
        duration,
        released
    };
    videos.push(video);
    return video;
}

const nodeInterface = new GraphQLInputObjectType({              // Lesson 12
    name: 'Node',
    fields: {
        id: {
            type: new GraphQLNonNull(GraphQLID)
        }
    },
    resolveType: (object) => {                                  // lesson 12
        if(object.title) {
            return videoType
        }
    }
})

// 1. Define schema

// 1.1  Our end model in graphql Type.
videoType = new GraphQLObjectType( {
    name: 'Video',
    description: 'A video on Egghead.io',
    fields: {
        id: {
            type: new GraphQLNonNull(GraphQLID),                // Lesson 12
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
        released: {
            type: GraphQLBoolean,
            description: 'Whether or not the video has been released'
        }
    },
    interfaces: [nodeInterface]

})

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// 1.2 What can we query
queryType - new GraphQLObjectType ({
    name: 'QueryType',
    description: 'The root query type.',
    fields: {
        videos: {
            type: new GraphQLList(videoType),
            resolve: getVideos()
        },
        video: {
            type: videoType,
            args : {                                        // Argument
                id: {
                    type: new GraphQLNonNull(GraphQLID),            // Lesson 8  (without id: return erro)
                    description: 'The id of the video.'
                }
            },
            resolve: (_, args) => {
                return getVideoById(args.id)                // defined earlier.
            }
        }
    }
})

const videoInputType = new GraphQLInputObjectType({         // Lesson 11
    name: 'VideoInput',
    fields: {
        title: {
            type: new GraphQLNonNull(GraphQLString),
            description: 'The title of the video'
        },
        duration: {
            type: GraphQLInt,
            description: 'The duration of the video'
        },
        released: {
            type: new GraphQLNonNull(GraphQLBoolean),
            description: 'Whether or not the video is released'
        }
    }
});

const mutationType = new GraphQLObjectType( {                   // Lesson 10.
    name: 'Mutation',
    description: 'The root Mutation type',
    fields: {
        createVideo: {
            type: videoType,
            args: {
                video: {
                    type: new GraphQLNonNull(videoInputType),           // Lesson 11
                }
            },
            resolve: (_, args) => {
                return createVideo(args.video);                         // Lesson 11
            }
        }
    }
});

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const schema = new GraphQLSchema({
    query: queryType,
    mutation: mutationType,
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