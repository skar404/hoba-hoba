from datetime import datetime, time
from typing import Optional

from pydantic import BaseModel


class Podcast(BaseModel):
    RSS_feed: str

    # имя подкаста, зачем ?!
    name: Optional[str]
    # футь до лога
    logo_path: Optional[str]
    # предпочтительное время поста
    post_time: Optional[time]

    # эпизод c которого все начилось
    final_episode: Optional[str]

    # Telegram ID создателя
    owner_id: int

    # тех. параметры
    active: bool = True
    create_at: datetime = datetime.now()
    update_at: datetime = datetime.now()
