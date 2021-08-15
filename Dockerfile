FROM golang:1.16-alpine
WORKDIR /go/src/ocpnv
ENV CGO_ENABLED=0 \
    BIN_NAME='ocpnv'
#COPY *.go .
#COPY *.template .
ADD . .
RUN go mod init \
    && go get github.com/jpillora/opts \
    && go test -v -cover \
    && GOOS=linux GOARCH=amd64 go build -o /${BIN_NAME}-amd64-linux \
    && GOOS=darwin GOARCH=amd64 go build -o /${BIN_NAME}-amd64-darwin \
    && GOOS=windows GOARCH=amd64 go build -o /${BIN_NAME}-amd64.exe
