FROM golang:latest

WORKDIR /go/src/app

COPY . .
 
# RUN go get github.com/githubnemo/CompileDaemon
# RUN go install -mod=mod github.com/githubnemo/CompileDaemon
# ENTRYPOINT CompileDaemon --build="go build -o main main.go" --command=./main

RUN go build -o main main.go
CMD ["./main"]