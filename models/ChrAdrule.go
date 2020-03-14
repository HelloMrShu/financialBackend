package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type ChrAdrule struct {
	Id   		int32 `orm:"pk"`
	Platform 	string
	Bidmode 	string
	Name 		string
	Media  		string
	Special_type string
	Content 	string
	Is_delete	int8
}

func init() {
	orm.RegisterModel(new(ChrAdrule))	
}

func (m *ChrAdrule) Read() error {
	if err := orm.NewOrm().Read(m); err != nil {
		return err
	}
	return nil
}

func TestList(
	page int,
	page_size int,
	strCond map[string]string,
	intCond map[string]int) []orm.Params {
	
	offset := (page - 1) * page_size
	o := orm.NewOrm()
	qs := o.QueryTable("chr_adrule")

	if strCond["name"] != "" {
		qs = qs.Filter("name__contains", strCond["name"])
	}

	if strCond["special_type"] != "" {
		qs = qs.Filter("special_type__contains", strCond["special_type"])
	}

	if intCond["media"] != 0 {
		qs = qs.Filter("Media__exact", intCond["media"])
	}

	var rules []orm.Params
	qs.Limit(page_size, offset).Values(&rules)
	return rules
}

func Count(strCond map[string]string, intCond map[string]int) int {
	o := orm.NewOrm()
	qs := o.QueryTable("chr_adrule")
	if strCond["name"] != "" {
		qs = qs.Filter("name__contains", strCond["name"])
	}

	if intCond["media"] != 0 {
		qs = qs.Filter("Media__exact", intCond["media"])
	}

	if strCond["special_type"] != "" {
		qs = qs.Filter("special_type__contains", strCond["special_type"])
	}
	c,_ := qs.Count()
	return int(c)
}