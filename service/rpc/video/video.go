package main

import (
	"flag"
	"fmt"

	"github.com/RaymondCode/simple-demo/service/rpc/video/internal/config"
	"github.com/RaymondCode/simple-demo/service/rpc/video/internal/server"
	"github.com/RaymondCode/simple-demo/service/rpc/video/internal/svc"
	"github.com/RaymondCode/simple-demo/service/rpc/video/video"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/video.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		video.RegisterVideoServer(grpcServer, server.NewVideoServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
