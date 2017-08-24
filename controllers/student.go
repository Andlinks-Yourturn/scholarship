package controllers

import (
	"github.com/astaxie/beego"
	m "beeapi/middlewares"
	"io/ioutil"
	"fmt"
	utilPro "github.com/tendermint/basecoin/util"
)

// Operations about object
type StudentController struct {
	beego.Controller
}

// @Title Create
// @Description create student
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Param	body		body 	models.Student	true		"body for student content"
// @Success 200 {string} models.Student.Address
// @Failure 403 body is empty
// @router / [post]
func (s *StudentController) Create() {
	// 调用 生成address 的函数，生成address

	info, err := utilPro.NewKeys(s.GetString("username"),s.GetString("password"))
	if err !=nil {
		fmt.Println(err)
	}



	// 调用 ipfs 函数，
	err = ioutil.WriteFile("tmp/output.json",s.Ctx.Input.RequestBody,0666)
	if err != nil{
		fmt.Println(err)
	}
	hash, err := m.IpfsUpload("tmp/output.json")
	if err!=nil{
		fmt.Println(err)
	}
	err = m.InsertIntoMT(beego.AppConfig.String("MTUrl"),info.Address.String(),hash)
	if err != nil {
		s.Data["json"] = err
	}else {
		s.Data["json"] = info.Address
	}
	s.ServeJSON()

}

// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [get]
//func (o *StudentController) Get() {
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

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
//func (o *StudentController) GetAll() {
//	obs := models.GetAll()
//	o.Data["json"] = obs
//	o.ServeJSON()
//}


// 更新学生的ipfs 信息



// @Title Update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
//func (o *StudentController) Put() {
//	objectId := o.Ctx.Input.Param(":objectId")
//	var ob models.Object
//	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
//
//	err := models.Update(objectId, ob.Score)
//	if err != nil {
//		o.Data["json"] = err.Error()
//	} else {
//		o.Data["json"] = "update success!"
//	}
//	o.ServeJSON()
//}