package controllers

import (
	"github.com/astaxie/beego"
	m "scholarship/middlewares"
	"io/ioutil"
	"fmt"
	"scholarship/models"
	"encoding/hex"
	middleware "scholarship/middlewares"
)

// Operations about object
type StudentController struct {
	beego.Controller
}

// @Title Create
// @Description create student
// @Param	name		query 	string	true		"The username for login"
// @Param	tea_name		query 	string	true		"The tea_username for login"
// @Param	password		query 	string	true		"The password for student to operator"
// @Param	signature		query 	string	true		"The sign for login"
// @Param	message		query 	string	true		"the detail of student info"
// @Success 200 {string} models.Student.Address
// @Failure 403 body is empty
// @router / [post]
func (s *StudentController) Create() {

	//return two parameter  student address & file address of student info
	// 先verify  后 创建address
	//http://192.168.1.64:46600/
	// "http://localhost:46600/verify?name=&signature=&message="
	var result models.ApiResult
	name := s.GetString("name")
	tea_name := s.GetString("tea_name")
	signature := s.GetString("sign")
	password := s.GetString("password")
	message := s.GetString("message")
	messageByte, err := hex.DecodeString(message)

	body, err := middleware.Verify(tea_name, signature, message)
	fmt.Println(string(body))

	if err !=nil{
		result.Result = "false"
		result.Data = "verify false"
	}else{
		if string(body)!= "false" {
			body, err := middleware.Register(name, password)

			//add new account into online system by do a transaction with it
			originalUser := beego.AppConfig.String("originalUser")
			originalUserPassword := beego.AppConfig.String("originalUserPassword")
			middleware.SendTx(originalUser, originalUserPassword, string(body), "1mycoin")

			//get money back
			originalUserAddress := beego.AppConfig.String("originalUserAddress")
			middleware.SendTx(name, password, originalUserAddress, "1mycoin")

			if err !=nil{
				fmt.Println(err)
			}
			fmt.Println(hex.EncodeToString(body))
			address := hex.EncodeToString(body)

			// 调用 ipfs 函数，
			err = ioutil.WriteFile("tmp/output.json", messageByte, 0666)
			if err != nil{
				fmt.Println(err)
			}
			hash, err := m.IpfsUpload("tmp/output.json")
			if err!=nil{
				fmt.Println(err)
			}
			err = m.InsertIntoMT(beego.AppConfig.String("MTUrl"),address,hash)
			if err != nil {
				result.Result = "false"
			}else {
				var createStudent models.CreateProjectResult
				createStudent.DocAddress = hash
				createStudent.ProAddress = address
				result.Result = "true"
				result.Data = createStudent
			}
		}else{
			result.Result = "false"
			result.Data = ""

		}
	}
	s.Data["json"] = result
	s.ServeJSON()
}

// @Title Get
// @Description find json by ipfsId
// @Param	name		path 	string	true		"the ipfsId you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:name [get]
func (o *StudentController) GetRelateStudent() {

	objectId := o.Ctx.Input.Param(":name")
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
