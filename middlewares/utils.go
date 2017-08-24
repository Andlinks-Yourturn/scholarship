package middlewares

import (
	"bytes"
	"os/exec"
	"fmt"
	"strings"
)

// 上传ipfs   返回ipfs 地址

func IpfsUpload(path string)(string, error){
	var out bytes.Buffer
	var stderr bytes.Buffer
	args := []string{"add",path}
	cmd := exec.Command("ipfs" , args...)
	cmd.Stdout = &out  //结果
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return "",err
	}
	err = cmd.Run()
	str :=strings.Split(out.String(), " ")[1]
	return str, nil
}

//从ipfs下载 写入到 指定的路径 eg output.json
func IpfsDownload(hash string,output string)(error)  {
	var out,stderr bytes.Buffer
	args := []string{"get",hash,"-o",output}
	cmd := exec.Command("ipfs" , args...)
	cmd.Stdout = &out  //结果
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return err
	}
	return nil
}