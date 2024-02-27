FROM golang:1.22.0-alpine

WORKDIR /app
COPY go.mod go.sum /app/

RUN go mod download

COPY ./ /app
RUN go build -o /go-gemini-api-sample

EXPOSE 8080

CMD ["/go-gemini-api-sample"]
