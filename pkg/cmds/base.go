package cmds

import (
	"context"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
	"log"
	"time"
)

const CriAddr = "unix:///run/containerd/containerd.sock" //临时写死

var grpcClient  *grpc.ClientConn  // grpc连接

func initClient()  {
	grpcOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	conn, err := grpc.DialContext(ctx, CriAddr, grpcOpts...)
	if err != nil {
		log.Fatalln(err)
	}

	grpcClient = conn
}


func NewRuntimeService() v1alpha2.RuntimeServiceClient  {
	return v1alpha2.NewRuntimeServiceClient(grpcClient)
}

func NewImageService() v1alpha2.ImageServiceClient{
	return v1alpha2.NewImageServiceClient(grpcClient)
}

var TTY bool //终端模式


func RunCmd() {
	cmd := &cobra.Command{
		Use:          "jiangctl",
		Short:        "list images",
		Example:      "jiangctl list images",
		SilenceUsage: true,
	}
	initClient()// 初始化 grpc 客户端
	// 加入子命令
	containersExecCmd.Flags().BoolVarP(&TTY,"tty","t",false,"-t")
	cmd.AddCommand(versionCmd, imagesCmd, podsCmd, containersCmd, containersListCmd, containersExecCmd)
	err := cmd.Execute()
	if err!=nil{
		log.Fatalln(err)
	}
}