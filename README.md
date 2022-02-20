# hoba-hoba

Hoba bot by Hoba podcast

# Architecture

## model:

collection format: 
```yaml
Content:
    uuid: uuid
    
    data:
        name: str 
        title: str
    
    feed:
        url: str

    admins:
      - uuid: uuid
        telegram_id: int64
        is_owner: bool
  
    publish:
      - telegram_id: int64

    crawler:
        time: datetime 
        index: str

Item:
    uuid: uuid 
    content_uuid: -> Content
    
    data:
        name: str
        title: str
        text: str
        url: str
        datetime: datetime
    
    state:
        telegram: bool
```

## app

```mermaid
graph LR
    subgraph Privat Network
        db[(Database)]

        tg_app[Telegram APP]
        public_api[Public API]
        crawler[Crawler]
        
        tg_app --> db
        public_api --> db
        crawler --> db
    end 

    web_client[Web Client]
    tg_api[Telegram API]
    rss_feed[RSS Feed]

    tg_api --> tg_app
    web_client --> public_api
    crawler ---> rss_feed

```
