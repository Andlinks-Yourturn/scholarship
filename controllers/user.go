package controllers

import (
	"github.com/astaxie/beego"
	"scholarship/models"
	"fmt"
	middleware "scholarship/middlewares"
)

type  UserController struct {
	beego.Controller
}

// @Title Get
// @Description register user
// @Param  name query string true "userName of the user"
// @Param  password query string true "password for basecoin"
// @Success 200 {object} models.user
// @Failure 403 :objectId is empty
// @router / [get]
func (user * UserController)Get(){
	var result models.ApiResult
	fmt.Println("post")
	name := user.GetString("name")
	password := user.GetString("password")
	var userInstance models.User
	//info, err := util.NewKeys(userName, password)
	//"http://localhost:46600/register?name=&password="

	body, err := middleware.Register(name, password)
	if err != nil{
		result.Result = "false"
	}else if string(body) == "false" {
		result.Result = "false"
	}else{
		// success to register
		userInstance.Username = name
		userInstance.Address = string(body)

		//add new account into online system by do a transaction with it
		originalUser := beego.AppConfig.String("originalUser")
		originalUserPassword := beego.AppConfig.String("originalUserPassword")
		middleware.SendTx(originalUser, originalUserPassword, string(body), "1mycoin")

		//get money back
		originalUserAddress := beego.AppConfig.String("originalUserAddress")
		middleware.SendTx(name, password, originalUserAddress, "1mycoin")


		result.Result = "true"
	}
	result.Data = userInstance
	user.ServeJSON()

}
func (register * UserController)GetAll(){

}