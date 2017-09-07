package controllers

import (
	"github.com/astaxie/beego"
	"scholarship/models"
	"scholarship/middlewares"
)

type AccountBalanceController struct {
	beego.Controller
}

// @Title Get
// @Description  test search for the relate account
// @Param  address query string true "address of the user"
// @Success 200 {object} models.ApiResult
// @Failure 403 :objectId is empty
// @router / [get]

func (account *AccountBalanceController)GetBalance(){
	var result models.ApiResult
	address := account.GetString("address")

	body, err := middlewares.GetBalance(address)
	if err != nil{
		result.Result = "false"
	}else if string(body) == "false"{
		result.Result = "false"
	}else {
		result.Result = "true"
	}
	result.Data = body
	account.Data["json"] = result
	account.ServeJSON()
}
