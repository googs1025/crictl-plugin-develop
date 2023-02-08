package cmds

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
	"k8s.io/klog/v2"
	"time"
)

var versionCmd= &cobra.Command{
	Use:          "version",
	Run: func(c *cobra.Command, args []string) {

		req := &v1alpha2.VersionRequest{}
		ctx, cancel := context.WithTimeout(context.Background(),time.Second*3)
		defer cancel()
		runtimeService := v1alpha2.NewRuntimeServiceClient(grpcClient)

		rsp, err := runtimeService.Version(ctx,req)
		if err != nil {
			klog.Error(err)
			return
		}
		fmt.Println("Version:", rsp.Version)
		fmt.Println("RuntimeName:", rsp.RuntimeName)
		fmt.Println("RuntimeVersion:", rsp.RuntimeVersion)
		fmt.Println("RuntimeApiVersion:", rsp.RuntimeApiVersion)
	},
}