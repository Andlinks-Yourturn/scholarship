package controllers

import (
	m "scholarship/middlewares"
	"github.com/astaxie/beego"
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/hex"
)

// Operations about object
type ProjectController struct {
	beego.Controller
}


// @Title Create
// @Description create project
// @Param	name		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Param	body		body 	models.Project	true		"body for project content"
// @Success 200 {string} models.Project.Address
// @Failure 403 body is empty
// @router / [post]
func (p *ProjectController) Post() {

	// 调用 生成address 的函数，生成address

	//"http://localhost:46600/register?name=&password="
	name := p.GetString("name")
	password := p.GetString("password")
	url := beego.AppConfig.String("BasecoinUrl")+"/register?name="+name +"&password="+password
	fmt.Println(url)
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, err := client.Do(reqest)
	//info, err := utilPro.NewKeys(p.GetString("username"),p.GetString("password"))
	//if err !=nil {
	//	fmt.Println(err)
	//}
	// 调用 ipfs 函数，
	err = ioutil.WriteFile("tmp/output.json",p.Ctx.Input.RequestBody,0666)
	if err != nil{
		fmt.Println(err)
	}
	hash, err := m.IpfsUpload("tmp/output.json")
	if err!=nil{
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Println(hex.EncodeToString(body))
	address := hex.EncodeToString(body)
	err = m.InsertIntoMT(beego.AppConfig.String("MTUrl"), address,hash)
	if err != nil {
		p.Data["json"] = err
	}else {
		p.Data["json"] = address
	}
	p.ServeJSON()
}

//// @Title Get
//// @Description find object by objectid
//// @Param	objectId		path 	string	true		"the objectid you want to get"
//// @Success 200 {object} models.Object
//// @Failure 403 :objectId is empty
//// @router /:objectId [get]
//func (o *ProjectController) Get() {
//	objectId := o.Ctx.Input.Param(":objectId")
//	if objectId != "" {
//		ob, err := models.GetOne(objectId)
//		if err != nil {
//			o.Data["json"] = err.Error()
//		} else {
//			o.Data["json"] = ob
//		}
//	}
//	o.ServeJSON()
//}

