package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"io/ioutil"
	"net/http"
	"scholarship/models"
)

type RechargeController struct {
	beego.Controller
}

// @Title Get
// @Description  recharge hide two Paramter userFrom & password
// @Param  userTo  query string  true  "userName of who receive donation"
// @Param  money    query string true "the amount of money to  send "
// @Success 200 {object} models.user
// @Failure 403 :objectId is emptya
// @router / [get]
func (recharge * RechargeController) Get(){
	var result models.ApiResult
	userFrom := beego.AppConfig.String("originalUser")
	password := beego.AppConfig.String("originaUserPassword")
	userTo := recharge.GetString("userTo")
	money := recharge.GetString("money")

	//"http://localhost:46600/sendTx?userFrom=&password=&money=&userToAddress"

	url :=   beego.AppConfig.String("BasecoinUrl")+"/sendTx?userFrom="+userFrom+
		"&password="+password+"&money="+money+"&userToAddress="+userTo
	fmt.Println(url)


	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	response, err := client.Do(reqest)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil{
		result.Result = "false"
		result.Data = "error in http response "
		recharge.Data["json"] = result
	}else if string(body) == "false"{
		result.Result = "false"
		result.Data = "fail to sendTx"
		recharge.Data["json"] =  result
	}else {
		result.Result = "true"
		recharge.Data["json"] = result
	}
	recharge.ServeJSON()
}