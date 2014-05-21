package routers

import (
	"mx-zip-codes/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.RESTRouter("/zip_codes", &controllers.MxZipCodesController{})
}