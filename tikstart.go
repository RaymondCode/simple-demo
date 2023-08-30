package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	error2 "tikstart/common/error"
	"tikstart/internal/types"

	"tikstart/internal/config"
	"tikstart/internal/handler"
	"tikstart/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/tikstart.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandler(errHandler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func errHandler(err error) (int, interface{}) {
	switch e := err.(type) {
	case error2.ApiError:
		return e.StatusCode, e.Response()
	case error2.ServerError:
		fmt.Printf("%s: %s\n", e, e.Detail)
		return e.StatusCode, e.Response()
	default:
		fmt.Printf("Internal Server Error: %s\n", e)
		return http.StatusInternalServerError, &types.BasicResponse{
			StatusCode: 50000,
			StatusMsg:  "Internal Server Error",
		}
	}
}
