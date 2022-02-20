# hoba-hoba

Hoba bot by Hoba podcast

# Architecture

## model:

collection format: 
```text
Content {
    uuid uuid
    
    data {
        name str 
        title str
    
    feed {
        url

    admins [
        uuid uuid
        telegram_id int64
        is_owner bool

    crawler {
        time datetime 
        index str

Item {
    uuid uuid 
    content_uuid -> Content
    
    data {
        name str
        title str
        text str
        url str
        datetime datetime
    
    state {
        telegram bool
```
