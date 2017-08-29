package controllers

import (
	"github.com/astaxie/beego"
	m "scholarship/middlewares"
	"io/ioutil"
	"fmt"
	"scholarship/models"
	"net/http"
	"errors"
	"encoding/hex"
)

// Operations about object
type StudentController struct {
	beego.Controller
}

// @Title Create
// @Description create student
// @Param	name		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Param	sign		query 	string	true		"The sign for login"
// @Param	body		body 	models.Student	true		"body for student content"
// @Success 200 {string} models.Student.Address
// @Failure 403 body is empty
// @router / [post]
func (s *StudentController) Create() {

	// 先verify  后 创建address

	// "http://localhost:46600/verify?name=&signature=&message="
	name := s.GetString("name")
	sign := s.GetString("sign")
	password := s.GetString("password")

	url := beego.AppConfig.String("BasecoinUrl")+"/verify?name="+name+"&sign="+ sign +"&message="+string(s.Ctx.Input.RequestBody)
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, err := client.Do(reqest)
	if err !=nil{
		s.Data["json"] = err
	}else{
		if(response.Status=="true"){
			// 调用 生成address 的函数，生成address
			//info, err := utilPro.NewKeys(s.GetString("username"),s.GetString("password"))
			//if err !=nil {
			//	fmt.Println(err)
			//}
			url := beego.AppConfig.String("BasecoinUrl")+"/register?name="+name +"&password="+password
			fmt.Println(url)
			client := &http.Client{}
			reqest, err := http.NewRequest("GET", url, nil)
			if err != nil {
				panic(err)
			}
			response, err := client.Do(reqest)
			body, err := ioutil.ReadAll(response.Body)
			if err !=nil{
				fmt.Println(err)
			}
			fmt.Println(hex.EncodeToString(body))
			address := hex.EncodeToString(body)

			// 调用 ipfs 函数，
			err = ioutil.WriteFile("tmp/output.json",s.Ctx.Input.RequestBody,0666)
			if err != nil{
				fmt.Println(err)
			}
			hash, err := m.IpfsUpload("tmp/output.json")
			if err!=nil{
				fmt.Println(err)
			}
			err = m.InsertIntoMT(beego.AppConfig.String("MTUrl"),address,hash)
			if err != nil {
				s.Data["json"] = err
			}else {
				s.Data["json"] = address
			}
		}else{
			s.Data["json"] = errors.New("验证失败")
		}
	}
	s.ServeJSON()
}

// @Title Get
// @Description find json by ipfsId
// @Param	objectId		path 	string	true		"the ipfsId you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:ipfsId [get]
func (o *StudentController) Get() {

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

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
//func (o *StudentController) GetAll() {
//	obs := models.GetAll()
//	o.Data["json"] = obs
//	o.ServeJSON()
//}


// 更新学生的ipfs 信息



// @Title Update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
//func (o *StudentController) Put() {
//	objectId := o.Ctx.Input.Param(":objectId")
//	var ob models.Object
//	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
//
//	err := models.Update(objectId, ob.Score)
//	if err != nil {
//		o.Data["json"] = err.Error()
//	} else {
//		o.Data["json"] = "update success!"
//	}
//	o.ServeJSON()
//}