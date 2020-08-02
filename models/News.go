package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type News struct {
	Id   	int32 `orm:"pk"`
	Title 	string
	Content string
	Created int32
	Updated int32
}

func init() {
	orm.RegisterModel(new(News))	
}

func NewsList(page int, page_size int, id int) ([]orm.Params, int) {
	
	offset := (page - 1) * page_size
	o := orm.NewOrm()
	qs := o.QueryTable("news")

	var news []orm.Params
	if id != 0 {
		qs = qs.Filter("id", id)
	}
	qs.OrderBy("-id").Limit(page_size, offset).Values(&news)

	total,_ := qs.Count()
	return news, int(total)
}

func NewsSave(id int32, title string, content string) int64 {
	
	timestamp := int32(time.Now().Unix())
	
	orm := orm.NewOrm()

	var new News
	if id != 0 {
		new = News{Id:id, Title:title, Content:content, Updated:timestamp}
		sid, err := orm.Update(&new)
		if err != nil {
			return sid
		}
	} else {
		new = News{Title:title, Content:content, Created:timestamp}
		sid, err := orm.Insert(&new)
		if err != nil {
			return sid
		}
	}
	return 0
}

func NewsDelete(id int32) bool {
	
	new := News{Id:id}
	orm := orm.NewOrm()
	_, err := orm.Delete(&new)
	if err != nil {
		return false
	}
	return true
}