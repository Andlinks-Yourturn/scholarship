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
// @Success 200 {string}
// @Failure 403 sign is fail
// @router / [get]
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
	addr := hex.EncodeToString(body)
	if addr =="false"{
		s.Data["json"]="sign  失败"

	}else{
		res.Result = "true"
		res.Data = hex.EncodeToString(body)
	}
	s.Data["json"] = res
	s.ServeJSON()
}

// @Title Get
// @Description find object by objectid
// @Param	name		path 	string	true		"the username you want to verify"
// @Param	password		path 	string	true		"the password you want to verify"
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} verify success!
// @Failure 403 {string} verify failure!
// @router /:username/:password [get]
func (s *SignatureController) Verify() {
	// "http://localhost:46600/verify?name=&signature=&message="
	name := s.GetString("name")
	signature := s.GetString("signature")

	v := &url.Values{}
	v.Set("name",name)
	v.Set("signature",signature)
	v.Set("message",string(s.Ctx.Input.RequestBody))
	url := beego.AppConfig.String("BasecoinUrl")+ "/verify?" + v.Encode()
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
		s.Data["json"] = err
	}else if string(body) == "false"{
		s.Data["json"] = "false"
	}else{
		s.Data["josn"] = "true"
	}
	s.ServeJSON()
}


// @Title Get
// @Description find json by ipfsId
// @Param	objectId		path 	string	true		"the ipfsId you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:ipfsId [get]
func (s *SignatureController) Get() {

	objectId := s.Ctx.Input.Param(":ipfsId")
	if objectId != "" {
		ob, err := models.GetOne(objectId)
		if err != nil {
			s.Data["json"] = err.Error()
		} else {
			s.Data["json"] = ob
		}
	}
	s.ServeJSON()
}
