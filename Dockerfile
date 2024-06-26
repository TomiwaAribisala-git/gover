FROM golang:latest as builder

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . . 

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main 



FROM alpine 

COPY --from=builder /app/main /main

EXPOSE 4000

ENTRYPOINT [ "/app/main" ]
