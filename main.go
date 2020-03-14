package main

import (
	_ "template/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/config"
	"fmt"
	"strings"
)

var BConfig config.Configer

func init() {
	var err error
	BConfig, err = config.NewConfig("ini", "conf/app.conf")
	if err != nil{
			fmt.Println("config init error:", err)
	}

	dbuser := BConfig.String("dev::dbuser")
	dbpwd := BConfig.String("dev::dbpass")
	dbhost := BConfig.String("dev::dbhost")
	aedbname := BConfig.String("dev::aedbname")
	dbport := BConfig.String("dev::dbport")

	aeDsn := genDsn(dbuser,dbpwd,dbhost,dbport,aedbname)
	exdbname := BConfig.String("dev::exdbname")
	exDsn :=genDsn(dbuser,dbpwd,dbhost,dbport,exdbname)

	orm.Debug = true
	dbdriver := BConfig.String("dev::dbdriver")
	orm.RegisterDataBase("default", dbdriver, aeDsn, 30)
	orm.RegisterDataBase("exchange", dbdriver, exDsn, 30)
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

func main() {
	beego.Run()
}