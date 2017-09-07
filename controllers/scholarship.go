package controllers

import (
	middleware "scholarship/middlewares"
	"scholarship/models"
	"github.com/astaxie/beego"
	"fmt"
)

// Operations about object
type ScholarshipController struct {
	beego.Controller
}

// @Title Create
// @Description create object
// @Param	sAddr		query 	string	true		"The student address"
// @Param	pAddr		query 	string	true		"The project address"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [get]
func (s *ScholarshipController) Match() {
	var result models.ApiResult
	sAddr := s.GetString("sAddr")
	pAddr := s.GetString("pAddr")
	fmt.Println("sAddr "+sAddr+"  pAddr"+pAddr)
	str := middleware.Comparefiles(beego.AppConfig.String("MTUrl"),pAddr,sAddr)
	if str == "failure" {
		result.Result = "false"
		result.Data = ""
		s.Data["json"] = result
	}else if str == "success"{
		result.Result = "true"
		result.Data = ""
		s.Data["json"] = "success"
	}else {
		result.Result = "false"
		result.Data = "fail to match"
		s.Data["json"] = result
	}
	s.ServeJSON()
}


