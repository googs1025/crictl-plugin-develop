package cmds

import (
	"container_cri_demo/pkg/utils"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
	"log"
	"time"
)



var podsCmd = &cobra.Command{
	Use:          "runp",  //单创建 pod
	Run: func(c *cobra.Command, args []string) {
		if len(args)==0{
			log.Fatalln("请指定POD配置文件")
		}
		config:=&v1alpha2.PodSandboxConfig{}
		err:=utils.YamlFile2Struct(args[0],config)
		if err!=nil{
			log.Fatalln(err)
		}
		ctx, cancel := context.WithTimeout(context.Background(),time.Second*20)
		defer cancel()

		req:=&v1alpha2.RunPodSandboxRequest{Config: config}
		rsp,err:=NewRuntimeService().RunPodSandbox(ctx,req)
		if err!=nil{
			log.Fatalln(err)
		}
		fmt.Println(rsp.PodSandboxId)



	},

}
//本课程来自 程序员在囧途(www.jtthink.com) 咨询群：98514334