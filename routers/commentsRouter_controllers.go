package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["beeapi/controllers:IpfsController"] = append(beego.GlobalControllerRouter["beeapi/controllers:IpfsController"],
		beego.ControllerComments{
			Method: "Uplaod",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beeapi/controllers:IpfsController"] = append(beego.GlobalControllerRouter["beeapi/controllers:IpfsController"],
		beego.ControllerComments{
			Method: "Download",
			Router: `/:ipfsId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beeapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["beeapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beeapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["beeapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beeapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["beeapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beeapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["beeapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beeapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["beeapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beeapi/controllers:RegisterController"] = append(beego.GlobalControllerRouter["beeapi/controllers:RegisterController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beeapi/controllers:ScholarshipController"] = append(beego.GlobalControllerRouter["beeapi/controllers:ScholarshipController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beeapi/controllers:ScholarshipController"] = append(beego.GlobalControllerRouter["beeapi/controllers:ScholarshipController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beeapi/controllers:ScholarshipController"] = append(beego.GlobalControllerRouter["beeapi/controllers:ScholarshipController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beeapi/controllers:ScholarshipController"] = append(beego.GlobalControllerRouter["beeapi/controllers:ScholarshipController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beeapi/controllers:SendTxController"] = append(beego.GlobalControllerRouter["beeapi/controllers:SendTxController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beeapi/controllers:SendTxController"] = append(beego.GlobalControllerRouter["beeapi/controllers:SendTxController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beeapi/controllers:SignatureController"] = append(beego.GlobalControllerRouter["beeapi/controllers:SignatureController"],
		beego.ControllerComments{
			Method: "Sign",
			Router: `/:username/:password`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beeapi/controllers:SignatureController"] = append(beego.GlobalControllerRouter["beeapi/controllers:SignatureController"],
		beego.ControllerComments{
			Method: "Verify",
			Router: `/:username/:password`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beeapi/controllers:StudentController"] = append(beego.GlobalControllerRouter["beeapi/controllers:StudentController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
