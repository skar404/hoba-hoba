from modules.podcast import create_new_podcast


async def create_podcast(request):
    await create_new_podcast(request.app, p=None)
    return
