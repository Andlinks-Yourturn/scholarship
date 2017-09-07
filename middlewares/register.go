package middlewares

import (
	"github.com/astaxie/beego"
)

func Register(name string, password string) ([]byte, error){
	url := beego.AppConfig.String("BasecoinUrl")+"/register?name="+name+"&password="+password

	body, err := SendRequest(url)

	if err != nil {
		return  []byte("error in  register http "), err
	}else if string(body) == "false"{
		return []byte("false"), nil
	}else{
		return body, nil
	}
}

