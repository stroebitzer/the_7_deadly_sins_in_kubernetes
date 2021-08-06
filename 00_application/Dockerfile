FROM golang:alpine3.14 as builder
WORKDIR /src
COPY go.mod .
# COPY go.sum .
RUN go mod download
COPY *.go ./
RUN go build -o application

FROM alpine:3.14.0
USER nobody:nobody 
WORKDIR /app
COPY --from=builder /src/application /app/application
ENTRYPOINT ["./application"]
