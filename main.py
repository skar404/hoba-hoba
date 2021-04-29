import asyncio

import aiohttp_cors
import uvloop
from aiohttp import web
from aiohttp_graphql import GraphQLView
from graphql.execution.executors.asyncio import AsyncioExecutor
from motor.motor_asyncio import AsyncIOMotorClient

from core import App
from handlers.schema import MainSchema


async def setup_db():
    db = AsyncIOMotorClient(
        'mongodb://root:example@localhost:27017/?authSource=admin'
    ).hoba
    return db


def main():
    asyncio.set_event_loop_policy(uvloop.EventLoopPolicy())
    loop = asyncio.get_event_loop()

    app = App()

    gql_view = GraphQLView(
        schema=MainSchema,
        graphiql=True,
        enable_async=True,
        batch=True,
        executor=AsyncioExecutor(loop=asyncio.get_event_loop()))

    for m in ['GET', 'POST']:
        app.router.add_route(m, '/graphql', gql_view, name='graphql')

    app.DB = loop.run_until_complete(setup_db())

    cors = aiohttp_cors.setup(app, defaults={
        '*': aiohttp_cors.ResourceOptions(
            allow_credentials=True,
            expose_headers='*',
            allow_headers=('Content-Type',),
        )
    })
    for route in list(app.router.routes()):
        cors.add(route)

    web.run_app(app, port=8000)


if __name__ == '__main__':
    main()
