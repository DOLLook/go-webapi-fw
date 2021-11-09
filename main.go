package main

import (
	"go-webapi-fw/common/mqutils"
	"go-webapi-fw/mqconsumers"
	"go-webapi-fw/web/iris_srv"
)

func main() {
	mqconsumers.Initialize()
	mqutils.BindConsumer()
	iris_srv.Start()
}