FROM golang:latest

ENV GO111MODULE=on

RUN go install github.com/a-h/templ/cmd/templ@v0.0.113

WORKDIR /app
COPY go.mod /app/.
COPY go.sum /app/.

RUN go mod download

COPY . /app/.

RUN templ generate

RUN go mod verify
RUN go mod vendor

RUN go build -o /app/tmp/cmd ./cmd
EXPOSE 80
CMD ["/app/tmp/cmd"]
