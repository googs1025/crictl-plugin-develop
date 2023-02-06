package cmds

import (
	"container_cri_demo/pkg/utils"
	"context"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
	"log"
	"os"
	"time"
)

// 镜像相关的 显示和处理
var imagesCmd= &cobra.Command{
	Use:          "images",
	Run: func(c *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(),time.Second*3)
		defer cancel()

		// 请求image req
		req := &v1alpha2.ListImagesRequest{}
		rsp, err := NewImageService().ListImages(ctx, req)
		if err!=nil{
			log.Fatalln(err)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"镜像", "标签", "ID", "大小"})
		for _, img := range rsp.GetImages() {
			imageName, _ := utils.ParseRepoDigest(img.RepoDigests)		//取到镜像名
			repotag := utils.ParseRepoTag(img.RepoTags,imageName)[0] 	//取到 镜像名和标签 切片
			row := []string{
				imageName,
				repotag[1],
				utils.ParseImageID(img.Id),
				utils.ParseSize(img.Size_),
			}
			table.Append(row)
		}
		utils.SetTable(table)
		table.Render()

	},

}