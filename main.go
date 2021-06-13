package main

import (
	_ "embed"
	api "gitlab.com/insanitywholesale/lister/grpc"
	pb "gitlab.com/insanitywholesale/lister/proto/v1"
	"gitlab.com/insanitywholesale/lister/rest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"os"
)

var (
	//go:embed openapiv2/v1/lister.swagger.json
	openapiDocs []byte

	commitHash string
	commitDate string
)

func getDocs(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write(openapiDocs)
	}
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte("commitHash: " + commitHash + "\n"))
		w.Write([]byte("commitDate: " + commitDate + "\n"))
	}
}

func main() {
	grpcport := os.Getenv("LISTER_GRPC_PORT")
	if grpcport == "" {
		grpcport = "15200"
	}

	listener, err := net.Listen("tcp", ":"+grpcport)
	if err != nil {
		log.Fatalf("listen failed %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterListerServer(grpcServer, api.Server{})
	reflection.Register(grpcServer)
	log.Println("grpc started on port", grpcport)
	go grpcServer.Serve(listener)

	restport := os.Getenv("LISTER_REST_PORT")
	if restport == "" {
		restport = "9392"
	}

	gw := rest.RunGateway(grpcport, restport)
	mux := http.NewServeMux()
	mux.HandleFunc("/info/", getInfo)
	mux.HandleFunc("/docs/", getDocs)
	mux.Handle("/", gw)

	s := &http.Server{
		Addr:    ":" + restport,
		Handler: mux,
	}

	log.Println("rest starting on port", restport)
	log.Fatal(s.ListenAndServe())
}
