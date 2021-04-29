from graphql import GraphQLObjectType, GraphQLField, GraphQLString, GraphQLInt, GraphQLID, GraphQLSchema, GraphQLList


async def user(_obj, info):
    return {
        "id": "ac46ef5a-4728-4955-a724-32da3ea39974",
        "name": "Denis",
        "telegram_id": 10000123,
        "telegram_login": "@Denis",
    }


async def podcast(_obj, info):
    return {
        "id": "bf81bd23-417a-45ba-bcd0-cde9c1bb374f",
        "name": "Хоба!",
    }


async def send_post(_obj, info):
    return [{
        "id": "bf81bd23-417a-45ba-bcd0-cde9c1bb374f",
    }]


MainSchema = GraphQLSchema(GraphQLObjectType(
    name="AsyncQueryType",
    fields={
        "user": GraphQLField(
            GraphQLObjectType(
                name="user",
                fields={
                    "id": GraphQLField(GraphQLID),
                    "name": GraphQLField(GraphQLString),
                    "telegram_id": GraphQLField(GraphQLInt),
                    "telegram_login": GraphQLField(GraphQLString),
                },
            ),
            resolver=user),
        "podcast": GraphQLField(
            GraphQLObjectType(
                name="podcast",
                fields={
                    "id": GraphQLField(GraphQLID),
                    "name": GraphQLField(GraphQLString),
                    "post": GraphQLField(GraphQLList(GraphQLObjectType(
                        name="send_post",
                        fields={
                            "id": GraphQLField(GraphQLID),
                        }
                    )), resolver=send_post)
                },
            ),
            resolver=podcast),
    }))
