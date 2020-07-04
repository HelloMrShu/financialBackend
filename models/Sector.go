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
	Updated int32
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

func SectorList(page int, page_size int, id int) ([]orm.Params, int) {
	
	offset := (page - 1) * page_size
	o := orm.NewOrm()
	qs := o.QueryTable("sector")

	var sectors []orm.Params
	if id != 0 {
		qs = qs.Filter("id", id)
	}
	qs.OrderBy("-id").Limit(page_size, offset).Values(&sectors)

	total,_ := qs.Count()
	return sectors, int(total)
}

func SectorSave(id int32, name string, intro string) int64 {
	
	timestamp := int32(time.Now().Unix())
	
	orm := orm.NewOrm()

	var sector Sector
	if id != 0 {
		sector = Sector{Id:id, Name:name, Intro:intro, Updated:timestamp}
		sid, err := orm.Update(&sector)
		if err != nil {
			return sid
		}
	} else {
		sector = Sector{Name:name, Intro:intro, Created:timestamp}
		sid, err := orm.Insert(&sector)
		if err != nil {
			return sid
		}
	}
	
	return 0
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