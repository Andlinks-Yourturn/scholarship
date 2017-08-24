package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"os/exec"
	"fmt"
	"os"
	"bytes"
	"strings"
	"github.com/bitly/go-simplejson"
)

// Operations about Ipfs
type IpfsController struct {
	beego.Controller
}

var (
	cmdOut []byte
	err    error
)

// @Title upload
// @Description upload json to IPFS
// @Param	body		body 	models.Ipfs	true		"body for user content"
// @Success 200 {int} models.Ipfs.IpfsId
// @Failure 403 body is empty
// @router / [post]
func (i *IpfsController) Uplaod() {

	//var cmdStr = "ipfs add "+ filePath;

	// 将接受到json 保存为 ipfs对象中
	//var ipfs models.Ipfs = models.Ipfs{}
	//if err = json.Unmarshal(i.Ctx.Input.RequestBody, &ipfs);err!=nil{
	//	fmt.Println("Unmarshal:",err.Error())
	//}
	err = ioutil.WriteFile("tmp/output.json",i.Ctx.Input.RequestBody,0666)
	if err != nil{
		fmt.Println(err)
	}
	// 上传至ipfs
	args := []string{"add","tmp/output.json"}
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("ipfs" , args...)
	cmd.Stdout = &out  //结果
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	}
	err = cmd.Run()
	i.Data["json"]=strings.Split(out.String(), " ")[1]
	i.ServeJSON()
}


// @Title download
// @Description download json from IPFS
// @Param	ipfsId		path 	string	true		"ipfsId for json"
// @Success 200 {object} models.Ipfs
// @Failure 403 ipfsId is error
// @router /:ipfsId [get]
func (i *IpfsController) Download(){
	ipfsId := i.Ctx.Input.Param(":ipfsId")

	// 从ipfs 下载json
	args := []string{"get",ipfsId,"-o","tmp/output.json"}
	if cmdOut, err = exec.Command("ipfs", args...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git rev-parse command: ", err)
	}
	fmt.Println(string(cmdOut))
	// 解析json
	data, err := ioutil.ReadFile("tmp/output.json")
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
	}
	js, err :=simplejson.NewJson([]byte(data))
	if err != nil {
		fmt.Println("newJson: ", err.Error())
	}
	i.Data["json"] = js
	i.ServeJSON()
}





