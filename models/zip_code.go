package models

import(
	"github.com/astaxie/beego/orm"
)

type ZipCode struct {
	Id int64
	Codigo string `orm:"column(codigo)"`
	Asentamiento string
	Municipio string
	Ciudad string
	Estado string
}

func init() {
	orm.RegisterModel(new(ZipCode))
}

func (u *ZipCode) TableName() string {
	return "zip_codes"
}