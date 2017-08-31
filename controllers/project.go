package controllers

import (
	m "scholarship/middlewares"
	"github.com/astaxie/beego"
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/hex"
	"scholarship/models"
	"net/url"
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

	var result models.ApiResult
	name := p.GetString("name")
	password := p.GetString("password")
	url1 := beego.AppConfig.String("BasecoinUrl")+"/register?name="+name +"&password="+password
	fmt.Println(url1)
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url1, nil)
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
		result.Result = "false"
		result.Data = err
		p.Data["json"] = result
	}else {
		// 开始转账
		//"http://localhost:46600/sendTx?userFrom=&password=&money=&userToAddress"
		//the value of money should be like "1000mycoin"
		value := p.GetString("value")
		self_name := p.GetString("self_name")
		self_password := p.GetString("self_password")

		//whether balance is enough to create
		urlForBalance := beego.AppConfig.String("BasecoinUrl")+"/query/account/balance?address="


		v := &url.Values{}
		v.Set("userToAddress",address)
		v.Set("userFrom",self_name)
		v.Set("password",self_password)
		v.Set("money",value)
		url2  := beego.AppConfig.String("BasecoinUrl")+"/sendTx?"+v.Encode()
		fmt.Println(url2)
		request, err := http.NewRequest("GET", url2, nil)
		if err != nil {
			panic(err)
		}
		//处理返回结果
		response, err := client.Do(request)
		if err != nil{
			fmt.Println(err)
		}
		body, err := ioutil.ReadAll(response.Body)
		if string(body) =="true"{
			result.Result = "true"
			result.Data = address
			p.Data["json"] = result
		}else{
			fmt.Println("转账失败")
			result.Result = "false"
			result.Data = address
			p.Data["json"] = address
		}
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

