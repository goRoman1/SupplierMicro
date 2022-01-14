//go:generate protoc -I=./proto --go_out=./proto ./proto/supplier.proto --go-grpc_out=./proto ./proto/supplierMicro.proto
//go:generate openssl req -new -x509 -sha256 -key problserv.key -out supserv.crt -days 3650 -subj "/CN=suppliermicroservice" -addext "subjectAltName = DNS:suppliermicroservice"
//go:generate openssl req -new -x509 -sha256 -key problserv.key -out suplocal.crt -days 3650 -subj "/CN=localhost" -addext "subjectAltName = DNS:localhost"
package main

import (
	"SupplierMicro/configs"
	"SupplierMicro/proto"
	"SupplierMicro/service"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os/exec"
)

const MicroName = "supplier_service"

func main() {
	connectionDB := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configs.PG_HOST,
		configs.PG_PORT,
		configs.POSTGRES_USER,
		configs.POSTGRES_PASSWORD,
		configs.POSTGRES_DB)

	db, err := sql.Open("postgres", connectionDB)
	if err != nil {
		log.Panicf("%s: failed to open db connection - %v", MicroName, err)
	}
	defer db.Close()

	service := service.NewSupplierMicroService(db)

	listener, err := net.Listen("tcp", net.JoinHostPort("", configs.GRPC_PORT))
	if err != nil {
		log.Panicf("%s: failed to listen on port - %v", MicroName, err)
	}

	cmd, _ := exec.Command("pwd").Output()
	fmt.Println(string(cmd))
	creds, err := credentials.NewServerTLSFromFile(configs.CERTIFICATE, configs.KEY_PRIVATE)
	if err != nil {
		log.Panicf("%s: can't load TLS keys : %v", MicroName, err)
	}

	server := grpc.NewServer(grpc.Creds(creds))
	defer server.GracefulStop()
	proto.RegisterSupplierMicroServiceServer(server, service)
	reflection.Register(server)

	if err := server.Serve(listener); err != nil {
		log.Panicf("%s: failed to start grpc - %v", MicroName, err)
	}
}
