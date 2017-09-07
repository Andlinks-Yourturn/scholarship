package middlewares

import (
	"github.com/astaxie/beego"
)

func Sign(name string, password string, message string)([]byte, error){
	url := beego.AppConfig.String("BasecoinUrl")+ "/sign?name="+name+"&password="+password+"&message="+message

	return SendRequest(url)
}

func Verify(name string, signature string, message string)([]byte, error){
	url := beego.AppConfig.String("BasecoinUrl")+ "/verify?name="+name+"&signature="+signature+"&message="+message

	return SendRequest(url)

}
