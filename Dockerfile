# Dockerfile References: https://docs.docker.com/engine/reference/builder/

FROM golang:1.15.2-buster AS builder

RUN mkdir /app

# Set the Current Working Directory inside the container
WORKDIR /app

ENV GO111MODULE=on

#azure repo auth
ARG pat
ENV pat=$pat

RUN git config \
  --global \
  url."https://anything:${pat}@powerpilot.visualstudio.com".insteadOf \
  "https://powerpilot.visualstudio.com"

COPY go.mod .
COPY go.sum .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy go files
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o service ./cmd

######## Start a new stage from scratch #######
FROM alpine:3.12.0

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/service .

# Expose port to the outside world
EXPOSE 80

# Command to run the executable
CMD ["./service"]
