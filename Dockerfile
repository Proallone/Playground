FROM golang:1.21.4

WORKDIR /go/src/app

COPY . .

# RUN go build -o main main.go
RUN go get github.com/githubnemo/CompileDaemon
RUN go install -mod=mod github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -o main main.go" --command=./main
# CMD ["./main"]