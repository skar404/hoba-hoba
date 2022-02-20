from urllib.parse import quote_plus

from motor import motor_asyncio
from motor.core import AgnosticDatabase, AgnosticClient

MONGO_DETAILS = f"mongodb://{quote_plus('root')}:{quote_plus('example')}@{'localhost:27017'}"
client: AgnosticDatabase = motor_asyncio.AsyncIOMotorClient(MONGO_DETAILS)
DB: AgnosticClient = client.hoba_bot
