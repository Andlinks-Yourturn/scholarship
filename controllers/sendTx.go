package controllers

import (
	"github.com/astaxie/beego"
	util  "github.com/tendermint/basecoin/util"
	"fmt"
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

	err :=  util.SendTx(userFrom, password, money, userTo,"test_chain_id")
	util.QueryBalance(userTo, "localhost", "46657")
	if err != nil{
		sendTX.Data["json"] = err
	}else{
		sendTX.Data["json"] = "OK"
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
	userFrom := sendTX.GetString("userFrom")
	password := sendTX.GetString("password")
	userTo := sendTX.GetString("userTo")
	money := sendTX.GetString("money")
	//util.SendTx("ligang", "1234567890", "1mycoin", "19D4B36BAAA7B203B301CB86F543EB2F49E34D39", "test_chain_id")
	fmt.Println(userFrom+userTo+password+money)
	err :=  util.SendTx(userFrom, password, money, userTo,"test_chain_id")
	fmt.Println("query balance")
	util.QueryBalance(userTo, "localhost", "46657")
	if err != nil{
		sendTX.Data["json"] = err
	}else{
		sendTX.Data["json"] = "OK"
	}
	sendTX.ServeJSON()

}
