package controllers

import (
	"github.com/astaxie/beego"
	"scholarship/models"
	middleware "scholarship/middlewares"
)

type SendTxController struct {
	beego.Controller
}



// @Title Post
// @Description  sendTx
// @Param  userFrom query string true "userName of the user who donate money"
// @Param  password query string true "password for basecoin"
// @Param  userTo  query string  true  "userName of who receive donation"
// @Param  money    query string true "the amount of money to  send "
// @Success 200 {object} models.user
// @Failure 403 :objectId is empty
// @router / [post]
func (sendTX * SendTxController) Post(){
	var result models.ApiResult
	userFrom := sendTX.GetString("userFrom")
	password := sendTX.GetString("password")
	userTo := sendTX.GetString("userTo")
	money := sendTX.GetString("money")


	body, err := middleware.SendTx(userFrom, password, userTo, money)
	if err != nil{
		result.Result = "false"
	}else if string(body) == "false"{
		result.Result = "false"
	}else {
		result.Result = "true"
	}
	result.Data = body
	sendTX.Data["json"] = result
	sendTX.ServeJSON()
}



// @Title Get
// @Description  sendTx
// @Param  userFrom query string true "userName of the user who donate money"
// @Param  password query string true "password for basecoin"
// @Param  userTo   query string true "userName of who receive donation"
// @Param  money    query string true "the amount of money to  send "
// @Success 200 {object} models.user
// @Failure 403 :objectId is empty
// @router / [get]
func (sendTX * SendTxController) Get(){
	var result models.ApiResult
	userFrom := sendTX.GetString("userFrom")
	password := sendTX.GetString("password")
	userTo := sendTX.GetString("userTo")
	money := sendTX.GetString("money")
	body, err := middleware.SendTx(userFrom, password, userTo, money)
	if err != nil{
		result.Result = "false"
	}else if string(body) == "false"{
		result.Result = "false"
	}else {
		result.Result = "true"
	}
	result.Data = body
	sendTX.Data["json"] = result
	sendTX.ServeJSON()
}
