package middlewares

import (
	"github.com/astaxie/beego"
)

func GetBalance(address string) ([]byte, error){
	url := beego.AppConfig.String("BasecoinUrl")+"/query/account/balance?address="+address
	return SendRequest(url)
}

func FindRelateAccount(address string) ([]byte, error){
	url := beego.AppConfig.String("BasecoinUrl")+"/query/account/relateAccount?address="+address
	return SendRequest(url)
}