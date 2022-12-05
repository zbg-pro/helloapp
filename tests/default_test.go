package test

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/smartystreets/goconvey/convey"
	"helloapp/models"
	_ "helloapp/routers"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"
	"time"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestGet is a sample to run an endpoint test
func TestGet(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/object", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	logs.Info("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestDateTime(t *testing.T) {
	u := "user_" + strconv.FormatInt(time.Now().UnixMilli(), 10)
	fmt.Println(u)
	Pwd1 := "admin"
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(Pwd1))
	cipherStr := md5Ctx.Sum(nil)
	Pwd1 = hex.EncodeToString(cipherStr)
	fmt.Println(Pwd1)
	fmt.Println(string(cipherStr))

	orm.RegisterDataBase("default", "mysql", "root:iPYDU0o3MRQOreEW@tcp@tcp(172.16.100.130:3306)/c2c?charset=utf8")
	o := orm.NewOrm()

	//orm.NewOrmUsingDB("dbname")
	var user1 models.User
	o.QueryTable("user").Filter("username", "").One(&user1, "Id")
}

func TestMysqlConn(t *testing.T) {
	//https://www.cnblogs.com/qidaii/articles/15633605.html

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql",
		"root:iPYDU0o3MRQOreEW@tcp(172.16.100.130:3306)/c2c?charset=utf8")

	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("show tables").Values(&maps)
	if err == nil {
		fmt.Println(num)
		fmt.Println(maps)
	}

	var list []orm.ParamsList
	num1, err1 := o.Raw("select * from activity_bill").ValuesList(&list)
	if err1 == nil && num1 > 0 {
		fmt.Println(list) // slene
	}
	/*a := o.Raw("select * from apikey where 1=1")
	fmt.Println(a)

	var ak apikey
	o.QueryTable("apikey").Filter("username", "122").One(&ak, "Id")*/

}

type apikey struct {
	id       int
	userId   string
	userName string
	ipaddrs  string
}
