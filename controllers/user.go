package controllers

import (
	"github.com/astaxie/beego"
 	"net/http"
	"scholarship/models"
	"fmt"
	"io/ioutil"
	"encoding/hex"
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
	//info, err := util.NewKeys(userName, password)
	//"http://localhost:46600/register?name=&password="

	url := beego.AppConfig.String("BasecoinUrl")+"/register?name="+name+"&password="+password
	fmt.Println(url)

	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, err := client.Do(reqest)
	body, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

	if err != nil{
		result.Result = "false"
		result.Data = ""
		user.Data["json"] = result
	}else {
		if string(body) == "false" {
			result.Result = "false"
			result.Data = ""
			user.Data["json"] = result
		}else{
			var userInstance models.User
			userInstance.Username = name
			userInstance.Address = hex.EncodeToString(body)
			result.Result = "true"
			result.Data = userInstance
			user.Data["json"] = result
		}
	}
	user.ServeJSON()

}
func (register * UserController)GetAll(){

}