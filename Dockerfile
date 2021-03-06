FROM golang:1.12.0-alpine3.9 as builder

WORKDIR /app/chatbot

RUN apk add git
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o chatbot

FROM scratch

WORKDIR /app

COPY --from=builder /app/chatbot/chatbot /app/chatbot

ENTRYPOINT [ "/app/chatbot" ]
