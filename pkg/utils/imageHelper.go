package utils

import (
	"fmt"
	"strings"
)

const noneString="<none>"

// 解析repo和digest
// 格式类似[docker.io/library/alpine@sha256:xxxxoooo]
// 返回两个值：第一个是imageName 第二个是 digest
func ParseRepoDigest(repoDigests []string) (string, string) {
	if len(repoDigests) == 0 {
		return noneString, noneString
	}
	repoDigestPair := strings.Split(repoDigests[0], "@")
	if len(repoDigestPair) != 2 {
		return "errImage", "errDigest"
	}
	return repoDigestPair[0], repoDigestPair[1]
}


//解析镜像和tag
//内容类似[docker.io/library/alpine:3.12]   也是个数组。 如果镜像出错，譬如打包出错，中途终止等
//返回值 是一个二维 string切片([][]string{})。 前端显示时 只需要显示第一个
//每一个 子切片 是一个 string{}   。包含两个值：镜像名称和tag
func ParseRepoTag(repoTags []string, imageName string) (repoTagPairs [][]string) {
	if len(repoTags) == 0 {
		repoTagPairs = append(repoTagPairs, []string{imageName, noneString})
		return
	}
	for _, repoTag := range repoTags {
		idx := strings.LastIndex(repoTag, ":")
		if idx <0 { //解析出错了， 直接返回errTag，
			repoTagPairs = append(repoTagPairs, []string{"errTag", "errTag"})
			continue
		}
		name := repoTag[:idx]
		if name == noneString {
			name = imageName
		}
		repoTagPairs = append(repoTagPairs, []string{name, repoTag[idx+1:]})
	}
	return
}

// 解析 size
func ParseSize(size uint64) string{
	return fmt.Sprintf("%.2fm",float64(size)/1024/1024) //单位是 m
}

// 截取ID  （13位)
func ParseImageID(id string) string  {
	idstr:=strings.Split(id,":")[1]
	return idstr[:13]
}


func ParseContainerID(id string) string {
	return id[:13]
}