from typing import Optional

from bson import ObjectId
from pydantic import BaseModel, Field

from backend.utils.mango import PyObjectId


class Setting(BaseModel):
    pass


class State(BaseModel):
    is_stop: bool = False
    setting: Setting


class User(BaseModel):
    tg_id: int
    active: bool = True

    state: Optional[State]

    class Config:
        arbitrary_types_allowed = True
        json_encoders = {
            ObjectId: str
        }
