# Telegram bot for monitoring rates of cryptocurrencies

## Supports the following commands:
- /start
- /help
- /list
- /rates
- /add
- /delete
- /deleteall

## Telegram API Token:
Create file configs/token.txt and paste there your token. File should consist of only 1 line.
Then run/build bot directly or via docker using Makefile:

### Building with Makefile:

> [!IMPORTANT] 
> Before building image create empty folders `bin` and `db` for databse and binary respectively

> [!WARNING]
> If creating an image finishes with error, try this commands:
```
docker pull golang:1.19-alpine
docker pull alpine
``` 

To create golang-alpine image with bot
```
make image
```
To run it via docker compose 
```
make up
```
