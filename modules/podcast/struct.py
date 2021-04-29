from datetime import datetime, time
from typing import Optional

from pydantic import BaseModel


class Podcast(BaseModel):
    RSS_feed: str
    # время между запросами к RSS
    sleep_time: Optional[int]

    # имя подкаста, зачем ?!
    name: Optional[str]
    # футь до лога
    logo_path: Optional[str]
    # предпочтительное время поста
    post_time: Optional[time]
    # токен бота от которого бдует создаваться пост
    telegram_bot_token: Optional[str]

    # эпизод c которого все начилось
    final_episode: Optional[str]

    # Telegram ID создателя
    owner_id: int

    # тех. параметры
    active: bool = True
    create_at: datetime = datetime.now()
    update_at: datetime = datetime.now()
