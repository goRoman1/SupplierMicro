FROM golang:1.17-alpine3.13 as builder
WORKDIR /go/src/SupplierMicro
COPY . .
ENV GOPROXY https://proxy.golang.org,direct
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o build/suppliermicroservice

FROM scratch
COPY --from=builder /go/src/SupplierMicro/build/suppliermicroservice /usr/bin/suppliermicroservice
COPY --from=builder /go/src/SupplierMicro/supserv.* /home/certificates/
ENTRYPOINT [ "/usr/bin/suppliermicroservice" ]