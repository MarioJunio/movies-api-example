FROM golang:1.16-alpine

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app/movies-api

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Build the Go app
RUN go build -o ./out/movies-api .

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./out/movies-api"]