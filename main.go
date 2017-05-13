package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"glass/models"
	_ "glass/routers"
)

func init() {
	orm.RegisterDataBase("default", "sqlite3", "glass")
	orm.RunSyncdb("default", false, true)
}

func main() {
	if _, ok, _ := models.CheckEmployeeByAccount("huanhuan", "123"); ok {
		fmt.Println("login success")
	} else {
		fmt.Println("login false")
	}
	beego.Run()
}
