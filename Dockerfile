FROM golang:alpine

WORKDIR /work
COPY . .

RUN go get entgo.io/ent/cmd/ent \
    && go run entgo.io/ent/cmd/ent generate \
#    && go mod tidy \

RUN go build -a -ldflags '-s' -installsuffix cgo -o app

CMD ["./app"]