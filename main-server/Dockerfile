FROM golang:1.19-alpine AS builder

WORKDIR /build

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd/ cmd/
COPY pkg/ pkg/

RUN CGO_ENABLED=0 GOOS=linux go build -a -o app cmd/web/main.go

FROM alpine:3.16

WORKDIR /cfrss
COPY --from=builder /build/app bin/app

EXPOSE 8000

ENTRYPOINT [ "bin/app" ]