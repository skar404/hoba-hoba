import asyncio
import os
from dataclasses import dataclass

import aiomonitor
import uvloop


@dataclass
class Content:
    url: str

    is_run: bool = False
    is_active: bool = True


class Config:
    class Sleep:
        feed_time = 1

        main_task = 0.5


class DB:
    items = [
        Content(url='http://google.com'),
        Content(url='http://yandex.com'),
        Content(url='http://github.com'),
    ]

    @classmethod
    def add(cls, url):
        cls.items.append(Content(url=url))

    @classmethod
    def delete(cls, _id):
        cls.items[_id].is_active = False
        del cls.items[_id]


async def get_feed(c: Content):
    while c.is_active:
        print(c.url)
        await asyncio.sleep(Config.Sleep.feed_time)


async def main():
    print(f'run app pid={os.getpid()} db_id={id(DB)}')

    while True:
        for i in DB.items:
            if not i.is_run:
                print(f'run crawler {i.url}')
                asyncio.ensure_future(get_feed(i))
                i.is_run = True

        await asyncio.sleep(Config.Sleep.main_task)


def debug_db(db_id):
    """
    Debug function
    :param db_id:
    :return:
    """
    # from crawler import debug_db;db = debug_db(5100517232)
    import ctypes
    return ctypes.cast(db_id, ctypes.py_object).value


if __name__ == '__main__':
    uvloop.install()

    loop = asyncio.new_event_loop()
    with aiomonitor.start_monitor(loop=loop):
        loop.create_task(main())
        loop.run_forever()
