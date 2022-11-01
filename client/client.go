package main

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/lqf1215/kitex-demo/kitex_gen/service"
	"github.com/lqf1215/kitex-demo/kitex_gen/service/userservice"
	"log"
)

func main() {
	client, err := userservice.NewClient("userservice", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}

	req := &service.GetUserListReq{Page: 1, PageSize: 10, Username: "张三"}
	resp, err := client.GetUserList(context.Background(), req)
	if err != nil {
		log.Fatal("err1", err.Error())
	}
	log.Printf("resp:%v", resp)
	req1 := &service.GetUserByIdReq{Id: 1}
	resp1, err1 := client.GetUserById(context.Background(), req1)
	if err1 != nil {
		log.Fatal("err1", err1.Error())
	}
	log.Printf("resp1:%v", resp1)
}
