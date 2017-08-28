package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about object
type SignatureController struct {
	beego.Controller
}

// @Title Sign
// @Description sign json
// @Param	body		body 	models.Object	true		"The object content"
// @Param	username		path 	string	true		"the username you want to sign"
// @Param	password		path 	string	true		"the password you want to sign"
// @Success 200 {string}
// @Failure 403 sign is fail
// @router /:username/:password [post]
func (s *SignatureController) Sign() {

}

// @Title Get
// @Description find object by objectid
// @Param	username		path 	string	true		"the username you want to verify"
// @Param	password		path 	string	true		"the password you want to verify"
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} verify success!
// @Failure 403 {string} verify failure!
// @router /:username/:password [get]
func (o *SignatureController) Verify() {

}


