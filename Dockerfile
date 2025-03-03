FROM golang:1.24-alpine


WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

ENV HOST=0.0.0.0
ENV PORT=8080


RUN go build -o main api/cmd/weather-api-server/main.go

EXPOSE 8080


CMD ["./main", "--scheme=http"]
