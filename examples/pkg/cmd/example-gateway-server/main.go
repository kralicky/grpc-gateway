/*
Command example-gateway-server is an example reverse-proxy implementation
whose HTTP handler is generated by grpc-gateway.
*/
package main

import (
	"context"
	"flag"

	"github.com/kralicky/grpc-gateway/v2/examples/pkg/gateway"
	"google.golang.org/grpc/grpclog"
)

var (
	endpoint   = flag.String("endpoint", "localhost:9090", "endpoint of the gRPC service")
	network    = flag.String("network", "tcp", `one of "tcp" or "unix". Must be consistent to -endpoint`)
	openAPIDir = flag.String("openapi_dir", "examples/internal/proto/examplepb", "path to the directory which contains OpenAPI definitions")
)

func main() {
	flag.Parse()

	ctx := context.Background()
	opts := gateway.Options{
		Addr: ":8080",
		GRPCServer: gateway.Endpoint{
			Network: *network,
			Addr:    *endpoint,
		},
		OpenAPIDir: *openAPIDir,
	}
	if err := gateway.Run(ctx, opts); err != nil {
		grpclog.Fatal(err)
	}
}
