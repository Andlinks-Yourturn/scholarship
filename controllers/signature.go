package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"net/http"
	"io/ioutil"
)

// Operations about object
type SignatureController struct {
	beego.Controller
}

// @Title Sign
// @Description sign json
// @Param	body		body 	models.student	true		"The object content"
// @Param	username		path 	string	true		"the username you want to sign"
// @Param	password		path 	string	true		"the password you want to sign"
// @Success 200 {string}
// @Failure 403 sign is fail
// @router /:username/:password [post]
func (s *SignatureController) Sign() {
	//  "http://localhost:46600/sign?name=ligang&password=1234567890&message=hello"
	name := s.GetString("name")
	password := s.GetString("password")
	//get private key from remote basecli
	url := beego.AppConfig.String("BasecliServiceUrl")
	client := &http.Client{}
	url += beego.AppConfig.String("BasecoinUrl")+ "/sign?name="+name+"&message="+ string(s.Ctx.Input.RequestBody) +"&password="+password
	fmt.Println(url)
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, err := client.Do(reqest)
	body, err := ioutil.ReadAll(response.Body)
	s.Data["json"]=body
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


