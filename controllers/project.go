package controllers

import (
	middleware "scholarship/middlewares"
	"github.com/astaxie/beego"
	"fmt"
	"io/ioutil"
	"encoding/hex"
	"scholarship/models"
)

// Operations about object
type ProjectController struct {
	beego.Controller
}


// @Title Create
// @Description create project
// @Param	name		query 	string	true		"The name for project"
// @Param	password		query 	string	true		"The password for project"
// @Param	self_name		query 	string	true		"The name for himself"
// @Param	self_password		query 	string	true		"The password for himself"
// @Param	value		query 	string	true		"The value for project, like "1000mycoin""
// @Param	projectInfo		query 	string	true		"body for project content"
// @Success 200 {string} models.Project.Address
// @Failure 403 body is empty
// @router / [get]
func (p *ProjectController) CreateProject() {

	// 调用 生成address 的函数，生成address
	var result models.ApiResult
	//"http://localhost:46600/register?name=&password="
	name := p.GetString("name")
	password := p.GetString("password")
	projectInfo := p.GetString("projectInfo")
	projectInfoForByte, err := hex.DecodeString(projectInfo)
	if err != nil{
		fmt.Println(err)
		result.Result = "false"
		result.Data = " hex.DecodeString error"
		p.Data["json"] = result
		p.ServeJSON()
	}


	body, err := middleware.Register(name, password)

	//add new account into online system by do a transaction with it
	originalUser := beego.AppConfig.String("originalUser")
	originalUserPassword := beego.AppConfig.String("originalUserPassword")
	middleware.SendTx(originalUser, originalUserPassword, string(body), "1mycoin")

	//get money back
	originalUserAddress := beego.AppConfig.String("originalUserAddress")
	middleware.SendTx(name, password, originalUserAddress, "1mycoin")

	address := hex.EncodeToString(body)

	// 调用 ipfs 函数，
	err = ioutil.WriteFile("tmp/output.json", projectInfoForByte,0666)
	hash, err := middleware.IpfsUpload("tmp/output.json")

	if err !=nil{
		fmt.Println(err)
		result.Result = "false"
		result.Data = "client.Do(reqest) or   ioutil.WriteFile IpfsUpload or ioutil.ReadAll error"
		p.Data["json"] = result
		p.ServeJSON()
	}


	err = middleware.InsertIntoMT(beego.AppConfig.String("MTUrl"), address,hash)
	if err != nil {
		result.Result = "false"
		result.Data = " InsertIntoMT error"
		p.Data["json"] = result
	}else {
		// 开始转账
		//"http://localhost:46600/sendTx?userFrom=&password=&money=&userToAddress"
		//the value of money should be like "1000mycoin"
		value := p.GetString("value")
		self_name := p.GetString("self_name")
		self_password := p.GetString("self_password")
		body, err := middleware.SendTx(self_name, self_password, address, value)

		if err != nil{
			result.Result = "false"
			result.Data = "err for result"
		}else if string(body) =="true"{
			var createProjectResult models.CreateProjectResult
			result.Result = "true"
			createProjectResult.DocAddress = hash
			createProjectResult.ProAddress = address
			result.Data = createProjectResult
		}else{
			fmt.Println("转账失败")
			result.Result = "false"
			result.Data = "fail to sendTx"
		}
	}


	p.Data["json"] = result
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

