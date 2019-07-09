# Start from golang:1.11-alpine base image
FROM golang

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/daniel-vera-g/go-server-template

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
# https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
# TODO remove -t in production & spearate docker files
RUN go get -d -v -t ./...

# Install the package and create test binary
RUN go install -v ./... && \
    cd ./controllers && \
    CGO_ENABLED=0 GOOS=linux go test -c

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["go", "run", "main.go"]
