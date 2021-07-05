package main

import (
	"context"
	"gomicro_note/p12/models"
	"log"

	httpServer "github.com/asim/go-micro/plugins/client/http/v3"
	etcd "github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/selector"
)

// etcd 通过轮询获取服务
// 调用http api json tag不一致处理
// 使用第三方包 github.com/favadi/protoc-go-inject-tag
func callAPI(s selector.Selector) {
	myClient := httpServer.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
	)
	req := myClient.NewRequest("ProdSrv", "/v1/prods",
		models.ProdRequest{Size: 6})
	var rsp models.ProdListResponse
	err := myClient.Call(context.Background(), req, &rsp)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(rsp.GetData())
}

func main() {
	// etcd 连接句柄
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"))

	sel := selector.NewSelector(
		selector.Registry(etcdReg),
		selector.SetStrategy(selector.RoundRobin),
	)
	callAPI(sel)

}
