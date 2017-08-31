package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"net/http"
	"io/ioutil"
	"scholarship/models"
	"net/url"
	"encoding/hex"
)

// Operations about object
type SignatureController struct {
	beego.Controller
}

// @Title Sign
// @Description sign json
// @Param	name		query 	string	true		"the username you want to sign"
// @Param	password		query 	string	true		"the password you want to sign"
// @Param	body		body 	models.Student	true		"The object content"
// @Success 200 {object} models.ApiResult
// @Failure 403 {object} models.ApiResult
// @router / [post]
func (s *SignatureController) Sign() {
	//  "http://localhost:46600/sign?name=ligang&password=1234567890&message=hello"
	name := s.GetString("name")
	password := s.GetString("password")
	fmt.Println(name,password)
	//get private key from remote basecli
	client := &http.Client{}
	v := &url.Values{}
	v.Set("name",name)
	v.Set("password",password)
	v.Set("message",string(s.Ctx.Input.RequestBody))
    url := beego.AppConfig.String("BasecoinUrl")+ "/sign?" + v.Encode()
	fmt.Println(url)
	fmt.Println(string(s.Ctx.Input.RequestBody))
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, err := client.Do(reqest)
	if err != nil{
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	//res :=&models.ApiResult{}
	var res models.ApiResult
	if string(body) =="false"{
		res.Result = "fa"
		res.Data = "sign fail"
	}else{
		res.Result = "true"
		res.Data = hex.EncodeToString(body)
	}
	s.Data["json"] = res
	s.ServeJSON()
}

// @Title Get
// @Description find object by objectid
// @Param	username		path 	string	true		"the username you want to verify"
// @Param	password		path 	string	true		"the password you want to verify"
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} verify success!
// @Failure 403 {string} verify failure!
// @router /:username/:password [get]
func (o *SignatureController) Verify() {

}


// @Title Get
// @Description find json by ipfsId
// @Param	objectId		path 	string	true		"the ipfsId you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:ipfsId [get]
func (o *SignatureController) Get() {

	objectId := o.Ctx.Input.Param(":ipfsId")
	if objectId != "" {
		ob, err := models.GetOne(objectId)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}
