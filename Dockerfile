FROM golang

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

COPY app app
COPY go.mod .
COPY go.sum .

RUN ls

RUN echo $GOPATH

RUN go mod edit -replace github.com/voltento/users-info/app/config=/app
RUN go build -o main app/app.go

WORKDIR /app

RUN pwd

RUN cp /build/main app.go

COPY config/config_docker.json config.json

CMD ["/app/app.go", "--config", "config.json"]
