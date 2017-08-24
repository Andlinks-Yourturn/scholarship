package controllers

import (
	"github.com/astaxie/beego"
	util  "github.com/tendermint/basecoin/util"
	"beeapi/models"
	"fmt"
)

type  RegisterController struct {
	beego.Controller
}

// @Title Get
// @Description  register user
// @Param  userName query string true "userName of the user"
// @Param  password query string true "password for basecoin"
// @Success 200 {object} models.user
// @Failure 403 :objectId is empty
// @router / [get]
func (register * RegisterController)Get(){
	fmt.Println("post")
	userName := register.GetString("userName")
	password := register.GetString("password")
	info, err := util.NewKeys(userName, password)
	if err!= nil {
		register.Data["json"] = err.Error()
	}else {
		var user models.User
		user.Username = userName
		user.Address = info.Address.String()
		register.Data["json"] = user
	}
	register.ServeJSON()
}
func (register * RegisterController)GetAll(){

}