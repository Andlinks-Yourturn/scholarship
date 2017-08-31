package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"fmt"
	"scholarship/models"
	"net/http"
)

type AccountController struct {
	beego.Controller
}

// @Title Get
// @Descripti  search for the relate account
// @Param  address query string true "address of the user"
// @Success 200 {object} models.user
// @Failure 403 :objectId is empty
// @router / [get]

func (account *AccountController) Search(){
	var result models.ApiResult
	address := account.GetString("address")
	fmt.Println("Search")

	url := beego.AppConfig.String("BasecoinUrl")+"/query/account/relateAccount?address="+address

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
		account.Data["json"] = result
	}else {
		if string(body) == "false" {
			result.Result = "false"
			result.Data = ""
			account.Data["json"] = result
		}else{
			result.Result = "true"
			result.Data = body
			account.Data["json"] = result
		}
	}
	account.ServeJSON()
}
