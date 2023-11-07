package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"

	"douniu/server/comment/api/internal/config"
	"douniu/server/comment/api/internal/handler"
	"douniu/server/comment/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/comment-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	httpx.SetValidator(svc.NewValidator())
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
