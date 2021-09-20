from pydantic import BaseModel


class Podcast(BaseModel):
    feed: str
