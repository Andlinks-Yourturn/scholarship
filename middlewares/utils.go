package middlewares

import (
	"bytes"
	"os/exec"
	"fmt"
	"strings"
	"io/ioutil"
	"net/http"
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

func SendRequest(url string) ([]byte, error){
	fmt.Println(url)
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		//		panic(err)
	}
	//处理返回结果
	response, err := client.Do(reqest)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil{
		return []byte("error"), err
	}else if string(body) == "false"{
		return []byte("false"), nil
	}else{
		return body, nil
	}
}