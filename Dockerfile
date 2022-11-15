FROM golang:1.18

WORKDIR /avito_app
COPY ./ ./

RUN go mod download
RUN go build -o ./cmd/app ./cmd/main.go

WORKDIR /avito_app/cmd
CMD ["./app"]