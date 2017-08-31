package controllers

import (
	m "scholarship/middlewares"

	"github.com/astaxie/beego"
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

	sAddr := s.GetString("sAddr")
	pAddr := s.GetString("pAddr")
	str := m.Comparefiles(beego.AppConfig.String("MTUrl"),pAddr,sAddr)
	if str == "failure" {
		s.Data["json"] = "failure"
	}else if str == "success"{
		s.Data["json"] = "success"
	}else {
		s.Data["json"] = "Matched error"
	}
	s.ServeJSON()
}


