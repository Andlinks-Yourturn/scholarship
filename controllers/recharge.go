package controllers

import (
	"github.com/astaxie/beego"
	middleware "scholarship/middlewares"
	"scholarship/models"
)

type RechargeController struct {
	beego.Controller
}

// @Title Get
// @Description   test2 recharge hide two Paramter userFrom & password
// @Param  userTo  query string  true  "userName of who receive donation"
// @Param  money    query string true "the amount of money to  send "
// @Success 200 {object} models.user
// @Failure 403 :objectId is emptya
// @router / [get]
func (recharge * RechargeController) Get(){
	var result models.ApiResult
	userFrom := beego.AppConfig.String("originalUser")
	password := beego.AppConfig.String("originalUserPassword")
	userTo := recharge.GetString("userTo")
	money := recharge.GetString("money")

	//"http://localhost:46600/sendTx?userFrom=&password=&money=&userToAddress"


	body, err := middleware.SendTx(userFrom, password, userTo, money)
	if err != nil{
		result.Result = "false"
	}else if string(body) == "false"{
		result.Result = "false"
	}else {
		result.Result = "true"
	}
	result.Data = body
	recharge.Data["json"] = result
	recharge.ServeJSON()
}
