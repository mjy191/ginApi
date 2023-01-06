package main

import (
	"context"
	"fmt"
	"ginApi/proto/goodsService"
	"google.golang.org/grpc"
	"net"
)

type Goods struct{}

func (this Goods) AddGoods(c context.Context, req *goodsService.AddGoodsReq) (*goodsService.AddGoodsRes, error) {
	fmt.Println(req)
	return &goodsService.AddGoodsRes{
		Success: true,
		Msg:     "成功",
	}, nil
}

func main() {

	//初始化grpc
	grpcServer := grpc.NewServer()
	//注册服务
	goodsService.RegisterGoodsServiceServer(grpcServer, &Goods{})
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println(err)
		return
	}
	//启动服务
	grpcServer.Serve(listener)
}
