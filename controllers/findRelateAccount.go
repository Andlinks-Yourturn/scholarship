package controllers

import (
	"github.com/astaxie/beego"
	"scholarship/models"
	"scholarship/middlewares"
)

type AccountController struct {
	beego.Controller
}

// @Title Get
// @Description search for the relate account
// @Param  address query string true "address of the user"
// @Success 200 {object} models.user
// @Failure 403 :objectId is empty
// @router / [get]

func (account *AccountController) Search(){
	var result models.ApiResult
	address := account.GetString("address")

	body, err := middlewares.FindRelateAccount(address)

	if err != nil{
		result.Result = "false"
		result.Data = "error"
		account.Data["json"] = result
	}else {
		if string(body) == "false" {
			result.Result = "false"
			result.Data = "cannot findRelateAccount"
			account.Data["json"] = result
		}else{
			result.Result = "true"
			result.Data = string(body)
			account.Data["json"] = result
		}
	}
	account.ServeJSON()
}
