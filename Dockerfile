FROM golang:1.16 AS api

WORKDIR /buildapp
COPY . .
ADD go.mod go.sum /buildapp/

RUN CGO_ENABLED=0 GOOS=linux go build -o diploma ./main.go

RUN CGO_ENABLED=0 GOOS=linux go build -o ./internal/generator/generator ./internal/generator/main.go

EXPOSE 8282

CMD ["./start.sh"]

