package main

import (
	_ "github.com/awethome/server/rest/docs"
	_ "github.com/awethome/server/rest/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/awethome_server")
}

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

