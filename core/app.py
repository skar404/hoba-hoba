from aiohttp.web_app import Application
from motor.motor_asyncio import AsyncIOMotorDatabase


class App(Application):
    DB: AsyncIOMotorDatabase
