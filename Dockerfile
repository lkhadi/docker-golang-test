FROM golang:1.19.0-alpine3.16
WORKDIR /go/src/app
COPY main.go .
RUN go mod init
RUN go build -o main .
#EXPOSE 8990
CMD ["./main"]
