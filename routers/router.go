// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"scholarship/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/register",
			beego.NSInclude(
				&controllers.RegisterController{},
			),
		),
		beego.NSNamespace("/sendTx",
			beego.NSInclude(
				&controllers.SendTxController{},
			),
		),
		beego.NSNamespace("/ipfs",
			beego.NSInclude(
				&controllers.IpfsController{},
			),
		),
		beego.NSNamespace("/student",
			beego.NSInclude(
				&controllers.StudentController{},
			),
		),
		beego.NSNamespace("/signature",
			beego.NSInclude(
				&controllers.SignatureController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
