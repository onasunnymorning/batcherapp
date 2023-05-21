VERSION 0.6
FROM golang:latest

WORKDIR /

requirements:
    COPY go.mod ./
    COPY go.sum ./
    RUN go mod download
    SAVE IMAGE batcherapp:requirements

build:
    FROM +requirements
    COPY ./batch ./batch
    COPY ./app ./app
    COPY ./cmd ./cmd

    RUN go build -o batcherapp ./cmd/api/v1/main.go
    
    # Removing source code after build
    RUN rm -Rdf ./app ./cmd ./batch ./go.mod ./go.sum

    EXPOSE 8080

    CMD [ "/batcherapp" ]

    SAVE IMAGE --push batcherapp:latest