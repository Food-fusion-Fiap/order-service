FROM golang:1.22.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /go/bin/app .

FROM golang:1.22.1-alpine

COPY --from=builder /go/bin/app /go/bin/app
COPY --from=builder /app/infra/cert/global-bundle.pem /infra/cert/global-bundle.pem

EXPOSE 8080

CMD ["/go/bin/app"]