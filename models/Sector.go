package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Sector struct {
	Id   	int32 `orm:"pk"`
	Name 	string
	Intro 	string
	Created int32
}

func init() {
	orm.RegisterModel(new(Sector))	
}

func (m *Sector) Read() error {
	if err := orm.NewOrm().Read(m); err != nil {
		return err
	}
	return nil
}

func SectorList(page int, page_size int) ([]orm.Params, int) {
	
	offset := (page - 1) * page_size
	o := orm.NewOrm()
	qs := o.QueryTable("sector")

	var sectors []orm.Params
	qs.Limit(page_size, offset).Values(&sectors)

	total,_ := qs.Count()
	return sectors, int(total)
}

func SectorSave(name string, intro string) bool {
	
	timestamp := int32(time.Now().Unix())
	sector := Sector{Name:name, Intro:intro, Created:timestamp}

	orm := orm.NewOrm()
	_, err := orm.Insert(&sector)
	if err != nil {
		return false
	}
	return true
}

func SectorDelete(id int32) bool {
	
	sector := Sector{Id:id}
	orm := orm.NewOrm()
	_, err := orm.Delete(&sector)
	if err != nil {
		return false
	}
	return true
}

func Count() int {
	o := orm.NewOrm()
	qs := o.QueryTable("sector")
	c,_ := qs.Count()
	return int(c)
}