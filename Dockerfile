# Stage 1: Build the application
FROM golang:1.23.4-alpine as builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Install build tools
RUN apk add --no-cache curl git

# Install Atlas CLI
RUN curl -s https://atlasgo.sh | sh

# Create app directory
WORKDIR /app

# Copy Go modules and source files
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the binary
RUN go build -o hr-system main.go

# Stage 2: Run the application
FROM alpine:3.19

# Install dependencies
RUN apk --no-cache add ca-certificates curl

# Install Atlas CLI
RUN curl -s https://atlasgo.sh | sh

# Set working directory and copy binary from builder
WORKDIR /root/
COPY --from=builder /app/hr-system .
COPY --from=builder /app/config/*.json config/
COPY --from=builder /app/config/*.env config/
COPY --from=builder /app/migrate/ migrate/
COPY --from=builder /app/docs/ docs/

ENV DB_HOST=$DB_HOST
ENV DB_PORT=$DB_PORT
ENV DB_USERNAME=$DB_USERNAME
ENV DB_PASSWORD=$DB_PASSWORD
ENV DB_NAME=$DB_NAME

#WORKDIR migrate
#RUN atlas migrate apply --url "mysql://$DB_USERNAME:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME" --dir 'file://mysql'

# Expose server port
EXPOSE 8080

# Command to run the application
#CMD ["./hr-system"]
#CMD ["sh", "-c", "cd migrate && ls"]
CMD ["sh", "-c", "cd migrate && atlas migrate apply --url \"mysql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}\" --dir 'file://mysql' && cd .. && ./hr-system"]
