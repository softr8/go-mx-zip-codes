package main

import (
	"fmt"
	_ "mx-zip-codes/routers"
	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
	"github.com/astaxie/beego/orm"
	"os"
)

func main() {
    os.Setenv("PGSSLMODE", "disable")
	os.Setenv("PGDATABASE", "mx-zip-codes")
	orm.RegisterDriver("pq", orm.DR_Postgres)
	orm.RegisterDataBase("default", "postgres", "")
	orm.Debug = true

    _, err := orm.GetDB("default")
	if err != nil {
		fmt.Println(err)
	}
	
	beego.Run()
}
