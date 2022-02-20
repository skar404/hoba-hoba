from fastapi import FastAPI

from backend.models.user import User

app = FastAPI()

# noinspection PyUnresolvedReferences
from backend.database import *


@app.get('/ping', response_model=User)
async def ping() -> User:
    return User(tg_id=1)
