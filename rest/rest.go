package rest

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/cors"
	gw "gitlab.com/insanitywholesale/lister/proto/v1"
	"google.golang.org/grpc"
	"net/http"
	"strings"
)

var (
	openapiDocs []byte

	commitHash string
	commitDate string
)

func SaveVars(oaD []byte, cH string, cD string) {
	openapiDocs = oaD
	commitHash = cH
	commitDate = cD
}

func fallback(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("idk about that fam\n"))
	return
}

func pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong\n"))
	return
}

func getDocs(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write(openapiDocs)
		return
	}
	return
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte("commitHash: " + commitHash + "\n"))
		w.Write([]byte("commitDate: " + commitDate + "\n"))
		return
	}
	return
}

func RunGateway(grpcport string, restport string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gwmux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterListerHandlerFromEndpoint(ctx, gwmux, ":"+grpcport, opts)
	if err != nil {
		return err
	}

	handler := cors.Default().Handler(gwmux)

	gwServer := &http.Server{
		Addr: ":" + restport,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/api") {
				handler.ServeHTTP(w, r)
				return
			}
			if strings.HasPrefix(r.URL.Path, "/docs") {
				getDocs(w, r)
				return
			}
			if strings.HasPrefix(r.URL.Path, "/info") {
				getInfo(w, r)
				return
			}
			if strings.HasPrefix(r.URL.Path, "/ping") {
				pong(w, r)
				return
			}
			fallback(w, r)
			return
		}),
	}
	return gwServer.ListenAndServe()
}
