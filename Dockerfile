# Start from latest golang base image
FROM golang:latest as builder

# Set the current directory inside the container
WORKDIR /app

# Copy sources inside the docker
COPY . .

# install the dependencies
RUN go mod tidy

# Build the binaries from the source
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

###### Start a new stage from scratch #######
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

# Set the necessary environment variables
ENV ADMIN_USER=admin
ENV ADMIN_PASS=demo
ENV SIGNING_KEY=veryverysecretkey

# Expose port 8080 to the outside container
EXPOSE 8080

# Declare entry point of the docker command
ENTRYPOINT ["./main"]

# Run the binary program produced by `go build`
CMD ["start","-a"]

