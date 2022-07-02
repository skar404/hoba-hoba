import asyncio
import os
from asyncio import Task
from typing import Optional
from dataclasses import dataclass

import aiohttp
import aiomonitor
import uvloop


@dataclass
class Content:
    url: str

    index: Optional[int] = None
    task: Optional[Task] = None
    is_run: bool = False
    is_active: bool = True


class Config:
    class Sleep:
        feed_time = 0.01
        main_task = 0.1


class DB:
    items = [
        Content(url='http://httpstat.us/200'),
        Content(url='http://httpstat.us/201'),
        Content(url='http://httpstat.us/202'),
        Content(url='http://httpstat.us/203'),
        Content(url='http://httpstat.us/204'),
        Content(url='http://httpstat.us/205'),
        Content(url='http://httpstat.us/206'),
        Content(url='http://httpstat.us/207'),
        Content(url='http://httpstat.us/208'),
        Content(url='http://httpstat.us/300'),
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
        try:
            async with aiohttp.ClientSession() as session:
                async with session.get(c.url) as response:
                    print(c.url, response.status)
            await asyncio.sleep(Config.Sleep.feed_time)
        except Exception as ex:
            print(f'get_feed {c.index=} ex=%s', ex)

    print(f'stop crawler {c.index=} task_id={id(c.task)}')


async def main():
    loop = asyncio.get_event_loop()

    with aiomonitor.start_monitor(loop=loop):
        print(f'run app pid={os.getpid()} db_id={id(DB)}')

        index = 0
        while True:
            for i in DB.items:
                # проверка что задачка не упала
                if i.task and i.task.cancelled():
                    print(f'task is cancelled {i.url} {i.index=}')
                    i.is_run = False

                # запускаем задачки
                if not i.is_run:
                    task = asyncio.ensure_future(get_feed(i))
                    print(f'run crawler {i.url} {index=} task_id={id(task)}')
                    task.set_name(f'crawler {i.url} {index=}')
                    i.task = task
                    i.is_run = True
                    i.index = index
                    index += 1

            await asyncio.sleep(Config.Sleep.main_task)


def debug_db(db_id):
    """
    Debug function
    :param db_id:
    :return:
    """
    # from crawler import debug_db;db = debug_db(5057659424)
    import ctypes
    return ctypes.cast(db_id, ctypes.py_object).value


if __name__ == '__main__':
    uvloop.install()
    asyncio.run(main())
