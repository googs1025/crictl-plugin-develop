package utils

import (
	"github.com/olekukonko/tablewriter"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"k8s.io/klog/v2"
	"os"
)

// YamlFile2Struct 读取文件内容 且反序列为struct
func YamlFile2Struct(path string,obj interface{}) error{
	b, err := GetFileContent(path)
	if err != nil {
		klog.Error("开启文件错误：", err)
		return err
	}
	err = yaml.Unmarshal(b, obj)
	if err != nil {
		klog.Error("解析yaml文件错误：", err)
		return err
	}
	return nil
}


// GetFileContent 文件读取函数
func GetFileContent(path string) ([]byte,error){
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// item is in []string{}
func InArray(arr []string,item string ) bool  {
	for _,p:=range arr{
		if p==item{
			return true
		}
	}
	return false
}

// SetTable 设置table的样式
func SetTable(table *tablewriter.Table){
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)
}

