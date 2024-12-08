FROM golang:1.19-alpine AS builder 

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev

# dependencies
COPY ["./go.mod", "./go.sum", "./"]
RUN go mod download

# build
COPY bot ./bot
RUN go build -o ./bin/bot bot/cmd/bot/main.go


FROM alpine:latest AS runner   

WORKDIR /tg_bot
COPY --from=builder /usr/local/src/bin/bot ./
COPY configs ./configs
COPY db ./db

CMD ["./bot"]
