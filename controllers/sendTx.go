package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"io/ioutil"
	"net/http"
	"scholarship/models"
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
	userFrom := sendTX.GetString("userFrom")
	password := sendTX.GetString("password")
	userTo := sendTX.GetString("userTo")
	money := sendTX.GetString("money")

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
		sendTX.Data["json"] = err
	}else if string(body) == "false"{
		sendTX.Data["json"] = "fail to sendTx"
	}else {
		sendTX.Data["json"] = "true"
	}
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
		sendTX.Data["json"] = result
	}else if string(body) == "false"{
		result.Result = "false"
		result.Data = "fail to sendTx"
		sendTX.Data["json"] = result
	}else {
		result.Result = "true"
		sendTX.Data["json"] = result
	}
	sendTX.ServeJSON()
}
