build:	
	@go build -o ./bin/bot ./bot/cmd/bot/main.go 
run: build
	@./bin/bot
image:
	@docker build -t tg_bot .
up:
	@docker compose up -d
