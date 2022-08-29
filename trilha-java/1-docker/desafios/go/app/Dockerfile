FROM golang:1.18 AS builder

WORKDIR /usr/src/app

COPY . ./

RUN go mod download && go mod verify

RUN apt-get install git

RUN go build -v -o . ./...
RUN cd /usr/src/app && \
    go run .

FROM scratch
COPY --from=builder /usr/src/app/hello /
CMD ["/hello"]
