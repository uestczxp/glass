package models

import (
	//    "error"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

//店员

type Employee struct {
	Account  string `orm:"column(account);pk"`
	Name     string `orm:"size(20)"`
	Password string `orm:"size(20)"`
	Type     int32
}

func init() {
	//	orm.RegisterDriver("mysql", orm.DRMySQL)
	//	orm.RegisterDataBase("default", "mysql", "root:zxp914@192.168.1.165/glass?charset=utf8")
	orm.RegisterModel(new(Employee))
	orm.Debug = true
	//	orm.RunSyncdb("default", false, true)
}

func (e *Employee) TableName() string {
	return "employee"
}

func CheckEmployeeByAccount(account, password string) (Employee, bool, error) {
	o := orm.NewOrm()
	var (
		e   Employee
		err error
		ok  bool
	)
	ok = false
	o.Using("default")
	cond := orm.NewCondition()
	cond = cond.And("Account", account).And("Password", password)
	qs := o.QueryTable(&e)
	qs = qs.SetCond(cond)
	if err = qs.One(&e); err == nil {
		ok = true
	}
	return e, ok, err
}
