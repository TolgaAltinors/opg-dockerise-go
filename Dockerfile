FROM golang:1.14-alpine AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy source to container
COPY app/. .

# Build the app to container location
RUN go build -o main .

# Expose port 8080, so it can be reached
EXPOSE 8080

# the command that will run
CMD ["./main"]

