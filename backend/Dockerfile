    FROM golang:1.24-alpine AS builder

    WORKDIR /app
    
    COPY go.mod go.sum ./
    RUN go mod download
    
    COPY . .
    
    RUN go build -o app ./cmd
    
    
    FROM alpine:latest
    
    WORKDIR /app
    
    COPY --from=builder /app/app .
    
    RUN apk --no-cache add ca-certificates tzdata
    
    EXPOSE 9090
    
    ENTRYPOINT ["./app"]
    