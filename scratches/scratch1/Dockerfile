FROM golang:latest
# RUN mkdir -p /go/src/github.com/kardianos/service
# RUN mkdir -p /go/src/github.com/coreos/go-log/log
WORKDIR /go/src/app
COPY . .
RUN go build -o main .
CMD ["/go/src/app/main"]
EXPOSE 8181

