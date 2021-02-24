FROM golang:latest
ENV GO111MODULE=auto
WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["main"]
