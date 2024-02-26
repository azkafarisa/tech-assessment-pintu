FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o go-app

EXPOSE 4000

CMD [ "./go-app" ]
