package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["scholarship/controllers:IpfsController"] = append(beego.GlobalControllerRouter["scholarship/controllers:IpfsController"],
		beego.ControllerComments{
			Method: "Uplaod",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["scholarship/controllers:IpfsController"] = append(beego.GlobalControllerRouter["scholarship/controllers:IpfsController"],
		beego.ControllerComments{
			Method: "Download",
			Router: `/:ipfsId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["scholarship/controllers:ObjectController"] = append(beego.GlobalControllerRouter["scholarship/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["scholarship/controllers:ObjectController"] = append(beego.GlobalControllerRouter["scholarship/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["scholarship/controllers:ObjectController"] = append(beego.GlobalControllerRouter["scholarship/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["scholarship/controllers:ObjectController"] = append(beego.GlobalControllerRouter["scholarship/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["scholarship/controllers:ObjectController"] = append(beego.GlobalControllerRouter["scholarship/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["scholarship/controllers:ProjectController"] = append(beego.GlobalControllerRouter["scholarship/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "CreateProject",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["scholarship/controllers:RechargeController"] = append(beego.GlobalControllerRouter["scholarship/controllers:RechargeController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["scholarship/controllers:ScholarshipController"] = append(beego.GlobalControllerRouter["scholarship/controllers:ScholarshipController"],
		beego.ControllerComments{
			Method: "Match",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["scholarship/controllers:SendTxController"] = append(beego.GlobalControllerRouter["scholarship/controllers:SendTxController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["scholarship/controllers:SendTxController"] = append(beego.GlobalControllerRouter["scholarship/controllers:SendTxController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["scholarship/controllers:SignatureController"] = append(beego.GlobalControllerRouter["scholarship/controllers:SignatureController"],
		beego.ControllerComments{
			Method: "Sign",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["scholarship/controllers:SignatureController"] = append(beego.GlobalControllerRouter["scholarship/controllers:SignatureController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:ipfsId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["scholarship/controllers:SignatureController"] = append(beego.GlobalControllerRouter["scholarship/controllers:SignatureController"],
		beego.ControllerComments{
			Method: "Verify",
			Router: `/:username/:password`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["scholarship/controllers:StudentController"] = append(beego.GlobalControllerRouter["scholarship/controllers:StudentController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["scholarship/controllers:StudentController"] = append(beego.GlobalControllerRouter["scholarship/controllers:StudentController"],
		beego.ControllerComments{
			Method: "GetRelateStudent",
			Router: `/:name`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["scholarship/controllers:UserController"] = append(beego.GlobalControllerRouter["scholarship/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})
	 beego.GlobalControllerRouter["scholarship/controllers:AccountController"] = append(beego.GlobalControllerRouter["scholarship/controllers:AccountController"],
                beego.ControllerComments{
                        Method: "Search",
                        Router: `/`,
                        AllowHTTPMethods: []string{"get"},
                        MethodParams: param.Make(),
                        Params: nil})	
	beego.GlobalControllerRouter["scholarship/controllers:AccountBalanceController"] = append(beego.GlobalControllerRouter["scholarship/controllers:AccountBalanceController"],
                beego.ControllerComments{
                        Method: "GetBalance",
                        Router: `/`,
                        AllowHTTPMethods: []string{"get"},
                        MethodParams: param.Make(),
                        Params: nil}) 
}
