package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"mx-zip-codes/models"
)

type MxZipCodesController struct {
	beego.Controller
}

func (this *MxZipCodesController) Get() {

	o := orm.NewOrm()
	o.Using("default")
	
	id, _ := this.Ctx.Input.Params[":objectId"]

    zipCode := new(models.ZipCode)
	qs := o.QueryTable(zipCode)
	
	var zipCodes []*models.ZipCode
	
	_, err := qs.Filter("Codigo", id).OrderBy("Asentamiento").All(&zipCodes)
	
	if err != nil {
	  fmt.Println(err)
	  this.Abort("404")
	} else {
		type FormattedZipCodes struct {
			Codigo string
			Colonias []string
			Municipio string
			Ciudad string
			Estado string
		}
		
		formattedZipCodes := FormattedZipCodes{}
		for _, local := range zipCodes {
			formattedZipCodes.Codigo = local.Codigo
			formattedZipCodes.Ciudad = local.Ciudad
			formattedZipCodes.Municipio = local.Municipio
			formattedZipCodes.Estado = local.Estado
			formattedZipCodes.Colonias = append(formattedZipCodes.Colonias, local.Asentamiento)
		}

		this.Data["json"] = &formattedZipCodes
		this.ServeJson()		
	}
}
