package middlewares

import (
	"github.com/astaxie/beego"
)

func SendTx(userFrom string, password string, userTo string, money string)([]byte,  error){
	url :=   beego.AppConfig.String("BasecoinUrl")+"/sendTx?userFrom="+userFrom+
		"&password="+password+"&money="+money+"&userToAddress="+userTo
	return SendRequest(url)
}
