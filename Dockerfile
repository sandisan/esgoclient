FROM golang:1.11
RUN mkdir -p /go/src/app
RUN go get github.com/nvn1729/badimportdemo/dontimportme
WORKDIR /go/src/app
COPY . /go/src/app
EXPOSE 8080
RUN go build
CMD ["./app"]
