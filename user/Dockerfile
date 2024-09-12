#build stage
FROM golang:alpine AS builder

WORKDIR /app
# COPY go.mod .
# COPY go.sum .
COPY . .
RUN go mod download

RUN go build -o user main.go

#run stage
FROM alpine 
WORKDIR /app
COPY --from=builder /app/user .
# COPY app.env .
COPY .env .

EXPOSE 4000
CMD ["/app/user"]
