package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type ChrAdrule struct {
	Id   		int32 `orm:"column(Id)"`
	Platform 	string `orm:"column(Platform)"`
	Bidmode 	string `orm:"column(Bidmode)"`
	Name 		string `orm:"column(Name)"`
	Media  		string `orm:"column(Media)"`
	Special_type string `orm:"column(Special_type)"`
	Content 	string `orm:"column(Content)"`
	Is_delete	int8 `orm:"column(Is_delete)"`
}

var rules []ChrAdrule

func init() {
    orm.Debug = true
    orm.RegisterDataBase("default", "mysql", "dbuser:dbuser@tcp(10.19.37.10:3306)/berry?charset=utf8", 30)
    orm.RegisterModel(new(ChrAdrule))
}

func (m *ChrAdrule) Read() error {
	if err := orm.NewOrm().Read(m); err != nil {
		return err
	}
	return nil
}

func TestList(page int, page_size int) []orm.Params {
	
	offset := (page - 1) * page_size
	o := orm.NewOrm()
	qs := o.QueryTable("chr_adrule")

	var rules []orm.Params
	qs.Limit(page_size, offset).Values(&rules)
	return rules
}

func Count() int {
	o := orm.NewOrm()
	qs := o.QueryTable("chr_adrule")
	c,_ := qs.Count()
	return int(c)
}