package main

import (
	"douniu/server/favorite/mq/internal/config"
	"douniu/server/favorite/mq/internal/service"
	"flag"
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/favorite.yaml", "the etc file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	s := service.NewService(c)

	queue := kq.MustNewQueue(c.KafkaConf, kq.WithHandle(s.Consume))
	defer queue.Stop()

	fmt.Println("favorite-mq started!!!")
	queue.Start()
}
