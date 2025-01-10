FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o crm.shopGO.com ./cmd/server

FROM scratch

COPY ./config /config

COPY  --from=builder /build/crm.shopGO.com /

ENTRYPOINT [ "/crm.shopGO.com", "config/local.yaml" ]