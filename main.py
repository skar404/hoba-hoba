import asyncio

import uvloop
from aiohttp import web
from motor.motor_asyncio import AsyncIOMotorClient

from core import App
from handlers import api


async def setup_db():
    db = AsyncIOMotorClient(
        'mongodb://root:example@localhost:27017/?authSource=admin'
    ).hoba
    return db


def main():
    asyncio.set_event_loop_policy(uvloop.EventLoopPolicy())

    loop = asyncio.get_event_loop()

    app = App()
    app.add_routes([
        # CRUD по работе с подкастом
        web.post('/api/podcasts', api.create_podcast),
        web.patch('/api/podcasts', api.create_podcast),
        web.delete('/api/podcasts', api.create_podcast),
        web.get('/api/podcasts', api.create_podcast),
    ])

    app.DB = loop.run_until_complete(setup_db())

    web.run_app(app)


if __name__ == '__main__':
    main()
