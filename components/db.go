package components

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/config"
	"strings"
	"fmt"
)

func InitDB()  {
	var err error
	BConfig, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil{
			fmt.Println("config init error:", err)
	}

	dbuser := BConfig.String("dev::dbuser")
	dbpwd := BConfig.String("dev::dbpass")
	dbhost := BConfig.String("dev::dbhost")
	dbname := BConfig.String("dev::dbname")
	dbport := BConfig.String("dev::dbport")

	dsn := genDsn(dbuser,dbpwd,dbhost,dbport,dbname)

	orm.Debug = true
	dbdriver := BConfig.String("dev::dbdriver")
	orm.RegisterDataBase("default", dbdriver, dsn, 30)
}

func genDsn(dbuser,dbpwd,dbhost,dbport,dbname string) string {

	ds := make([]string, 0)
	ds = append(ds, dbuser + ":")
	ds = append(ds, dbpwd + "@tcp(")
	ds = append(ds, dbhost + ":")
	ds = append(ds, dbport + ")/")
	ds = append(ds, dbname + "?charset=utf8")

	datasourcename := strings.Join(ds, "")
	return datasourcename
}