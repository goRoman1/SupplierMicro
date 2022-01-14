FROM golang:1.17-alpine as builder
WORKDIR /go/src/SupplierMicro
COPY . .
ENV GO111MODULE=on
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/suppliermicroservice

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
WORKDIR /usr/bin
COPY --from=builder /go/src/SupplierMicro/build/suppliermicroservice ./suppliermicroservice
ENTRYPOINT [ "/usr/bin/suppliermicroservice" ]