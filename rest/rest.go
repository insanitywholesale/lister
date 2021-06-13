package rest

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/cors"
	gw "gitlab.com/insanitywholesale/lister/proto/v1"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func RunGateway(grpcport string, restport string) http.Handler {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterListerHandlerFromEndpoint(ctx, mux, ":"+grpcport, opts)
	if err != nil {
		log.Fatal(err)
	}

	return cors.Default().Handler(mux)
}
