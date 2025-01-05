FROM golang:1.23-alpine AS server-builder

WORKDIR /server

COPY . .

RUN go mod tidy 
RUN go build -o server cmd/server/main.go

FROM alpine:latest

WORKDIR /server

COPY --from=server-builder /server .

RUN chmod +x ./server

CMD ["./server"]
