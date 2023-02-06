package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
	"log"
	"time"
)

func main() {

	grpcOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// 请求的路径写死，因为模拟kubelet调用cri接口。
	addr := "unix:///run/containerd/containerd.sock"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	conn, err := grpc.DialContext(ctx, addr, grpcOpts...)
	if err != nil {
		log.Fatalln(err)
	}

	var req = &v1alpha2.VersionRequest{}
	var ret = &v1alpha2.VersionResponse{}

	// 类似请求路径：v1alpha2 版本RuntimeService Group里面的version接口。
	err = conn.Invoke(ctx,"/runtime.v1alpha2.RuntimeService/Version", req, ret)

	fmt.Println(ret)

	if err != nil {
		log.Fatalln(err)
	}


}
