package main

import (
	"douniu/server/video/mq/internal/config"
	"douniu/server/video/mq/internal/service"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"

	"github.com/zeromicro/go-queue/kq"
)

var configFile = flag.String("f", "etc/video.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	srv := service.NewService(c)
	queue := kq.MustNewQueue(c.KqConsumerConf, kq.WithHandle(srv.Consume))
	defer queue.Stop()

	fmt.Println("seckill started!!!")
	queue.Start()
}
