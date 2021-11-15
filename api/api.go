package api

import (
	"github.com/NpoolPlatform/sphinx-plugin/message/npool"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// api层面需要用其他模块部署

type Server struct {
	npool.UnimplementedPluginServer
}

func Register(server grpc.ServiceRegistrar) {
	if false {
		npool.RegisterPluginServer(server, &Server{})
	}
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
