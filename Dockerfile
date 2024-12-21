# Base image
FROM golang:1.21-alpine

# Set environment variables for Go
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set working directory
WORKDIR /app

# Install necessary dependencies, including postgresql-client
RUN apk add --no-cache git postgresql-client

# Copy module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code
COPY . .


ENV DB_HOST=database-1.cluster-ct8q64aoo2er.eu-north-1.rds.amazonaws.com
ENV DB_USER=postgres
ENV DB_PASSWORD=qazwsxedc12A!!!
ENV DB_NAME=postgres
ENV DB_PORT=5432
ENV DB_SSLMODE=disable
ENV DB_TIMEZONE=Asia/Almaty



# Build the application
RUN go build -o main .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./main"]
