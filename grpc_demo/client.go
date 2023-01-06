package main

import (
	"context"
	"fmt"
	"ginApi/proto/goodsService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	//初始化grpc
	conn, err := grpc.Dial("127.0.0.1:8800", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return
	}
	//注册客户端
	client := goodsService.NewGoodsServiceClient(conn)
	//调用远程服务
	res, err1 := client.AddGoods(context.Background(), &goodsService.AddGoodsReq{
		Title:   "标题",
		Price:   100,
		Content: "内容",
	})
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Printf("%#v", res)
}
