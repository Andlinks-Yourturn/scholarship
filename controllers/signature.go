package controllers

import (
	"github.com/astaxie/beego"
	"scholarship/models"
	"scholarship/middlewares"
)

// Operations about object
type SignatureController struct {
	beego.Controller
}

// @Title Sign
// @Description sign json
// @Param	name		query 	string	true		"the username you want to sign"
// @Param	password		query 	string	true		"the password you want to sign"
// @Param	message		query 	string	true		"The  content  of hex"
// @Success 200 {string}
// @Failure 403 sign is fail
// @router / [get]
func (s *SignatureController) Sign() {
	//  "http://localhost:46600/sign?name=ligang&password=1234567890&message=hello"
	var result models.ApiResult
	name := s.GetString("name")
	password := s.GetString("password")
	message := s.GetString("message")


	body, err := middlewares.Sign(name, password, message)
	if err != nil{
		result.Result = "false"
	}else if string(body) == "false"{
		result.Result = "false"
	}else {
		result.Result = "true"
	}
	result.Data = body
	s.Data["json"] = result
	s.ServeJSON()
}

// @Title Get
// @Description find object by objectid
// @Param	name		path 	string	true		"the username you want to verify"
// @Param	password		path 	string	true		"the password you want to verify"
// @Param	message		message 	query 	string	true		"The  content  of hex"
// @Success 200 {string} verify success!
// @Failure 403 {string} verify failure!
// @router /:username/:password [get]
func (s *SignatureController) Verify() {
	// "http://localhost:46600/verify?name=&signature=&message="
	var result models.ApiResult
	name := s.GetString("name")
	signature := s.GetString("signature")
	message := s.GetString("message")

	body, err := middlewares.Verify(name, signature, message)
	if err != nil{
		result.Result = "false"
	}else if string(body) == "false"{
		result.Result = "false"
	}else {
		result.Result = "true"
	}
	result.Data = body
	s.Data["json"] = result
	s.ServeJSON()
}


// @Title Get
// @Description find json by ipfsId
// @Param	objectId		path 	string	true		"the ipfsId you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:ipfsId [get]
func (s *SignatureController) Get() {

	objectId := s.Ctx.Input.Param(":ipfsId")
	if objectId != "" {
		ob, err := models.GetOne(objectId)
		if err != nil {
			s.Data["json"] = err.Error()
		} else {
			s.Data["json"] = ob
		}
	}
	s.ServeJSON()
}
