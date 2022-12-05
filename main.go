package main

import (
	"github.com/beego/beego/v2/client/orm"
	_ "helloapp/routers"
	"path/filepath"

	beego "github.com/beego/beego/v2/server/web"
)

var (
	ConfigurationFile = "./conf/app.conf"
	WorkingDirectory  = "./"
	LogFile           = "./runtime/logs"
	BaseUrl           = ""
	AutoLoadDelay     = 0
)

func WorkingDir(elem ...string) string {

	elems := append([]string{WorkingDirectory}, elem...)

	return filepath.Join(elems...)
}

func init() {
	if p, err := filepath.Abs("./conf/app.conf"); err == nil {
		ConfigurationFile = p
	}
	if p, err := filepath.Abs("./"); err == nil {
		WorkingDirectory = p
	}
	if p, err := filepath.Abs("./runtime/logs"); err == nil {
		LogFile = p
	}
}

func main() {
	//配置读取
	// photopath, err := beego.AppConfig.String("photopath")

	orm.RegisterDataBase("default", "mysql", "root:iPYDU0o3MRQOreEW@tcp@tcp(172.16.100.130:3306)/c2c?charset=utf8")

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	} else if beego.BConfig.RunMode == "pro" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/static"] = "static" //这里如果访问http://127.0.0.1:8080/static/ 会访问到static目录
	}

	//beego.BConfig.WebConfig.ViewsPath = WorkingDir("views")
	beego.Run()
}
